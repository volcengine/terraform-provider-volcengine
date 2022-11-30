---
subcategory: "AUTOSCALING"
layout: "volcengine"
page_title: "Volcengine: volcengine_scaling_policy"
sidebar_current: "docs-volcengine-resource-scaling_policy"
description: |-
  Provides a resource to manage scaling policy
---
# volcengine_scaling_policy
Provides a resource to manage scaling policy
## Example Usage
```hcl
resource "volcengine_scaling_policy" "foo" {
  active                                     = false
  scaling_group_id                           = "scg-ybqm0b6kcigh9zu9ce6t"
  scaling_policy_name                        = "tf-test"
  scaling_policy_type                        = "Alarm"
  adjustment_type                            = "QuantityChangeInCapacity"
  adjustment_value                           = 100
  cooldown                                   = 10
  scheduled_policy_launch_time               = "2022-07-09T09:59Z"
  scheduled_policy_recurrence_end_time       = "2022-07-24T09:25Z"
  scheduled_policy_recurrence_type           = "Daily"
  scheduled_policy_recurrence_value          = 10
  alarm_policy_rule_type                     = "Static"
  alarm_policy_evaluation_count              = 1
  alarm_policy_condition_metric_name         = "Instance_CpuBusy_Avg"
  alarm_policy_condition_metric_unit         = "Percent"
  alarm_policy_condition_comparison_operator = "="
  alarm_policy_condition_threshold           = 100
}
```
## Argument Reference
The following arguments are supported:
* `adjustment_type` - (Required) The adjustment type of the scaling policy. Valid values: QuantityChangeInCapacity, PercentChangeInCapacity, TotalCapacity.
* `adjustment_value` - (Required) The adjustment value of the scaling policy.
* `scaling_group_id` - (Required, ForceNew) The id of the scaling group to which the scaling policy belongs.
* `scaling_policy_name` - (Required) The name of the scaling policy.
* `scaling_policy_type` - (Required, ForceNew) The type of scaling policy. Valid values: Scheduled, Recurrence, Alarm.
* `active` - (Optional) The active flag of the scaling policy. [Warning] the scaling policy can be active only when the scaling group be active otherwise will fail.
* `alarm_policy_condition_comparison_operator` - (Optional) The comparison operator of the alarm policy condition of the scaling policy. Valid values: >, <, =.
* `alarm_policy_condition_metric_name` - (Optional) The metric name of the alarm policy condition of the scaling policy. Valid values: CpuTotal_Max, CpuTotal_Min, CpuTotal_Avg, MemoryUsedUtilization_Max, MemoryUsedUtilization_Min, MemoryUsedUtilization_Avg, Instance_CpuBusy_Max, Instance_CpuBusy_Min, Instance_CpuBusy_Avg.
* `alarm_policy_condition_metric_unit` - (Optional) The comparison operator of the alarm policy condition of the scaling policy.
* `alarm_policy_condition_threshold` - (Optional) The threshold of the alarm policy condition of the scaling policy.
* `alarm_policy_evaluation_count` - (Optional) The evaluation count of the alarm policy of the scaling policy.
* `alarm_policy_rule_type` - (Optional) The rule type of the alarm policy of the scaling policy. Valid value: Static.
* `cooldown` - (Optional) The cooldown of the scaling policy. Default value is the cooldown time of the scaling group.
* `scheduled_policy_launch_time` - (Optional) The launch time of the scheduled policy of the scaling policy.
* `scheduled_policy_recurrence_end_time` - (Optional) The recurrence end time of the scheduled policy of the scaling policy.
* `scheduled_policy_recurrence_type` - (Optional) The recurrence type the scheduled policy of the scaling policy. Valid values: Daily, Weekly, Monthly, Cron.
* `scheduled_policy_recurrence_value` - (Optional) The recurrence value the scheduled policy of the scaling policy.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `status` - The status of the scaling policy. Valid values: Active, InActive.


## Import
ScalingPolicy can be imported using the ScalingGroupId:ScalingPolicyId, e.g.
```
$ terraform import volcengine_scaling_policy.default scg-yblfbfhy7agh9zn72iaz:sp-yblf9l4fvcl8j1prohsp
```

