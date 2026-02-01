---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_etl_tasks"
sidebar_current: "docs-volcengine-datasource-tls_etl_tasks"
description: |-
  Use this data source to query detailed information of tls etl tasks
---
# volcengine_tls_etl_tasks
Use this data source to query detailed information of tls etl tasks
## Example Usage
```hcl
data "volcengine_tls_etl_tasks" "foo" {

}
```
## Argument Reference
The following arguments are supported:
* `iam_project_name` - (Optional) Specify the IAM project name to query the data processing tasks under the specified IAM project.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `project_id` - (Optional) Specify the log item ID for querying the data processing tasks under the specified log item.
* `project_name` - (Optional) Specify the name of the log item for querying the data processing tasks under the specified log item. Support fuzzy query.
* `source_topic_id` - (Optional) Specify the log topic ID for querying the data processing tasks related to this log topic.
* `source_topic_name` - (Optional) Specify the name of the log topic for querying the data processing tasks related to this log topic. Support fuzzy matching.
* `status` - (Optional) Specify the processing task status for querying data processing tasks in this status.
* `task_id` - (Optional) The ID of the processing task.
* `task_name` - (Optional) The name of the processing task.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `tasks` - Detailed information of the processing task.
    * `create_time` - Processing task creation time.
    * `description` - A simple description of the processing task.
    * `dsl_type` - DSL type, fixed as NORMAL.
    * `enable` - The running status of the processing task.
    * `etl_status` - Task scheduling status.
    * `from_time` - The start time of the data to be processed.
    * `last_enable_time` - Recent startup time.
    * `modify_time` - The most recent modification time of the processing task.
    * `name` - The name of the processing task.
    * `project_id` - Specify the log item ID for querying the data processing tasks under the specified log item.
    * `project_name` - Specify the name of the log item for querying the data processing tasks under the specified log item. Support fuzzy query.
    * `script` - Processing rules.
    * `source_topic_id` - The log topic ID where the log to be processed is located.
    * `source_topic_name` - The name of the log topic where the log to be processed is located.
    * `target_resources` - Output the relevant information of the target.
        * `alias` - Customize the name of the output target, which needs to be used to refer to the output target in the data processing rules.
        * `project_id` - The log item ID used for storing the processed logs.
        * `project_name` - The name of the log item used for storing the processed logs.
        * `topic_id` - Log topics used for storing processed logs.
        * `topic_name` - The name of the log topic used for storing the processed logs.
    * `task_id` - The ID of the processing task.
    * `task_type` - The task type is fixed as Resident.
    * `to_time` - The end time of the data to be processed.
* `total_count` - The total count of query.


