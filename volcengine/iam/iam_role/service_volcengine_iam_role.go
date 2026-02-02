package iam_role

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineIamRoleService struct {
	Client *ve.SdkClient
}

func NewIamRoleService(c *ve.SdkClient) *VolcengineIamRoleService {
	return &VolcengineIamRoleService{
		Client: c,
	}
}

func (s *VolcengineIamRoleService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineIamRoleService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageOffsetQuery(m, "Limit", "Offset", 100, 0, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListRoles"
		logger.Debug(logger.ReqFormat, action, condition)
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

		logger.Debug(logger.RespFormat, action, condition, *resp)

		results, err = ve.ObtainSdkValue("Result.RoleMetadata", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.RoleMetadata is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineIamRoleService) ReadResource(resourceData *schema.ResourceData, roleId string) (data map[string]interface{}, err error) {
	var (
		result interface{}
		ok     bool
	)
	if roleId == "" {
		roleId = s.ReadResourceId(resourceData.Id())
	}
	condition := map[string]interface{}{
		"RoleName": roleId,
	}
	action := "GetRole"
	logger.Debug(logger.ReqFormat, action, condition)
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, condition, *resp)

	result, err = ve.ObtainSdkValue("Result.Role", *resp)
	if err != nil {
		return data, err
	}
	if data, ok = result.(map[string]interface{}); !ok {
		return data, errors.New("value is not map")
	}
	if len(data) == 0 {
		return data, fmt.Errorf("Role %s not exist ", roleId)
	}
	return data, err
}

func (s *VolcengineIamRoleService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (s *VolcengineIamRoleService) WithResourceResponseHandlers(role map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		role["Id"] = role["RoleName"]
		return role, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineIamRoleService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	createIamRoleCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateRole",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"tags": {
					TargetField: "Tags",
					ConvertType: ve.ConvertListN,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				roleName, err := ve.ObtainSdkValue("Result.Role.RoleName", *resp)
				if err != nil {
					return err
				}
				d.SetId(roleName.(string))
				return nil
			},
		},
	}
	return []ve.Callback{createIamRoleCallback}
}

func (s *VolcengineIamRoleService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	updateRoleCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateRole",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"role_name": {
					TargetField: "NewRoleName",
					ConvertType: ve.ConvertDefault,
				},
				"display_name": {
					TargetField: "NewDisplayName",
					ConvertType: ve.ConvertDefault,
				},
				"description": {
					TargetField: "NewDescription",
					ConvertType: ve.ConvertDefault,
				},
				"max_session_duration": {
					TargetField: "MaxSessionDuration",
					ConvertType: ve.ConvertDefault,
				},
				"tags": {
					Ignore: true,
				},
			},
			RequestIdField: "RoleName",
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	callbacks := []ve.Callback{updateRoleCallback}
	setResourceTagsCallbacks := s.setResourceTags(data, "Role", callbacks)
	return setResourceTagsCallbacks
}

func (s *VolcengineIamRoleService) RemoveResource(data *schema.ResourceData, r *schema.Resource) []ve.Callback {
	deleteRoleCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteRole",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["RoleName"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
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
							return resource.NonRetryableError(fmt.Errorf("error on reading iam role on delete %q, %w", d.Id(), callErr))
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
	return []ve.Callback{deleteRoleCallback}
}

func (s *VolcengineIamRoleService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"query": {
				TargetField: "Query",
			},
		},
		ResponseConverts: map[string]ve.ResponseConvert{
			"RoleName": {
				TargetField: "role_name",
			},
			"RoleId": {
				TargetField: "role_id",
			},
			"IsServiceLinkedRole": {
				TargetField: "is_service_linked_role",
			},
			"DisplayName": {
				TargetField: "display_name",
			},
			"MaxSessionDuration": {
				TargetField: "max_session_duration",
			},
			"Tags": {
				TargetField: "tags",
			},
			"Trn": {
				TargetField: "trn",
			},
			"Description": {
				TargetField: "description",
			},
			"TrustPolicyDocument": {
				TargetField: "trust_policy_document",
			},
			"CreateDate": {
				TargetField: "create_date",
			},
			"UpdateDate": {
				TargetField: "update_date",
			},
		},
		NameField:    "RoleName",
		IdField:      "RoleName",
		CollectField: "roles",
	}
}

func (s *VolcengineIamRoleService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "iam",
		Action:      actionName,
		Version:     "2018-01-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		RegionType:  ve.Global,
	}
}

func (s *VolcengineIamRoleService) setResourceTags(resourceData *schema.ResourceData, resourceType string, callbacks []ve.Callback) []ve.Callback {
	addedTags, removedTags, _, _ := ve.GetSetDifference("tags", resourceData, ve.TagsHash, false)

	removeCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UntagResources",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if removedTags != nil && len(removedTags.List()) > 0 {
					(*call.SdkParam)["ResourceNames.1"] = resourceData.Id()
					(*call.SdkParam)["ResourceType"] = resourceType
					for index, tag := range removedTags.List() {
						(*call.SdkParam)["TagKeys."+strconv.Itoa(index+1)] = tag.(map[string]interface{})["key"].(string)
					}
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	callbacks = append(callbacks, removeCallback)

	addCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "TagResources",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if addedTags != nil && len(addedTags.List()) > 0 {
					(*call.SdkParam)["ResourceNames.1"] = resourceData.Id()
					(*call.SdkParam)["ResourceType"] = resourceType
					for index, tag := range addedTags.List() {
						(*call.SdkParam)["Tags."+strconv.Itoa(index+1)+".Key"] = tag.(map[string]interface{})["key"].(string)
						(*call.SdkParam)["Tags."+strconv.Itoa(index+1)+".Value"] = tag.(map[string]interface{})["value"].(string)
					}
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	callbacks = append(callbacks, addCallback)

	return callbacks
}
