package scaling_policy

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
	"github.com/volcengine/terraform-provider-vestack/logger"
)

type VestackScalingPolicyService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewScalingPolicyService(c *ve.SdkClient) *VestackScalingPolicyService {
	return &VestackScalingPolicyService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VestackScalingPolicyService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VestackScalingPolicyService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		autoScalingClient := s.Client.AutoScalingClient
		action := "DescribeScalingPolicies"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = autoScalingClient.DescribeScalingPoliciesCommon(nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = autoScalingClient.DescribeScalingPoliciesCommon(&condition)
			if err != nil {
				return data, err
			}
		}
		logger.Debug(logger.RespFormat, action, action, *resp)
		results, err = ve.ObtainSdkValue("Result.ScalingPolicies", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.ScalingPolicies is not Slice")
		}
		return data, err
	})
}

func (s *VestackScalingPolicyService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	req := map[string]interface{}{
		"ScalingPolicyIds.1": ids[1],
		"ScalingGroupId":     ids[0],
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
		return data, fmt.Errorf("ScalingPolicy %s not exist ", id)
	}
	return data, err
}

func (s *VestackScalingPolicyService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				demo       map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "Error")
			demo, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status", demo)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("ScalingPolicy Status error, status:%s", status.(string))
				}
			}
			//注意 返回的第一个参数不能为空 否则会一直等下去
			return demo, status.(string), err
		},
	}

}

func (VestackScalingPolicyService) WithResourceResponseHandlers(scalingPolicy map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		scalingPolicy["active"] = scalingPolicy["Status"].(string) == "Active"
		return scalingPolicy, nil, nil
	}

	return []ve.ResourceResponseHandler{handler}

}

func (s *VestackScalingPolicyService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateScalingPolicy",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"scheduled_policy_launch_time": {
					TargetField: "ScheduledPolicy.LaunchTime",
				},
				"scheduled_policy_recurrence_end_time": {
					TargetField: "ScheduledPolicy.RecurrenceEndTime",
				},
				"scheduled_policy_recurrence_type": {
					TargetField: "ScheduledPolicy.RecurrenceType",
				},
				"scheduled_policy_recurrence_value": {
					TargetField: "ScheduledPolicy.RecurrenceValue",
				},
				"alarm_policy_rule_type": {
					TargetField: "AlarmPolicy.RuleType",
				},
				"alarm_policy_evaluation_count": {
					TargetField: "AlarmPolicy.EvaluationCount",
				},
				"alarm_policy_condition_metric_name": {
					TargetField: "AlarmPolicy.Condition.MetricName",
				},
				"alarm_policy_condition_metric_unit": {
					TargetField: "AlarmPolicy.Condition.MetricUnit",
				},
				"alarm_policy_condition_comparison_operator": {
					TargetField: "AlarmPolicy.Condition.ComparisonOperator",
				},
				"alarm_policy_condition_threshold": {
					TargetField: "AlarmPolicy.Condition.Threshold",
				},
				"cooldown": {
					TargetField: "Cooldown",
					Convert: func(data *schema.ResourceData, i interface{}) interface{} {
						return i
					},
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.AutoScalingClient.CreateScalingPolicyCommon(call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				//注意 获取内容 这个地方不能是指针 需要转一次
				id, _ := ve.ObtainSdkValue("Result.ScalingPolicyId", *resp)
				d.SetId(fmt.Sprintf("%v:%v", d.Get("scaling_group_id"), id))
				if resourceData.Get("active") != nil && resourceData.Get("active").(bool) {
					//action := "EnableScalingGroup"
					//param := &map[string]interface{}{"ScalingGroupId": resourceData.Get("scaling_group_id")}
					//logger.Debug(logger.RespFormat, action, param)
					//if _, err := s.Client.AutoScalingClient.EnableScalingGroupCommon(param); err != nil {
					//	logger.Debug(logger.ErrFormat, action, param, err)
					//	return err
					//}

					action := "EnableScalingPolicy"
					param := &map[string]interface{}{"ScalingPolicyId": id}
					logger.Debug(logger.RespFormat, action, param)
					if _, err := s.Client.AutoScalingClient.EnableScalingPolicyCommon(param); err != nil {
						logger.Debug(logger.ErrFormat, action, param, err)
						return err
					}
				}
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"InActive", "Active"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}

}

