resource "vestack_scaling_policy" "foo" {
  active = false
  scaling_group_id = "scg-ybqm0b6kcigh9zu9ce6t"
  scaling_policy_name = "tf-test"
  scaling_policy_type = "Alarm"
  adjustment_type = "QuantityChangeInCapacity"
  adjustment_value = 100
  cooldown = 10
  scheduled_policy_launch_time = "2022-07-09T09:59Z"
  scheduled_policy_recurrence_end_time = "2022-07-24T09:25Z"
  scheduled_policy_recurrence_type = "Daily"
  scheduled_policy_recurrence_value = 10
  alarm_policy_rule_type = "Static"
  alarm_policy_evaluation_count = 1
  alarm_policy_condition_metric_name = "Instance_CpuBusy_Avg"
  alarm_policy_condition_metric_unit = "Percent"
  alarm_policy_condition_comparison_operator = "="
  alarm_policy_condition_threshold = 100
}

