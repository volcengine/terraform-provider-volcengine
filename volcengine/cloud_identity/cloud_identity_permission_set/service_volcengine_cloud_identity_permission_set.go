package cloud_identity_permission_set

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/mitchellh/copystructure"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineCloudIdentityPermissionSetService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewCloudIdentityPermissionSetService(c *ve.SdkClient) *VolcengineCloudIdentityPermissionSetService {
	return &VolcengineCloudIdentityPermissionSetService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineCloudIdentityPermissionSetService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineCloudIdentityPermissionSetService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		newCondition map[string]interface{}
		resp         *map[string]interface{}
		results      interface{}
		policies     []interface{}
		ok           bool
	)
	data, err = ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListPermissionSets"

		deepCopyValue, err := copystructure.Copy(condition)
		if err != nil {
			return data, fmt.Errorf(" DeepCopy condition error: %v ", err)
		}
		if newCondition, ok = deepCopyValue.(map[string]interface{}); !ok {
			return data, fmt.Errorf(" DeepCopy condition error: newCondition is not map ")
		}

		// 处理 PermissionSetIds，逗号分离
		if ids, exists := condition["PermissionSetIds"]; exists {
			idsArr, ok := ids.([]interface{})
			if !ok {
				return data, fmt.Errorf(" PermissionSetIds is not slice ")
			}
			permissionSetIds := make([]string, 0)
			for _, id := range idsArr {
				permissionSetIds = append(permissionSetIds, id.(string))
			}
			newCondition["PermissionSetIds"] = strings.Join(permissionSetIds, ",")
		}

		bytes, _ := json.Marshal(newCondition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
		if newCondition == nil {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &newCondition)
			if err != nil {
				return data, err
			}
		}
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, newCondition, string(respBytes))

		results, err = ve.ObtainSdkValue("Result.PermissionSets", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.PermissionSets is not Slice")
		}
		return data, err
	})
	if err != nil {
		return data, err
	}

	for _, v := range data {
		permission, ok := v.(map[string]interface{})
		if !ok {
			return data, fmt.Errorf("Value is not map ")
		}

		req := map[string]interface{}{
			"PermissionSetId": permission["PermissionSetId"].(string),
		}
		// query policy info in permission set
		policies, err = ve.WithPageNumberQuery(req, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
			action := "ListPermissionPoliciesInPermissionSet"

			bytes, _ := json.Marshal(condition)
			logger.Debug(logger.ReqFormat, action, string(bytes))
			resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				logger.Debug(logger.RespFormat, action, condition, err)
				return policies, err
			}
			respBytes, _ := json.Marshal(resp)
			logger.Debug(logger.RespFormat, action, condition, string(respBytes))

			result, err := ve.ObtainSdkValue("Result.PermissionPolicies", *resp)
			if err != nil {
				return policies, err
			}
			if result == nil {
				result = []interface{}{}
			}
			if policies, ok = result.([]interface{}); !ok {
				return policies, errors.New("Result.PermissionPolicies is not Slice")
			}
			return policies, err
		})
		if err != nil {
			return data, err
		}
		permission["PermissionPolicies"] = policies
	}

	return data, err
}

func (s *VolcengineCloudIdentityPermissionSetService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	req := map[string]interface{}{
		"PermissionSetIds": []interface{}{id},
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
		return data, fmt.Errorf("cloud_identity_permission_set %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineCloudIdentityPermissionSetService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineCloudIdentityPermissionSetService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"PermissionPolicyDocument": {
				TargetField: "inline_policy_document",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineCloudIdentityPermissionSetService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreatePermissionSet",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"name": {
					TargetField: "Name",
				},
				"description": {
					TargetField: "Description",
				},
				"relay_state": {
					TargetField: "RelayState",
				},
				"session_duration": {
					TargetField: "SessionDuration",
				},
				"permission_policies": {
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
				id, _ := ve.ObtainSdkValue("Result.PermissionSetId", *resp)
				d.SetId(id.(string))
				return nil
			},
			// 必须顺序执行，否则并发失败
			LockId: func(d *schema.ResourceData) string {
				return "lock-CloudIdentity"
			},
		},
	}
	callbacks = append(callbacks, callback)

	if v, exist := resourceData.GetOk("permission_policies"); exist {
		if policySet, ok := v.(*schema.Set); ok {
			policyArr := policySet.List()
			for _, policy := range policyArr {
				if policyMap, ok := policy.(map[string]interface{}); ok {
					policyType := policyMap["permission_policy_type"].(string)
					if policyType == "System" {
						callbacks = append(callbacks, s.policyActionCallback(resourceData, "AddSystemPolicyToPermissionSet", policyMap))
					} else if policyType == "Inline" {
						callbacks = append(callbacks, s.policyActionCallback(resourceData, "AddInlinePolicyToPermissionSet", policyMap))
					}
				}
			}
		}
	}

	return callbacks
}

