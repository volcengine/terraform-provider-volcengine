---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_alarms"
sidebar_current: "docs-volcengine-datasource-tls_alarms"
description: |-
  Use this data source to query detailed information of tls alarms
---
# volcengine_tls_alarms
Use this data source to query detailed information of tls alarms
## Example Usage
```hcl
data "volcengine_tls_alarms" "default" {
}
```
## Argument Reference
The following arguments are supported:
* `alarm_id` - (Optional) The alarm id.
* `alarm_name` - (Optional) The alarm name.
* `output_file` - (Optional) File name where to save data source results.
* `project_id` - (Optional) The project id.
* `status` - (Optional) The status.
* `topic_id` - (Optional) The topic id.
* `topic_name` - (Optional) The topic name.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `alarms` - The list of alarms.
    * `alarm_id` - The alarm id.
    * `alarm_name` - The name of the alarm.
    * `alarm_notify_group` - List of notification groups corresponding to the alarm.
        * `alarm_notify_group_id` - The id of the notify group.
        * `alarm_notify_group_name` - Name of the notification group.
        * `create_time` - The create time the notification.
        * `iam_project_name` - The iam project name.
        * `modify_time` - The modification time the notification.
        * `notify_type` - The notify group type.
        * `receivers` - List of IAM users to receive alerts.
            * `end_time` - The end time.
            * `receiver_channels` - The list of the receiver channels.
            * `receiver_names` - List of the receiver names.
            * `receiver_type` - The receiver type.
            * `start_time` - The start time.
    * `alarm_period_detail` - Period for sending alarm notifications. When the number of continuous alarm triggers reaches the specified limit (TriggerPeriod), Log Service will send alarm notifications according to the specified period.
        * `email` - Email alarm period, the unit is minutes, and the value range is 1~1440.
        * `general_webhook` - Customize the webhook alarm period, the unit is minutes, and the value range is 1~1440.
        * `phone` - Telephone alarm cycle, the unit is minutes, and the value range is 10~1440.
        * `sms` - SMS alarm cycle, the unit is minutes, and the value range is 10~1440.
    * `alarm_period` - Period for sending alarm notifications. When the number of continuous alarm triggers reaches the specified limit (TriggerPeriod), Log Service will send alarm notifications according to the specified period.
    * `condition` - Alarm trigger condition.
    * `create_time` - The create time.
    * `join_configurations` - The list of join configurations.
        * `condition` - The condition.
        * `set_operation_type` - The set operation type.
    * `modify_time` - The modify time.
    * `project_id` - The project id.
    * `query_request` - Search and analyze sentences, 1~3 can be configured.
        * `end_time_offset_unit` - The end time offset unit.
        * `end_time_offset` - The end time of the query range is relative to the current historical time. The unit is minutes. The value is not positive and must be greater than StartTimeOffset. The maximum value is 0 and the minimum value is -1440.
        * `number` - Alarm object sequence number; increments from 1.
        * `query` - Query statement, the maximum supported length is 1024.
        * `start_time_offset_unit` - The start time offset unit.
        * `start_time_offset` - The start time of the query range is relative to the current historical time, in minutes. The value is non-positive, the maximum value is 0, and the minimum value is -1440.
        * `time_span_type` - The time span type.
        * `topic_id` - The id of the topic.
        * `topic_name` - The name of the topic.
        * `truncated_time` - The truncated time.
    * `request_cycle` - The execution period of the alarm task.
        * `cron_tab` - The cron tab.
        * `time` - The cycle of alarm task execution, or the time point of periodic execution. The unit is minutes, and the value range is 1~1440.
        * `type` - Execution cycle type.
    * `send_resolved` - Whether to send resolved.
    * `severity` - The severity of the alarm.
    * `status` - Whether to enable the alert policy. The default value is true, that is, on.
    * `trigger_conditions` - The list of trigger conditions.
        * `condition` - The condition.
        * `count_condition` - The count condition.
        * `no_data` - The no data.
        * `severity` - The severity.
    * `trigger_period` - Continuous cycle. The alarm will be issued after the trigger condition is continuously met for TriggerPeriod periods; the minimum value is 1, the maximum value is 10, and the default value is 1.
    * `user_define_msg` - Customize the alarm notification content.
* `total_count` - The total count of query.


