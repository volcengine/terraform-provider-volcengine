---
subcategory: "AUTOSCALING"
layout: "volcengine"
page_title: "Volcengine: volcengine_scaling_lifecycle_hooks"
sidebar_current: "docs-volcengine-datasource-scaling_lifecycle_hooks"
description: |-
  Use this data source to query detailed information of scaling lifecycle hooks
---
# volcengine_scaling_lifecycle_hooks
Use this data source to query detailed information of scaling lifecycle hooks
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
  count                  = 3
  lifecycle_hook_name    = "acc-test-lifecycle-${count.index}"
  lifecycle_hook_policy  = "CONTINUE"
  lifecycle_hook_timeout = 30
  lifecycle_hook_type    = "SCALE_IN"
  scaling_group_id       = volcengine_scaling_group.foo.id
}

data "volcengine_scaling_lifecycle_hooks" "foo" {
  ids              = volcengine_scaling_lifecycle_hook.foo[*].lifecycle_hook_id
  scaling_group_id = volcengine_scaling_group.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `scaling_group_id` - (Required) An id of scaling group id.
* `ids` - (Optional) A list of lifecycle hook ids.
* `lifecycle_hook_names` - (Optional) A list of lifecycle hook names.
* `name_regex` - (Optional) A Name Regex of lifecycle hook.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `lifecycle_hooks` - The collection of lifecycle hook query.
    * `id` - The id of the lifecycle hook.
    * `lifecycle_command` - Batch job command.
        * `command_id` - Batch job command ID, which indicates the batch job command to be executed after triggering the lifecycle hook and installed in the instance.
        * `parameters` - Parameters and parameter values in batch job commands.
The number of parameters ranges from 0 to 60.
    * `lifecycle_hook_id` - The id of the lifecycle hook.
    * `lifecycle_hook_name` - The name of the lifecycle hook.
    * `lifecycle_hook_policy` - The policy of the lifecycle hook.
    * `lifecycle_hook_timeout` - The timeout of the lifecycle hook.
    * `lifecycle_hook_type` - The type of the lifecycle hook.
    * `scaling_group_id` - The id of the scaling group.
* `total_count` - The total count of lifecycle hook query.


