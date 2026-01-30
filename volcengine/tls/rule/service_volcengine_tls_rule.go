package rule

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

type VolcengineTlsRuleService struct {
	Client *ve.SdkClient
}

func (v *VolcengineTlsRuleService) GetClient() *ve.SdkClient {
	return v.Client
}

func (v *VolcengineTlsRuleService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp        *map[string]interface{}
		results     interface{}
		transResult []interface{}
		ok          bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeRules"
		req := map[string]interface{}{}
		if v, ok := condition["project_id"]; ok {
			req["ProjectId"] = v
		}
		if v, ok := condition["project_name"]; ok {
			req["ProjectName"] = v
		}
		if v, ok := condition["iam_project_name"]; ok {
			req["IamProjectName"] = v
		}
		if v, ok := condition["rule_id"]; ok {
			req["RuleId"] = v
		}
		if v, ok := condition["rule_name"]; ok {
			req["RuleName"] = v
		}
		if v, ok := condition["topic_id"]; ok {
			req["TopicId"] = v
		}
		if v, ok := condition["topic_name"]; ok {
			req["TopicName"] = v
		}
		if v, ok := condition["log_type"]; ok {
			req["LogType"] = v
		}
		if v, ok := condition["pause"]; ok {
			req["Pause"] = v
		}
		if v, ok := condition["PageNumber"]; ok {
			req["PageNumber"] = v
		}
		if v, ok := condition["PageSize"]; ok {
			req["PageSize"] = v
		}

		logger.Debug(logger.ReqFormat, action, req)
		resp, err = v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			ContentType: ve.Default,
			HttpMethod:  ve.GET,
			Path:        []string{action},
			Client:      v.Client.BypassSvcClient.NewTlsClient(),
		}, &req)
		logger.Debug(logger.RespFormat, action, resp)
		results, err = ve.ObtainSdkValue("RESPONSE.RuleInfos", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.RuleInfos is not Slice")
		}
		for _, d := range data {
			dataMap, ok := d.(map[string]interface{})
			if !ok {
				return data, errors.New("value is not map")
			}
			userRuleMap, ok := dataMap["UserDefineRule"].(map[string]interface{})
			if !ok {
				return data, errors.New("value is not map")
			}
			plugin, ok := userRuleMap["Plugin"].(map[string]interface{})
			if !ok {
				if plugin == nil {
					continue
				}
				return data, errors.New("value is not map")
			}
			if len(plugin) == 0 {
				continue
			}
			// 接口中 processors 为小写
			processors, ok := plugin["processors"]
			if !ok {
				continue
			}
			logger.DebugInfo("plugin ori : ", processors)
			arr := make([]string, 0)
			if _, ok = processors.([]interface{}); ok {
				for _, processor := range processors.([]interface{}) {
					p, _ := json.Marshal(processor)
					pStr := string(p)
					arr = append(arr, pStr)
				}
			}
			plugin["processors"] = arr
			userRuleMap["Plugin"] = plugin
			dataMap["UserDefineRule"] = userRuleMap
			transResult = append(transResult, dataMap)
		}
		return data, err
	})
}

func (v *VolcengineTlsRuleService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		resp *map[string]interface{}
		ok   bool
	)
	if id == "" {
		id = v.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"RuleId": id,
	}
	// First try to use DescribeRuleV2
	action := "DescribeRuleV2"
	logger.Debug(logger.ReqFormat, action, req)
	resp, err = v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
		ContentType: ve.Default,
		HttpMethod:  ve.GET,
		Path:        []string{action},
		Client:      v.Client.BypassSvcClient.NewTlsClient(),
	}, &req)
	if err != nil {
		return data, err
	}
	// If DescribeRuleV2 succeeds, use its results
	logger.Debug(logger.RespFormat, action, resp)
	resultsInterface, err := ve.ObtainSdkValue("RESPONSE", *resp)
	if err != nil {
		return data, err
	}
	if data, ok = resultsInterface.(map[string]interface{}); !ok {
		return data, fmt.Errorf("read resource value is not map")
	}
	if len(data) == 0 {
		return data, fmt.Errorf("tls rule %s not exist", id)
	}
	return dataMapTransToList(data), nil
}

