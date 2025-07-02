package apig_route

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

type VolcengineApigRouteService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewApigRouteService(c *ve.SdkClient) *VolcengineApigRouteService {
	return &VolcengineApigRouteService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineApigRouteService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineApigRouteService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListRoutes"

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

		results, err = ve.ObtainSdkValue("Result.Items", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Items is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineApigRouteService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		result interface{}
		ok     bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	action := "GetRoute"
	req := map[string]interface{}{
		"Id": id,
	}
	logger.Debug(logger.ReqFormat, action, req)
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, req, *resp)

	result, err = ve.ObtainSdkValue("Result.Route", *resp)
	if err != nil {
		return data, err
	}

	if data, ok = result.(map[string]interface{}); !ok {
		return data, errors.New("Value is not map ")
	}
	if len(data) == 0 {
		return data, fmt.Errorf("apig_gateway_route %s not exist ", id)
	}

	if v, exist := data["UpstreamList"]; exist {
		if upstreamList, ok := v.([]interface{}); ok {
			for _, upstream := range upstreamList {
				if upstreamMap, ok := upstream.(map[string]interface{}); ok {
					if aiProviderSettings, exist := upstreamMap["AIProviderSettings"]; exist {
						if aiProviderSettingsMap, ok := aiProviderSettings.(map[string]interface{}); ok {
							upstreamMap["AIProviderSettings"] = []interface{}{aiProviderSettingsMap}
						}
					}
				}
			}
		}
	}
	if matchRule, exist := data["MatchRule"]; exist {
		if matchRuleMap, ok := matchRule.(map[string]interface{}); ok {
			if path, exist := matchRuleMap["Path"]; exist {
				if pathMap, ok := path.(map[string]interface{}); ok {
					matchRuleMap["Path"] = []interface{}{pathMap}
				}
			}
			if queryString, exist := matchRuleMap["QueryString"]; exist {
				if queryStringList, ok := queryString.([]interface{}); ok {
					for _, queryStringItem := range queryStringList {
						if queryStringMap, ok := queryStringItem.(map[string]interface{}); ok {
							if value, exist := queryStringMap["Value"]; exist {
								if valueMap, ok := value.(map[string]interface{}); ok {
									queryStringMap["Value"] = []interface{}{valueMap}
								}
							}
						}
					}
				}
			}
			if header, exist := matchRuleMap["Header"]; exist {
				if headerList, ok := header.([]interface{}); ok {
					for _, headerItem := range headerList {
						if headerMap, ok := headerItem.(map[string]interface{}); ok {
							if value, exist := headerMap["Value"]; exist {
								if valueMap, ok := value.(map[string]interface{}); ok {
									headerMap["Value"] = []interface{}{valueMap}
								}
							}
						}
					}
				}
			}
		}
	}
	if advancedSetting, exist := data["AdvancedSetting"]; exist {
		if advancedSettingMap, ok := advancedSetting.(map[string]interface{}); ok {
			if timeoutSetting, exist := advancedSettingMap["TimeoutSetting"]; exist {
				if timeoutSettingMap, ok := timeoutSetting.(map[string]interface{}); ok {
					if len(timeoutSettingMap) == 0 {
						timeoutSettingMap["Enable"] = false
					}
					advancedSettingMap["TimeoutSetting"] = []interface{}{timeoutSettingMap}
				}
			}
			if corsPolicySetting, exist := advancedSettingMap["CorsPolicySetting"]; exist {
				if corsPolicySettingMap, ok := corsPolicySetting.(map[string]interface{}); ok {
					if len(corsPolicySettingMap) == 0 {
						corsPolicySettingMap["Enable"] = false
					}
					advancedSettingMap["CorsPolicySetting"] = []interface{}{corsPolicySettingMap}
				}
			}
			if urlRewriteSetting, exist := advancedSettingMap["URLRewriteSetting"]; exist {
				if urlRewriteSettingMap, ok := urlRewriteSetting.(map[string]interface{}); ok {
					if len(urlRewriteSettingMap) == 0 {
						urlRewriteSettingMap["Enable"] = false
					}
					advancedSettingMap["URLRewriteSetting"] = []interface{}{urlRewriteSettingMap}
				}
			}
			if retryPolicySetting, exist := advancedSettingMap["RetryPolicySetting"]; exist {
				if retryPolicySettingMap, ok := retryPolicySetting.(map[string]interface{}); ok {
					if len(retryPolicySettingMap) == 0 {
						retryPolicySettingMap["Enable"] = false
					}
					advancedSettingMap["RetryPolicySetting"] = []interface{}{retryPolicySettingMap}
				}
			}
			if mirrorPolicies, exist := advancedSettingMap["MirrorPolicies"]; exist {
				if mirrorPoliciesList, ok := mirrorPolicies.([]interface{}); ok {
					for _, mirrorPolicy := range mirrorPoliciesList {
						if mirrorPolicyMap, ok := mirrorPolicy.(map[string]interface{}); ok {
							if upstream, exist := mirrorPolicyMap["Upstream"]; exist {
								if upstreamMap, ok := upstream.(map[string]interface{}); ok {
									mirrorPolicyMap["Upstream"] = []interface{}{upstreamMap}
								}
							}
							if percent, exist := mirrorPolicyMap["Percent"]; exist {
								if percentMap, ok := percent.(map[string]interface{}); ok {
									mirrorPolicyMap["Percent"] = []interface{}{percentMap}
								}
							}
						}
					}
				}
			}
		}
	}

	return data, err
}

