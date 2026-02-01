---
subcategory: "ECS"
layout: "volcengine"
page_title: "Volcengine: volcengine_ecs_invocation"
sidebar_current: "docs-volcengine-resource-ecs_invocation"
description: |-
  Provides a resource to manage ecs invocation
---
# volcengine_ecs_invocation
Provides a resource to manage ecs invocation
## Example Usage
```hcl
resource "volcengine_ecs_invocation" "foo" {
  command_id             = "cmd-ychkepkhtim0tr3b****"
  instance_ids           = ["i-ychmz92487l8j00o****"]
  invocation_name        = "tf-test"
  invocation_description = "tf"
  username               = "root"
  timeout                = 90
  working_dir            = "/home"
  repeat_mode            = "Rate"
  frequency              = "5m"
  launch_time            = "2023-06-20T09:48:00Z"
  recurrence_end_time    = "2023-06-20T09:59:00Z"
  project_name           = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
  parameters {
    name  = "test_str"
    value = "tf"
  }
}
```
## Argument Reference
The following arguments are supported:
* `command_id` - (Required, ForceNew) The command id of the ecs invocation.
* `instance_ids` - (Required, ForceNew) The list of ECS instance IDs.
* `invocation_name` - (Required, ForceNew) The name of the ecs invocation.
* `username` - (Required, ForceNew) The username of the ecs command. When this field is not specified, use the value of the field with the same name in ecs command as the default value.
* `frequency` - (Optional, ForceNew) The frequency of the ecs invocation. This field is valid and required when the value of the repeat_mode field is `Rate`.
* `invocation_description` - (Optional, ForceNew) The description of the ecs invocation.
* `launch_time` - (Optional, ForceNew) The launch time of the ecs invocation. RFC3339 format. This field is valid and required when the value of the repeat_mode field is `Rate` or `Fixed`.
* `parameters` - (Optional, ForceNew) The custom parameters of the ecs command. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `project_name` - (Optional) The project name of the ecs command.
* `recurrence_end_time` - (Optional, ForceNew) The recurrence end time of the ecs invocation. RFC3339 format. This field is valid and required when the value of the repeat_mode field is `Rate`.
* `repeat_mode` - (Optional, ForceNew) The repeat mode of the ecs invocation. Valid values: `Once`, `Rate`, `Fixed`. Default is `Once`.
* `tags` - (Optional) Tags.
* `timeout` - (Optional, ForceNew) The timeout of the ecs command. Unit: seconds. Valid value range: 30~86400. Default is 60.
* `working_dir` - (Optional, ForceNew) The working directory of the ecs invocation. When this field is not specified, use the value of the field with the same name in ecs command as the default value.

The `parameters` object supports the following:

* `name` - (Required, ForceNew) The name of the parameter.
* `value` - (Required, ForceNew) The value of the parameter.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `end_time` - The end time of the ecs invocation.
* `invocation_status` - The status of the ecs invocation.
* `start_time` - The start time of the ecs invocation.


## Import
EcsInvocation can be imported using the id, e.g.
```
$ terraform import volcengine_ecs_invocation.default ivk-ychnxnm45dl8j0mm****
```

