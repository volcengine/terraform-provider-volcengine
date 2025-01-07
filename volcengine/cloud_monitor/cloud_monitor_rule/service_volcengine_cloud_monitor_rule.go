package cloud_monitor_rule

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

type VolcengineCloudMonitorRuleService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewCloudMonitorRuleService(c *ve.SdkClient) *VolcengineCloudMonitorRuleService {
	return &VolcengineCloudMonitorRuleService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineCloudMonitorRuleService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineCloudMonitorRuleService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListRules"
		if condition != nil {
			if ids, exist := condition["Ids"]; exist && len(ids.([]interface{})) != 0 {
				action = "ListRulesByIds"
			}
		}

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

		for _, v := range data {
			ruleMap, ok := v.(map[string]interface{})
			if !ok {
				return data, fmt.Errorf("Result.Data Rule is not map")
			}
			dimensionArr := make([]interface{}, 0)
			if dimensions, exist := ruleMap["OriginalDimensions"]; exist {
				dimensionMap, ok := dimensions.(map[string]interface{})
				if !ok {
					return data, fmt.Errorf("OriginalDimensions is not map")
				}
				for key, value := range dimensionMap {
					dimensionArr = append(dimensionArr, map[string]interface{}{
						"Key":   key,
						"Value": value,
					})
				}
			}
			ruleMap["OriginalDimensions"] = dimensionArr
		}

		return data, err
	})
}

