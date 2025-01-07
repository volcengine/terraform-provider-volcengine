resource "volcengine_cloud_monitor_rule" "foo" {
  rule_name = "acc-test-rule"
  description = "acc-test"
  namespace = "VCM_ECS"
  sub_namespace = "Storage"
  level = "warning"
  enable_state = "disable"
  evaluation_count = 5
  effect_start_at = "00:15"
  effect_end_at = "22:55"
  silence_time = 5
  alert_methods = ["Email", "Webhook"]
#  web_hook = "http://alert.volc.com/callback"
  webhook_ids = ["187655704106731****", "187655712542447****"]
  contact_group_ids = ["174284616403161****"]
  multiple_conditions = true
  condition_operator = "||"
  regions = ["cn-beijing", "cn-shanghai"]
  original_dimensions {
    key = "ResourceID"
    value = ["*"]
  }
  original_dimensions {
    key = "DiskName"
    value = ["vda", "vda1"]
  }
  conditions {
    metric_name = "DiskUsageAvail"
    metric_unit = "Megabytes"
    statistics = "avg"
    comparison_operator = ">"
    threshold = "100"
  }
  conditions {
    metric_name = "DiskUsageUtilization"
    metric_unit = "Percent"
    statistics = "avg"
    comparison_operator = ">"
    threshold = "90"
  }
  recovery_notify {
    enable = true
  }
}