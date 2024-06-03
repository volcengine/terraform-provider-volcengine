package cloud_identity_permission_set_provisioning

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineCloudIdentityPermissionSetProvisioningService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewCloudIdentityPermissionSetProvisioningService(c *ve.SdkClient) *VolcengineCloudIdentityPermissionSetProvisioningService {
	return &VolcengineCloudIdentityPermissionSetProvisioningService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineCloudIdentityPermissionSetProvisioningService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineCloudIdentityPermissionSetProvisioningService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListPermissionSetProvisionings"

		bytes, _ := json.Marshal(condition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
		if condition == nil {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return data, err
			}
		}
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))

		results, err = ve.ObtainSdkValue("Result.PermissionSetProvisionings", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.PermissionSetProvisionings is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineCloudIdentityPermissionSetProvisioningService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	ids := strings.Split(id, ":")
	if len(ids) != 2 {
		return map[string]interface{}{}, fmt.Errorf("invalid permission set provisioning id: %v", id)
	}
	permissionSetId := ids[0]
	targetId := ids[1]

	req := map[string]interface{}{
		"PermissionSetId": permissionSetId,
		"TargetId":        targetId,
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("cloud_identity_permission_set_provisioning %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineCloudIdentityPermissionSetProvisioningService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				d          map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "Failed")
			d, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("ProvisioningStatus", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("cloud_identity_user_provisioning status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (VolcengineCloudIdentityPermissionSetProvisioningService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineCloudIdentityPermissionSetProvisioningService) getTaskStatus(taskId string) (data map[string]interface{}, err error) {
	var (
		result interface{}
		ok     bool
	)

	action := "GetTaskStatus"
	req := map[string]interface{}{
		"TaskId": taskId,
	}
	logger.Debug(logger.ReqFormat, action, req)
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, resp)
	result, err = ve.ObtainSdkValue("Result", *resp)
	if err != nil {
		return data, err
	}
	if data, ok = result.(map[string]interface{}); !ok {
		return data, errors.New("Value is not map ")
	}
	if len(data) == 0 {
		return data, fmt.Errorf("cloud_identity_permission_set_privisioning task %s not exist ", taskId)
	}
	return data, nil
}

func (s *VolcengineCloudIdentityPermissionSetProvisioningService) buildTaskStateConf(resourceData *schema.ResourceData, taskId string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     []string{"Success"},
		Timeout:    resourceData.Timeout(schema.TimeoutCreate),
		Refresh: func() (result interface{}, state string, err error) {
			var (
				d          map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "Failed")
			d, err = s.getTaskStatus(taskId)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("cloud_identity_permission_set_privisioning task status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineCloudIdentityPermissionSetProvisioningService) CreateResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	permissionSetId := resourceData.Get("permission_set_id")
	targetId := resourceData.Get("target_id")

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ProvisionPermissionSet",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"PermissionSetId": permissionSetId,
				"TargetId":        targetId,
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				// refresh task status
				taskId, err := ve.ObtainSdkValue("Result.TaskId", *resp)
				if err != nil {
					return err
				}
				stateConf := s.buildTaskStateConf(resourceData, taskId.(string))
				_, err = stateConf.WaitForState()
				if err != nil {
					return err
				}

				d.SetId(fmt.Sprintf("%v:%v", (*call.SdkParam)["PermissionSetId"], (*call.SdkParam)["TargetId"]))
				return nil
			},
			//Refresh: &ve.StateRefresh{
			//	Target:  []string{"Provisioned"},
			//	Timeout: resourceData.Timeout(schema.TimeoutCreate),
			//},
			// 必须顺序执行，否则并发失败
			LockId: func(d *schema.ResourceData) string {
				return "lock-CloudIdentity"
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineCloudIdentityPermissionSetProvisioningService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	if resourceData.HasChange("provisioning_status") {
		permissionSetId := resourceData.Get("permission_set_id")
		targetId := resourceData.Get("target_id")

		callback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ProvisionPermissionSet",
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				SdkParam: &map[string]interface{}{
					"PermissionSetId": permissionSetId,
					"TargetId":        targetId,
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					// 实时查询 PermissionSet 状态，如果不为 Provisioned 则重新执行发布
					data, err := s.ReadResource(resourceData, d.Id())
					if err != nil {
						if ve.ResourceNotFoundError(err) {
							return true, nil
						} else {
							return false, err
						}
					}
					status, err := ve.ObtainSdkValue("ProvisioningStatus", data)
					if err != nil {
						return false, err
					}

					if status.(string) != "Provisioned" {
						return true, nil
					}
					return false, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp, err)
					return resp, err
				},
				AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
					// refresh task status
					taskId, err := ve.ObtainSdkValue("Result.TaskId", *resp)
					if err != nil {
						return err
					}
					stateConf := s.buildTaskStateConf(resourceData, taskId.(string))
					_, err = stateConf.WaitForState()
					if err != nil {
						return err
					}
					return nil
				},
				//Refresh: &ve.StateRefresh{
				//	Target:  []string{"Provisioned"},
				//	Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				//},
				// 必须顺序执行，否则并发失败
				LockId: func(d *schema.ResourceData) string {
					return "lock-CloudIdentity"
				},
			},
		}
		callbacks = append(callbacks, callback)
	}

	return callbacks
}

func (s *VolcengineCloudIdentityPermissionSetProvisioningService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	permissionSetId := resourceData.Get("permission_set_id")
	targetId := resourceData.Get("target_id")

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeprovisionPermissionSet",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"PermissionSetId": permissionSetId,
				"TargetId":        targetId,
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				// refresh task status
				taskId, err := ve.ObtainSdkValue("Result.TaskId", *resp)
				if err != nil {
					return err
				}
				stateConf := s.buildTaskStateConf(resourceData, taskId.(string))
				_, err = stateConf.WaitForState()
				if err != nil {
					return err
				}

				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
			// 必须顺序执行，否则并发失败
			LockId: func(d *schema.ResourceData) string {
				return "lock-CloudIdentity"
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineCloudIdentityPermissionSetProvisioningService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "Ids",
				ConvertType: ve.ConvertWithN,
			},
		},
		NameField:    "Name",
		IdField:      "Id",
		CollectField: "instances",
		ResponseConverts: map[string]ve.ResponseConvert{
			"Id": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineCloudIdentityPermissionSetProvisioningService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "cloudidentity",
		Version:     "2023-01-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}

func getPostUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "cloudidentity",
		Version:     "2023-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
