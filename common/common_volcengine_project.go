package common

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type Project struct {
	Client *SdkClient
}

type ProjectTrn struct {
	ResourceType string
	ResourceID   string
	ServiceName  string
}

func NewProjectService(c *SdkClient) *Project {
	return &Project{
		Client: c,
	}
}

func (p *Project) ModifyProject(trn ProjectTrn, resourceData *schema.ResourceData, resource *schema.Resource, key string, sr *StateRefresh) []Callback {
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
		Version:     "2021-08-01",
		HttpMethod:  GET,
		Action:      actionName,
	}
}
