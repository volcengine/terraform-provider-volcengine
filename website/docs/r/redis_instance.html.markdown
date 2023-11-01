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
  zone_ids            = [data.volcengine_zones.foo.zones[0].id]
  instance_name       = "tf-test"
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
}
```
## Argument Reference
The following arguments are supported:
* `engine_version` - (Required, ForceNew) The engine version of redis instance. Valid value: `4.0`, `5.0`, `6.0`.
* `node_number` - (Required) The number of nodes in each shard, the valid value range is `1-6`. When the value is 1, it means creating a single node instance, and this field can not be modified. When the value is greater than 1, it means creating a primary and secondary instance, and this field can be modified.
* `password` - (Required) The account password. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `shard_capacity` - (Required) The memory capacity of each shard, unit is MiB. The valid value range is as fallows: When the value of `ShardedCluster` is 0: 256, 1024, 2048, 4096, 8192, 16384, 32768, 65536. When the value of `ShardedCluster` is 1: 1024, 2048, 4096, 8192, 16384. When the value of `node_number` is 1, the value of this field can not be 256.
* `sharded_cluster` - (Required, ForceNew) Whether enable sharded cluster for the current redis instance. Valid values: 0, 1. 0 means disable, 1 means enable.
* `subnet_id` - (Required) The subnet id of the redis instance. The specified subnet id must belong to the zone ids.
* `zone_ids` - (Required, ForceNew) The list of zone IDs of instance. When creating a single node instance, only one zone id can be specified.
* `apply_immediately` - (Optional) Whether to apply the instance configuration change operation immediately. The value of this field is false, means that the change operation will be applied within maintenance time.
* `auto_renew` - (Optional) Whether to enable automatic renewal. This field is valid only when `ChargeType` is `PrePaid`, the default value is false. 
When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `backup_active` - (Optional) Whether enable auto backup for redis instance. This field is valid and required when updating the backup plan of primary and secondary instance.
* `backup_hour` - (Optional) The time period to start performing the backup. The valid value range is any integer between 0 and 23, where 0 means that the system will perform the backup in the period of 00:00~01:00, 1 means that the backup will be performed in the period of 01:00~02:00, and so on. 
This field is valid and required when updating the backup plan of primary and secondary instance.
* `backup_period` - (Optional) The backup period. The valid value can be any integer between 1 and 7. Among them, 1 means backup every Monday, 2 means backup every Tuesday, and so on. 
This field is valid and required when updating the backup plan of primary and secondary instance.
* `charge_type` - (Optional) The charge type of redis instance. Valid value: `PostPaid`, `PrePaid`.
* `create_backup` - (Optional) Whether to create a final backup when modify the instance configuration or destroy the redis instance.
* `deletion_protection` - (Optional) Whether enable deletion protection for redis instance. Valid values: `enabled`, `disabled`(default).
* `instance_name` - (Optional) The name of the redis instance.
* `param_values` - (Optional) The configuration item information to be modified. This field can only be added or modified. Deleting this field is invalid.
When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields, or use the command `terraform apply` to perform a modification operation.
* `port` - (Optional, ForceNew) The port of custom define private network address. The valid value range is `1024-65535`. The default value is `6379`.
* `project_name` - (Optional, ForceNew) The project name to which the redis instance belongs, if this parameter is empty, the new redis instance will not be added to any project.
* `purchase_months` - (Optional) The purchase months of redis instance, the unit is month. the valid value range is as fallows: `1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36`. This field is valid and required when `ChargeType` is `Prepaid`. 
When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `shard_number` - (Optional) The number of shards in redis instance, the valid value range is `2-256`. This field is valid and required when the value of `ShardedCluster` is 1.
* `tags` - (Optional, ForceNew) Tags.
* `vpc_auth_mode` - (Optional) Whether to enable password-free access when connecting to an instance through a private network. Valid values: `open`, `close`.

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

