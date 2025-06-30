---
subcategory: "VMP"
layout: "volcengine"
page_title: "Volcengine: volcengine_vmp_notify_group_policies"
sidebar_current: "docs-volcengine-datasource-vmp_notify_group_policies"
description: |-
  Use this data source to query detailed information of vmp notify group policies
---
# volcengine_vmp_notify_group_policies
Use this data source to query detailed information of vmp notify group policies
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

data "volcengine_vmp_notify_group_policies" "foo" {
  ids = [volcengine_vmp_notify_group_policy.foo.id]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of notify group policy ids.
* `name` - (Optional) The name of notify group policy.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `notify_policies` - The list of notify group policies.
    * `create_time` - The create time of notify group policy.
    * `description` - The description of notify group policy.
    * `id` - The id of the notify group policy.
    * `levels` - The levels of the notify group policy.
        * `group_by` - The aggregate dimension.
        * `group_interval` - The aggregation cycle.
        * `group_wait` - The wait time.
        * `level` - The level of the policy.
        * `repeat_interval` - The notification cycle.
    * `name` - The name of notify group policy.
* `total_count` - The total count of query.


