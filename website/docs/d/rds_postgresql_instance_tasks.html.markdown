---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_instance_tasks"
sidebar_current: "docs-volcengine-datasource-rds_postgresql_instance_tasks"
description: |-
  Use this data source to query detailed information of rds postgresql instance tasks
---
# volcengine_rds_postgresql_instance_tasks
Use this data source to query detailed information of rds postgresql instance tasks
## Example Usage
```hcl
data "volcengine_rds_postgresql_instance_tasks" "example" {
  # Choose one of TaskId or time window
  # task_id = "202512121649255DCB10D567104F714DDE-660239"

  # Or filter by time window (≤ 7 days)
  instance_id         = "postgres-72715e0d9f58"
  creation_start_time = "2025-12-10T21:30:00Z"
  creation_end_time   = "2025-12-15T23:40:00Z"

  # Optional filters
  task_action  = "ModifyDBEndpointReadWriteFlag"
  task_status  = ["Running", "Success"]
  project_name = "default"
}
```
## Argument Reference
The following arguments are supported:
* `creation_end_time` - (Optional) Task end time. Format: yyyy-MM-ddTHH:mm:ssZ (UTC). Note: The maximum interval between creation_start_time and creation_end_time cannot exceed 7 days.
* `creation_start_time` - (Optional) Task start time. Format: yyyy-MM-ddTHH:mm:ssZ (UTC). Note: One of TaskId or task time (creation_start_time and creation_end_time) must be specified.
* `instance_id` - (Optional) The id of the PostgreSQL instance.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) Project name.
* `task_action` - (Optional) Task action.
* `task_id` - (Optional) Task ID. Note: One of TaskId or task time (creation_start_time and creation_end_time) must be specified.
* `task_status` - (Optional) Task status. Values: Canceled, WaitStart, WaitSwitch, Running, Running_BeforeSwitch, Running_Switching, Running_AfterSwitch, Success, Failed, Timeout, Rollbacking, RollbackFailed, Paused.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `task_infos` - Task list.
    * `cost_time_ms` - Task execution time in milliseconds.
    * `create_time` - Task creation time. Format: yyyy-MM-ddTHH:mm:ssZ (UTC).
    * `finish_time` - Task finish time. Format: yyyy-MM-ddTHH:mm:ssZ (UTC).
    * `instance_id` - Instance ID.
    * `project_name` - Project name.
    * `region` - Region.
    * `scheduled_switch_end_time` - The scheduled end time for the switch. The time format is yyyy-MM-ddTHH:mm:ssZ (UTC time). Note: This field will only be returned for tasks in the "Waiting to Start", "Waiting to Execute", or "Waiting to Switch" states.
    * `scheduled_switch_start_time` - The start time of the scheduled switch. The time format is yyyy-MM-ddTHH:mm:ssZ (UTC time). Note: This field will only be returned for tasks in the "Waiting to Start", "Waiting to Execute", or "Waiting to Switch" states.
    * `task_action` - Task action.
    * `task_id` - Task ID.
    * `task_params` - Task parameters in JSON string.
    * `task_status` - Task status.
* `total_count` - The total count of query.