func (v *VolcengineTlsRuleService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, s string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (v *VolcengineTlsRuleService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return m, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (v *VolcengineTlsRuleService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateRule",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"paths": {
					ConvertType: ve.ConvertJsonArray,
				},
				"exclude_paths": {
					ConvertType: ve.ConvertJsonObjectArray,
					NextLevelConvert: map[string]ve.RequestConvert{
						"type": {
							TargetField: "Type",
						},
						"value": {
							TargetField: "Value",
						},
					},
				},
				"extract_rule": {
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"quote": {
							ConvertType: ve.ConvertDefault,
						},
						"time_zone": {
							ConvertType: ve.ConvertDefault,
						},
						"keys": {
							ConvertType: ve.ConvertJsonArray,
						},
						"filter_key_regex": {
							ConvertType: ve.ConvertJsonObjectArray,
							NextLevelConvert: map[string]ve.RequestConvert{
								"key": {
									TargetField: "Key",
								},
								"regex": {
									TargetField: "Regex",
								},
							},
						},
						"log_template": {
							ConvertType: ve.ConvertJsonObject,
						},
					},
				},
				"user_define_rule": {
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"fields": {
							ConvertType: ve.ConvertJsonObject,
						},
						"parse_path_rule": {
							ConvertType: ve.ConvertJsonObject,
							NextLevelConvert: map[string]ve.RequestConvert{
								"keys": {
									ConvertType: ve.ConvertJsonArray,
								},
							},
						},
						"shard_hash_key": {
							ConvertType: ve.ConvertJsonObject,
						},
						"plugin": {
							ConvertType: ve.ConvertJsonObject,
						},
						"advanced": {
							ConvertType: ve.ConvertJsonObject,
							NextLevelConvert: map[string]ve.RequestConvert{
								"close_eof": {
									TargetField: "CloseEOF",
								},
							},
						},
					},
				},
				"container_rule": {
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"include_container_label_regex": {
							ConvertType: ve.ConvertJsonObject,
						},
						"exclude_container_label_regex": {
							ConvertType: ve.ConvertJsonObject,
						},
						"include_container_env_regex": {
							ConvertType: ve.ConvertJsonObject,
						},
						"exclude_container_env_regex": {
							ConvertType: ve.ConvertJsonObject,
						},
						"env_tag": {
							ConvertType: ve.ConvertJsonObject,
						},
						"kubernetes_rule": {
							ConvertType: ve.ConvertJsonObject,
							NextLevelConvert: map[string]ve.RequestConvert{
								"include_pod_label_regex": {
									ConvertType: ve.ConvertJsonObject,
								},
								"exclude_pod_label_regex": {
									ConvertType: ve.ConvertJsonObject,
								},
								"label_tag": {
									ConvertType: ve.ConvertJsonObject,
								},
								"annotation_tag": {
									ConvertType: ve.ConvertJsonObject,
								},
							},
						},
					},
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				logger.DebugInfo("Sdk ori param: ", call.SdkParam)
				processors, ok := d.GetOk("user_define_rule.0.plugin.0.processors")
				if ok {
					arr := make([]interface{}, 0)
					if _, ok = processors.(*schema.Set); ok {
						for _, p := range processors.(*schema.Set).List() {
							logger.DebugInfo("processors ori: ", p.(string))
							var pJson map[string]interface{}
							err := json.Unmarshal([]byte(p.(string)), &pJson)
							if err != nil {
								return false, errors.New("failed to unmarshal processors")
							}
							arr = append(arr, pJson)
						}
					}
					// 接口中 processors 为小写
					(*call.SdkParam)["UserDefineRule.Plugin.processors"] = arr
				}
				logger.DebugInfo("Sdk unmarshal param: ", call.SdkParam)
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				// 接口中 processors 为小写，需删除自动转换的大写key
				if _, ok := (*call.SdkParam)["UserDefineRule"]; ok {
					if userDefineMap, ok := (*call.SdkParam)["UserDefineRule"].(map[string]interface{}); ok {
						if plugin, ok := userDefineMap["Plugin"]; ok {
							if pluginMap, ok := plugin.(map[string]interface{}); ok {
								if _, ok = pluginMap["Processors"]; ok {
									delete((*call.SdkParam)["UserDefineRule"].(map[string]interface{})["Plugin"].(map[string]interface{}), "Processors")
								}
							}
						}
					}
				}
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.POST,
					Path:        []string{call.Action},
					Client:      v.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("RESPONSE.RuleId", *resp)
				d.SetId(id.(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcengineTlsRuleService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyRule",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"rule_name": {
					ConvertType: ve.ConvertDefault,
				},
				"paths": {
					ConvertType: ve.ConvertJsonArray,
				},
				"log_type": {
					ConvertType: ve.ConvertDefault,
				},
				"exclude_paths": {
					ConvertType: ve.ConvertJsonObjectArray,
					NextLevelConvert: map[string]ve.RequestConvert{
						"type": {
							TargetField: "Type",
						},
						"value": {
							TargetField: "Value",
						},
					},
				},
				"extract_rule": {
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"quote": {
							ConvertType: ve.ConvertDefault,
							ForceGet:    true,
						},
						"delimiter": {
							ConvertType: ve.ConvertDefault,
							ForceGet:    true,
						},
						"begin_regex": {
							ConvertType: ve.ConvertDefault,
							ForceGet:    true,
						},
						"log_regex": {
							ConvertType: ve.ConvertDefault,
							ForceGet:    true,
						},
						"time_key": {
							ConvertType: ve.ConvertDefault,
							ForceGet:    true,
						},
						"time_format": {
							ConvertType: ve.ConvertDefault,
							ForceGet:    true,
						},
						"un_match_up_load_switch": {
							ConvertType: ve.ConvertDefault,
							ForceGet:    true,
						},
						"un_match_log_key": {
							ConvertType: ve.ConvertDefault,
							ForceGet:    true,
						},
						"time_zone": {
							ConvertType: ve.ConvertDefault,
							ForceGet:    true,
						},
						"keys": {
							ConvertType: ve.ConvertJsonArray,
							ForceGet:    true,
						},
						"filter_key_regex": {
							ConvertType: ve.ConvertJsonObjectArray,
							ForceGet:    true,
							NextLevelConvert: map[string]ve.RequestConvert{
								"key": {
									TargetField: "Key",
								},
								"regex": {
									TargetField: "Regex",
								},
							},
						},
						"log_template": {
							ConvertType: ve.ConvertJsonObject,
							ForceGet:    true,
							NextLevelConvert: map[string]ve.RequestConvert{
								"type": {
									ConvertType: ve.ConvertDefault,
									ForceGet:    true,
								},
								"value": {
									ConvertType: ve.ConvertDefault,
									ForceGet:    true,
								},
							},
						},
					},
				},
				"log_sample": {
					ConvertType: ve.ConvertDefault,
				},
				"input_type": {
					ConvertType: ve.ConvertDefault,
				},
				"user_define_rule": {
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"fields": {
							ConvertType: ve.ConvertJsonObject,
							ForceGet:    true,
						},
						"parse_path_rule": {
							ConvertType: ve.ConvertJsonObject,
							ForceGet:    true,
							NextLevelConvert: map[string]ve.RequestConvert{
								"keys": {
									ConvertType: ve.ConvertJsonArray,
								},
							},
						},
						"shard_hash_key": {
							ConvertType: ve.ConvertJsonObject,
							ForceGet:    true,
							NextLevelConvert: map[string]ve.RequestConvert{
								"hash_key": {
									ConvertType: ve.ConvertDefault,
									ForceGet:    true,
								},
							},
						},
						"plugin": {
							ConvertType: ve.ConvertJsonObject,
						},
						"advanced": {
							ConvertType: ve.ConvertJsonObject,
							ForceGet:    true,
							NextLevelConvert: map[string]ve.RequestConvert{
								"close_eof": {
									TargetField: "CloseEOF",
								},
							},
						},
						"tail_files": {
							ConvertType: ve.ConvertDefault,
							ForceGet:    true,
						},
						"enable_raw_log": {
							ConvertType: ve.ConvertDefault,
							ForceGet:    true,
						},
					},
				},
				"container_rule": {
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"stream": {
							ConvertType: ve.ConvertDefault,
							ForceGet:    true,
						},
						"container_name_regex": {
							ConvertType: ve.ConvertDefault,
							ForceGet:    true,
						},
						"include_container_label_regex": {
							ForceGet:    true,
							ConvertType: ve.ConvertJsonObject,
						},
						"exclude_container_label_regex": {
							ForceGet:    true,
							ConvertType: ve.ConvertJsonObject,
						},
						"include_container_env_regex": {
							ForceGet:    true,
							ConvertType: ve.ConvertJsonObject,
						},
						"exclude_container_env_regex": {
							ForceGet:    true,
							ConvertType: ve.ConvertJsonObject,
						},
						"env_tag": {
							ForceGet:    true,
							ConvertType: ve.ConvertJsonObject,
						},
						"kubernetes_rule": {
							ForceGet:    true,
							ConvertType: ve.ConvertJsonObject,
							NextLevelConvert: map[string]ve.RequestConvert{
								"namespace_name_regex": {
									ConvertType: ve.ConvertDefault,
									ForceGet:    true,
								},
								"workload_type": {
									ConvertType: ve.ConvertDefault,
									ForceGet:    true,
								},
								"pod_name_regex": {
									ConvertType: ve.ConvertDefault,
									ForceGet:    true,
								},
								"workload_name_regex": {
									ConvertType: ve.ConvertDefault,
									ForceGet:    true,
								},
								"include_pod_label_regex": {
									ForceGet:    true,
									ConvertType: ve.ConvertJsonObject,
								},
								"exclude_pod_label_regex": {
									ForceGet:    true,
									ConvertType: ve.ConvertJsonObject,
								},
								"label_tag": {
									ForceGet:    true,
									ConvertType: ve.ConvertJsonObject,
								},
								"annotation_tag": {
									ForceGet:    true,
									ConvertType: ve.ConvertJsonObject,
								},
							},
						},
					},
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["RuleId"] = d.Id()
				logger.DebugInfo("Sdk ori param: ", call.SdkParam)
				processors, ok := d.GetOk("user_define_rule.0.plugin.0.processors")
				if ok {
					arr := make([]interface{}, 0)
					if _, ok = processors.(*schema.Set); ok {
						for _, p := range processors.(*schema.Set).List() {
							logger.DebugInfo("processors ori: ", p.(string))
							var pJson map[string]interface{}
							err := json.Unmarshal([]byte(p.(string)), &pJson)
							if err != nil {
								return false, errors.New("failed to unmarshal processors")
							}
							arr = append(arr, pJson)
						}
					}
					// 接口中 processors 为小写
					(*call.SdkParam)["UserDefineRule.Plugin.processors"] = arr
				}
				logger.DebugInfo("Sdk unmarshal param: ", call.SdkParam)
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				if userDefineRule, ok := (*call.SdkParam)["UserDefineRule"]; ok {
					if userDefineMap, ok := userDefineRule.(map[string]interface{}); ok {
						if plugin, ok := userDefineMap["Plugin"]; ok {
							if pluginMap, ok := plugin.(map[string]interface{}); ok {
								if _, ok := pluginMap["Processors"]; ok {
									delete(pluginMap, "Processors")
								}
							}
						}
					}
				}
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.PUT,
					Path:        []string{call.Action},
					Client:      v.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcengineTlsRuleService) RemoveResource(data *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteRule",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"RuleId": data.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.DELETE,
					Path:        []string{call.Action},
					Client:      v.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := v.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading tls rule on delete %q, %w", d.Id(), callErr))
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
				return ve.CheckResourceUtilRemoved(d, v.ReadResource, 3*time.Minute)
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcengineTlsRuleService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		CollectField: "rules",
		IdField:      "RuleId",
		NameField:    "RuleName",
		ResponseConverts: map[string]ve.ResponseConvert{
			"close_eof": {
				TargetField: "Close_EOF",
			},
		},
	}
}