func (s *VolcengineApigRouteService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
			status, err = ve.ObtainSdkValue("Status", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("apig_route status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (VolcengineApigRouteService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"AIProviderSettings": {
				TargetField: "ai_provider_settings",
			},
			"URLRewriteSetting": {
				TargetField: "url_rewrite_setting",
			},
			"URLRewrite": {
				TargetField: "url_rewrite",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineApigRouteService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateRoute",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"enable": {
					TargetField: "Enable",
					ForceGet:    true,
				},
				"upstream_list": {
					TargetField: "UpstreamList",
					ConvertType: ve.ConvertJsonObjectArray,
					NextLevelConvert: map[string]ve.RequestConvert{
						"ai_provider_settings": {
							TargetField: "AIProviderSettings",
							ConvertType: ve.ConvertJsonObject,
						},
					},
				},
				"match_rule": {
					TargetField: "MatchRule",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"path": {
							TargetField: "Path",
							ConvertType: ve.ConvertJsonObject,
						},
						"method": {
							TargetField: "Method",
							ConvertType: ve.ConvertJsonArray,
						},
						"query_string": {
							TargetField: "QueryString",
							ConvertType: ve.ConvertJsonObjectArray,
							NextLevelConvert: map[string]ve.RequestConvert{
								"value": {
									TargetField: "Value",
									ConvertType: ve.ConvertJsonObject,
								},
							},
						},
						"header": {
							TargetField: "Header",
							ConvertType: ve.ConvertJsonObjectArray,
							NextLevelConvert: map[string]ve.RequestConvert{
								"value": {
									TargetField: "Value",
									ConvertType: ve.ConvertJsonObject,
								},
							},
						},
					},
				},
				"advanced_setting": {
					TargetField: "AdvancedSetting",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"timeout_setting": {
							TargetField: "TimeoutSetting",
							ConvertType: ve.ConvertJsonObject,
						},
						"cors_policy_setting": {
							TargetField: "CorsPolicySetting",
							ConvertType: ve.ConvertJsonObject,
						},
						"url_rewrite_setting": {
							TargetField: "URLRewriteSetting",
							ConvertType: ve.ConvertJsonObject,
						},
						"retry_policy_setting": {
							TargetField: "RetryPolicySetting",
							ConvertType: ve.ConvertJsonObject,
							NextLevelConvert: map[string]ve.RequestConvert{
								"retry_on": {
									TargetField: "RetryOn",
									ConvertType: ve.ConvertJsonArray,
								},
								"http_codes": {
									TargetField: "HttpCodes",
									ConvertType: ve.ConvertJsonArray,
								},
							},
						},
						"header_operations": {
							TargetField: "HeaderOperations",
							ConvertType: ve.ConvertJsonObjectArray,
						},
						"mirror_policies": {
							TargetField: "MirrorPolicies",
							ConvertType: ve.ConvertJsonObjectArray,
							NextLevelConvert: map[string]ve.RequestConvert{
								"upstream": {
									TargetField: "Upstream",
									ConvertType: ve.ConvertJsonObject,
								},
								"percent": {
									TargetField: "Percent",
									ConvertType: ve.ConvertJsonObject,
								},
							},
						},
					},
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.Id", *resp)
				d.SetId(id.(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineApigRouteService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateRoute",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"name": {
					TargetField: "Name",
					ForceGet:    true,
				},
				"enable": {
					TargetField: "Enable",
					ForceGet:    true,
				},
				"priority": {
					TargetField: "Priority",
					ForceGet:    true,
				},
				"upstream_list": {
					TargetField: "UpstreamList",
					ConvertType: ve.ConvertJsonObjectArray,
					ForceGet:    true,
					NextLevelConvert: map[string]ve.RequestConvert{
						"ai_provider_settings": {
							TargetField: "AIProviderSettings",
							ConvertType: ve.ConvertJsonObject,
						},
					},
				},
				"match_rule": {
					TargetField: "MatchRule",
					ConvertType: ve.ConvertJsonObject,
					ForceGet:    true,
					NextLevelConvert: map[string]ve.RequestConvert{
						"path": {
							TargetField: "Path",
							ConvertType: ve.ConvertJsonObject,
						},
						"method": {
							TargetField: "Method",
							ConvertType: ve.ConvertJsonArray,
						},
						"query_string": {
							TargetField: "QueryString",
							ConvertType: ve.ConvertJsonObjectArray,
							NextLevelConvert: map[string]ve.RequestConvert{
								"value": {
									TargetField: "Value",
									ConvertType: ve.ConvertJsonObject,
								},
							},
						},
						"header": {
							TargetField: "Header",
							ConvertType: ve.ConvertJsonObjectArray,
							NextLevelConvert: map[string]ve.RequestConvert{
								"value": {
									TargetField: "Value",
									ConvertType: ve.ConvertJsonObject,
								},
							},
						},
					},
				},
				"advanced_setting": {
					TargetField: "AdvancedSetting",
					ConvertType: ve.ConvertJsonObject,
					ForceGet:    true,
					NextLevelConvert: map[string]ve.RequestConvert{
						"timeout_setting": {
							TargetField: "TimeoutSetting",
							ConvertType: ve.ConvertJsonObject,
						},
						"cors_policy_setting": {
							TargetField: "CorsPolicySetting",
							ConvertType: ve.ConvertJsonObject,
						},
						"url_rewrite_setting": {
							TargetField: "URLRewriteSetting",
							ConvertType: ve.ConvertJsonObject,
						},
						"retry_policy_setting": {
							TargetField: "RetryPolicySetting",
							ConvertType: ve.ConvertJsonObject,
							NextLevelConvert: map[string]ve.RequestConvert{
								"retry_on": {
									TargetField: "RetryOn",
									ConvertType: ve.ConvertJsonArray,
								},
								"http_codes": {
									TargetField: "HttpCodes",
									ConvertType: ve.ConvertJsonArray,
								},
							},
						},
						"header_operations": {
							TargetField: "HeaderOperations",
							ConvertType: ve.ConvertJsonObjectArray,
						},
						"mirror_policies": {
							TargetField: "MirrorPolicies",
							ConvertType: ve.ConvertJsonObjectArray,
							NextLevelConvert: map[string]ve.RequestConvert{
								"upstream": {
									TargetField: "Upstream",
									ConvertType: ve.ConvertJsonObject,
								},
								"percent": {
									TargetField: "Percent",
									ConvertType: ve.ConvertJsonObject,
								},
							},
						},
					},
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["Id"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineApigRouteService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteRoute",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"Id": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(5*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading apig route on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineApigRouteService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"name": {
				TargetField: "Filter.Name",
			},
			"path": {
				TargetField: "Filter.Path",
			},
		},
		NameField:    "Name",
		IdField:      "Id",
		CollectField: "routes",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"AIProviderSettings": {
				TargetField: "ai_provider_settings",
			},
			"URLRewriteSetting": {
				TargetField: "url_rewrite_setting",
			},
			"URLRewrite": {
				TargetField: "url_rewrite",
			},
		},
	}
}

func (s *VolcengineApigRouteService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "apig",
		Version:     "2022-11-12",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
