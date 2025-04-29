package common

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type ProjectUpdateEnabled interface {
	ProjectTrn() *ProjectTrn
}

type Project struct {
	Client *SdkClient
}

type ProjectTrn struct {
	ResourceType         string
	ResourceID           string
	ServiceName          string
	ProjectSchemaField   string
	ProjectResponseField string
}

func NewProjectService(c *SdkClient) *Project {
	return &Project{
		Client: c,
	}
}

func (p *Project) ModifyProject(trn *ProjectTrn, resourceData *schema.ResourceData, r *schema.Resource, service ResourceService) []Callback {
	var call []Callback
	id := service.ReadResourceId(resourceData.Id())
	if resourceData.HasChange(trn.ProjectSchemaField) {
		modifyProject := Callback{
			Call: SdkCall{
				Action:      "MoveProjectResource",
				ConvertMode: RequestConvertInConvert,
				Convert: map[string]RequestConvert{
					trn.ProjectSchemaField: {
						ConvertType: ConvertDefault,
						TargetField: "TargetProjectName",
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *SdkClient, call SdkCall) (bool, error) {
					if (*call.SdkParam)["TargetProjectName"] == nil || (*call.SdkParam)["TargetProjectName"] == "" {
						return false, fmt.Errorf("Could set ProjectName to empty ")
					}
					//获取用户ID
					input := map[string]interface{}{
						"ProjectName": (*call.SdkParam)["TargetProjectName"],
					}
					logger.Debug(logger.ReqFormat, "GetProject", input)
					out, err := p.Client.UniversalClient.DoCall(p.getUniversalInfo("GetProject"), &input)
					if err != nil {
						return false, err
					}
					accountId, err := ObtainSdkValue("Result.AccountID", *out)
					if err != nil {
						return false, err
					}

					var trnStr string
					if trn.ServiceName == "tos" && trn.ResourceType == "bucket" {
						// tos bucket 特殊处理
						trnStr = fmt.Sprintf("trn:%s:%s:%d:%s", trn.ServiceName, p.Client.Region, int(accountId.(float64)), id)
					} else if trn.ServiceName == "transitrouter" && trn.ResourceType == "transitrouterbandwidthpackage" {
						// transit router bandwidth package 特殊处理
						trnStr = fmt.Sprintf("trn:%s:%s:%d:%s/%s", trn.ServiceName, "", int(accountId.(float64)),
							trn.ResourceType, id)
					} else if trn.ServiceName == "dns" && trn.ResourceType == "zone" {
						// dns zone 特殊处理
						trnStr = fmt.Sprintf("trn:%s:%s:%d:%s/%s", trn.ServiceName, "", int(accountId.(float64)),
							trn.ResourceType, id)
					} else if trn.ServiceName == "cr" && trn.ResourceType == "repository" {
						// cr namespace 特殊处理
						ids := strings.Split(id, ":")
						if len(ids) != 2 {
							return false, fmt.Errorf("invalid cr namespace id:%s", id)
						}
						newId := strings.Join(ids, "/")
						trnStr = fmt.Sprintf("trn:%s:%s:%d:%s/%s", trn.ServiceName, p.Client.Region, int(accountId.(float64)),
							trn.ResourceType, newId)
					} else {
						trnStr = fmt.Sprintf("trn:%s:%s:%d:%s/%s", trn.ServiceName, p.Client.Region, int(accountId.(float64)),
							trn.ResourceType, id)
					}
					(*call.SdkParam)["ResourceTrn.1"] = trnStr
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *SdkClient, call SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return p.Client.UniversalClient.DoCall(p.getUniversalInfo(call.Action), call.SdkParam)
				},
				Refresh: &StateRefresh{
					Target:  []string{resourceData.Get(trn.ProjectSchemaField).(string)},
					Timeout: resourceData.Timeout(schema.TimeoutCreate),
				},
				refreshState: func(data *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
					return &resource.StateChangeConf{
						Pending:    []string{},
						Delay:      1 * time.Second,
						MinTimeout: 1 * time.Second,
						Target:     target,
						Timeout:    timeout,
						Refresh: func() (result interface{}, state string, err error) {
							var (
								d    map[string]interface{}
								name interface{}
							)
							d, err = service.ReadResource(resourceData, service.ReadResourceId(id))
							if err != nil {
								return nil, "", err
							}
							name, err = ObtainSdkValue(trn.ProjectResponseField, d)
							if err != nil {
								return nil, "", err
							}

							return d, name.(string), err
						},
					}
				},
			},
		}
		call = append(call, modifyProject)
	}
	return call
}

func (p *Project) ModifyProjectOld(trn ProjectTrn, resourceData *schema.ResourceData, resource *schema.Resource, key string, sr *StateRefresh) []Callback {
	var call []Callback
	if resourceData.HasChange(key) {
		modifyProject := Callback{
			Call: SdkCall{
				Action:      "MoveProjectResource",
				ConvertMode: RequestConvertInConvert,
				Convert: map[string]RequestConvert{
					key: {
						ConvertType: ConvertDefault,
						TargetField: "TargetProjectName",
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *SdkClient, call SdkCall) (bool, error) {
					if (*call.SdkParam)["TargetProjectName"] == nil || (*call.SdkParam)["TargetProjectName"] == "" {
						return false, fmt.Errorf("Could set ProjectName to empty ")
					}
					//获取用户ID
					input := map[string]interface{}{
						"ProjectName": (*call.SdkParam)["TargetProjectName"],
					}
					logger.Debug(logger.ReqFormat, "GetProject", input)
					out, err := p.Client.UniversalClient.DoCall(p.getUniversalInfo("GetProject"), &input)
					if err != nil {
						return false, err
					}
					accountId, err := ObtainSdkValue("Result.AccountID", *out)
					if err != nil {
						return false, err
					}
					trnStr := fmt.Sprintf("trn:%s:%s:%d:%s/%s", trn.ServiceName, p.Client.Region, int(accountId.(float64)),
						trn.ResourceType, trn.ResourceID)
					(*call.SdkParam)["ResourceTrn.1"] = trnStr
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *SdkClient, call SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return p.Client.UniversalClient.DoCall(p.getUniversalInfo(call.Action), call.SdkParam)
				},
				Refresh: sr,
			},
		}
		call = append(call, modifyProject)
	}
	return call
}

func (p *Project) getUniversalInfo(actionName string) UniversalInfo {
	return UniversalInfo{
		ServiceName: "iam",
		Action:      actionName,
		Version:     "2021-08-01",
		HttpMethod:  GET,
		ContentType: Default,
		RegionType:  Global,
	}
}
