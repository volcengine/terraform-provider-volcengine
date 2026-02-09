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

func (v *VolcengineTlsRuleService) ReadResources(m map[string]interface{}) ([]interface{}, error) {
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		var (
			resp        *map[string]interface{}
			results     interface{}
			err         error
			transResult []interface{}
		)
		action := "DescribeRules"

		logger.Debug(logger.ReqFormat, action, condition)
		resp, err = v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			ContentType: ve.Default,
			HttpMethod:  ve.GET,
			Path:        []string{action},
			Client:      v.Client.BypassSvcClient.NewTlsClient(),
		}, &condition)
		if err != nil {
			return nil, err
		}
		logger.Debug(logger.RespFormat, action, resp)

		// 第一步：获取原始结果，判断是否为空
		results, err = ve.ObtainSdkValue("RESPONSE.RuleInfos", *resp)
		if err != nil {
			return nil, err
		}
		if results == nil {
			return []interface{}{}, nil
		}

		// 第二步：校验结果类型并转换
		data, ok := results.([]interface{})
		if !ok {
			return nil, errors.New("Result.RuleInfos is not Slice")
		}
		if len(data) == 0 {
			return []interface{}{}, nil
		}

		// 第三步：数据转换
		for _, d := range data {
			dataMap, ok := d.(map[string]interface{})
			if !ok {
				return nil, errors.New("value is not map")
			}

			transResult = append(transResult, dataMapTransToList(dataMap))
		}

		return transResult, nil
	})
}

