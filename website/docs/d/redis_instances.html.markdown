---
subcategory: "REDIS"
layout: "volcengine"
page_title: "Volcengine: volcengine_redis_instances"
sidebar_current: "docs-volcengine-datasource-redis_instances"
description: |-
  Use this data source to query detailed information of redis instances
---
# volcengine_redis_instances
Use this data source to query detailed information of redis instances
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

resource "volcengine_redis_instance" "foo" {
  zone_ids            = [data.volcengine_zones.foo.zones[0].id]
  instance_name       = "acc-test-tf-redis"
  sharded_cluster     = 1
  password            = "1qaz!QAZ12"
  node_number         = 2
  shard_capacity      = 1024
  shard_number        = 2
  engine_version      = "5.0"
  subnet_id           = volcengine_subnet.foo.id
  deletion_protection = "disabled"
  vpc_auth_mode       = "close"
  charge_type         = "PostPaid"
  port                = 6381
  project_name        = "default"
}

data "volcengine_redis_instances" "foo" {
  instance_id = volcengine_redis_instance.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `charge_type` - (Optional) The charge type of redis instance to query. Valid values: `PostPaid`, `PrePaid`.
* `engine_version` - (Optional) The engine version of redis instance to query. Valid values: `4.0`, `5.0`, `6.0`.
* `instance_id` - (Optional) The id of redis instance to query. This field supports fuzzy queries.
* `instance_name` - (Optional) The name of redis instance to query. This field supports fuzzy queries.
* `name_regex` - (Optional) A name regex of redis.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name of redis instance to query.
* `sharded_cluster` - (Optional) Whether enable sharded cluster for redis instance. Valid values: 0, 1.
* `status` - (Optional) The status of redis instance to query.
* `tags` - (Optional) Tags.
* `vpc_id` - (Optional) The vpc id of redis instance to query. This field supports fuzzy queries.
* `zone_id` - (Optional) The zone id of redis instance to query. This field supports fuzzy queries.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `instances` - The collection of redis instances query.
    * `backup_plan` - The list of backup plans.
        * `active` - Whether enable auto backup.
        * `backup_hour` - The time period to start performing the backup. The value range is any integer between 0 and 23, where 0 means that the system will perform the backup in the period of 00:00~01:00, 1 means that the backup will be performed in the period of 01:00~02:00, and so on.
        * `backup_type` - The backup type.
        * `expect_next_backup_time` - The expected time for the next backup to be performed.
        * `instance_id` - The instance ID.
        * `last_update_time` - The last time the backup policy was modified.
        * `period` - The backup cycle. The value can be any integer between 1 and 7. Among them, 1 means backup every Monday, 2 means backup every Tuesday, and so on.
        * `ttl` - The number of days to keep backups, the default is 7 days.
    * `capacity` - The memory capacity information.
        * `total` - The total memory capacity of the redis instance. Unit: MiB.
        * `used` - The used memory capacity of the redis instance. Unit: MiB.
    * `charge_type` - The charge type of the redis instance.
    * `configure_nodes` - Set the list of available zones to which the node belongs.
        * `az` - Set the availability zone to which the node belongs. The number of nodes of an instance (i.e., NodeNumber) and the availability zone deployment scheme (i.e., the value of the MultiAZ parameter) will affect the filling of the current parameter. Among them:
 When a new instance is a single-node instance (i.e., the value of NodeNumber is 1), only a single availability zone deployment scheme is allowed (i.e., the value of MultiAZ must be disabled). At this time, only one availability zone needs to be passed in AZ, and all nodes in the instance will be deployed in this availability zone. When creating a new instance as a primary-standby instance (that is, when the value of NodeNumber is greater than or equal to 2), the number of availability zones passed in must be equal to the number of nodes in a single shard (that is, the value of the NodeNumber parameter), and the value of AZ must comply with the multi-availability zone deployment scheme rules. The specific rules are as follows: If the primary-standby instance selects the multi-availability zone deployment scheme (that is, the value of MultiAZ is enabled), then at least two different availability zone IDs must be passed in in AZ, and the first availability zone is the availability zone where the primary node is located. If the primary and standby instances choose a single availability zone deployment scheme (that is, the value of MultiAZ is disabled), then the availability zones passed in for each node must be the same.
    * `create_time` - The creation time of the redis instance.
    * `deletion_protection` - whether enable deletion protection.
    * `engine_version` - The engine version of the redis instance.
    * `expired_time` - The expire time of the redis instance, valid when charge type is `PrePaid`.
    * `id` - The id of the redis instance.
    * `instance_id` - The id of the redis instance.
    * `instance_name` - The name of the redis instance.
    * `maintenance_time` - The maintainable time of the redis instance.
    * `multi_az` - Set the availability zone deployment scheme for the instance. The value range is as follows: 
disabled: Single availability zone deployment scheme.
 enabled: Multi-availability zone deployment scheme.
 Description:
 When the newly created instance is a single-node instance (that is, when the value of NodeNumber is 1), only the single availability zone deployment scheme is allowed. At this time, the value of MultiAZ must be disabled.
    * `node_ids` - The list of redis instance node IDs.
    * `node_number` - The number of nodes in each shard.
    * `params` - The list of params.
        * `current_value` - Current value of the configuration parameter.
        * `default_value` - Default value of the configuration parameter.
        * `description` - The description of the configuration parameter.
        * `editable_for_instance` - Whether the current redis instance supports editing this parameter.
        * `need_reboot` - Whether need to reboot the redis instance when modifying this parameter.
        * `options` - The list of options. Valid when the configuration parameter type is `Radio`.
            * `description` - The description of this option item.
            * `value` - The Optional item for `Radio` type parameters.
        * `param_name` - The name of the configuration parameter.
        * `range` - The valid value range of the numeric type configuration parameter.
        * `type` - The type of the configuration parameter.
        * `unit` - The unit of the numeric type configuration parameter.
    * `project_name` - The project name of the redis instance.
    * `region_id` - The region id of the redis instance.
    * `shard_capacity` - The memory capacity of each shard. Unit: GiB.
    * `shard_number` - The number of shards in the redis instance.
    * `sharded_cluster` - Whether enable sharded cluster for the redis instance.
    * `status` - The status of the redis instance.
    * `subnet_id` - The subnet id of the redis instance.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `visit_addrs` - The list of connection information.
        * `addr_type` - The connection address type.
        * `address` - The connection address.
        * `eip_id` - The EIP ID bound to the instance's public network address.
        * `port` - The connection port.
        * `vip_v6` - The ipv6 address of the connection address.
        * `vip` - The ipv4 address of the connection address.
    * `vpc_auth_mode` - Whether to enable password-free access when connecting to an instance through a private network.
    * `vpc_id` - The vpc ID of the redis instance.
    * `zone_ids` - The list of zone ID which the redis instance belongs.
* `total_count` - The total count of redis instances query.