func (s *VestackScalingPolicyService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callbacks := make([]ve.Callback, 0)
	ids := strings.Split(resourceData.Id(), ":")

	// 修改伸缩规则属性
	modifyScalingPolicyCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyScalingPolicy",
			ConvertMode: ve.RequestConvertInConvert,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) == 0 {
					return false, nil
				}

				// 修改定时规则时，需要传整个对象
				if resourceData.HasChange("scheduled_policy_launch_time") ||
					resourceData.HasChange("scheduled_policy_recurrence_end_time") ||
					resourceData.HasChange("scheduled_policy_recurrence_type") ||
					resourceData.HasChange("scheduled_policy_recurrence_value") {
					(*call.SdkParam)["ScheduledPolicy.LaunchTime"] = d.Get("scheduled_policy_launch_time")
					(*call.SdkParam)["ScheduledPolicy.RecurrenceEndTime"] = d.Get("scheduled_policy_recurrence_end_time")
					(*call.SdkParam)["ScheduledPolicy.RecurrenceType"] = d.Get("scheduled_policy_recurrence_type")
					(*call.SdkParam)["ScheduledPolicy.RecurrenceValue"] = d.Get("scheduled_policy_recurrence_value")
				}

				if resourceData.HasChange("alarm_policy_rule_type") ||
					resourceData.HasChange("alarm_policy_evaluation_count") ||
					resourceData.HasChange("alarm_policy_condition_metric_name") ||
					resourceData.HasChange("alarm_policy_condition_metric_unit") ||
					resourceData.HasChange("alarm_policy_condition_comparison_operator") ||
					resourceData.HasChange("alarm_policy_condition_threshold") {
					(*call.SdkParam)["AlarmPolicy.RuleType"] = d.Get("alarm_policy_rule_type")
					(*call.SdkParam)["AlarmPolicy.EvaluationCount"] = d.Get("alarm_policy_evaluation_count")
					(*call.SdkParam)["AlarmPolicy.Condition.MetricName"] = d.Get("alarm_policy_condition_metric_name")
					(*call.SdkParam)["AlarmPolicy.Condition.MetricUnit"] = d.Get("alarm_policy_condition_metric_unit")
					(*call.SdkParam)["AlarmPolicy.Condition.ComparisonOperator"] = d.Get("alarm_policy_condition_comparison_operator")
					(*call.SdkParam)["AlarmPolicy.Condition.Threshold"] = d.Get("alarm_policy_condition_threshold")
				}

				(*call.SdkParam)["ScalingPolicyId"] = ids[1]
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.AutoScalingClient.ModifyScalingPolicyCommon(call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return nil
			},
			Convert: map[string]ve.RequestConvert{
				"scheduled_policy_launch_time": {
					TargetField: "ScheduledPolicy.LaunchTime",
				},
				"scheduled_policy_recurrence_end_time": {
					TargetField: "ScheduledPolicy.RecurrenceEndTime",
				},
				"scheduled_policy_recurrence_type": {
					TargetField: "ScheduledPolicy.RecurrenceType",
				},
				"scheduled_policy_recurrence_value": {
					TargetField: "ScheduledPolicy.RecurrenceValue",
				},
				"alarm_policy_rule_type": {
					TargetField: "AlarmPolicy.RuleType",
				},
				"alarm_policy_evaluation_count": {
					TargetField: "AlarmPolicy.EvaluationCount",
				},
				"alarm_policy_condition_metric_name": {
					TargetField: "AlarmPolicy.Condition.MetricName",
				},
				"alarm_policy_condition_metric_unit": {
					TargetField: "AlarmPolicy.Condition.MetricUnit",
				},
				"alarm_policy_condition_comparison_operator": {
					TargetField: "AlarmPolicy.Condition.ComparisonOperator",
				},
				"alarm_policy_condition_threshold": {
					TargetField: "AlarmPolicy.Condition.Threshold",
				},
				"scaling_policy_name": {
					ConvertType: ve.ConvertDefault,
				},
				"adjustment_type": {
					ConvertType: ve.ConvertDefault,
				},
				"adjustment_value": {
					ConvertType: ve.ConvertDefault,
				},
				"cooldown": {
					TargetField: "Cooldown",
					Convert: func(data *schema.ResourceData, i interface{}) interface{} {
						return i
					},
				},
			},
		},
	}
	callbacks = append(callbacks, modifyScalingPolicyCallback)

	// 使能伸缩规则
	if resourceData.HasChange("active") {
		// 开启伸缩规则之前，必须使能伸缩组
		//if resourceData.Get("active").(bool) {
		//	callbacks = append(callbacks, ve.Callback{
		//		Call: ve.SdkCall{
		//			Action: "EnableScalingGroup",
		//			SdkParam: &map[string]interface{}{
		//				"ScalingGroupId": resourceData.Get("scaling_group_id"),
		//			},
		//			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
		//				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
		//				return s.Client.AutoScalingClient.EnableScalingGroupCommon(call.SdkParam)
		//			},
		//		},
		//	})
		//}

		// 伸缩规则状态变更
		callbacks = append(callbacks, s.enableOrDisablePolicyCallback(ids[1], resourceData))
	}

	return callbacks
}

