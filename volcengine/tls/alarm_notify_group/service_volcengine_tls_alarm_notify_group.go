package alarm_notify_group

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

type VolcengineTlsAlarmNotifyGroupService struct {
	Client *ve.SdkClient
}

func (v *VolcengineTlsAlarmNotifyGroupService) GetClient() *ve.SdkClient {
	return v.Client
}

func (v *VolcengineTlsAlarmNotifyGroupService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeAlarmNotifyGroups"
		logger.Debug(logger.ReqFormat, action, condition)
		resp, err = v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			ContentType: ve.Default,
			HttpMethod:  ve.GET,
			Path:        []string{action},
			Client:      v.Client.BypassSvcClient.NewTlsClient(),
		}, &condition)
		logger.Debug(logger.RespFormat, action, resp)
		results, err = ve.ObtainSdkValue("RESPONSE.AlarmNotifyGroups", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.AlarmNotifyGroups is not Slice")
		}
		return data, err
	})
}

func (v *VolcengineTlsAlarmNotifyGroupService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = v.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"AlarmNotifyGroupId": id,
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
		return data, fmt.Errorf("tls alarm notify group %s not exist ", id)
	}

	// ========== 核心修改1：仅保留 NoticeRules 的过滤逻辑，但不再解析 NotifyType/Receivers ==========
	hclHasNoticeRules := false
	if v, ok := resourceData.GetOk("notice_rules"); ok {
		if l, ok := v.([]interface{}); ok && len(l) > 0 {
			hclHasNoticeRules = true
		}
	}

	hclHasNotifyType := false
	if v, ok := resourceData.GetOk("notify_type"); ok {
		if s, okSet := v.(*schema.Set); okSet && s.Len() > 0 {
			hclHasNotifyType = true
		} else if l, okList := v.([]interface{}); okList && len(l) > 0 {
			// In some cases (like import), it might be a list
			hclHasNotifyType = true
		}
	}

	hclHasReceivers := false
	if v, ok := resourceData.GetOk("receivers"); ok {
		if l, ok := v.([]interface{}); ok && len(l) > 0 {
			hclHasReceivers = true
		}
	}

	// 检查是否处于 Import 状态
	// 在 Import 时，resourceData 只有 Id，Required 字段 (alarm_notify_group_name) 为空
	// 在 Create/Update/Read 时，Required 字段一定有值
	_, hasName := resourceData.GetOk("alarm_notify_group_name")
	isImport := !hasName

	if !isImport && !hclHasNoticeRules && (hclHasReceivers || hclHasNotifyType) {
		// Only recover NotifyType and Receivers from NoticeRules if they are empty in top level
		if nr, ok := data["NoticeRules"].([]interface{}); ok && len(nr) > 0 {
			rule, ok := nr[0].(map[string]interface{})
			if ok {
				logger.DebugInfo("NoticeRule for recovery check: %+v", rule)

				// Recover NotifyType if missing or empty
				shouldRecoverNotifyType := true
				if val, ok := data["NotifyType"]; ok {
					if slice, okSlice := val.([]interface{}); okSlice && len(slice) > 0 {
						shouldRecoverNotifyType = false
					}
				}

				if shouldRecoverNotifyType {
					if node, okRN := rule["RuleNode"].(map[string]interface{}); okRN {
						extractedNotifyTypes := extractNotifyTypesFromRuleNode(node)
						logger.DebugInfo("Extracted NotifyTypes: %+v", extractedNotifyTypes)
						if len(extractedNotifyTypes) > 0 {
							data["NotifyType"] = extractedNotifyTypes
						}
					}
				}

				// Recover Receivers if missing or empty
				shouldRecoverReceivers := true
				if val, ok := data["Receivers"]; ok {
					if slice, okSlice := val.([]interface{}); okSlice && len(slice) > 0 {
						shouldRecoverReceivers = false
					}
				}

				if shouldRecoverReceivers {
					if infos, okRI := rule["ReceiverInfos"]; okRI {
						logger.DebugInfo("Found ReceiverInfos: %+v", infos)
						data["Receivers"] = infos
					}
				}
			}
		}
		delete(data, "NoticeRules")
	}

	// Filter out fields that are not present in HCL, unless it is an import operation
	if !isImport {
		if !hclHasNoticeRules {
			delete(data, "NoticeRules")
		}
		if !hclHasNotifyType {
			delete(data, "NotifyType")
		}
		if !hclHasReceivers {
			delete(data, "Receivers")
		}
	}

	// ========== 核心修改2：确保 NotifyType 字段类型统一且非 nil ==========
	// 处理 NotifyType 类型转换（统一为 []interface{}）
	if val, ok := data["NotifyType"]; ok {
		switch t := val.(type) {
		case []string:
			// 转换 []string 到 []interface{}
			strs := make([]interface{}, len(t))
			for i, s := range t {
				strs[i] = s
			}
			data["NotifyType"] = strs
		case []interface{}:
			// 已经是 []interface{}，无需转换
			break
		default:
			// 未知类型，置为空列表
			logger.Info("Unknown NotifyType type: %T, value: %v", val, val)
			data["NotifyType"] = []interface{}{}
		}
	} else {
		// 字段不存在时，置为空列表
		data["NotifyType"] = []interface{}{}
	}

	// ========== 核心修改3：确保 Receivers 字段非 nil（可选，防止其他差异） ==========
	if _, ok := data["Receivers"]; !ok {
		data["Receivers"] = []interface{}{}
	}

	logger.DebugInfo("Final data map from ReadResource: %+v", data)
	return data, err
}

