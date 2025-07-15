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
* `project_name` - (Optional) The project name of ecs command.
* `tags` - (Optional) Tags.
* `type` - (Optional) The type of ecs command. Valid values: `Shell`.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `commands` - The collection of query.
    * `command_content` - The base64 encoded content of the ecs command.
    * `command_id` - The id of the ecs command.
    * `command_provider` - The provider of the public command.
    * `content_encoding` - Whether the command content is base64 encoded. Valid values: `Base64`, `PlainText`. Default is `Base64`.
    * `created_at` - The create time of the ecs command.
    * `description` - The description of the ecs command.
    * `enable_parameter` - Whether to enable custom parameter. Default is `false`.
    * `id` - The id of the ecs command.
    * `invocation_times` - The invocation times of the ecs command. Public commands do not display the invocation times.
    * `name` - The name of the ecs command.
    * `parameter_definitions` - The custom parameter definitions of the ecs command.
        * `decimal_precision` - The decimal precision of the custom parameter. This field is required when the parameter type is `Digit`.
        * `default_value` - The default value of the custom parameter.
        * `max_length` - The maximum length of the custom parameter. This field is required when the parameter type is `String`.
        * `max_value` - The maximum value of the custom parameter. This field is required when the parameter type is `Digit`.
        * `min_length` - The minimum length of the custom parameter. This field is required when the parameter type is `String`.
        * `min_value` - The minimum value of the custom parameter. This field is required when the parameter type is `Digit`.
        * `name` - The name of the custom parameter.
        * `required` - Whether the custom parameter is required.
        * `type` - The type of the custom parameter. Valid values: `String`, `Digit`.
    * `project_name` - The project name of the ecs command.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `timeout` - The timeout of the ecs command.
    * `type` - The type of the ecs command.
    * `updated_at` - The update time of the ecs command.
    * `username` - The username of the ecs command.
    * `working_dir` - The working directory of the ecs command.
* `total_count` - The total count of query.


