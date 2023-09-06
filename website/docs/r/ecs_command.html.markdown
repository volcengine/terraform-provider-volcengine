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
  command_content = "IyEvYmluL2Jhc2gKCgplY2hvICJvcGVyYXRpb24gc3VjY2VzcyEi"
}
```
## Argument Reference
The following arguments are supported:
* `command_content` - (Required) The base64 encoded content of the ecs command.
* `name` - (Required) The name of the ecs command.
* `description` - (Optional) The description of the ecs command.
* `timeout` - (Optional) The timeout of the ecs command. Valid value range: 10-600.
* `username` - (Optional) The username of the ecs command.
* `working_dir` - (Optional) The working directory of the ecs command.

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

