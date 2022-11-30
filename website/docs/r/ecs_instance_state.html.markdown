---
subcategory: "ECS"
layout: "volcengine"
page_title: "Volcengine: volcengine_ecs_instance_state"
sidebar_current: "docs-volcengine-resource-ecs_instance_state"
description: |-
  Provides a resource to manage ecs instance state
---
# volcengine_ecs_instance_state
Provides a resource to manage ecs instance state
## Example Usage
```hcl
resource "volcengine_ecs_instance_state" "foo" {
  instance_id = "i-l8u2ai4j0fauo6mrpgk8"
  action      = "Start"
}
```
## Argument Reference
The following arguments are supported:
* `action` - (Required) Start or Stop of Instance Action, the value can be `Start` or `Stop`.
* `instance_id` - (Required, ForceNew) Id of Instance.
* `stopped_mode` - (Optional) Stop Mode of Instance, the value can be `KeepCharging` or `StopCharging`, default `KeepCharging`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `status` - Status of Instance.


## Import
State Instance can be imported using the id, e.g.
```
$ terraform import volcengine_ecs_instance_state.default state:i-mizl7m1kqccg5smt1bdpijuj
```