func (v *VolcengineTlsRuleService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	if id == "" {
		id = v.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"RuleId": id,
	}
	// Use DescribeRules with RuleId filter for consistency
	action := "DescribeRules"
	logger.Debug(logger.ReqFormat, action, req)
	resp, err := v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
		ContentType: ve.Default,
		HttpMethod:  ve.GET,
		Path:        []string{action},
		Client:      v.Client.BypassSvcClient.NewTlsClient(),
	}, &req)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, resp)

	results, err := ve.ObtainSdkValue("RESPONSE.RuleInfos", *resp)
	if err != nil {
		return data, err
	}
	resultsList, ok := results.([]interface{})
	if !ok || len(resultsList) == 0 {
		return data, fmt.Errorf("tls rule %s not exist", id)
	}

	data, ok = resultsList[0].(map[string]interface{})
	if !ok {
		return data, fmt.Errorf("read resource value is not map")
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

// dataMapTransToList 处理 RuleInfo 中的复杂结构，确保其符合 Terraform Schema 要求
func dataMapTransToList(data map[string]interface{}) map[string]interface{} {
	if data == nil {
		return data
	}

	// 1. 基本字段转换 (API 大写 -> Schema 小写)
	res := map[string]interface{}{
		"topic_id":    data["TopicId"],
		"rule_id":     data["RuleId"],
		"rule_name":   data["RuleName"],
		"log_type":    data["LogType"],
		"log_sample":  data["LogSample"],
		"input_type":  data["InputType"],
		"create_time": data["CreateTime"],
		"modify_time": data["ModifyTime"],
	}

	// 2. Paths (TypeSet)
	if v, ok := data["Paths"]; ok && v != nil {
		res["paths"] = v
	} else {
		res["paths"] = []interface{}{}
	}

	// 3. ExcludePaths (TypeSet of Resource)
	if v, ok := data["ExcludePaths"]; ok && v != nil {
		if list, ok := v.([]interface{}); ok {
			var eps []interface{}
			for _, item := range list {
				if m, ok := item.(map[string]interface{}); ok {
					eps = append(eps, map[string]interface{}{
						"type":  m["Type"],
						"value": m["Value"],
					})
				}
			}
			res["exclude_paths"] = eps
		}
	} else {
		res["exclude_paths"] = []interface{}{}
	}

	// 4. ExtractRule (TypeList MaxItems:1)
	if v, ok := data["ExtractRule"]; ok && v != nil {
		if m, ok := v.(map[string]interface{}); ok {
			er := map[string]interface{}{
				"delimiter":               m["Delimiter"],
				"begin_regex":             m["BeginRegex"],
				"log_regex":               m["LogRegex"],
				"time_key":                m["TimeKey"],
				"time_format":             m["TimeFormat"],
				"un_match_up_load_switch": m["UnMatchUpLoadSwitch"],
				"un_match_log_key":        m["UnMatchLogKey"],
				"quote":                   m["Quote"],
				"time_zone":               m["TimeZone"],
			}
			if keys, ok := m["Keys"]; ok && keys != nil {
				er["keys"] = keys
			} else {
				er["keys"] = []interface{}{}
			}
			if fkr, ok := m["FilterKeyRegex"]; ok && fkr != nil {
				if list, ok := fkr.([]interface{}); ok {
					var fkrs []interface{}
					for _, item := range list {
						if itemMap, ok := item.(map[string]interface{}); ok {
							fkrs = append(fkrs, map[string]interface{}{
								"key":   itemMap["Key"],
								"regex": itemMap["Regex"],
							})
						}
					}
					er["filter_key_regex"] = fkrs
				}
			} else {
				er["filter_key_regex"] = []interface{}{}
			}
			if lt, ok := m["LogTemplate"]; ok && lt != nil {
				if ltMap, ok := lt.(map[string]interface{}); ok {
					er["log_template"] = []interface{}{
						map[string]interface{}{
							"type":   ltMap["Type"],
							"format": ltMap["Value"],
						},
					}
				}
			}
			res["extract_rule"] = []interface{}{er}
		}
	}

	// 5. UserDefineRule (TypeList MaxItems:1)
	if v, ok := data["UserDefineRule"]; ok && v != nil {
		if m, ok := v.(map[string]interface{}); ok {
			udr := map[string]interface{}{
				"enable_raw_log": m["EnableRawLog"],
				"tail_files":     m["TailFiles"],
			}
			if fields, ok := m["Fields"]; ok && fields != nil {
				udr["fields"] = fields
			} else {
				udr["fields"] = map[string]interface{}{}
			}
			if ppr, ok := m["ParsePathRule"]; ok && ppr != nil {
				if pprMap, ok := ppr.(map[string]interface{}); ok {
					pprRes := map[string]interface{}{
						"path_sample": pprMap["PathSample"],
						"regex":       pprMap["Regex"],
					}
					if keys, ok := pprMap["Keys"]; ok && keys != nil {
						pprRes["keys"] = keys
					} else {
						pprRes["keys"] = []interface{}{}
					}
					udr["parse_path_rule"] = []interface{}{pprRes}
				}
			}
			if shk, ok := m["ShardHashKey"]; ok && shk != nil {
				if shkMap, ok := shk.(map[string]interface{}); ok {
					udr["shard_hash_key"] = []interface{}{
						map[string]interface{}{
							"hash_key": shkMap["HashKey"],
						},
					}
				}
			}
			if plugin, ok := m["Plugin"]; ok && plugin != nil {
				if pluginMap, ok := plugin.(map[string]interface{}); ok {
					pRes := map[string]interface{}{}
					processors, ok := pluginMap["processors"]
					if !ok {
						processors = pluginMap["Processors"]
					}
					if procList, ok := processors.([]interface{}); ok {
						arr := make([]interface{}, 0)
						for _, processor := range procList {
							jsonBytes, _ := json.Marshal(processor)
							arr = append(arr, string(jsonBytes))
						}
						pRes["processors"] = arr
					} else {
						pRes["processors"] = []interface{}{}
					}
					udr["plugin"] = []interface{}{pRes}
				}
			}
			if adv, ok := m["Advanced"]; ok && adv != nil {
				if advMap, ok := adv.(map[string]interface{}); ok {
					udr["advanced"] = []interface{}{
						map[string]interface{}{
							"close_inactive": advMap["CloseInactive"],
							"close_removed":  advMap["CloseRemoved"],
							"close_renamed":  advMap["CloseRenamed"],
							"close_eof":      advMap["CloseEOF"],
							"close_timeout":  advMap["CloseTimeout"],
						},
					}
				}
			}
			res["user_define_rule"] = []interface{}{udr}
		}
	}

	// 6. ContainerRule (TypeList MaxItems:1)
	if v, ok := data["ContainerRule"]; ok && v != nil {
		if m, ok := v.(map[string]interface{}); ok {
			cr := map[string]interface{}{
				"stream":               m["Stream"],
				"container_name_regex": m["ContainerNameRegex"],
			}
			maps := []string{
				"IncludeContainerLabelRegex", "ExcludeContainerLabelRegex",
				"IncludeContainerEnvRegex", "ExcludeContainerEnvRegex", "EnvTag",
			}
			for _, key := range maps {
				target := ve.HumpToDownLine(key)
				if val, ok := m[key]; ok && val != nil {
					cr[target] = val
				} else {
					cr[target] = map[string]interface{}{}
				}
			}
			if kr, ok := m["KubernetesRule"]; ok && kr != nil {
				if krMap, ok := kr.(map[string]interface{}); ok {
					krRes := map[string]interface{}{
						"namespace_name_regex": krMap["NamespaceNameRegex"],
						"workload_type":        krMap["WorkloadType"],
						"workload_name_regex":  krMap["WorkloadNameRegex"],
						"pod_name_regex":       krMap["PodNameRegex"],
					}
					krMaps := []string{"IncludePodLabelRegex", "ExcludePodLabelRegex", "LabelTag", "AnnotationTag"}
					for _, key := range krMaps {
						target := ve.HumpToDownLine(key)
						if val, ok := krMap[key]; ok && val != nil {
							krRes[target] = val
						} else {
							krRes[target] = map[string]interface{}{}
						}
					}
					cr["kubernetes_rule"] = []interface{}{krRes}
				}
			}
			res["container_rule"] = []interface{}{cr}
		}
	}

	return res
}
