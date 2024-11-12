---
subcategory: "REDIS"
layout: "volcengine"
page_title: "Volcengine: volcengine_redis_instance"
sidebar_current: "docs-volcengine-resource-redis_instance"
description: |-
  Provides a resource to manage redis instance
---
# volcengine_redis_instance
Provides a resource to manage redis instance
## Notice
When Destroy this resource,If the resource charge type is PrePaid,Please unsubscribe the resource 
in  [Volcengine Console](https://console.volcengine.com/finance/unsubscribe/),when complete console operation,yon can
use 'terraform state rm ${resourceId}' to remove.
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
  instance_name       = "tf-test2"
  sharded_cluster     = 1
  password            = "1qaz!QAZ12"
  node_number         = 4
  shard_capacity      = 1024
  shard_number        = 2
  engine_version      = "5.0"
  subnet_id           = volcengine_subnet.foo.id
  deletion_protection = "disabled"
  vpc_auth_mode       = "close"
  charge_type         = "PostPaid"
  port                = 6381
  project_name        = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
  tags {
    key   = "k3"
    value = "v3"
  }

  param_values {
    name  = "active-defrag-cycle-min"
    value = "5"
  }
  param_values {
    name  = "active-defrag-cycle-max"
    value = "28"
  }

  backup_period = [1, 2, 3]
  backup_hour   = 6
  backup_active = true

  create_backup     = false
  apply_immediately = true

  multi_az = "enabled"
  configure_nodes {
    az = "cn-guilin-a"
  }
  configure_nodes {
    az = "cn-guilin-b"
  }
  configure_nodes {
    az = "cn-guilin-c"
  }
  configure_nodes {
    az = "cn-guilin-b"
  }
  #additional_bandwidth = 12
}
```
## Argument Reference
The following arguments are supported:
* `engine_version` - (Required, ForceNew) The engine version of redis instance. Valid value: `5.0`, `6.0`, `7.0`.
* `node_number` - (Required) The number of nodes in each shard, the valid value range is `1-6`. When the value is 1, it means creating a single node instance, and this field can not be modified. When the value is greater than 1, it means creating a primary and secondary instance, and this field can be modified.
* `shard_capacity` - (Required) The memory capacity of each shard, unit is MiB. The valid value range is as fallows: When the value of `ShardedCluster` is 0: 256, 1024, 2048, 4096, 8192, 16384, 32768, 65536. When the value of `ShardedCluster` is 1: 1024, 2048, 4096, 8192, 16384. When the value of `node_number` is 1, the value of this field can not be 256.
* `sharded_cluster` - (Required) Whether enable sharded cluster for the current redis instance. Valid values: 0, 1. 0 means disable, 1 means enable.
* `subnet_id` - (Required) The subnet id of the redis instance. The specified subnet id must belong to the zone ids.
* `additional_bandwidth` - (Optional) Modify the single-shard additional bandwidth of the target Redis instance. Set the additional bandwidth of a single shard, that is, the bandwidth that needs to be additionally increased on the basis of the default bandwidth. Unit: MB/s. The value of additional bandwidth needs to meet the following conditions at the same time: It must be greater than or equal to 0. When the value is 0, it means that no additional bandwidth is added, and the bandwidth of a single shard is the default bandwidth. The sum of additional bandwidth and default bandwidth cannot exceed the upper limit of bandwidth that can be modified for the current instance. Different specification nodes have different upper limits of bandwidth that can be modified. For more details, please refer to bandwidth modification range. The upper limits of the total write bandwidth and the total read bandwidth of an instance are both 2048MB/s.
* `apply_immediately` - (Optional) Whether to apply the instance configuration change operation immediately. The value of this field is false, means that the change operation will be applied within maintenance time.
* `auto_renew` - (Optional) Whether to enable automatic renewal. This field is valid only when `ChargeType` is `PrePaid`, the default value is false. 
When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `backup_active` - (Optional) Whether enable auto backup for redis instance. This field is valid and required when updating the backup plan of primary and secondary instance.
* `backup_hour` - (Optional) The time period to start performing the backup. The valid value range is any integer between 0 and 23, where 0 means that the system will perform the backup in the period of 00:00~01:00, 1 means that the backup will be performed in the period of 01:00~02:00, and so on. 
This field is valid and required when updating the backup plan of primary and secondary instance.
* `backup_period` - (Optional) The backup period. The valid value can be any integer between 1 and 7. Among them, 1 means backup every Monday, 2 means backup every Tuesday, and so on. 
This field is valid and required when updating the backup plan of primary and secondary instance.
* `charge_type` - (Optional) The charge type of redis instance. Valid value: `PostPaid`, `PrePaid`.
* `configure_nodes` - (Optional) Set the list of available zones to which the node belongs.
* `create_backup` - (Optional) Whether to create a final backup when modify the instance configuration or destroy the redis instance.
* `deletion_protection` - (Optional) Whether enable deletion protection for redis instance. Valid values: `enabled`, `disabled`(default).
* `instance_name` - (Optional) The name of the redis instance.
* `multi_az` - (Optional) Set the availability zone deployment scheme for the instance. The value range is as follows: 
disabled: Single availability zone deployment scheme.
 enabled: Multi-availability zone deployment scheme.
 Description:
 When the newly created instance is a single-node instance (that is, when the value of NodeNumber is 1), only the single availability zone deployment scheme is allowed. At this time, the value of MultiAZ must be disabled.
* `param_values` - (Optional) The configuration item information to be modified. This field can only be added or modified. Deleting this field is invalid.
When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields, or use the command `terraform apply` to perform a modification operation.
* `password` - (Optional) The account password. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields. If this parameter is left blank, it means that no password is set for the default account. At this time, the system will automatically generate a password for the default account to ensure instance access security. No account can obtain this random password. Therefore, before connecting to the instance, you need to reset the password of the default account through the ModifyDBAccount interface.You can also set a new account and password through the CreateDBAccount interface according to business needs. If you need to use password-free access function, you need to enable password-free access first through the ModifyDBInstanceVpcAuthMode interface.
* `port` - (Optional, ForceNew) The port of custom define private network address. The valid value range is `1024-65535`. The default value is `6379`.
* `project_name` - (Optional) The project name to which the redis instance belongs, if this parameter is empty, the new redis instance will be added to the `default` project.
* `purchase_months` - (Optional) The purchase months of redis instance, the unit is month. the valid value range is as fallows: `1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36`. This field is valid and required when `ChargeType` is `Prepaid`. 
When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `shard_number` - (Optional) The number of shards in redis instance, the valid value range is `2-256`. This field is valid and required when the value of `ShardedCluster` is 1.
* `tags` - (Optional) Tags.
* `vpc_auth_mode` - (Optional) Whether to enable password-free access when connecting to an instance through a private network. Valid values: `open`, `close`.
* `zone_ids` - (Optional, **Deprecated**) This field has been deprecated after version-0.0.152. Please use multi_az and configure_nodes to specify the availability zone. The list of zone IDs of instance. When creating a single node instance, only one zone id can be specified.

The `configure_nodes` object supports the following:

* `az` - (Required) Set the availability zone to which the node belongs. The number of nodes of an instance (i.e., NodeNumber) and the availability zone deployment scheme (i.e., the value of the MultiAZ parameter) will affect the filling of the current parameter. Among them:
 When a new instance is a single-node instance (i.e., the value of NodeNumber is 1), only a single availability zone deployment scheme is allowed (i.e., the value of MultiAZ must be disabled). At this time, only one availability zone needs to be passed in AZ, and all nodes in the instance will be deployed in this availability zone. When creating a new instance as a primary-standby instance (that is, when the value of NodeNumber is greater than or equal to 2), the number of availability zones passed in must be equal to the number of nodes in a single shard (that is, the value of the NodeNumber parameter), and the value of AZ must comply with the multi-availability zone deployment scheme rules. The specific rules are as follows: If the primary-standby instance selects the multi-availability zone deployment scheme (that is, the value of MultiAZ is enabled), then at least two different availability zone IDs must be passed in in AZ, and the first availability zone is the availability zone where the primary node is located. If the primary and standby instances choose a single availability zone deployment scheme (that is, the value of MultiAZ is disabled), then the availability zones passed in for each node must be the same.

The `param_values` object supports the following:

* `name` - (Required) The name of configuration parameter.
* `value` - (Required) The value of configuration parameter.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
redis instance can be imported using the id, e.g.
```
$ terraform import volcengine_redis_instance.default redis-n769ewmjjqyqh5dv
```
Adding or removing nodes and migrating availability zones for multiple AZ instances are not supported to be orchestrated simultaneously, but it is possible for single AZ instances.

