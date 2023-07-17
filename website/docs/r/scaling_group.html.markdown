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
resource "volcengine_scaling_group" "foo" {
  scaling_group_name        = "test-tf"
  subnet_ids                = ["subnet-2fe79j7c8o5c059gp68ksxr93"]
  multi_az_policy           = "BALANCE"
  desire_instance_number    = 0
  min_instance_number       = 0
  max_instance_number       = 1
  instance_terminate_policy = "OldestInstance"
  default_cooldown          = 10
  project_name              = "default"
  tags {
    key   = "tf-key1"
    value = "tf-value1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `max_instance_number` - (Required) The max instance number of the scaling group. Value range: 0 ~ 100.
* `min_instance_number` - (Required) The min instance number of the scaling group. Value range: 0 ~ 100.
* `scaling_group_name` - (Required) The name of the scaling group.
* `subnet_ids` - (Required) The list of the subnet id to which the ENI is connected.
* `default_cooldown` - (Optional) The default cooldown interval of the scaling group. Value range: 5 ~ 86400, unit: second. Default value: 300.
* `desire_instance_number` - (Optional) The desire instance number of the scaling group.
* `instance_terminate_policy` - (Optional) The instance terminate policy of the scaling group. Valid values: OldestInstance, NewestInstance, OldestScalingConfigurationWithOldestInstance, OldestScalingConfigurationWithNewestInstance. Default value: OldestScalingConfigurationWithOldestInstance.
* `launch_template_id` - (Optional) The ID of the launch template bound to the scaling group. The launch template and scaling configuration cannot take effect at the same time.
* `launch_template_version` - (Optional) The version of the launch template bound to the scaling group. Valid values are the version number, Latest, or Default.
* `multi_az_policy` - (Optional) The multi az policy of the scaling group. Valid values: PRIORITY, BALANCE. Default value: PRIORITY.
* `project_name` - (Optional) The ProjectName of the scaling group.
* `server_group_attributes` - (Optional) The load balancer server group attributes of the scaling group.
* `tags` - (Optional) Tags.

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
* `db_instance_ids` - The list of db instance ids.
* `lifecycle_state` - The lifecycle state of the scaling group.
* `scaling_group_id` - The id of the scaling group.
* `total_instance_count` - The total instance count of the scaling group.
* `updated_at` - The create time of the scaling group.
* `vpc_id` - The VPC id of the scaling group.


## Import
ScalingGroup can be imported using the id, e.g.
```
$ terraform import volcengine_scaling_group.default scg-mizl7m1kqccg5smt1bdpijuj
```

