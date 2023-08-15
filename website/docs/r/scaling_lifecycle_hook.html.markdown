---
subcategory: "AUTOSCALING"
layout: "volcengine"
page_title: "Volcengine: volcengine_scaling_lifecycle_hook"
sidebar_current: "docs-volcengine-resource-scaling_lifecycle_hook"
description: |-
  Provides a resource to manage scaling lifecycle hook
---
# volcengine_scaling_lifecycle_hook
Provides a resource to manage scaling lifecycle hook
## Example Usage
```hcl
data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_scaling_group" "foo" {
  scaling_group_name        = "acc-test-scaling-group-lifecycle"
  subnet_ids                = [volcengine_subnet.foo.id]
  multi_az_policy           = "BALANCE"
  desire_instance_number    = 0
  min_instance_number       = 0
  max_instance_number       = 1
  instance_terminate_policy = "OldestInstance"
  default_cooldown          = 10
}

resource "volcengine_scaling_lifecycle_hook" "foo" {
  lifecycle_hook_name    = "acc-test-lifecycle"
  lifecycle_hook_policy  = "CONTINUE"
  lifecycle_hook_timeout = 30
  lifecycle_hook_type    = "SCALE_IN"
  scaling_group_id       = volcengine_scaling_group.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `lifecycle_hook_name` - (Required, ForceNew) The name of the lifecycle hook.
* `lifecycle_hook_policy` - (Required) The policy of the lifecycle hook. Valid values: CONTINUE, REJECT.
* `lifecycle_hook_timeout` - (Required) The timeout of the lifecycle hook.
* `lifecycle_hook_type` - (Required) The type of the lifecycle hook. Valid values: SCALE_IN, SCALE_OUT.
* `scaling_group_id` - (Required, ForceNew) The id of the scaling group.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `lifecycle_hook_id` - The id of the lifecycle hook.


## Import
ScalingLifecycleHook can be imported using the ScalingGroupId:LifecycleHookId, e.g.
```
$ terraform import volcengine_scaling_lifecycle_hook.default scg-yblfbfhy7agh9zn72iaz:sgh-ybqholahe4gso0ee88sd
```

