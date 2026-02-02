package host_group

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

type Service struct {
	Client *ve.SdkClient
}

func NewService(c *ve.SdkClient) *Service {
	return &Service{
		Client: c,
	}
}

func (s *Service) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *Service) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		req := map[string]interface{}{}
		if v, ok := condition["host_group_id"]; ok {
			req["HostGroupId"] = v
		}
		if v, ok := condition["host_group_name"]; ok {
			req["HostGroupName"] = v
		}
		if v, ok := condition["host_identifier"]; ok {
			req["HostIdentifier"] = v
		}
		if v, ok := condition["iam_project_name"]; ok {
			req["IamProjectName"] = v
		}
		if v, ok := condition["auto_update"]; ok {
			req["AutoUpdate"] = v
		}
		if v, ok := condition["service_logging"]; ok {
			req["ServiceLogging"] = v
		}
		if v, ok := condition["hidden"]; ok {
			req["Hidden"] = v
		}
		if v, ok := condition["PageSize"]; ok {
			req["PageSize"] = v
		}
		if v, ok := condition["PageNumber"]; ok {
			req["PageNumber"] = v
		}

		// 默认行为：调用 DescribeHostGroupsV2
		action := "DescribeHostGroupsV2"
		logger.Debug(logger.ReqFormat, action, req)
		resp, err = s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			ContentType: ve.Default,
			HttpMethod:  ve.GET,
			Path:        []string{action},
			Client:      s.Client.BypassSvcClient.NewTlsClient(),
		}, &req)
		if err != nil {
			return data, err
		}
		logger.Debug(logger.RespFormat, action, req, *resp)
		results, err = ve.ObtainSdkValue("RESPONSE.HostGroupHostsRulesInfos", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.HostGroupHostsRulesInfo is not Slice")
		}
		var res []interface{}
		for _, ele := range data {
			hostGroupInfo := ele.(map[string]interface{})["HostGroupInfo"].(map[string]interface{})
			newItem := map[string]interface{}{}

			newItem["host_group_info"] = []interface{}{
				map[string]interface{}{
					"host_group_id":     hostGroupInfo["HostGroupId"],
					"host_group_name":   hostGroupInfo["HostGroupName"],
					"host_group_type":   hostGroupInfo["HostGroupType"],
					"host_identifier":   hostGroupInfo["HostIdentifier"],
					"host_count":        hostGroupInfo["HostCount"],
					"rule_count":        hostGroupInfo["RuleCount"],
					"create_time":       hostGroupInfo["CreateTime"],
					"modify_time":       hostGroupInfo["ModifyTime"],
					"iam_project_name":  hostGroupInfo["IamProjectName"],
					"update_start_time": hostGroupInfo["UpdateStartTime"],
					"update_end_time":   hostGroupInfo["UpdateEndTime"],
					"auto_update":       hostGroupInfo["AutoUpdate"],
					"service_logging":   hostGroupInfo["ServiceLogging"],
				},
			}
			res = append(res, newItem)
		}
		return res, nil
	})
}

func (s *Service) DatasourceRules(d *schema.ResourceData, r *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		IdField:      "RuleId",
		CollectField: "rule_infos",
	}
}

func (s *Service) ReadRules(m map[string]interface{}) (data []interface{}, err error) {
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		req := map[string]interface{}{
			"PageNumber": condition["PageNumber"],
			"PageSize":   condition["PageSize"],
		}
		if v, ok := condition["HostGroupId"]; ok {
			req["HostGroupId"] = v
		} else if v, ok := condition["host_group_id"]; ok {
			req["HostGroupId"] = v
		}

		action := "DescribeHostGroupRules"
		logger.Debug(logger.ReqFormat, action, req)
		resp, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			ContentType: ve.Default,
			HttpMethod:  ve.GET,
			Path:        []string{action},
			Client:      s.Client.BypassSvcClient.NewTlsClient(),
		}, &req)
		if err != nil {
			return nil, err
		}
		logger.Debug(logger.RespFormat, action, req, *resp)
		results, err := ve.ObtainSdkValue("RESPONSE.RuleInfos", *resp)
		if err != nil {
			return nil, err
		}
		if results == nil {
			return []interface{}{}, nil
		}
		rawList, ok := results.([]interface{})
		if !ok {
			return nil, errors.New("Result.RuleInfos is not Slice")
		}
		var res []interface{}
		for _, item := range rawList {
			if itemMap, ok := item.(map[string]interface{}); ok {
				res = append(res, transRuleMap(itemMap))
			}
		}
		return res, nil
	})
}