func (s *VolcengineCloudMonitorRuleService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"Ids": []interface{}{id},
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("cloud_monitor_rule %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineCloudMonitorRuleService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
					return nil, "", fmt.Errorf("cloud_monitor_rule status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (VolcengineCloudMonitorRuleService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineCloudMonitorRuleService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateRule",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"recovery_notify": {
					TargetField: "RecoveryNotify",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"enable": {
							TargetField: "Enable",
						},
					},
				},
				"alert_methods": {
					TargetField: "AlertMethods",
					ConvertType: ve.ConvertJsonArray,
				},
				"contact_group_ids": {
					TargetField: "ContactGroupIds",
					ConvertType: ve.ConvertJsonArray,
				},
				"regions": {
					TargetField: "Regions",
					ConvertType: ve.ConvertJsonArray,
				},
				"conditions": {
					TargetField: "Conditions",
					ConvertType: ve.ConvertJsonObjectArray,
					NextLevelConvert: map[string]ve.RequestConvert{
						"metric_name": {
							TargetField: "MetricName",
						},
						"metric_unit": {
							TargetField: "MetricUnit",
						},
						"statistics": {
							TargetField: "Statistics",
						},
						"comparison_operator": {
							TargetField: "ComparisonOperator",
						},
						"threshold": {
							TargetField: "Threshold",
						},
					},
				},
				"original_dimensions": {
					Ignore: true,
				},
				"webhook_ids": {
					TargetField: "WebhookIds",
					ConvertType: ve.ConvertJsonArray,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["RuleType"] = "static"

				dimensions := d.Get("original_dimensions").(*schema.Set).List()
				dimensionMap := make(map[string]interface{})
				for _, v := range dimensions {
					dimension, ok := v.(map[string]interface{})
					if !ok {
						return false, fmt.Errorf("dimension is not map")
					}
					value := dimension["value"].(*schema.Set).List()
					dimensionMap[dimension["key"].(string)] = value
				}
				(*call.SdkParam)["OriginalDimensions"] = dimensionMap

				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				data, err := ve.ObtainSdkValue("Result.Data", *resp)
				if err != nil {
					return err
				}
				dataArr, ok := data.([]interface{})
				if !ok || len(dataArr) == 0 {
					return fmt.Errorf("create cloud monitor rule failed")
				}
				d.SetId(dataArr[0].(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineCloudMonitorRuleService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateRule",
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
				"namespace": {
					TargetField: "Namespace",
					ForceGet:    true,
				},
				"sub_namespace": {
					TargetField: "SubNamespace",
					ForceGet:    true,
				},
				"level": {
					TargetField: "Level",
					ForceGet:    true,
				},
				"enable_state": {
					TargetField: "EnableState",
					ForceGet:    true,
				},
				"evaluation_count": {
					TargetField: "EvaluationCount",
					ForceGet:    true,
				},
				"effect_start_at": {
					TargetField: "EffectStartAt",
					ForceGet:    true,
				},
				"effect_end_at": {
					TargetField: "EffectEndAt",
					ForceGet:    true,
				},
				"silence_time": {
					TargetField: "SilenceTime",
					ForceGet:    true,
				},
				"web_hook": {
					//TargetField: "WebHook",
					Ignore: true,
				},
				"multiple_conditions": {
					TargetField: "MultipleConditions",
					ForceGet:    true,
				},
				"condition_operator": {
					TargetField: "ConditionOperator",
					ForceGet:    true,
				},
				"recovery_notify": {
					TargetField: "RecoveryNotify",
					ConvertType: ve.ConvertJsonObject,
					ForceGet:    true,
					NextLevelConvert: map[string]ve.RequestConvert{
						"enable": {
							TargetField: "Enable",
						},
					},
				},
				"alert_methods": {
					TargetField: "AlertMethods",
					ForceGet:    true,
					ConvertType: ve.ConvertJsonArray,
				},
				"contact_group_ids": {
					//TargetField: "ContactGroupIds",
					//ConvertType: ve.ConvertJsonArray,
					Ignore: true,
				},
				"regions": {
					TargetField: "Regions",
					ForceGet:    true,
					ConvertType: ve.ConvertJsonArray,
				},
				"conditions": {
					TargetField: "Conditions",
					ForceGet:    true,
					ConvertType: ve.ConvertJsonObjectArray,
					NextLevelConvert: map[string]ve.RequestConvert{
						"metric_name": {
							TargetField: "MetricName",
						},
						"metric_unit": {
							TargetField: "MetricUnit",
						},
						"statistics": {
							TargetField: "Statistics",
						},
						"comparison_operator": {
							TargetField: "ComparisonOperator",
						},
						"threshold": {
							TargetField: "Threshold",
						},
					},
				},
				"original_dimensions": {
					Ignore: true,
				},
				"webhook_ids": {
					TargetField: "WebhookIds",
					ConvertType: ve.ConvertJsonArray,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				// 将 original_dimensions 转为 map 形式
				dimensions := d.Get("original_dimensions").(*schema.Set).List()
				dimensionMap := make(map[string]interface{})
				for _, v := range dimensions {
					dimension, ok := v.(map[string]interface{})
					if !ok {
						return false, fmt.Errorf("dimension is not map")
					}
					value := dimension["value"].(*schema.Set).List()
					dimensionMap[dimension["key"].(string)] = value
				}
				(*call.SdkParam)["OriginalDimensions"] = dimensionMap

				methods := d.Get("alert_methods").(*schema.Set).List()
				// alert_methods 包含 Webhook
				if contains("Webhook", methods) {
					if webhook, ok := d.GetOk("web_hook"); ok {
						(*call.SdkParam)["WebHook"] = webhook
					}
				}
				// alert_methods 包含 Phone、Email、SMS
				if contains("Phone", methods) || contains("Email", methods) || contains("SMS", methods) {
					if groupIds, ok := d.GetOk("contact_group_ids"); ok {
						(*call.SdkParam)["ContactGroupIds"] = groupIds.(*schema.Set).List()
					}
				}

				(*call.SdkParam)["Id"] = d.Id()
				(*call.SdkParam)["RuleType"] = "static"
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

func (s *VolcengineCloudMonitorRuleService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteRulesByIds",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"Ids": []string{resourceData.Id()},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
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
							return resource.NonRetryableError(fmt.Errorf("error on reading cloud monitor rule on delete %q, %w", d.Id(), callErr))
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
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineCloudMonitorRuleService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "Ids",
				ConvertType: ve.ConvertJsonArray,
			},
			"alert_state": {
				TargetField: "AlertState",
				ConvertType: ve.ConvertJsonArray,
			},
			"namespace": {
				TargetField: "Namespace",
				ConvertType: ve.ConvertJsonArray,
			},
			"level": {
				TargetField: "Level",
				ConvertType: ve.ConvertJsonArray,
			},
			"enable_state": {
				TargetField: "EnableState",
				ConvertType: ve.ConvertJsonArray,
			},
		},
		NameField:    "RuleName",
		IdField:      "Id",
		CollectField: "rules",
		ContentType:  ve.ContentTypeJson,
	}
}

func (s *VolcengineCloudMonitorRuleService) ReadResourceId(id string) string {
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
