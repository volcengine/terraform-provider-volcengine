---
subcategory: "ECS"
layout: "volcengine"
page_title: "Volcengine: volcengine_ecs_commands"
sidebar_current: "docs-volcengine-datasource-ecs_commands"
description: |-
  Use this data source to query detailed information of ecs commands
---
# volcengine_ecs_commands
Use this data source to query detailed information of ecs commands
## Example Usage
```hcl
data "volcengine_ecs_commands" "default" {
  command_id = "cmd-ychkepkhtim0tr3b****"
}
```
## Argument Reference
The following arguments are supported:
* `command_id` - (Optional) The id of ecs command.
* `command_provider` - (Optional) The provider of public command. When this field is not specified, query for custom commands.
* `name_regex` - (Optional) A Name Regex of Resource.
* `name` - (Optional) The name of ecs command. This field support fuzzy query.
* `order` - (Optional) The order of ecs command query result.
* `output_file` - (Optional) File name where to save data source results.
* `type` - (Optional) The type of ecs command. Valid values: `Shell`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `commands` - The collection of query.
    * `command_content` - The base64 encoded content of the ecs command.
    * `command_id` - The id of the ecs command.
    * `command_provider` - The provider of the public command.
    * `created_at` - The create time of the ecs command.
    * `description` - The description of the ecs command.
    * `id` - The id of the ecs command.
    * `invocation_times` - The invocation times of the ecs command. Public commands do not display the invocation times.
    * `name` - The name of the ecs command.
    * `timeout` - The timeout of the ecs command.
    * `type` - The type of the ecs command.
    * `updated_at` - The update time of the ecs command.
    * `username` - The username of the ecs command.
    * `working_dir` - The working directory of the ecs command.
* `total_count` - The total count of query.