// transRuleMap 处理 RuleInfo 中的复杂结构，确保其符合 Terraform Schema 要求
func transRuleMap(data map[string]interface{}) map[string]interface{} {
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

func (s *Service) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	// Use DescribeHostGroupV2
	action := "DescribeHostGroupV2"
	req := map[string]interface{}{
		"HostGroupId": id,
	}
	logger.Debug(logger.ReqFormat, action, req)
	resp, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
		ContentType: ve.Default,
		HttpMethod:  ve.GET,
		Path:        []string{action},
		Client:      s.Client.BypassSvcClient.NewTlsClient(),
	}, &req)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, req, *resp)

	result, err := ve.ObtainSdkValue("RESPONSE.HostGroupHostsRulesInfo.HostGroupInfo", *resp)
	if err != nil {
		// Fallback to try without HostGroupHostsRulesInfo wrapper if needed,
		// but assuming standard structure similar to list item based on ReadResources.
		// If DescribeHostGroupV2 returns directly HostGroupInfo under RESPONSE or Result, adjust path.
		// Trying "RESPONSE.HostGroupInfo" as a likely alternative if the above fails or returns nil.
		result, err = ve.ObtainSdkValue("RESPONSE.HostGroupInfo", *resp)
		if err != nil {
			return data, err
		}
	}
	if result == nil {
		return data, fmt.Errorf("Host Group %s not exist ", id)
	}

	hostGroupInfo, ok := result.(map[string]interface{})
	if !ok {
		return data, errors.New("RESPONSE.HostGroupInfo is not map")
	}

	// Manual mapping from API PascalCase to Terraform snake_case
	data = map[string]interface{}{
		"host_group_id":     hostGroupInfo["HostGroupId"],
		"host_group_name":   hostGroupInfo["HostGroupName"],
		"host_group_type":   hostGroupInfo["HostGroupType"],
		"host_identifier":   hostGroupInfo["HostIdentifier"],
		"host_count":        hostGroupInfo["HostCount"],
		"rule_count":        hostGroupInfo["RuleCount"],
		"create_time":       hostGroupInfo["CreateTime"],
		"modify_time":       hostGroupInfo["ModifyTime"],
		"iam_project_name":  hostGroupInfo["IamProjectName"],
		"update_start_time": hostGroupInfo["UpdateStartTime"],
		"update_end_time":   hostGroupInfo["UpdateEndTime"],
		"auto_update":       hostGroupInfo["AutoUpdate"],
		"service_logging":   hostGroupInfo["ServiceLogging"],
	}

	return data, nil
}

func (s *Service) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (Service) WithResourceResponseHandlers(data map[string]interface{}) []ve.ResourceResponseHandler {
	return []ve.ResourceResponseHandler{}

}

func (s *Service) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateHostGroup",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"host_group_name": {
					TargetField: "HostGroupName",
				},
				"host_group_type": {
					TargetField: "HostGroupType",
				},
				"host_identifier": {
					TargetField: "HostIdentifier",
				},
				"host_ip_list": {
					TargetField: "HostIpList",
					ConvertType: ve.ConvertJsonArray,
				},
				"auto_update": {
					TargetField: "AutoUpdate",
				},
				"update_start_time": {
					TargetField: "UpdateStartTime",
				},
				"update_end_time": {
					TargetField: "UpdateEndTime",
				},
				"service_logging": {
					TargetField: "ServiceLogging",
				},
				"iam_project_name": {
					TargetField: "IamProjectName",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.POST,
					Path:        []string{call.Action},
					Client:      s.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
				if err != nil {
					return nil, err
				}
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam, *resp)
				return resp, nil
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("RESPONSE.HostGroupId", *resp)
				d.SetId(id.(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}

}

func (s *Service) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyHostGroup",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"host_group_name": {
					TargetField: "HostGroupName",
				},
				"auto_update": {
					TargetField: "AutoUpdate",
				},
				"update_start_time": {
					TargetField: "UpdateStartTime",
				},
				"update_end_time": {
					TargetField: "UpdateEndTime",
				},
				"service_logging": {
					TargetField: "ServiceLogging",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["HostGroupId"] = d.Id()

				if d.HasChanges("host_ip_list", "host_identifier", "host_group_type") {
					(*call.SdkParam)["HostGroupType"] = d.Get("host_group_type")
					if d.Get("host_group_type").(string) == "IP" {
						(*call.SdkParam)["HostIpList"] = d.Get("host_ip_list").(*schema.Set).List()
					} else {
						(*call.SdkParam)["HostIdentifier"] = d.Get("host_identifier")
					}
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.PUT,
					Path:        []string{call.Action},
					Client:      s.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
				if err != nil {
					return nil, err
				}
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam, *resp)
				return resp, nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *Service) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteHostGroup",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"HostGroupId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.DELETE,
					Path:        []string{call.Action},
					Client:      s.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
				if err != nil {
					return nil, err
				}
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam, *resp)
				return resp, nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *Service) DatasourceResources(d *schema.ResourceData, r *schema.Resource) ve.DataSourceInfo {
	if _, ok := r.Schema["describe_host_group_rules"]; ok {
		return ve.DataSourceInfo{
			IdField:      "host_group_id",
			CollectField: "rule_infos",
		}
	}
	return ve.DataSourceInfo{
		NameField:    "HostGroupName",
		IdField:      "HostGroupId",
		CollectField: "infos",
	}
}

func (s *Service) ReadResourceId(id string) string {
	return id
}

func (s *Service) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "tls",
		ResourceType:         "hostgroup",
		ProjectSchemaField:   "iam_project_name",
		ProjectResponseField: "IamProjectName",
	}
}
