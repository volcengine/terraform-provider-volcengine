---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_alarm"
sidebar_current: "docs-volcengine-resource-tls_alarm"
description: |-
  Provides a resource to manage tls alarm
---
# volcengine_tls_alarm
Provides a resource to manage tls alarm
## Example Usage
```hcl
resource "volcengine_tls_alarm" "foo" {
  alarm_name     = "test-terraform-tf"
  project_id     = "88d31abb-62c7-40f5-998e-889747c2a116"
  status         = false
  trigger_period = 2
  #alarm_period = 60
  alarm_notify_group = ["bf3ecf26-2081-4e27-ae18-f44dbe5c6138"]
  user_define_msg    = "test for terraform"

  query_request {
    number                 = 1
    topic_id               = "a690a9b8-72c1-40a3-b8c6-f89a81d3748e"
    start_time_offset      = -15
    end_time_offset        = 0
    query                  = "Failed | select count(*) as errNum"
    time_span_type         = "Relative"
    truncated_time         = "Minute"
    end_time_offset_unit   = "Minute"
    start_time_offset_unit = "Minute"
  }

  request_cycle {
    type = "Period"
    time = 20
    // cron_tab = "0 18 * * *" # If type is Cron
  }

  alarm_period_detail {
    sms             = 20
    phone           = 20
    email           = 20
    general_webhook = 20
  }

  # Condition and Severity are ignored if TriggerConditions is used
  # condition = "$1.errNum>0"
  # severity = "critical"

  trigger_conditions {
    condition       = "$1.errNum>0"
    severity        = "critical"
    count_condition = "__count__ > 0"
    no_data         = false
  }


  send_resolved = true
}
```
## Argument Reference
The following arguments are supported:
* `alarm_name` - (Required) The name of the alarm.
* `alarm_notify_group` - (Required, ForceNew) List of notification groups corresponding to the alarm.
* `project_id` - (Required, ForceNew) The project id.
* `query_request` - (Required) Search and analyze sentences, 1~3 can be configured.
* `request_cycle` - (Required) The execution period of the alarm task.
* `trigger_period` - (Required) Continuous cycle. The alarm will be issued after the trigger condition is continuously met for TriggerPeriod periods; the minimum value is 1, the maximum value is 10, and the default value is 1.
* `alarm_period_detail` - (Optional) Period for sending alarm notifications. When the number of continuous alarm triggers reaches the specified limit (TriggerPeriod), Log Service will send alarm notifications according to the specified period.
* `alarm_period` - (Optional) Period for sending alarm notifications. When the number of continuous alarm triggers reaches the specified limit (TriggerPeriod), Log Service will send alarm notifications according to the specified period.
* `condition` - (Optional) Alarm trigger condition.
* `join_configurations` - (Optional) The list of join configurations.
* `send_resolved` - (Optional) Whether to send resolved.
* `severity` - (Optional) The severity of the alarm.
* `status` - (Optional) Whether to enable the alert policy. The default value is true, that is, on.
* `trigger_conditions` - (Optional) The list of trigger conditions.
* `user_define_msg` - (Optional) Customize the alarm notification content.

The `alarm_period_detail` object supports the following:

* `email` - (Required) Email alarm period, the unit is minutes, and the value range is 1~1440.
* `general_webhook` - (Required) Customize the webhook alarm period, the unit is minutes, and the value range is 1~1440.
* `phone` - (Required) Telephone alarm cycle, the unit is minutes, and the value range is 10~1440.
* `sms` - (Required) SMS alarm cycle, the unit is minutes, and the value range is 10~1440.

The `join_configurations` object supports the following:

* `condition` - (Optional) The condition.
* `set_operation_type` - (Optional) The set operation type.

The `query_request` object supports the following:

* `end_time_offset` - (Required) The end time of the query range is relative to the current historical time. The unit is minutes. The value is not positive and must be greater than StartTimeOffset. The maximum value is 0 and the minimum value is -1440.
* `number` - (Required) Alarm object sequence number; increments from 1.
* `query` - (Required) Query statement, the maximum supported length is 1024.
* `start_time_offset` - (Required) The start time of the query range is relative to the current historical time, in minutes. The value is non-positive, the maximum value is 0, and the minimum value is -1440.
* `topic_id` - (Required) The id of the topic.
* `end_time_offset_unit` - (Optional) The end time offset unit.
* `start_time_offset_unit` - (Optional) The start time offset unit.
* `time_span_type` - (Optional) The time span type.
* `truncated_time` - (Optional) The truncated time.

The `request_cycle` object supports the following:

* `cron_tab` - (Optional) The cron tab.
* `time` - (Optional) The cycle of alarm task execution, or the time point of periodic execution. The unit is minutes, and the value range is 1~1440.
* `type` - (Optional) Execution cycle type.

The `trigger_conditions` object supports the following:

* `condition` - (Optional) The condition.
* `count_condition` - (Optional) The count condition.
* `no_data` - (Optional) The no data.
* `severity` - (Optional) The severity.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `alarm_id` - The alarm id.


## Import
tls alarm can be imported using the id and project id, e.g.
```
$ terraform import volcengine_tls_alarm.default projectId:fc************
```

