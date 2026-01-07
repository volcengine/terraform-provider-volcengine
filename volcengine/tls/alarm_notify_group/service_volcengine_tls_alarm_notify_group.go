package alarm_notify_group

import (
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
		}, &m)
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
	return data, err
}

func (v *VolcengineTlsAlarmNotifyGroupService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, s string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (v *VolcengineTlsAlarmNotifyGroupService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return m, map[string]ve.ResponseConvert{
			"RuleNode": {
				Convert: func(v interface{}) interface{} {
					if v == nil {
						return nil
					}
					return []interface{}{v}
				},
			},
			"Children": {
				// Chain: "RuleNode",
				Convert: func(v interface{}) interface{} {
					return v
				},
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (v *VolcengineTlsAlarmNotifyGroupService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateAlarmNotifyGroup",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"notify_type": {
					ConvertType: ve.ConvertJsonArray,
				},
				"receivers": {
					ConvertType: ve.ConvertJsonObjectArray,
					NextLevelConvert: map[string]ve.RequestConvert{
						"receiver_names": {
							ConvertType: ve.ConvertJsonArray,
						},
						"receiver_channels": {
							ConvertType: ve.ConvertJsonArray,
						},
						"general_webhook_url": {
							TargetField: "GeneralWebhookUrl",
						},
						"general_webhook_body": {
							TargetField: "GeneralWebhookBody",
						},
						"alarm_webhook_at_users": {
							TargetField: "AlarmWebhookAtUsers",
							ConvertType: ve.ConvertJsonArray,
						},
						"alarm_webhook_is_at_all": {
							TargetField: "AlarmWebhookIsAtAll",
						},
						"general_webhook_headers": {
							TargetField: "GeneralWebhookHeaders",
							ConvertType: ve.ConvertJsonObjectArray,
							ForceGet:    true,
							NextLevelConvert: map[string]ve.RequestConvert{
								"key": {
									TargetField: "key",
								},
								"value": {
									TargetField: "value",
								},
							},
						},
						"general_webhook_method": {
							TargetField: "GeneralWebhookMethod",
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
					ConvertType: ve.ConvertJsonObjectArray,
					NextLevelConvert: map[string]ve.RequestConvert{
						"has_next": {
							TargetField: "HasNext",
						},
						"rule_node": {
							TargetField: "RuleNode",
							Convert: func(d *schema.ResourceData, v interface{}) interface{} {
								return convertRuleNode(v)
							},
						},
						"has_end_node": {
							TargetField: "HasEndNode",
						},
						"receiver_infos": {
							TargetField: "ReceiverInfos",
							ConvertType: ve.ConvertJsonObjectArray,
							NextLevelConvert: map[string]ve.RequestConvert{
								"receiver_names": {
									ConvertType: ve.ConvertJsonArray,
								},
								"receiver_channels": {
									ConvertType: ve.ConvertJsonArray,
								},
								"receiver_type": {
									ConvertType: ve.ConvertDefault,
									ForceGet:    true,
								},
								"start_time": {
									ConvertType: ve.ConvertDefault,
									ForceGet:    true,
								},
								"end_time": {
									ConvertType: ve.ConvertDefault,
									ForceGet:    true,
								},
								"general_webhook_url": {
									TargetField: "GeneralWebhookUrl",
								},
								"general_webhook_body": {
									TargetField: "GeneralWebhookBody",
								},
								"alarm_webhook_at_users": {
									TargetField: "AlarmWebhookAtUsers",
									ConvertType: ve.ConvertJsonArray,
								},
								"alarm_webhook_is_at_all": {
									TargetField: "AlarmWebhookIsAtAll",
								},
								"general_webhook_headers": {
									TargetField: "GeneralWebhookHeaders",
									ConvertType: ve.ConvertJsonObjectArray,
									ForceGet:    true,
									NextLevelConvert: map[string]ve.RequestConvert{
										"key": {
											TargetField: "key",
										},
										"value": {
											TargetField: "value",
										},
									},
								},
								"general_webhook_method": {
									TargetField: "GeneralWebhookMethod",
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
				id, _ := ve.ObtainSdkValue("RESPONSE.AlarmNotifyGroupId", *resp)
				d.SetId(id.(string))
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
					ConvertType: ve.ConvertDefault,
				},
				"notify_type": {
					ConvertType: ve.ConvertJsonArray,
				},
				"iam_project_name": {
					Ignore: true,
				},
				"receivers": {
					ConvertType: ve.ConvertJsonObjectArray,
					NextLevelConvert: map[string]ve.RequestConvert{
						"receiver_names": {
							ConvertType: ve.ConvertJsonArray,
							ForceGet:    true,
						},
						"receiver_channels": {
							ConvertType: ve.ConvertJsonArray,
							ForceGet:    true,
						},
						"receiver_type": {
							ConvertType: ve.ConvertDefault,
							ForceGet:    true,
						},
						"start_time": {
							ConvertType: ve.ConvertDefault,
							ForceGet:    true,
						},
						"end_time": {
							ConvertType: ve.ConvertDefault,
							ForceGet:    true,
						},
						"general_webhook_url": {
							TargetField: "GeneralWebhookUrl",
						},
						"general_webhook_body": {
							TargetField: "GeneralWebhookBody",
						},
						"alarm_webhook_at_users": {
							TargetField: "AlarmWebhookAtUsers",
							ConvertType: ve.ConvertJsonArray,
						},
						"alarm_webhook_is_at_all": {
							TargetField: "AlarmWebhookIsAtAll",
						},
						"general_webhook_headers": {
							TargetField: "GeneralWebhookHeaders",
							ConvertType: ve.ConvertJsonObjectArray,
							ForceGet:    true,
							NextLevelConvert: map[string]ve.RequestConvert{
								"key": {
									TargetField: "key",
								},
								"value": {
									TargetField: "value",
								},
							},
						},
						"general_webhook_method": {
							TargetField: "GeneralWebhookMethod",
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
					ConvertType: ve.ConvertJsonObjectArray,
					NextLevelConvert: map[string]ve.RequestConvert{
						"has_next": {
							TargetField: "HasNext",
						},
						"rule_node": {
							TargetField: "RuleNode",
							Convert: func(d *schema.ResourceData, v interface{}) interface{} {
								return convertRuleNode(v)
							},
						},
						"has_end_node": {
							TargetField: "HasEndNode",
						},
						"receiver_infos": {
							TargetField: "ReceiverInfos",
							ConvertType: ve.ConvertJsonObjectArray,
							NextLevelConvert: map[string]ve.RequestConvert{
								"receiver_names": {
									ConvertType: ve.ConvertJsonArray,
									ForceGet:    true,
								},
								"receiver_channels": {
									ConvertType: ve.ConvertJsonArray,
									ForceGet:    true,
								},
								"receiver_type": {
									ConvertType: ve.ConvertDefault,
									ForceGet:    true,
								},
								"start_time": {
									ConvertType: ve.ConvertDefault,
									ForceGet:    true,
								},
								"end_time": {
									ConvertType: ve.ConvertDefault,
									ForceGet:    true,
								},
								"general_webhook_url": {
									TargetField: "GeneralWebhookUrl",
								},
								"general_webhook_body": {
									TargetField: "GeneralWebhookBody",
								},
								"alarm_webhook_at_users": {
									TargetField: "AlarmWebhookAtUsers",
									ConvertType: ve.ConvertJsonArray,
								},
								"alarm_webhook_is_at_all": {
									TargetField: "AlarmWebhookIsAtAll",
								},
								"general_webhook_headers": {
									TargetField: "GeneralWebhookHeaders",
									ConvertType: ve.ConvertJsonObjectArray,
									ForceGet:    true,
									NextLevelConvert: map[string]ve.RequestConvert{
										"key": {
											TargetField: "key",
										},
										"value": {
											TargetField: "value",
										},
									},
								},
								"general_webhook_method": {
									TargetField: "GeneralWebhookMethod",
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
