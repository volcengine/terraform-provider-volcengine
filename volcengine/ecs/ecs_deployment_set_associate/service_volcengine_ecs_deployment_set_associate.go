package ecs_deployment_set_associate

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineEcsDeploymentSetAssociateService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewEcsDeploymentSetAssociateService(c *ve.SdkClient) *VolcengineEcsDeploymentSetAssociateService {
	return &VolcengineEcsDeploymentSetAssociateService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineEcsDeploymentSetAssociateService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineEcsDeploymentSetAssociateService) ReadResources(condition map[string]interface{}) ([]interface{}, error) {
	return nil, nil
}

func (s *VolcengineEcsDeploymentSetAssociateService) ReadResource(resourceData *schema.ResourceData, tmpId string) (data map[string]interface{}, err error) {
	var (
		resp             *map[string]interface{}
		results          interface{}
		ok               bool
		deploymentSetId  string
		targetInstanceId string
		ids              []string
		dep              []interface{}
		instanceIds      []interface{}
	)

	if tmpId == "" {
		tmpId = s.ReadResourceId(resourceData.Id())
	}

	ids = strings.Split(tmpId, ":")
	deploymentSetId = ids[0]
	targetInstanceId = ids[1]

	req := map[string]interface{}{
		"DeploymentSetIds.1": deploymentSetId,
	}
	client := s.Client.UniversalClient
	action := "DescribeDeploymentSets"
	logger.Debug(logger.ReqFormat, action, req)
	resp, err = client.DoCall(getUniversalInfo(action), &req)
	if err != nil {
		return data, err
	}

	results, err = ve.ObtainSdkValue("Result.DeploymentSets", *resp)
	if err != nil {
		return data, err
	}
	if dep, ok = results.([]interface{}); !ok {
		return data, errors.New("Result.DeploymentSets is not Slice")
	}
	if len(dep) == 0 {
		return data, fmt.Errorf("Ecs DeploymentSet %s not exist ", deploymentSetId)
	}
	results, err = ve.ObtainSdkValue("InstanceIds", dep[0])
	if instanceIds, ok = results.([]interface{}); !ok {
		return data, errors.New("InstanceIds is not Slice")
	}

	for _, id := range instanceIds {
		if id.(string) == targetInstanceId {
			data = make(map[string]interface{})
			data["DeploymentSetId"] = deploymentSetId
			data["InstanceId"] = targetInstanceId
			return data, err
		}
	}

	return data, err
}

func (s *VolcengineEcsDeploymentSetAssociateService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      5 * time.Second,
		MinTimeout: 5 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				data map[string]interface{}
			)

			if err = resource.Retry(3*time.Minute, func() *resource.RetryError {
				data, err = s.ReadResource(resourceData, id)
				if err != nil {
					return resource.NonRetryableError(err)
				}
				if len(data) == 0 {
					return resource.RetryableError(fmt.Errorf("Retry "))
				}
				return nil
			}); err != nil {
				return nil, "error", err
			}
			return data, "success", err
		},
	}
}

func (s *VolcengineEcsDeploymentSetAssociateService) WithResourceResponseHandlers(deploymentSet map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return deploymentSet, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineEcsDeploymentSetAssociateService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyInstanceDeployment",
			ConvertMode: ve.RequestConvertAll,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId((*call.SdkParam)["DeploymentSetId"].(string) + ":" + (*call.SdkParam)["InstanceId"].(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"success"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineEcsDeploymentSetAssociateService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineEcsDeploymentSetAssociateService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyInstanceDeployment",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"InstanceId":      strings.Split(resourceData.Id(), ":")[1],
				"DeploymentSetId": "",
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//删除部署集
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on  reading deployment set associate on delete %q, %w", d.Id(), callErr))
						}
					}

					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineEcsDeploymentSetAssociateService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineEcsDeploymentSetAssociateService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "ecs",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		Action:      actionName,
	}
}