func (v *VolcengineTlsAlarmNotifyGroupService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, s string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (v *VolcengineTlsAlarmNotifyGroupService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		result := make(map[string]interface{})
		if val, ok := m["AlarmNotifyGroupId"]; ok {
			result["alarm_notify_group_id"] = val
		}
		if val, ok := m["AlarmNotifyGroupName"]; ok {
			result["alarm_notify_group_name"] = val
		}
		if val, ok := m["IamProjectName"]; ok {
			result["iam_project_name"] = val
		}

		if val, ok := m["NoticeRules"]; ok && val != nil {
			if sliceVal, ok := val.([]interface{}); ok && len(sliceVal) > 0 {
				result["notice_rules"] = convertNoticeRulesResponse(val)
			}
		}

		if val, ok := m["NotifyType"]; ok {
			result["NotifyType"] = val
		}
		if val, ok := m["Receivers"]; ok {
			result["Receivers"] = val
		}

		return result, map[string]ve.ResponseConvert{
			"NotifyType": {
				TargetField: "notify_type",
				Convert: func(v interface{}) interface{} {
					// 确保 []string/[]interface{} 都正确转换
					switch val := v.(type) {
					case []interface{}:
						strs := make([]interface{}, 0, len(val))
						for _, item := range val {
							if s, ok := item.(string); ok {
								strs = append(strs, s)
							} else {
								strs = append(strs, fmt.Sprintf("%v", item))
							}
						}
						return strs
					case []string:
						strs := make([]interface{}, len(val))
						for i, s := range val {
							strs[i] = s
						}
						return strs
					default:
						logger.Info("Unknown NotifyType type: %T, value: %v", v, v)
						return []interface{}{}
					}
				},
			},
			"Receivers": {
				TargetField: "receivers",
				Convert: func(v interface{}) interface{} {
					return convertReceiversResponse(v)
				},
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func extractNotifyTypesFromRuleNode(m map[string]interface{}) []interface{} {
	var results []interface{}
	if m == nil {
		return results
	}

	// 检查是否为条件节点
	if t, ok := m["Type"].(string); ok && t == "Condition" {
		if val, okV := m["Value"].([]interface{}); okV && len(val) >= 3 {
			field, _ := val[0].(string)
			if field == "NotifyType" {
				jsonStr, _ := val[2].(string)
				var ids []interface{}
				if err := json.Unmarshal([]byte(jsonStr), &ids); err == nil {
					for _, idRaw := range ids {
						id := fmt.Sprintf("%v", idRaw)
						switch id {
						case "1":
							results = append(results, "Trigger")
						case "2":
							results = append(results, "Recovery")
						default:
							results = append(results, id)
						}
					}
				}
			}
		}
	}

	// 递归检查子节点
	if children, okC := m["Children"].([]interface{}); okC {
		for _, child := range children {
			if childMap, okCM := child.(map[string]interface{}); okCM {
				results = append(results, extractNotifyTypesFromRuleNode(childMap)...)
			}
		}
	}

	return results
}

func convertReceiversResponse(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	list, ok := v.([]interface{})
	if !ok || len(list) == 0 {
		return nil
	}
	result := make([]interface{}, 0, len(list))
	for _, item := range list {
		converted := convertReceiver(item)
		if converted != nil {
			result = append(result, converted)
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

func convertNoticeRulesResponse(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	list, ok := v.([]interface{})
	if !ok || len(list) == 0 {
		return nil
	}
	result := make([]interface{}, 0, len(list))
	for _, item := range list {
		converted := convertNoticeRule(item)
		if converted != nil {
			result = append(result, converted)
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

func convertReceiver(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	m, ok := v.(map[string]interface{})
	if !ok {
		return nil
	}
	result := make(map[string]interface{})
	if val, ok := m["ReceiverType"]; ok {
		result["receiver_type"] = val
	}
	if val, ok := m["ReceiverNames"]; ok {
		result["receiver_names"] = val
	}
	if val, ok := m["ReceiverChannels"]; ok {
		result["receiver_channels"] = val
	}
	if val, ok := m["StartTime"]; ok {
		result["start_time"] = val
	}
	if val, ok := m["EndTime"]; ok {
		result["end_time"] = val
	}
	if val, ok := m["GeneralWebhookUrl"]; ok {
		result["general_webhook_url"] = val
	}
	if val, ok := m["GeneralWebhookBody"]; ok {
		result["general_webhook_body"] = val
	}
	if val, ok := m["GeneralWebhookMethod"]; ok {
		result["general_webhook_method"] = val
	}
	if val, ok := m["GeneralWebhookHeaders"]; ok {
		result["general_webhook_headers"] = convertWebhookHeaders(val)
	}
	if val, ok := m["AlarmWebhookAtUsers"]; ok {
		result["alarm_webhook_at_users"] = val
	}
	if val, ok := m["AlarmWebhookIsAtAll"]; ok {
		result["alarm_webhook_is_at_all"] = val
	}
	if val, ok := m["AlarmContentTemplateId"]; ok {
		result["alarm_content_template_id"] = val
	}
	if val, ok := m["AlarmWebhookIntegrationId"]; ok {
		result["alarm_webhook_integration_id"] = val
	}
	if val, ok := m["AlarmWebhookIntegrationName"]; ok {
		result["alarm_webhook_integration_name"] = val
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

func convertNoticeRule(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	m, ok := v.(map[string]interface{})
	if !ok {
		return nil
	}
	result := make(map[string]interface{})
	if val, ok := m["HasNext"]; ok {
		result["has_next"] = val
	}
	if val, ok := m["HasEndNode"]; ok {
		result["has_end_node"] = val
	}
	if val, ok := m["RuleNode"]; ok && val != nil {
		result["rule_node"] = []interface{}{convertRuleNodeResponse(val)}
	}
	if val, ok := m["ReceiverInfos"]; ok {
		result["receiver_infos"] = convertReceiverInfosResponse(val)
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

func convertRuleNodeResponse(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	m, ok := v.(map[string]interface{})
	if !ok {
		return nil
	}
	result := make(map[string]interface{})
	if val, ok := m["Type"]; ok {
		result["type"] = val
	}
	if val, ok := m["Value"]; ok {
		result["value"] = val
	}
	if val, ok := m["Children"]; ok && val != nil {
		children, ok := val.([]interface{})
		if ok {
			convertedChildren := make([]interface{}, 0, len(children))
			for _, child := range children {
				converted := convertRuleNodeResponse(child)
				if converted != nil {
					convertedChildren = append(convertedChildren, converted)
				}
			}
			result["children"] = convertedChildren
		}
	}
	return result
}

func convertReceiverInfosResponse(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	list, ok := v.([]interface{})
	if !ok || len(list) == 0 {
		return nil
	}
	result := make([]interface{}, 0, len(list))
	for _, item := range list {
		converted := convertReceiver(item)
		if converted != nil {
			result = append(result, converted)
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

func convertWebhookHeaders(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	list, ok := v.([]interface{})
	if !ok || len(list) == 0 {
		return nil
	}
	result := make([]interface{}, 0, len(list))
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		header := make(map[string]interface{})
		for k, val := range m {
			switch k {
			case "key", "Key":
				header["key"] = val
			case "value", "Value":
				header["value"] = val
			}
		}
		result = append(result, header)
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

func (v *VolcengineTlsAlarmNotifyGroupService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateAlarmNotifyGroup",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"alarm_notify_group_name": {
					TargetField: "AlarmNotifyGroupName",
					ForceGet:    true,
				},
				"notify_type": {
					TargetField: "NotifyType",
					ConvertType: ve.ConvertJsonArray,
					ForceGet:    true,
				},
				"iam_project_name": {
					TargetField: "IamProjectName",
				},
				"receivers": {
					TargetField: "Receivers",
					ConvertType: ve.ConvertJsonObjectArray,
					ForceGet:    true,
					NextLevelConvert: map[string]ve.RequestConvert{
						"receiver_names": {
							TargetField: "ReceiverNames",
							ConvertType: ve.ConvertJsonArray,
							ForceGet:    true,
						},
						"receiver_channels": {
							TargetField: "ReceiverChannels",
							ConvertType: ve.ConvertJsonArray,
							ForceGet:    true,
						},
						"receiver_type": {
							TargetField: "ReceiverType",
							ForceGet:    true,
						},
						"start_time": {
							TargetField: "StartTime",
							ForceGet:    true,
						},
						"end_time": {
							TargetField: "EndTime",
							ForceGet:    true,
						},
						"general_webhook_url": {
							TargetField: "GeneralWebhookUrl",
							ForceGet:    true,
						},
						"general_webhook_body": {
							TargetField: "GeneralWebhookBody",
							ForceGet:    true,
						},
						"alarm_webhook_at_users": {
							TargetField: "AlarmWebhookAtUsers",
							ConvertType: ve.ConvertJsonArray,
						},
						"alarm_webhook_is_at_all": {
							TargetField: "AlarmWebhookIsAtAll",
							ForceGet:    true,
						},
						"general_webhook_headers": {
							TargetField: "GeneralWebhookHeaders",
							ConvertType: ve.ConvertJsonObjectArray,
							NextLevelConvert: map[string]ve.RequestConvert{
								"key": {
									TargetField: "key",
									ForceGet:    true,
								},
								"value": {
									TargetField: "value",
									ForceGet:    true,
								},
							},
						},
						"general_webhook_method": {
							TargetField: "GeneralWebhookMethod",
							ForceGet:    true,
						},
						"alarm_content_template_id": {
							TargetField: "AlarmContentTemplateId",
						},
						"alarm_webhook_integration_id": {
							TargetField: "AlarmWebhookIntegrationId",
						},
						"alarm_webhook_integration_name": {
							TargetField: "AlarmWebhookIntegrationName",
						},
					},
				},
				"notice_rules": {
					TargetField: "NoticeRules",
					ConvertType: ve.ConvertJsonObjectArray,
					ForceGet:    true,
					NextLevelConvert: map[string]ve.RequestConvert{
						"has_next": {
							TargetField: "HasNext",
							ForceGet:    true,
						},
						"rule_node": {
							TargetField: "RuleNode",
							Convert: func(d *schema.ResourceData, v interface{}) interface{} {
								return convertRuleNode(v)
							},
						},
						"has_end_node": {
							TargetField: "HasEndNode",
							ForceGet:    true,
						},
						"receiver_infos": {
							TargetField: "ReceiverInfos",
							ConvertType: ve.ConvertJsonObjectArray,
							ForceGet:    true,
							NextLevelConvert: map[string]ve.RequestConvert{
								"receiver_names": {
									TargetField: "ReceiverNames",
									ConvertType: ve.ConvertJsonArray,
									ForceGet:    true,
								},
								"receiver_channels": {
									TargetField: "ReceiverChannels",
									ConvertType: ve.ConvertJsonArray,
									ForceGet:    true,
								},
								"receiver_type": {
									TargetField: "ReceiverType",
									ForceGet:    true,
								},
								"start_time": {
									TargetField: "StartTime",
									ForceGet:    true,
								},
								"end_time": {
									TargetField: "EndTime",
									ForceGet:    true,
								},
								"general_webhook_url": {
									TargetField: "GeneralWebhookUrl",
									ForceGet:    true,
								},
								"general_webhook_body": {
									TargetField: "GeneralWebhookBody",
									ForceGet:    true,
								},
								"alarm_webhook_at_users": {
									TargetField: "AlarmWebhookAtUsers",
									ConvertType: ve.ConvertJsonArray,
								},
								"alarm_webhook_is_at_all": {
									TargetField: "AlarmWebhookIsAtAll",
									ForceGet:    true,
								},
								"general_webhook_headers": {
									TargetField: "GeneralWebhookHeaders",
									ConvertType: ve.ConvertJsonObjectArray,
									NextLevelConvert: map[string]ve.RequestConvert{
										"key": {
											TargetField: "key",
											ForceGet:    true,
										},
										"value": {
											TargetField: "value",
											ForceGet:    true,
										},
									},
								},
								"general_webhook_method": {
									TargetField: "GeneralWebhookMethod",
									ForceGet:    true,
								},
								"alarm_content_template_id": {
									TargetField: "AlarmContentTemplateId",
								},
								"alarm_webhook_integration_id": {
									TargetField: "AlarmWebhookIntegrationId",
								},
								"alarm_webhook_integration_name": {
									TargetField: "AlarmWebhookIntegrationName",
								},
							},
						},
					},
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.POST,
					Path:        []string{call.Action},
					Client:      v.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, err := ve.ObtainSdkValue("RESPONSE.AlarmNotifyGroupId", *resp)
				if err != nil {
					return err
				}
				if s, ok := id.(string); ok {
					d.SetId(s)
				} else {
					return fmt.Errorf("AlarmNotifyGroupId is not string")
				}
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcengineTlsAlarmNotifyGroupService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyAlarmNotifyGroup",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"alarm_notify_group_name": {
					TargetField: "AlarmNotifyGroupName",
					ConvertType: ve.ConvertDefault,
					ForceGet:    true,
				},
				"notify_type": {
					TargetField: "NotifyType",
					ConvertType: ve.ConvertJsonArray,
					ForceGet:    true,
				},
				"iam_project_name": {
					Ignore: true,
				},
				"receivers": {
					TargetField: "Receivers",
					ConvertType: ve.ConvertJsonObjectArray,
					ForceGet:    true,
					NextLevelConvert: map[string]ve.RequestConvert{
						"receiver_names": {
							TargetField: "ReceiverNames",
							ConvertType: ve.ConvertJsonArray,
							ForceGet:    true,
						},
						"receiver_channels": {
							TargetField: "ReceiverChannels",
							ConvertType: ve.ConvertJsonArray,
							ForceGet:    true,
						},
						"receiver_type": {
							TargetField: "ReceiverType",
							ConvertType: ve.ConvertDefault,
							ForceGet:    true,
						},
						"start_time": {
							TargetField: "StartTime",
							ConvertType: ve.ConvertDefault,
							ForceGet:    true,
						},
						"end_time": {
							TargetField: "EndTime",
							ConvertType: ve.ConvertDefault,
							ForceGet:    true,
						},
						"general_webhook_url": {
							TargetField: "GeneralWebhookUrl",
							ForceGet:    true,
						},
						"general_webhook_body": {
							TargetField: "GeneralWebhookBody",
							ForceGet:    true,
						},
						"alarm_webhook_at_users": {
							TargetField: "AlarmWebhookAtUsers",
							ConvertType: ve.ConvertJsonArray,
						},
						"alarm_webhook_is_at_all": {
							TargetField: "AlarmWebhookIsAtAll",
							ForceGet:    true,
						},
						"general_webhook_headers": {
							TargetField: "GeneralWebhookHeaders",
							ConvertType: ve.ConvertJsonObjectArray,
							ForceGet:    true,
							NextLevelConvert: map[string]ve.RequestConvert{
								"key": {
									TargetField: "key",
									ForceGet:    true,
								},
								"value": {
									TargetField: "value",
									ForceGet:    true,
								},
							},
						},
						"general_webhook_method": {
							TargetField: "GeneralWebhookMethod",
							ForceGet:    true,
						},
						"alarm_content_template_id": {
							TargetField: "AlarmContentTemplateId",
						},
						"alarm_webhook_integration_id": {
							TargetField: "AlarmWebhookIntegrationId",
						},
						"alarm_webhook_integration_name": {
							TargetField: "AlarmWebhookIntegrationName",
						},
					},
				},
				"notice_rules": {
					TargetField: "NoticeRules",
					ConvertType: ve.ConvertJsonObjectArray,
					ForceGet:    true,
					NextLevelConvert: map[string]ve.RequestConvert{
						"has_next": {
							TargetField: "HasNext",
							ForceGet:    true,
						},
						"rule_node": {
							TargetField: "RuleNode",
							Convert: func(d *schema.ResourceData, v interface{}) interface{} {
								return convertRuleNode(v)
							},
						},
						"has_end_node": {
							TargetField: "HasEndNode",
							ForceGet:    true,
						},
						"receiver_infos": {
							TargetField: "ReceiverInfos",
							ConvertType: ve.ConvertJsonObjectArray,
							ForceGet:    true,
							NextLevelConvert: map[string]ve.RequestConvert{
								"receiver_names": {
									TargetField: "ReceiverNames",
									ConvertType: ve.ConvertJsonArray,
									ForceGet:    true,
								},
								"receiver_channels": {
									TargetField: "ReceiverChannels",
									ConvertType: ve.ConvertJsonArray,
									ForceGet:    true,
								},
								"receiver_type": {
									TargetField: "ReceiverType",
									ConvertType: ve.ConvertDefault,
									ForceGet:    true,
								},
								"start_time": {
									TargetField: "StartTime",
									ConvertType: ve.ConvertDefault,
									ForceGet:    true,
								},
								"end_time": {
									TargetField: "EndTime",
									ConvertType: ve.ConvertDefault,
									ForceGet:    true,
								},
								"general_webhook_url": {
									TargetField: "GeneralWebhookUrl",
									ForceGet:    true,
								},
								"general_webhook_body": {
									TargetField: "GeneralWebhookBody",
									ForceGet:    true,
								},
								"alarm_webhook_at_users": {
									TargetField: "AlarmWebhookAtUsers",
									ConvertType: ve.ConvertJsonArray,
								},
								"alarm_webhook_is_at_all": {
									TargetField: "AlarmWebhookIsAtAll",
									ForceGet:    true,
								},
								"general_webhook_headers": {
									TargetField: "GeneralWebhookHeaders",
									ConvertType: ve.ConvertJsonObjectArray,
									ForceGet:    true,
									NextLevelConvert: map[string]ve.RequestConvert{
										"key": {
											TargetField: "key",
											ForceGet:    true,
										},
										"value": {
											TargetField: "value",
											ForceGet:    true,
										},
									},
								},
								"general_webhook_method": {
									TargetField: "GeneralWebhookMethod",
									ForceGet:    true,
								},
								"alarm_content_template_id": {
									TargetField: "AlarmContentTemplateId",
								},
								"alarm_webhook_integration_id": {
									TargetField: "AlarmWebhookIntegrationId",
								},
								"alarm_webhook_integration_name": {
									TargetField: "AlarmWebhookIntegrationName",
								},
							},
						},
					},
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.PUT,
					Path:        []string{call.Action},
					Client:      v.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["AlarmNotifyGroupId"] = d.Id()
				return true, nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcengineTlsAlarmNotifyGroupService) RemoveResource(data *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteAlarmNotifyGroup",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"AlarmNotifyGroupId": data.Id(),
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
							return resource.NonRetryableError(fmt.Errorf("error on reading tls alarm on delete %q, %w", d.Id(), callErr))
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

func (v *VolcengineTlsAlarmNotifyGroupService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		NameField:    "AlarmNotifyGroupName",
		IdField:      "AlarmNotifyGroupId",
		CollectField: "groups",
	}
}

func (v *VolcengineTlsAlarmNotifyGroupService) ReadResourceId(s string) string {
	return s
}

func NewTlsAlarmNotifyGroupService(client *ve.SdkClient) *VolcengineTlsAlarmNotifyGroupService {
	return &VolcengineTlsAlarmNotifyGroupService{
		Client: client,
	}
}

func (*VolcengineTlsAlarmNotifyGroupService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "tls",
		ResourceType:         "alarmnotifygroup",
		ProjectSchemaField:   "iam_project_name",
		ProjectResponseField: "IamProjectName",
	}
}

func convertRuleNode(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	list, ok := v.([]interface{})
	if !ok || len(list) == 0 {
		return nil
	}
	// Take first element
	nodeMap, ok := list[0].(map[string]interface{})
	if !ok {
		return nil
	}
	return convertRuleNodeMap(nodeMap)
}

func convertRuleNodeMap(m map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	if val, ok := m["type"]; ok {
		result["Type"] = val
	}
	if val, ok := m["value"]; ok {
		strList := convertToStringList(val)
		if len(strList) > 0 {
			result["Value"] = strList
		}
	}
	if val, ok := m["children"]; ok {
		childrenList, ok := val.([]interface{})
		if ok && len(childrenList) > 0 {
			convertedChildren := make([]interface{}, 0, len(childrenList))
			for _, child := range childrenList {
				childMap, ok := child.(map[string]interface{})
				if ok {
					converted := convertRuleNodeMap(childMap)
					if converted != nil {
						convertedChildren = append(convertedChildren, converted)
					}
				}
			}
			if len(convertedChildren) > 0 {
				result["Children"] = convertedChildren
			}
		}
	}
	return result
}

func convertToStringList(v interface{}) []string {
	if v == nil {
		return nil
	}
	list, ok := v.([]interface{})
	if !ok {
		return nil
	}
	result := make([]string, 0, len(list))
	for _, item := range list {
		if s, ok := item.(string); ok {
			result = append(result, s)
		}
	}
	return result
}
