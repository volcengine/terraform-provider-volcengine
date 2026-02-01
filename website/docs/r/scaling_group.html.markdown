---
subcategory: "AUTOSCALING"
layout: "volcengine"
page_title: "Volcengine: volcengine_scaling_group"
sidebar_current: "docs-volcengine-resource-scaling_group"
description: |-
  Provides a resource to manage scaling group
---
# volcengine_scaling_group
Provides a resource to manage scaling group
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
  count                     = 3
  scaling_group_name        = "acc-test-scaling-group-${count.index}"
  subnet_ids                = [volcengine_subnet.foo.id]
  multi_az_policy           = "BALANCE"
  desire_instance_number    = 0
  min_instance_number       = 0
  max_instance_number       = 10
  instance_terminate_policy = "OldestInstance"
  default_cooldown          = 30

  tags {
    key   = "k2"
    value = "v2"
  }

  tags {
    key   = "k1"
    value = "v1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `max_instance_number` - (Required) The max instance number of the scaling group. Value range: 0 ~ 100.
* `min_instance_number` - (Required) The min instance number of the scaling group. Value range: 0 ~ 100.
* `scaling_group_name` - (Required) The name of the scaling group.
* `subnet_ids` - (Required) The list of the subnet id to which the ENI is connected.
* `db_instance_ids` - (Optional, ForceNew) ID of the RDS database instance.
* `default_cooldown` - (Optional) The default cooldown interval of the scaling group. Value range: 5 ~ 86400, unit: second. Default value: 300.
* `desire_instance_number` - (Optional) The desire instance number of the scaling group.
* `ignore_failed_scaling_activities` - (Optional) Whether to ignore failed ASG scaling activities while waiting for capacity. Default is false.
* `instance_terminate_policy` - (Optional) The instance terminate policy of the scaling group. Valid values: OldestInstance, NewestInstance, OldestScalingConfigurationWithOldestInstance, OldestScalingConfigurationWithNewestInstance. Default value: OldestScalingConfigurationWithOldestInstance.
* `launch_template_id` - (Optional) The ID of the launch template bound to the scaling group. The launch template and scaling configuration cannot take effect at the same time.
* `launch_template_overrides` - (Optional) Specify instance specifications.
* `launch_template_version` - (Optional) The version of the launch template bound to the scaling group. Valid values are the version number, Latest, or Default.
* `multi_az_policy` - (Optional) The multi az policy of the scaling group. Valid values: PRIORITY, BALANCE. Default value: PRIORITY.
* `project_name` - (Optional) The ProjectName of the scaling group.
* `scaling_mode` - (Optional, ForceNew) Example recycling mode for the elastic group, with values:
release (default): Release mode.
recycle: Shutdown recycling mode.
* `server_group_attributes` - (Optional) The load balancer server group attributes of the scaling group.
* `tags` - (Optional) Tags.
* `wait_for_capacity_timeout` - (Optional) Maximum duration that Provider should wait for ASG instances to be InService before timing out. Setting this to "0" causes Provider to skip all Capacity Waiting behavior. Default is "0".

The `launch_template_overrides` object supports the following:

* `instance_type` - (Required) The instance type.

The `server_group_attributes` object supports the following:

* `port` - (Required) The port receiving request of the server group. Value range: 1 ~ 65535.
* `server_group_id` - (Required) The id of the server group.
* `weight` - (Required) The weight of the instance. Value range: 0 ~ 100.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `active_scaling_configuration_id` - The scaling configuration id which used by the scaling group.
* `created_at` - The create time of the scaling group.
* `health_check_type` - The health check type of the scaling group.
* `lifecycle_state` - The lifecycle state of the scaling group.
* `load_balancer_health_check_grace_period` - Grace period for health check of CLB instance in elastic group.
* `scaling_group_id` - The id of the scaling group.
* `stopped_instance_count` - The number of stopped instances.
* `total_instance_count` - The total instance count of the scaling group.
* `updated_at` - The create time of the scaling group.
* `vpc_id` - The VPC id of the scaling group.


## Import
ScalingGroup can be imported using the id, e.g.
```
$ terraform import volcengine_scaling_group.default scg-mizl7m1kqccg5smt1bdpijuj
```

