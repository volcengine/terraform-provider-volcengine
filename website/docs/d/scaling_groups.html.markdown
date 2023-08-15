---
subcategory: "AUTOSCALING"
layout: "volcengine"
page_title: "Volcengine: volcengine_scaling_groups"
sidebar_current: "docs-volcengine-datasource-scaling_groups"
description: |-
  Use this data source to query detailed information of scaling groups
---
# volcengine_scaling_groups
Use this data source to query detailed information of scaling groups
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

data "volcengine_scaling_groups" "default" {
  ids = volcengine_scaling_group.foo[*].id
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of scaling group ids.
* `name_regex` - (Optional) A Name Regex of scaling group.
* `output_file` - (Optional) File name where to save data source results.
* `scaling_group_names` - (Optional) A list of scaling group names.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `scaling_groups` - The collection of scaling group query.
    * `active_scaling_configuration_id` - The scaling configuration id which used by the scaling group.
    * `created_at` - The create time of the scaling group.
    * `db_instance_ids` - The list of db instance ids.
    * `default_cooldown` - The default cooldown interval of the scaling group.
    * `desire_instance_number` - The desire instance number of the scaling group.
    * `id` - The id of the scaling group.
    * `instance_terminate_policy` - The instance terminate policy of the scaling group.
    * `launch_template_id` - The ID of the launch template bound to the scaling group.
    * `launch_template_version` - The version of the launch template bound to the scaling group.
    * `lifecycle_state` - The lifecycle state of the scaling group.
    * `max_instance_number` - The max instance number of the scaling group.
    * `min_instance_number` - The min instance number of the scaling group.
    * `multi_az_policy` - The multi az policy of the scaling group. Valid values: PRIORITY, BALANCE.
    * `project_name` - The ProjectName of scaling group.
    * `scaling_group_id` - The id of the scaling group.
    * `scaling_group_name` - The name of the scaling group.
    * `server_group_attributes` - The list of server group attributes.
        * `load_balancer_id` - The load balancer id.
        * `port` - The port receiving request of the server group.
        * `server_group_id` - The server group id.
        * `weight` - The weight of the instance.
    * `subnet_ids` - The list of the subnet id to which the ENI is connected.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `total_instance_count` - The total instance count of the scaling group.
    * `updated_at` - The create time of the scaling group.
    * `vpc_id` - The VPC id of the scaling group.
* `total_count` - The total count of scaling group query.