func (s *VestackScalingPolicyService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	ids := strings.Split(resourceData.Id(), ":")
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteScalingPolicy",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"ScalingPolicyId": ids[1],
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.AutoScalingClient.DeleteScalingPolicyCommon(call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading ScalingPolicy on delete %q, %w", d.Id(), callErr))
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

func (s *VestackScalingPolicyService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "ScalingPolicyIds",
				ConvertType: ve.ConvertWithN,
			},
			"scaling_policy_names": {
				TargetField: "ScalingPolicyIds",
				ConvertType: ve.ConvertWithN,
			},
		},
		NameField:    "ScalingPolicyName",
		IdField:      "ScalingPolicyId",
		CollectField: "scaling_policies",
		ResponseConverts: map[string]ve.ResponseConvert{
			"ScalingPolicyId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"ScheduledPolicy.LaunchTime": {
				TargetField: "scheduled_policy_launch_time",
			},
			"ScheduledPolicy.RecurrenceStartTime": {
				TargetField: "scheduled_policy_recurrence_start_time",
			},
			"ScheduledPolicy.RecurrenceEndTime": {
				TargetField: "scheduled_policy_recurrence_end_time",
			},
			"ScheduledPolicy.RecurrenceType": {
				TargetField: "scheduled_policy_recurrence_type",
			},
			"ScheduledPolicy.RecurrenceValue": {
				TargetField: "scheduled_policy_recurrence_value",
			},
			"AlarmPolicy.RuleType": {
				TargetField: "alarm_policy_rule_type",
			},
			"AlarmPolicy.EvaluationCount": {
				TargetField: "alarm_policy_evaluation_count",
			},
			"AlarmPolicy.Condition.MetricName": {
				TargetField: "alarm_policy_condition_metric_name",
			},
			"AlarmPolicy.Condition.MetricUnit": {
				TargetField: "alarm_policy_condition_metric_unit",
			},
			"AlarmPolicy.Condition.ComparisonOperator": {
				TargetField: "alarm_policy_condition_comparison_operator",
			},
			"AlarmPolicy.Condition.Threshold": {
				TargetField: "alarm_policy_condition_threshold",
			},
		},
	}
}

func (s *VestackScalingPolicyService) ReadResourceId(id string) string {
	return id
}

func (s *VestackScalingPolicyService) enableOrDisablePolicyCallback(policyId string, d *schema.ResourceData) ve.Callback {
	param := &map[string]interface{}{
		"ScalingPolicyId": policyId,
	}
	enable := d.Get("active").(bool)
	if enable {
		return ve.Callback{
			Call: ve.SdkCall{
				Action:      "EnableScalingPolicy",
				ConvertMode: ve.RequestConvertIgnore,
				SdkParam:    param,
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					return s.Client.AutoScalingClient.EnableScalingPolicyCommon(call.SdkParam)
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Active"},
					Timeout: d.Timeout(schema.TimeoutUpdate),
				},
			},
		}
	} else {
		return ve.Callback{
			Call: ve.SdkCall{
				Action:      "DisableScalingPolicy",
				ConvertMode: ve.RequestConvertIgnore,
				SdkParam:    param,
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					return s.Client.AutoScalingClient.DisableScalingPolicyCommon(call.SdkParam)
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"InActive"},
					Timeout: d.Timeout(schema.TimeoutUpdate),
				},
			},
		}
	}
}
