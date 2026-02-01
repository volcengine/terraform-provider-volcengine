---
subcategory: "ECS"
layout: "volcengine"
page_title: "Volcengine: volcengine_ecs_command"
sidebar_current: "docs-volcengine-resource-ecs_command"
description: |-
  Provides a resource to manage ecs command
---
# volcengine_ecs_command
Provides a resource to manage ecs command
## Example Usage
```hcl
resource "volcengine_ecs_command" "foo" {
  name            = "tf-test"
  description     = "tf"
  working_dir     = "/home"
  username        = "root"
  timeout         = 100
  command_content = base64encode("#!/bin/bash\n\n\necho \"{{ test_str }} {{ test_num }} operation success!\"")
  project_name    = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
  enable_parameter = true
  parameter_definitions {
    name       = "test_str"
    type       = "String"
    required   = true
    min_length = 1
    max_length = 100
  }
  parameter_definitions {
    name      = "test_num"
    type      = "Digit"
    required  = false
    min_value = "-10"
    max_value = "100"
  }
}
```
## Argument Reference
The following arguments are supported:
* `command_content` - (Required) The base64 encoded content of the ecs command.
* `name` - (Required) The name of the ecs command.
* `description` - (Optional) The description of the ecs command.
* `enable_parameter` - (Optional) Whether to enable custom parameter. Default is `false`.
* `parameter_definitions` - (Optional) The custom parameter definitions of the ecs command.
* `project_name` - (Optional) The project name of the ecs command.
* `tags` - (Optional) Tags.
* `timeout` - (Optional) The timeout of the ecs command. Unit: seconds. Valid value range: 30~86400. Default is 300.
* `type` - (Optional) The type of the ecs command. Valid values: `Shell`, `Python`, `PowerShell`, `Bat`. Default is `Shell`.
* `username` - (Optional) The username of the ecs command.
* `working_dir` - (Optional) The working directory of the ecs command.

The `parameter_definitions` object supports the following:

* `name` - (Required) The name of the custom parameter.
* `type` - (Required) The type of the custom parameter. Valid values: `String`, `Digit`.
* `decimal_precision` - (Optional) The decimal precision of the custom parameter. This field is required when the parameter type is `Digit`.
* `default_value` - (Optional) The default value of the custom parameter.
* `max_length` - (Optional) The maximum length of the custom parameter. This field is required when the parameter type is `String`.
* `max_value` - (Optional) The maximum value of the custom parameter. This field is required when the parameter type is `Digit`.
* `min_length` - (Optional) The minimum length of the custom parameter. This field is required when the parameter type is `String`.
* `min_value` - (Optional) The minimum value of the custom parameter. This field is required when the parameter type is `Digit`.
* `required` - (Optional) Whether the custom parameter is required.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `created_at` - The create time of the ecs command.
* `invocation_times` - The invocation times of the ecs command. Public commands do not display the invocation times.
* `updated_at` - The update time of the ecs command.


## Import
EcsCommand can be imported using the id, e.g.
```
$ terraform import volcengine_ecs_command.default cmd-ychkepkhtim0tr3bcsw1
```

