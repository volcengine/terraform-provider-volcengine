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
  alarm_name = "test"
  project_id = "cc44f8b6-0328-4622-b043-023fca735cd4"
  //status = true
  //trigger_period = 1
  //alarm_period = 10
  alarm_notify_group = ["3019107f-28a2-4208-a2b6-c33fcb97ac3a"]
  user_define_msg    = "test for terraform"
  query_request {
    number            = 1
    topic_id          = "af1a2240-ba62-4f18-b421-bde2f9684e57"
    start_time_offset = -15
    query             = "Failed | select count(*) as errNum"
    end_time_offset   = 0
  }
  request_cycle {
    type = "Period"
    time = 11
  }
  alarm_period_detail {
    sms             = 10
    phone           = 10
    email           = 2
    general_webhook = 3
  }
  condition = "$1.errNum>0"
}
```
## Argument Reference
The following arguments are supported:
* `alarm_name` - (Required) The name of the alarm.
* `alarm_notify_group` - (Required, ForceNew) List of notification groups corresponding to the alarm.
* `condition` - (Required) Alarm trigger condition.
* `project_id` - (Required, ForceNew) The project id.
* `query_request` - (Required) Search and analyze sentences, 1~3 can be configured.
* `request_cycle` - (Required) The execution period of the alarm task.
* `alarm_period_detail` - (Optional) Period for sending alarm notifications. When the number of continuous alarm triggers reaches the specified limit (TriggerPeriod), Log Service will send alarm notifications according to the specified period.
* `alarm_period` - (Optional) Period for sending alarm notifications. When the number of continuous alarm triggers reaches the specified limit (TriggerPeriod), Log Service will send alarm notifications according to the specified period.
* `status` - (Optional) Whether to enable the alert policy. The default value is true, that is, on.
* `trigger_period` - (Optional) Continuous cycle. The alarm will be issued after the trigger condition is continuously met for TriggerPeriod periods; the minimum value is 1, the maximum value is 10, and the default value is 1.
* `user_define_msg` - (Optional) Customize the alarm notification content.

The `alarm_period_detail` object supports the following:

* `email` - (Required) Email alarm period, the unit is minutes, and the value range is 1~1440.
* `general_webhook` - (Required) Customize the webhook alarm period, the unit is minutes, and the value range is 1~1440.
* `phone` - (Required) Telephone alarm cycle, the unit is minutes, and the value range is 10~1440.
* `sms` - (Required) SMS alarm cycle, the unit is minutes, and the value range is 10~1440.

The `query_request` object supports the following:

* `end_time_offset` - (Required) The end time of the query range is relative to the current historical time. The unit is minutes. The value is not positive and must be greater than StartTimeOffset. The maximum value is 0 and the minimum value is -1440.
* `number` - (Required) Alarm object sequence number; increments from 1.
* `query` - (Required) Query statement, the maximum supported length is 1024.
* `start_time_offset` - (Required) The start time of the query range is relative to the current historical time, in minutes. The value is non-positive, the maximum value is 0, and the minimum value is -1440.
* `topic_id` - (Required) The id of the topic.

The `request_cycle` object supports the following:

* `time` - (Required) The cycle of alarm task execution, or the time point of periodic execution. The unit is minutes, and the value range is 1~1440.
* `type` - (Required) Execution cycle type.
Period: Periodic execution, which means executing once every certain period of time.
Fixed: Regular execution, which means executing at a fixed time point every day.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `alarm_id` - The alarm id.


## Import
tls alarm can be imported using the id and project id, e.g.
```
$ terraform import volcengine_tls_alarm.default projectId:fc************
```

