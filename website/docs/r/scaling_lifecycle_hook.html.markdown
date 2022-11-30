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
resource "volcengine_scaling_lifecycle_hook" "foo" {
  scaling_group_id       = "scg-ybru8pazhgl8j1di4tyd"
  lifecycle_hook_name    = "tf-test"
  lifecycle_hook_timeout = 30
  lifecycle_hook_type    = "SCALE_IN"
  lifecycle_hook_policy  = "CONTINUE"
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

