---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_schedule_sql_tasks"
sidebar_current: "docs-volcengine-datasource-tls_schedule_sql_tasks"
description: |-
  Use this data source to query detailed information of tls schedule sql tasks
---
# volcengine_tls_schedule_sql_tasks
Use this data source to query detailed information of tls schedule sql tasks
## Example Usage
```hcl
data "volcengine_tls_schedule_sql_tasks" "default" {
}
```
## Argument Reference
The following arguments are supported:
* `iam_project_name` - (Optional) IAM log project name.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `project_id` - (Optional) The log project ID to which the source log topic belongs.
* `project_name` - (Optional) The name of the log item to which the source log topic belongs.
* `source_topic_name` - (Optional) Source log topic name.
* `status` - (Optional) Timed SQL analysis task status.
* `task_id` - (Optional) Timed SQL analysis task ID.
* `task_name` - (Optional) Timed SQL analysis task name.
* `topic_id` - (Optional) Source log topic ID.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `tasks` - The List of timed SQL analysis tasks.
    * `create_time_stamp` - Set the creation time of timed SQL analysis tasks.
    * `description` - A simple description of the timed SQL analysis task.
    * `dest_project_id` - The log project ID to which the target log topic belongs.
    * `dest_region` - The region to which the target log project belongs.
    * `dest_topic_id` - The target log topic ID used for storing the result data of timed SQL analysis.
    * `dest_topic_name` - The name of the target log topic used for storing the data of the timed SQL analysis results.
    * `modify_time_stamp` - The most recent modification time of the scheduled SQL analysis task.
    * `process_end_time` - Schedule the end time of the timed SQL analysis task in the format of a second-level timestamp.
    * `process_sql_delay` - The delay time of each scheduling. The value range is from 0 to 120, and the unit is seconds.
    * `process_start_time` - The start time of the scheduled SQL task, that is, the start time when the first instance is scheduled. The format is a timestamp at the second level.
    * `process_time_window` - SQL time window, which refers to the time range for log retrieval and analysis when a timed SQL analysis task is running, is in a left-closed and right-open format.
    * `query` - Timed SQL analysis tasks are retrieval and analysis statements that are executed regularly.
    * `request_cycle` - The scheduling cycle of timed SQL analysis tasks.
        * `cron_tab` - Cron expression. The log service specifies the timed execution of alarm tasks through the Cron expression. The minimum granularity of Cron expressions is minutes, 24 hours. For example, 0 18 * * * indicates that an alarm task is executed exactly at 18:00 every day.
        * `cron_time_zone` - When setting the Type to Cron, the time zone also needs to be set.
        * `time` - The scheduling cycle or the time point of regular execution (the number of minutes away from 00:00), with a value range of 1 to 1440, and the unit is minutes.
        * `type` - The type of Scheduling cycle.
    * `source_project_id` - The log project ID to which the source log topic belongs.
    * `source_project_name` - The name of the log item to which the source log topic belongs.
    * `source_topic_id` - The source log topic ID where the original log for timed SQL analysis is located.
    * `source_topic_name` - The name of the source log topic where the original log for timed SQL analysis is located.
    * `status` - Whether to start the scheduled SQL analysis task immediately after completing the task configuration.
    * `task_id` - Timed SQL analysis task ID.
    * `task_name` - Timed SQL analysis task name.
* `total_count` - The total count of query.


