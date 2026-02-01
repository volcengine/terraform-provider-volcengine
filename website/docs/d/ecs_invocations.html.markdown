---
subcategory: "ECS"
layout: "volcengine"
page_title: "Volcengine: volcengine_ecs_invocations"
sidebar_current: "docs-volcengine-datasource-ecs_invocations"
description: |-
  Use this data source to query detailed information of ecs invocations
---
# volcengine_ecs_invocations
Use this data source to query detailed information of ecs invocations
## Example Usage
```hcl
data "volcengine_ecs_invocations" "default" {
  invocation_id     = "ivk-ych9y4vujvl8j01c****"
  invocation_status = ["Success"]
}
```
## Argument Reference
The following arguments are supported:
* `command_id` - (Optional) The id of ecs command.
* `command_name` - (Optional) The name of ecs command. This field support fuzzy query.
* `command_type` - (Optional) The type of ecs command. Valid values: `Shell`.
* `invocation_id` - (Optional) The id of ecs invocation.
* `invocation_name` - (Optional) The name of ecs invocation. This field support fuzzy query.
* `invocation_status` - (Optional) The list of status of ecs invocation. Valid values: `Pending`, `Scheduled`, `Running`, `Success`, `Failed`, `Stopped`, `PartialFailed`, `Finished`.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name of ecs invocation.
* `repeat_mode` - (Optional) The repeat mode of ecs invocation. Valid values: `Once`, `Rate`, `Fixed`.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `invocations` - The collection of query.
    * `command_content` - The base64 encoded content of the ecs command.
    * `command_description` - The description of the ecs command.
    * `command_id` - The id of the ecs command.
    * `command_name` - The name of the ecs command.
    * `command_provider` - The provider of the ecs command.
    * `command_type` - The type of the ecs command.
    * `end_time` - The end time of the ecs invocation.
    * `frequency` - The frequency of the ecs invocation.
    * `id` - The id of the ecs invocation.
    * `instance_ids` - The list of ECS instance IDs.
    * `instance_number` - The instance number of the ecs invocation.
    * `invocation_description` - The description of the ecs invocation.
    * `invocation_id` - The id of the ecs invocation.
    * `invocation_name` - The name of the ecs invocation.
    * `invocation_status` - The status of the ecs invocation.
    * `launch_time` - The launch time of the ecs invocation.
    * `parameters` - The custom parameters of the ecs invocation.
    * `project_name` - The project name of the ecs invocation.
    * `recurrence_end_time` - The recurrence end time of the ecs invocation.
    * `repeat_mode` - The repeat mode of the ecs invocation.
    * `start_time` - The start time of the ecs invocation.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `timeout` - The timeout of the ecs command.
    * `username` - The username of the ecs command.
    * `working_dir` - The working directory of the ecs command.
* `total_count` - The total count of query.


