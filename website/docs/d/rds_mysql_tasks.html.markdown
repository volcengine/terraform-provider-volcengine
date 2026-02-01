---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_mysql_tasks"
sidebar_current: "docs-volcengine-datasource-rds_mysql_tasks"
description: |-
  Use this data source to query detailed information of rds mysql tasks
---
# volcengine_rds_mysql_tasks
Use this data source to query detailed information of rds mysql tasks
## Example Usage
```hcl
data "volcengine_rds_mysql_tasks" "foo" {
  instance_id         = "mysql-b51d37110dd1"
  creation_start_time = "2025-06-21T00:00:00Z"
  creation_end_time   = "2025-06-23T00:00:00Z"
}
```
## Argument Reference
The following arguments are supported:
* `creation_end_time` - (Optional) The end time of the task. The time format is yyyy-MM-ddTHH:mm:ssZ (UTC time). Instructions: For the two groups of parameters, task time (CreationStartTime and CreationEndTime) and TaskId, one of them must be selected. The maximum time interval between the task start time (CreationStartTime) and the task end time (CreationEndTime) shall not exceed 7 days.
* `creation_start_time` - (Optional) The start time of the task. The time format is yyyy-MM-ddTHH:mm:ssZ (UTC time). Instructions: For the two groups of parameters, task time (CreationStartTime and CreationEndTime) and TaskId, one of them must be selected. The maximum time interval between the task start time (CreationStartTime) and the task end time (CreationEndTime) cannot exceed 7 days.
* `instance_id` - (Optional) Instance ID.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name.
* `task_action` - (Optional) Task name.
* `task_category` - (Optional) Task Category. The values are as shown in the following list, and multiple values can be selected: BackupRecoveryManagement: Backup and Recovery Management. DatabaseAdminManagement: Database Administration Management. DatabaseProxy: Database Proxy. HighAvailability: High Availability. InstanceAttribute: Instance Attribute. InstanceManagement: Instance Management. NetworkManagement: Network Management. SecurityManagement: Security Management. SystemMaintainManagement: System Operation and Maintenance Management. VersionUpgrade: Version Upgrade.
* `task_id` - (Optional) Task ID. Description: For the two groups of parameters, TaskId and task time (CreationStartTime and CreationEndTime), one of them must be selected.
* `task_source` - (Optional) Task source. Values: User: Tenant. System: System. SystemUser: Internal operation and maintenance. UserMaintain: Maintenance operations initiated by system/operation and maintenance administrators and visible to tenants.
* `task_status` - (Optional) Task status. The values are as shown in the following list, and multiple values can be selected: WaitSwitch: Waiting for switching. WaitStart: Waiting for execution. Canceled: Canceled. Stopped: Terminated. Running_BeforeSwitch: Running (before switching). Timeout: Execution Timeout. Success: Execution Success. Failed: Execution Failed. Running: In Execution. Stopping: In Termination.
* `task_type` - (Optional) Task type. Values: Web: Console request. OpenAPI: OpenAPI request. AssumeRole: Role - playing request. Other: Other requests.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `datas` - The collection of query.
    * `create_time` - The creation time of the task.
    * `finish_time` - The completion time of the task.
    * `progress` - Task progress. The unit is percentage. Description: Only tasks with a task status of In Progress, that is, tasks with a TaskStatus value of Running, will return the task progress.
    * `scheduled_execute_end_time` - The deadline for the planned startup. The time format is yyyy-MM-ddTHH:mm:ssZ (UTC time). Description: This field will only be returned for tasks in the "Waiting to Start", "Waiting to Execute", or "Waiting to Switch" states.
    * `scheduled_switch_end_time` - The scheduled end time for the switch. The time format is yyyy-MM-ddTHH:mm:ssZ (UTC time). Description: This field will only be returned for tasks in the "Waiting to Start", "Waiting to Execute", or "Waiting to Switch" states.
    * `scheduled_switch_start_time` - The start time of the scheduled switch. The time format is yyyy-MM-ddTHH:mm:ssZ (UTC time). Description: This field is returned only for tasks in the "Waiting to Start", "Waiting to Execute", or "Waiting to Switch" state.
    * `start_time` - The start time of the task.
    * `task_action` - Task name.
    * `task_category` - Task category.
    * `task_desc` - The description of the task.
    * `task_detail` - Detailed information of the task.
        * `check_item_log` - The log of inspection items for the instance major version upgrade.
        * `check_items` - Check results for major version upgrade.
            * `check_detail` - Details of the failed check items.
                * `impact` - The impact of the issue that caused the failure of the check item after the upgrade.
                * `issue` - Problems that caused the failure to pass the check items.
            * `description` - The description of the check item.
            * `item_name` - The name of the check item.
            * `risk_level` - The risk level of the failed check items. Values:
Notice: Attention.
Warning: Warning.
Error: Error.
        * `task_info` - Details of the task.
            * `create_time` - The creation time of the task.
            * `finish_time` - The completion time of the task.
            * `progress` - Task progress. The unit is percentage. Description: Only tasks with a task status of In Progress, that is, tasks with a TaskStatus value of Running, will return the task progress.
            * `related_instance_infos` - Instances related to the task.
                * `instance_id` - Instance ID.
    * `task_id` - Task ID.
    * `task_params` - Task parameters.
    * `task_progress` - Progress details.
        * `name` - Step Name. Values:
InstanceInitialization: Task initialization.
InstanceRecoveryPreparation Instance recovery preparation.
DataBackupImport: Cold backup import.
LogBackupBinlogAdd: Binlog playback.
TaskSuccessful: Task success.
        * `step_extra_info` - Specific information of the step.
            * `type` - Current stage. CostTime: The time taken for the current stage.
CurDataSize: The amount of data imported currently.
CurBinlog: The number of Binlog files being replayed currently.
RemainCostTime: The remaining time taken.
RemainDataSize: The remaining amount of data to be imported. RemainBinlog: The number of Binlog files remaining for playback.
            * `unit` - Unit. Values:
MS: Milliseconds.
Bytes: Bytes.
Files: Number of (files).
            * `value` - The specific value corresponding to the Type field.
        * `step_status` - Step status. Values:
Running: In progress.
Success: Successful.
Failed: Failed.
Unexecuted: Not executed.
    * `task_status` - Task status. The values are as shown in the following list, and multiple values can be selected: WaitSwitch: Waiting for switching. WaitStart: Waiting for execution. Canceled: Canceled. Stopped: Terminated. Running_BeforeSwitch: Running (before switching). Timeout: Execution Timeout. Success: Execution Success. Failed: Execution Failed. Running: In Execution. Stopping: In Termination.
* `total_count` - The total count of query.


