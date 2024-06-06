package cloud_identity_permission_set_assignment

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

type VolcengineCloudIdentityPermissionSetAssignmentService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewCloudIdentityPermissionSetAssignmentService(c *ve.SdkClient) *VolcengineCloudIdentityPermissionSetAssignmentService {
	return &VolcengineCloudIdentityPermissionSetAssignmentService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineCloudIdentityPermissionSetAssignmentService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineCloudIdentityPermissionSetAssignmentService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListAccountAssignments"

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

		results, err = ve.ObtainSdkValue("Result.AccountAssignments", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.AccountAssignments is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineCloudIdentityPermissionSetAssignmentService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	ids := strings.Split(id, ":")
	if len(ids) != 4 {
		return map[string]interface{}{}, fmt.Errorf("invalid permission set attachment id: %v", id)
	}
	permissionSetId := ids[0]
	targetId := ids[1]
	principalType := ids[2]
	principalId := ids[3]

	req := map[string]interface{}{
		"PermissionSetId": permissionSetId,
		"TargetId":        targetId,
		"PrincipalType":   principalType,
		"PrincipalId":     principalId,
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
		return data, fmt.Errorf("cloud_identity_permission_set_assignment %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineCloudIdentityPermissionSetAssignmentService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineCloudIdentityPermissionSetAssignmentService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineCloudIdentityPermissionSetAssignmentService) getTaskStatus(taskId string) (data map[string]interface{}, err error) {
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
		return data, fmt.Errorf("cloud_identity_permission_set_assignment task %s not exist ", taskId)
	}
	return data, nil
}

func (s *VolcengineCloudIdentityPermissionSetAssignmentService) buildTaskStateConf(resourceData *schema.ResourceData, taskId string) *resource.StateChangeConf {
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
					return nil, "", fmt.Errorf("cloud_identity_permission_set_assignment task status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineCloudIdentityPermissionSetAssignmentService) CreateResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	permissionSetId := resourceData.Get("permission_set_id")
	targetId := resourceData.Get("target_id")
	principalType := resourceData.Get("principal_type")
	principalId := resourceData.Get("principal_id")

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateAccountAssignment",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"PermissionSetId": permissionSetId,
				"TargetId":        targetId,
				"PrincipalType":   principalType,
				"PrincipalId":     principalId,
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

				d.SetId(fmt.Sprintf("%v:%v:%v:%v", (*call.SdkParam)["PermissionSetId"], (*call.SdkParam)["TargetId"], (*call.SdkParam)["PrincipalType"], (*call.SdkParam)["PrincipalId"]))
				return nil
			},
			// 必须顺序执行，否则并发失败
			LockId: func(d *schema.ResourceData) string {
				return "lock-CloudIdentity"
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineCloudIdentityPermissionSetAssignmentService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineCloudIdentityPermissionSetAssignmentService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	permissionSetId := resourceData.Get("permission_set_id")
	targetId := resourceData.Get("target_id")
	principalType := resourceData.Get("principal_type")
	principalId := resourceData.Get("principal_id")

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteAccountAssignment",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"PermissionSetId": permissionSetId,
				"TargetId":        targetId,
				"PrincipalType":   principalType,
				"PrincipalId":     principalId,
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
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading cloud identity permission set assignment on delete %q, %w", d.Id(), callErr))
						}
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
			// 必须顺序执行，否则并发失败
			LockId: func(d *schema.ResourceData) string {
				return "lock-CloudIdentity"
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineCloudIdentityPermissionSetAssignmentService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		NameField:    "PermissionSetName",
		IdField:      "PermissionSetId",
		CollectField: "assignments",
		ResponseConverts: map[string]ve.ResponseConvert{
			"PermissionSetId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineCloudIdentityPermissionSetAssignmentService) ReadResourceId(id string) string {
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