func (s *VolcengineCloudIdentityPermissionSetService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdatePermissionSet",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"name": {
					TargetField: "Name",
				},
				"description": {
					TargetField: "Description",
				},
				"relay_state": {
					TargetField: "RelayState",
				},
				"session_duration": {
					TargetField: "SessionDuration",
				},
				"permission_policies": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) > 0 {
					(*call.SdkParam)["PermissionSetId"] = d.Id()
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
			// 必须顺序执行，否则并发失败
			LockId: func(d *schema.ResourceData) string {
				return "lock-CloudIdentity"
			},
		},
	}
	callbacks = append(callbacks, callback)

	if resourceData.HasChange("permission_policies") {
		add, remove, _, _ := ve.GetSetDifference("permission_policies", resourceData, PermissionPolicyHash, false)
		for _, element := range add.List() {
			if policy, ok := element.(map[string]interface{}); ok {
				policyType := policy["permission_policy_type"].(string)
				if policyType == "System" {
					callbacks = append(callbacks, s.policyActionCallback(resourceData, "AddSystemPolicyToPermissionSet", element))
				} else if policyType == "Inline" {
					callbacks = append(callbacks, s.policyActionCallback(resourceData, "AddInlinePolicyToPermissionSet", element))
				}
			}
		}
		for _, element := range remove.List() {
			callbacks = append(callbacks, s.policyActionCallback(resourceData, "RemovePermissionPolicyFromPermissionSet", element))
		}
	}

	return callbacks
}

func (s *VolcengineCloudIdentityPermissionSetService) policyActionCallback(resourceData *schema.ResourceData, action string, element interface{}) ve.Callback {
	return ve.Callback{
		Call: ve.SdkCall{
			Action:      action,
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["PermissionSetId"] = d.Id()
				policy, ok := element.(map[string]interface{})
				if !ok {
					return false, fmt.Errorf("policy is not map")
				}
				policyType := policy["permission_policy_type"].(string)
				(*call.SdkParam)["PermissionPolicyType"] = policyType
				if policyType == "System" {
					(*call.SdkParam)["PermissionPolicyName"] = policy["permission_policy_name"]
				} else if policyType == "Inline" {
					(*call.SdkParam)["InlinePolicyDocument"] = policy["inline_policy_document"]
				}
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

func (s *VolcengineCloudIdentityPermissionSetService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeletePermissionSet",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"PermissionSetId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				// 已部署的 permission set 无法直接删除，需解除部署，直接返回接口报错
				if strings.Contains(baseErr.Error(), "The permission set has provision, please deprovision permission set first") {
					msg := fmt.Sprintf("error: %s.\nmsg: %s",
						baseErr.Error(),
						"If you want to remove this permission set, "+
							"please use `terraform import volcengine_cloud_identity_permission_set_provisioning.resource_name permission_set_id:target_id` command to import `permission_set_provisioning` resource firstly.\n"+
							"Then you can use `terraform destroy` command to delete both `permission_set_provisioning` and `permission_set`.")
					return fmt.Errorf(msg)
				}

				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading cloud identity permission set on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineCloudIdentityPermissionSetService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "PermissionSetIds",
				ConvertType: ve.ConvertJsonArray,
			},
		},
		ContentType:  ve.ContentTypeJson,
		NameField:    "Name",
		IdField:      "PermissionSetId",
		CollectField: "permission_sets",
		ResponseConverts: map[string]ve.ResponseConvert{
			"PermissionSetId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineCloudIdentityPermissionSetService) ReadResourceId(id string) string {
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
