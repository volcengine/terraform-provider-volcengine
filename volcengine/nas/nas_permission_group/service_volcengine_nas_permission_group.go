package nas_permission_group

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

type VolcengineNasPermissionGroupService struct {
	Client *ve.SdkClient
}

func (v *VolcengineNasPermissionGroupService) GetClient() *ve.SdkClient {
	return v.Client
}

func (v *VolcengineNasPermissionGroupService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribePermissionGroups"
		condition["FileSystemType"] = "Extreme"
		bytes, _ := json.Marshal(condition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
		resp, err = v.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
		if err != nil {
			return data, err
		}
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))
		results, err = ve.ObtainSdkValue("Result.PermissionGroups", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.PermissionGroups is not slice")
		}
		for index, ele := range data {
			action = "DescribePermissionRules"
			permissionGroup, ok := ele.(map[string]interface{})
			if !ok {
				return data, errors.New("Result.PermissionGroup is not a map")
			}
			permissionGroupId := permissionGroup["PermissionGroupId"]
			req := map[string]interface{}{
				"PermissionGroupId": permissionGroupId,
			}
			logger.Debug(logger.ReqFormat, action, req)
			resp, err = v.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
			if err != nil {
				return data, err
			}
			logger.Debug(logger.RespFormat, action, resp, *resp)
			results, err = ve.ObtainSdkValue("Result.PermissionRules", *resp)
			if err != nil {
				return data, err
			}
			data[index].(map[string]interface{})["PermissionRules"] = results
		}
		return data, err
	})
}

func (v *VolcengineNasPermissionGroupService) ReadResource(resData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = v.ReadResourceId(resData.Id())
	}
	req := map[string]interface{}{
		"Filters": []map[string]interface{}{
			{
				"Key":   "PermissionGroupId",
				"Value": id,
			},
		},
	}
	results, err = v.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, r := range results {
		if data, ok = r.(map[string]interface{}); !ok {
			return data, errors.New("value is not map")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("PermissionGroup %s does not exist", id)
	}
	return data, err
}

func (v *VolcengineNasPermissionGroupService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, s string) *resource.StateChangeConf {
	return nil
}

func (v *VolcengineNasPermissionGroupService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		if rules, ok := m["PermissionRules"].([]interface{}); ok {
			newRules := make([]interface{}, 0)
			for _, rule := range rules {
				ruleMap := rule.(map[string]interface{})
				if mode, ok := ruleMap["UserMode"]; ok {
					ruleMap["UseMode"] = mode
				}
				delete(ruleMap, "UserMode")
				newRules = append(newRules, ruleMap)
			}
			m["PermissionRules"] = newRules
		}
		return m, map[string]ve.ResponseConvert{}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (v *VolcengineNasPermissionGroupService) CreateResource(resData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callbacks := make([]ve.Callback, 0)
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreatePermissionGroup",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"permission_rules": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				// 仅支持Extreme
				(*call.SdkParam)["FileSystemType"] = "Extreme"
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.PermissionGroupId", *resp)
				d.SetId(id.(string))
				return nil
			},
		},
	}
	callbacks = append(callbacks, callback)
	if rules, ok := resData.GetOk("permission_rules"); ok {
		ruleCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "UpdatePermissionRule",
				ContentType: ve.ContentTypeJson,
				ConvertMode: ve.RequestConvertIgnore,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					rulesValue := make([]map[string]interface{}, 0)
					for _, rule := range rules.(*schema.Set).List() {
						ruleMapValue := make(map[string]interface{})
						if ruleMap, ok := rule.(map[string]interface{}); ok {
							ruleMapValue["CidrIp"] = ruleMap["cidr_ip"]
							ruleMapValue["RwMode"] = ruleMap["rw_mode"]
							ruleMapValue["UserMode"] = ruleMap["use_mode"]
						}
						rulesValue = append(rulesValue, ruleMapValue)
					}
					(*call.SdkParam)["PermissionRules"] = rulesValue
					(*call.SdkParam)["FileSystemType"] = "Extreme"
					(*call.SdkParam)["PermissionGroupId"] = resData.Id()
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return v.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
			},
		}
		callbacks = append(callbacks, ruleCallback)
	}
	return callbacks
}

func (v *VolcengineNasPermissionGroupService) ModifyResource(resData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callbacks := make([]ve.Callback, 0)
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdatePermissionGroup",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"permission_group_name": {
					ConvertType: ve.ConvertDefault,
				},
				"description": {
					ConvertType: ve.ConvertDefault,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["FileSystemType"] = "Extreme"
				(*call.SdkParam)["PermissionGroupId"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	callbacks = append(callbacks, callback)
	if resData.HasChange("permission_rules") {
		ruleCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "UpdatePermissionRule",
				ContentType: ve.ContentTypeJson,
				ConvertMode: ve.RequestConvertIgnore,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["FileSystemType"] = "Extreme"
					(*call.SdkParam)["PermissionGroupId"] = resData.Id()
					rulesValue := make([]map[string]interface{}, 0)
					rules, ok := d.GetOk("permission_rules")
					if ok {
						for _, rule := range rules.(*schema.Set).List() {
							ruleMapValue := make(map[string]interface{})
							if ruleMap, ok := rule.(map[string]interface{}); ok {
								ruleMapValue["CidrIp"] = ruleMap["cidr_ip"]
								ruleMapValue["RwMode"] = ruleMap["rw_mode"]
								ruleMapValue["UserMode"] = ruleMap["use_mode"]
							}
							rulesValue = append(rulesValue, ruleMapValue)
						}
					}
					(*call.SdkParam)["PermissionRules"] = rulesValue
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return v.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
			},
		}
		callbacks = append(callbacks, ruleCallback)
	}
	return callbacks
}

func (v *VolcengineNasPermissionGroupService) RemoveResource(resData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeletePermissionGroup",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"PermissionGroupId": resData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := v.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) && strings.Contains(callErr.Error(), resData.Id()) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading nas permission group on delete %q, %w", d.Id(), callErr))
						}
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, v.ReadResource, 10*time.Minute)
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcengineNasPermissionGroupService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"filters": {
				ConvertType: ve.ConvertJsonObjectArray,
			},
		},
		ContentType:  ve.ContentTypeJson,
		IdField:      "PermissionGroupId",
		NameField:    "PermissionGroupName",
		CollectField: "permission_groups",
	}
}

func (v *VolcengineNasPermissionGroupService) ReadResourceId(s string) string {
	return s
}

func NewVolcengineNasPermissionGroupService(client *ve.SdkClient) *VolcengineNasPermissionGroupService {
	return &VolcengineNasPermissionGroupService{
		Client: client,
	}
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "FileNAS",
		Action:      actionName,
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
	}
}
