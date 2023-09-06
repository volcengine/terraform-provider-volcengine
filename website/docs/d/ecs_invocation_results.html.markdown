---
subcategory: "ECS"
layout: "volcengine"
page_title: "Volcengine: volcengine_ecs_invocation_results"
sidebar_current: "docs-volcengine-datasource-ecs_invocation_results"
description: |-
  Use this data source to query detailed information of ecs invocation results
---
# volcengine_ecs_invocation_results
Use this data source to query detailed information of ecs invocation results
## Example Usage
```hcl
data "volcengine_ecs_invocation_results" "default" {
  invocation_id            = "ivk-ych9y4vujvl8j01c****"
  invocation_result_status = ["Success"]
}
```
## Argument Reference
The following arguments are supported:
* `invocation_id` - (Required) The id of ecs invocation.
* `command_id` - (Optional) The id of ecs command.
* `instance_id` - (Optional) The id of ecs instance.
* `invocation_result_status` - (Optional) The list of status of ecs invocation in a single instance. Valid values: `Pending`, `Running`, `Success`, `Failed`, `Timeout`.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `invocation_results` - The collection of query.
    * `command_id` - The id of the ecs command.
    * `end_time` - The end time of the ecs invocation in the instance.
    * `error_code` - The error code of the ecs invocation.
    * `error_message` - The error message of the ecs invocation.
    * `exit_code` - The exit code of the ecs command.
    * `id` - The id of the ecs invocation result.
    * `instance_id` - The id of the ecs instance.
    * `invocation_id` - The id of the ecs invocation.
    * `invocation_result_id` - The id of the ecs invocation result.
    * `invocation_result_status` - The status of ecs invocation in a single instance.
    * `output` - The base64 encoded output message of the ecs invocation.
    * `start_time` - The start time of the ecs invocation in the instance.
    * `username` - The username of the ecs command.
* `total_count` - The total count of query.


