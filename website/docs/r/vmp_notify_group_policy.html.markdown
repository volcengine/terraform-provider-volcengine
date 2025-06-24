---
subcategory: "VMP"
layout: "volcengine"
page_title: "Volcengine: volcengine_vmp_notify_group_policy"
sidebar_current: "docs-volcengine-resource-vmp_notify_group_policy"
description: |-
  Provides a resource to manage vmp notify group policy
---
# volcengine_vmp_notify_group_policy
Provides a resource to manage vmp notify group policy
## Example Usage
```hcl
resource "volcengine_vmp_notify_group_policy" "foo" {
  name        = "acc-test-1"
  description = "acc-test-1"
  levels {
    level           = "P2"
    group_by        = ["__rule__"]
    group_wait      = "35"
    group_interval  = "30"
    repeat_interval = "30"
  }
  levels {
    level           = "P0"
    group_by        = ["__rule__"]
    group_wait      = "30"
    group_interval  = "30"
    repeat_interval = "30"
  }
  levels {
    level           = "P1"
    group_by        = ["__rule__"]
    group_wait      = "40"
    group_interval  = "45"
    repeat_interval = "30"
  }
}
```
## Argument Reference
The following arguments are supported:
* `levels` - (Required) The levels of the notify group policy. Levels must be registered in three (`P0`, `P1`, `P2`) aggregation strategies, and `Level` cannot be repeated.
* `name` - (Required) The name of the notify group policy.
* `description` - (Optional) The description of the notify group policy.

The `levels` object supports the following:

* `group_by` - (Required) The aggregate dimension, the value can be `__rule__`.
* `group_interval` - (Required) The aggregation cycle. Integer form, unit is second.
* `group_wait` - (Required) The wait time. Integer form, unit is second.
* `level` - (Required) The level of the policy, the value can be one of the following: `P0`, `P1`, `P2`.
* `repeat_interval` - (Required) The notification cycle. Integer form, unit is second.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
VMP Notify Group Policy can be imported using the id, e.g.
```
$ terraform import volcengine_vmp_notify_group_policy.default 60dde3ca-951c-4c05-8777-e5a7caa07ad6
```

