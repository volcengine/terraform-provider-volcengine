package host_group

import (
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
		IdField:      "host_group_id",
		CollectField: "rule_infos",
	}
}

func (s *Service) ReadRules(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		req := map[string]interface{}{
			"HostGroupId": condition["host_group_id"],
			"PageNumber":  condition["PageNumber"],
			"PageSize":    condition["PageSize"],
		}
		action := "DescribeHostGroupRules"
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
		var res []interface{}
		for _, rule := range data {
			r := rule.(map[string]interface{})
			ruleMap := map[string]interface{}{
				"rule_id":     r["RuleId"],
				"rule_name":   r["RuleName"],
				"paths":       r["Paths"],
				"topic_id":    r["TopicId"],
				"topic_name":  r["TopicName"],
				"log_type":    r["LogType"],
				"input_type":  r["InputType"],
				"create_time": r["CreateTime"],
				"modify_time": r["ModifyTime"],
				"log_sample":  r["LogSample"],
			}

			if v, ok := r["ExtractRule"].(map[string]interface{}); ok {
				extractRule := map[string]interface{}{
					"delimiter":               v["Delimiter"],
					"begin_regex":             v["BeginRegex"],
					"log_regex":               v["LogRegex"],
					"keys":                    v["Keys"],
					"time_key":                v["TimeKey"],
					"time_format":             v["TimeFormat"],
					"un_match_up_load_switch": v["UnMatchUpLoadSwitch"],
					"un_match_log_key":        v["UnMatchLogKey"],
				}
				if fkr, ok := v["FilterKeyRegex"].([]interface{}); ok {
					var fkrList []interface{}
					for _, f := range fkr {
						fm := f.(map[string]interface{})
						fkrList = append(fkrList, map[string]interface{}{
							"key":   fm["Key"],
							"regex": fm["Regex"],
						})
					}
					extractRule["filter_key_regex"] = fkrList
				}
				ruleMap["extract_rule"] = []interface{}{extractRule}
			}

			if v, ok := r["ExcludePaths"].([]interface{}); ok {
				var epList []interface{}
				for _, ep := range v {
					epm := ep.(map[string]interface{})
					epList = append(epList, map[string]interface{}{
						"type":  epm["Type"],
						"value": epm["Value"],
					})
				}
				ruleMap["exclude_paths"] = epList
			}

			if v, ok := r["ContainerRule"].(map[string]interface{}); ok {
				ruleMap["container_rule"] = []interface{}{
					map[string]interface{}{
						"stream":               v["Stream"],
						"container_name_regex": v["ContainerNameRegex"],
					},
				}
			}

			if v, ok := r["UserDefineRule"].(map[string]interface{}); ok {
				ruleMap["user_define_rule"] = []interface{}{
					map[string]interface{}{
						"enable_raw_log": v["EnableRawLog"],
						"tail_files":     v["TailFiles"],
					},
				}
			}
			res = append(res, ruleMap)
		}
		return res, nil
	})
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
