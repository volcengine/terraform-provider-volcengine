---
subcategory: "CLOUD_MONITOR"
layout: "volcengine"
page_title: "Volcengine: volcengine_cloud_monitor_rule"
sidebar_current: "docs-volcengine-resource-cloud_monitor_rule"
description: |-
  Provides a resource to manage cloud monitor rule
---
# volcengine_cloud_monitor_rule
Provides a resource to manage cloud monitor rule
## Example Usage
```hcl
resource "volcengine_cloud_monitor_rule" "foo" {
  rule_name        = "acc-test-rule"
  description      = "acc-test"
  namespace        = "VCM_ECS"
  sub_namespace    = "Storage"
  level            = "warning"
  enable_state     = "disable"
  evaluation_count = 5
  effect_start_at  = "00:15"
  effect_end_at    = "22:55"
  silence_time     = 5
  alert_methods    = ["Email", "Webhook"]
  #  web_hook = "http://alert.volc.com/callback"
  webhook_ids         = ["187655704106731****", "187655712542447****"]
  contact_group_ids   = ["174284616403161****"]
  multiple_conditions = true
  condition_operator  = "||"
  regions             = ["cn-beijing", "cn-shanghai"]
  original_dimensions {
    key   = "ResourceID"
    value = ["*"]
  }
  original_dimensions {
    key   = "DiskName"
    value = ["vda", "vda1"]
  }
  conditions {
    metric_name         = "DiskUsageAvail"
    metric_unit         = "Megabytes"
    statistics          = "avg"
    comparison_operator = ">"
    threshold           = "100"
  }
  conditions {
    metric_name         = "DiskUsageUtilization"
    metric_unit         = "Percent"
    statistics          = "avg"
    comparison_operator = ">"
    threshold           = "90"
  }
  recovery_notify {
    enable = true
  }
}
```
## Argument Reference
The following arguments are supported:
* `alert_methods` - (Required) The alert methods of the cloud monitor rule. Valid values: `Email`, `Phone`, `SMS`, `Webhook`.
* `conditions` - (Required) The conditions of the cloud monitor rule.
* `effect_end_at` - (Required) The effect end time of the cloud monitor rule. The expression is `HH:MM`.
* `effect_start_at` - (Required) The effect start time of the cloud monitor rule. The expression is `HH:MM`.
* `enable_state` - (Required) The enable state of the cloud monitor rule. Valid values: `enable`, `disable`.
* `evaluation_count` - (Required) The evaluation count of the cloud monitor rule.
* `level` - (Required) The level of the cloud monitor rule. Valid values: `critical`, `warning`, `notice`.
* `namespace` - (Required, ForceNew) The namespace of the cloud monitor rule.
* `original_dimensions` - (Required) The original dimensions of the cloud monitor rule.
* `regions` - (Required, ForceNew) The region ids of the cloud monitor rule.
* `rule_name` - (Required) The name of the cloud monitor rule.
* `silence_time` - (Required) The silence time of the cloud monitor rule. Unit in minutes. Valid values: 5, 30, 60, 180, 360, 720, 1440.
* `sub_namespace` - (Required, ForceNew) The sub namespace of the cloud monitor rule.
* `condition_operator` - (Optional) The condition operator of the cloud monitor rule. Valid values: `&&`, `||`.
* `contact_group_ids` - (Optional) The contact group ids of the cloud monitor rule. When the alert method is `Email`, `SMS`, or `Phone`, This field must be specified.
* `description` - (Optional) The description of the cloud monitor rule.
* `multiple_conditions` - (Optional) Whether to enable the multiple conditions function of the cloud monitor rule.
* `recovery_notify` - (Optional) The recovery notify of the cloud monitor rule.
* `web_hook` - (Optional) The web hook of the cloud monitor rule. When the alert method is `Webhook`, one of `web_hook` and `webhook_ids` must be specified.
* `webhook_ids` - (Optional) The web hook id list of the cloud monitor rule. When the alert method is `Webhook`, one of `web_hook` and `webhook_ids` must be specified.

The `conditions` object supports the following:

* `comparison_operator` - (Required) The comparison operation of the cloud monitor rule. Valid values: `>`, `>=`, `<`, `<=`, `!=`, `=`.
* `metric_name` - (Required) The metric name of the cloud monitor rule.
* `metric_unit` - (Required) The metric unit of the cloud monitor rule.
* `statistics` - (Required) The statistics of the cloud monitor rule. Valid values: `avg`, `max`, `min`.
* `threshold` - (Required) The threshold of the cloud monitor rule.

The `original_dimensions` object supports the following:

* `key` - (Required) The key of the dimension.
* `value` - (Required) The value of the dimension.

The `recovery_notify` object supports the following:

* `enable` - (Optional) Whether to enable the recovery notify function.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `alert_state` - The alert state of the cloud monitor rule.
* `created_at` - The created time of the cloud monitor rule.
* `updated_at` - The updated time of the cloud monitor rule.


## Import
CloudMonitorRule can be imported using the id, e.g.
```
$ terraform import volcengine_cloud_monitor_rule.default 174284623567451****
```

