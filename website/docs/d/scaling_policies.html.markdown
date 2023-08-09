---
subcategory: "AUTOSCALING"
layout: "volcengine"
page_title: "Volcengine: volcengine_scaling_policies"
sidebar_current: "docs-volcengine-datasource-scaling_policies"
description: |-
  Use this data source to query detailed information of scaling policies
---
# volcengine_scaling_policies
Use this data source to query detailed information of scaling policies
## Example Usage
```hcl
data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_scaling_group" "foo" {
  scaling_group_name        = "acc-test-scaling-group"
  subnet_ids                = [volcengine_subnet.foo.id]
  multi_az_policy           = "BALANCE"
  desire_instance_number    = 0
  min_instance_number       = 0
  max_instance_number       = 1
  instance_terminate_policy = "OldestInstance"
  default_cooldown          = 10
}

resource "volcengine_scaling_policy" "foo" {
  count                                      = 3
  active                                     = false
  scaling_group_id                           = volcengine_scaling_group.foo.id
  scaling_policy_name                        = "acc-tf-sg-policy-test-${count.index}"
  scaling_policy_type                        = "Alarm"
  adjustment_type                            = "QuantityChangeInCapacity"
  adjustment_value                           = 100
  cooldown                                   = 10
  alarm_policy_rule_type                     = "Static"
  alarm_policy_evaluation_count              = 1
  alarm_policy_condition_metric_name         = "Instance_CpuBusy_Avg"
  alarm_policy_condition_metric_unit         = "Percent"
  alarm_policy_condition_comparison_operator = "="
  alarm_policy_condition_threshold           = 100
}

data "volcengine_scaling_policies" "foo" {
  ids              = volcengine_scaling_policy.foo[*].scaling_policy_id
  scaling_group_id = volcengine_scaling_group.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `scaling_group_id` - (Required) An id of the scaling group to which the scaling policy belongs.
* `ids` - (Optional) A list of scaling policy ids.
* `name_regex` - (Optional) A Name Regex of scaling policy.
* `output_file` - (Optional) File name where to save data source results.
* `scaling_policy_names` - (Optional) A list of scaling policy names.
* `scaling_policy_type` - (Optional) A type of scaling policy. Valid values: Scheduled, Recurrence, Manual, Alarm.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `scaling_policies` - The collection of scaling policy query.
    * `adjustment_type` - The adjustment type of the scaling policy.
    * `adjustment_value` - The adjustment value of the scaling policy.
    * `alarm_policy_condition_comparison_operator` - The comparison operator of the alarm policy condition of the scaling policy.
    * `alarm_policy_condition_metric_name` - The metric name of the alarm policy condition of the scaling policy.
    * `alarm_policy_condition_metric_unit` - The comparison operator of the alarm policy condition of the scaling policy.
    * `alarm_policy_condition_threshold` - The threshold of the alarm policy condition of the scaling policy.
    * `alarm_policy_evaluation_count` - The evaluation count of the alarm policy of the scaling policy.
    * `alarm_policy_rule_type` - The rule type of the alarm policy of the scaling policy.
    * `cooldown` - The cooldown of the scaling policy.
    * `id` - The id of the scaling policy.
    * `scaling_group_id` - The id of the scaling group to which the scaling policy belongs.
    * `scaling_policy_id` - The id of the scaling policy.
    * `scaling_policy_name` - The name of the scaling policy.
    * `scaling_policy_type` - The type of the scaling policy.
    * `scheduled_policy_launch_time` - The launch time of the scheduled policy of the scaling policy.
    * `scheduled_policy_recurrence_end_time` - The recurrence end time of the scheduled policy of the scaling policy.
    * `scheduled_policy_recurrence_start_time` - The recurrence start time of the scheduled policy of the scaling policy.
    * `scheduled_policy_recurrence_type` - The recurrence type of the scheduled policy of the scaling policy.
    * `scheduled_policy_recurrence_value` - The recurrence value of the scheduled policy of the scaling policy.
    * `status` - The status of the scaling policy.
* `total_count` - The total count of scaling policy query.


