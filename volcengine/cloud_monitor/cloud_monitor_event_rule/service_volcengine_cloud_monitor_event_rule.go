package cloud_monitor_event_rule

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

type VolcengineCloudMonitorEventRuleService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewCloudMonitorEventRuleService(c *ve.SdkClient) *VolcengineCloudMonitorEventRuleService {
	return &VolcengineCloudMonitorEventRuleService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineCloudMonitorEventRuleService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineCloudMonitorEventRuleService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListEventRules"
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
		results, err = ve.ObtainSdkValue("Result.Data", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Data is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineCloudMonitorEventRuleService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		temp    map[string]interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if temp, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		} else if temp["RuleId"].(string) == id {
			data = temp
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("cloud_monitor_event_rule %s not exist ", id)
	}
	data["Status"] = data["EnableState"]
	data["EventSource"] = data["Source"]
	startTime := data["EffectStartAt"]
	endTime := data["EffectEndAt"]
	data["EffectiveTime"] = map[string]interface{}{
		"StartTime": startTime,
		"EndTime":   endTime,
	}
	return data, err
}

func (s *VolcengineCloudMonitorEventRuleService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
					return nil, "", fmt.Errorf("cloud_monitor_event_rule status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineCloudMonitorEventRuleService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	status := resourceData.Get("status")
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateEventRule",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"event_type": {
					TargetField: "EventType",
					ConvertType: ve.ConvertJsonArray,
				},
				"filter_pattern": {
					TargetField: "FilterPattern",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"type": {
							TargetField: "Type",
							ConvertType: ve.ConvertJsonArray,
						},
						"source": {
							TargetField: "Source",
						},
					},
				},
				"effective_time": {
					TargetField: "EffectiveTime",
					ConvertType: ve.ConvertJsonObject,
				},
				"contact_methods": {
					TargetField: "ContactMethods",
					ConvertType: ve.ConvertJsonArray,
				},
				"contact_group_ids": {
					TargetField: "ContactGroupIds",
					ConvertType: ve.ConvertJsonArray,
				},
				"tls_target": {
					TargetField: "TLSTarget",
					ConvertType: ve.ConvertJsonObjectArray,
					NextLevelConvert: map[string]ve.RequestConvert{
						"region_name_en": {
							TargetField: "RegionNameEN",
						},
						"region_name_cn": {
							TargetField: "RegionNameCN",
						},
					},
				},
				"message_queue": {
					TargetField: "MessageQueue",
					ConvertType: ve.ConvertJsonObjectArray,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["EventBusName"] = "default"
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.Data.RuleId", *resp)
				d.SetId(id.(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{status.(string)},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineCloudMonitorEventRuleService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"TLSTarget": {
				TargetField: "tls_target",
			},
			"RegionNameEN": {
				TargetField: "region_name_en",
			},
			"RegionNameCN": {
				TargetField: "region_name_cn",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineCloudMonitorEventRuleService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	status := resourceData.Get("status")
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateEventRule",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"rule_name": {
					TargetField: "RuleName",
					ForceGet:    true,
				},
				"description": {
					TargetField: "Description",
				},
				"event_source": {
					TargetField: "EventSource",
					ForceGet:    true,
				},
				"event_type": {
					TargetField: "EventType",
					ForceGet:    true,
					ConvertType: ve.ConvertJsonArray,
				},
				"status": {
					TargetField: "Status",
					ForceGet:    true,
				},
				"level": {
					TargetField: "Level",
					ForceGet:    true,
				},
				"filter_pattern": {
					TargetField: "FilterPattern",
					ConvertType: ve.ConvertJsonObject,
					ForceGet:    true,
					NextLevelConvert: map[string]ve.RequestConvert{
						"type": {
							TargetField: "Type",
							ForceGet:    true,
							ConvertType: ve.ConvertJsonArray,
						},
						"source": {
							ForceGet:    true,
							TargetField: "Source",
						},
					},
				},
				"effective_time": {
					TargetField: "EffectiveTime",
					ForceGet:    true,
					ConvertType: ve.ConvertJsonObject,
				},
				"contact_methods": {
					TargetField: "ContactMethods",
					ForceGet:    true,
					ConvertType: ve.ConvertJsonArray,
				},
				"contact_group_ids": {
					Ignore: true,
				},
				"endpoint": {
					Ignore: true,
				},
				"tls_target": {
					Ignore: true,
				},
				"message_queue": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["RuleId"] = d.Id()
				(*call.SdkParam)["EventBusName"] = "default"
				methods := d.Get("contact_methods").(*schema.Set).List()
				// 如methods含TLS，传TlsTarget
				if contains("TLS", methods) {
					if tls, ok := d.GetOk("tls_target"); ok {
						tlsT := make([]interface{}, 0)
						for _, t := range tls.(*schema.Set).List() {
							tMap := t.(map[string]interface{})
							tlsT = append(tlsT, map[string]interface{}{
								"RegionNameEN": tMap["region_name_en"],
								"RegionNameCN": tMap["region_name_cn"],
								"ProjectName":  tMap["project_name"],
								"ProjectId":    tMap["project_id"],
								"TopicId":      tMap["topic_id"],
							})
						}
						(*call.SdkParam)["TLSTarget"] = tlsT
					}
				}
				// 如methods含MQ，传messageQueue
				if contains("MQ", methods) {
					if mq, ok := d.GetOk("message_queue"); ok {
						messageQ := make([]interface{}, 0)
						for _, m := range mq.(*schema.Set).List() {
							mMap := m.(map[string]interface{})
							messageQ = append(messageQ, map[string]interface{}{
								"InstanceId": mMap["instance_id"],
								"Region":     mMap["region"],
								"Topic":      mMap["topic"],
								"Type":       mMap["type"],
								"VpcId":      mMap["vpc_id"],
							})
						}
						(*call.SdkParam)["MessageQueue"] = messageQ
					}
				}
				// 如methods含Webhook，传endpoint
				if contains("Webhook", methods) {
					if endpoint, ok := d.GetOk("endpoint"); ok {
						(*call.SdkParam)["Endpoint"] = endpoint
					}
				}
				// 如method含Phone、Email、SMS，传contactGroupIds
				if contains("Phone", methods) || contains("Email", methods) || contains("SMS", methods) {
					if groupIds, ok := d.GetOk("contact_group_ids"); ok {
						(*call.SdkParam)["ContactGroupIds"] = groupIds.(*schema.Set).List()
					}
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{status.(string)},
				Timeout: resourceData.Timeout(schema.TimeoutUpdate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineCloudMonitorEventRuleService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteEventRule",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"RuleId":       []string{resourceData.Id()},
				"EventBusName": "default",
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineCloudMonitorEventRuleService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		NameField:    "RuleName",
		IdField:      "RuleId",
		CollectField: "rules",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"RuleId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"Source": {
				TargetField: "event_source",
			},
			"EnableState": {
				TargetField: "status",
			},
			"TLSTarget": {
				TargetField: "tls_target",
			},
			"RegionNameEN": {
				TargetField: "region_name_en",
			},
			"RegionNameCN": {
				TargetField: "region_name_cn",
			},
		},
	}
}

func (s *VolcengineCloudMonitorEventRuleService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "Volc_Observe",
		Version:     "2018-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