func (v *VolcengineTlsRuleService) ReadResourceId(s string) string {
	return s
}

func NewTlsRuleService(c *ve.SdkClient) *VolcengineTlsRuleService {
	return &VolcengineTlsRuleService{
		Client: c,
	}
}

// parse_path_rule, shard_hash_key, plugin, advanced, kubernetes_rule map 转 list
func dataMapTransToList(data map[string]interface{}) map[string]interface{} {
	userDefineRule, ok := data["UserDefineRule"].(map[string]interface{})
	if ok && userDefineRule != nil {
		parsePathRule, ok := userDefineRule["ParsePathRule"]
		if ok {
			userDefineRule["ParsePathRule"] = []interface{}{parsePathRule}
		}
		shardHashKey, ok := userDefineRule["ShardHashKey"]
		if ok {
			userDefineRule["ShardHashKey"] = []interface{}{shardHashKey}
		}
		plugin, ok := userDefineRule["Plugin"]
		if ok {
			userDefineRule["Plugin"] = []interface{}{plugin}
		}
		advanced, ok := userDefineRule["Advanced"]
		if ok {
			userDefineRule["Advanced"] = []interface{}{advanced}
		}
	}
	data["UserDefineRule"] = userDefineRule
	containerRule, ok := data["ContainerRule"].(map[string]interface{})
	if ok && containerRule != nil {
		kubernetesRule, ok := containerRule["KubernetesRule"]
		if ok {
			containerRule["KubernetesRule"] = []interface{}{kubernetesRule}
		}
	}
	data["ContainerRule"] = []interface{}{containerRule}

	// 处理 ExcludePaths
	if excludePaths, ok := data["ExcludePaths"].([]interface{}); ok {
		// 使用 Set 结构来存储 ExcludePaths，与 Schema 定义保持一致
		newExcludePaths := schema.NewSet(tlsRuleHash("type", "value"), nil)
		for _, item := range excludePaths {
			if m, ok := item.(map[string]interface{}); ok {
				newMap := make(map[string]interface{})
				if v, ok := m["Type"]; ok {
					newMap["type"] = v
				}
				if v, ok := m["Value"]; ok {
					newMap["value"] = v
				}
				newExcludePaths.Add(newMap)
			}
		}
		data["ExcludePaths"] = newExcludePaths
	}
	return data
}
