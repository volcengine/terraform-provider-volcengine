package cloud_identity_user_provisioning

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineCloudIdentityUserProvisioningService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewCloudIdentityUserProvisioningService(c *ve.SdkClient) *VolcengineCloudIdentityUserProvisioningService {
	return &VolcengineCloudIdentityUserProvisioningService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineCloudIdentityUserProvisioningService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineCloudIdentityUserProvisioningService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListUserProvisionings"

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

		results, err = ve.ObtainSdkValue("Result.UserProvisionings", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.UserProvisionings is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineCloudIdentityUserProvisioningService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		result    interface{}
		policy    interface{}
		policyArr []interface{}
		policyMap map[string]interface{}
		ok        bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	req := map[string]interface{}{
		"UserProvisioningId": id,
	}
	action := "GetUserProvisioning"
	logger.Debug(logger.ReqFormat, action, req)
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, *resp)
	result, err = ve.ObtainSdkValue("Result", *resp)
	if err != nil {
		return data, err
	}
	if data, ok = result.(map[string]interface{}); !ok {
		return data, errors.New("Value is not map ")
	}
	if len(data) == 0 {
		return data, fmt.Errorf("cloud_identity_user_provisioning %s not exist ", id)
	}

	policyReq := map[string]interface{}{
		"UserProvisioningIds": []string{id},
	}
	policyAction := "ListPoliciesByUserProvision"
	logger.Debug(logger.ReqFormat, policyAction, policyReq)
	policyResp, err := s.Client.UniversalClient.DoCall(getPostUniversalInfo(policyAction), &policyReq)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, policyAction, *policyResp)
	policy, err = ve.ObtainSdkValue("Result.Policy", *policyResp)
	if err != nil {
		return data, err
	}
	if policy == nil {
		policy = []interface{}{}
	}
	if policyArr, ok = policy.([]interface{}); !ok {
		return data, errors.New("Result.Policy is not slice ")
	}
	if len(policyArr) > 0 {
		if policyMap, ok = policyArr[0].(map[string]interface{}); !ok {
			return data, errors.New("Result.Policy Value is not map ")
		}
		if policyName, exist := policyMap["PolicyName"]; exist {
			data["PolicyName"] = []interface{}{policyName}
		}
	}

	return data, err
}

func (s *VolcengineCloudIdentityUserProvisioningService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
			status, err = ve.ObtainSdkValue("ProvisionStatus", d)
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

func (VolcengineCloudIdentityUserProvisioningService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineCloudIdentityUserProvisioningService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateUserProvisioning",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"principal_type": {
					TargetField: "PrincipalType",
				},
				"principal_id": {
					TargetField: "PrincipalId",
				},
				"target_id": {
					TargetField: "TargetId",
				},
				"description": {
					TargetField: "Description",
				},
				"identity_source_strategy": {
					TargetField: "IdentitySourceStrategy",
				},
				"duplication_strategy": {
					TargetField: "DuplicationStrategy",
				},
				"duplication_suffix": {
					TargetField: "DuplicationSuffix",
				},
				"deletion_strategy": {
					TargetField: "DeletionStrategy",
				},
				"policy_name": {
					Ignore: true,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.UserProvisioningId", *resp)
				d.SetId(id.(string))
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
	callbacks = append(callbacks, callback)

	if policy, exist := resourceData.GetOk("policy_name"); exist {
		if policySet, ok := policy.(*schema.Set); ok {
			for _, policyName := range policySet.List() {
				callbacks = append(callbacks, s.policyActionCallback(resourceData, "AttachPolicyToUserProvision", policyName))
			}
		}
	}

	return callbacks
}

func (s *VolcengineCloudIdentityUserProvisioningService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	if resourceData.HasChange("policy_name") {
		add, remove, _, _ := ve.GetSetDifference("policy_name", resourceData, schema.HashString, false)
		for _, element := range add.List() {
			callbacks = append(callbacks, s.policyActionCallback(resourceData, "AttachPolicyToUserProvision", element))
		}
		for _, element := range remove.List() {
			callbacks = append(callbacks, s.policyActionCallback(resourceData, "DetachPolicyToUserProvision", element))
		}
	}

	return callbacks
}

func (s *VolcengineCloudIdentityUserProvisioningService) policyActionCallback(resourceData *schema.ResourceData, action string, element interface{}) ve.Callback {
	return ve.Callback{
		Call: ve.SdkCall{
			Action:      action,
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["UserProvisioningId"] = d.Id()
				(*call.SdkParam)["PolicyName"] = element.(string)
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam, resp)
				return resp, err
			},
			// 必须顺序执行，否则并发失败
			LockId: func(d *schema.ResourceData) string {
				return "lock-CloudIdentity"
			},
		},
	}
}

func (s *VolcengineCloudIdentityUserProvisioningService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteUserProvisioning",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"UserProvisioningId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
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
							return resource.NonRetryableError(fmt.Errorf("error on reading cloud identity user provisioning on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineCloudIdentityUserProvisioningService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		IdField:      "UserProvisioningId",
		CollectField: "user_provisionings",
		ResponseConverts: map[string]ve.ResponseConvert{
			"UserProvisioningId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineCloudIdentityUserProvisioningService) ReadResourceId(id string) string {
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
