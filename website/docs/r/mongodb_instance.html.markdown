---
subcategory: "MONGODB"
layout: "volcengine"
page_title: "Volcengine: volcengine_mongodb_instance"
sidebar_current: "docs-volcengine-resource-mongodb_instance"
description: |-
  Provides a resource to manage mongodb instance
---
# volcengine_mongodb_instance
Provides a resource to manage mongodb instance
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

resource "volcengine_mongodb_instance" "foo" {
  zone_ids          = [data.volcengine_zones.foo.zones[0].id]
  db_engine_version = "MongoDB_4_0"
  instance_type     = "ReplicaSet"
  node_spec         = "mongo.2c4g"
  #  mongos_node_spec       = "mongo.mongos.2c4g"
  #  mongos_node_number     = 3
  #  shard_number           = 3
  storage_space_gb       = 20
  subnet_id              = volcengine_subnet.foo.id
  instance_name          = "acc-test-mongodb-replica"
  charge_type            = "PostPaid"
  super_account_password = "93f0cb0614Aab12"
  project_name           = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
  node_availability_zone {
    zone_id     = data.volcengine_zones.foo.zones[0].id
    node_number = 2
  }
  #  period_unit = "Month"
  #  period      = 1
  #  auto_renew  = false
  #  ssl_action  = "Close"
  #  lifecycle {
  #    ignore_changes = [
  #      super_account_password,
  #    ]
  #  }
}
```
## Argument Reference
The following arguments are supported:
* `node_spec` - (Required) The spec of node. When the instance_type is ReplicaSet, this parameter represents the computing node specification of the replica set instance. When the instance_type is ShardedCluster, this parameter represents the specification of the Shard node.
* `storage_space_gb` - (Required) The total storage space of a replica set instance, or the storage space of a single shard in a sharded cluster. Unit: GiB.
* `subnet_id` - (Required, ForceNew) The subnet id of instance.
* `auto_renew` - (Optional) Whether to enable automatic renewal. This parameter is required when the `ChargeType` is `Prepaid`.
* `charge_type` - (Optional) The charge type of instance, valid value contains `Prepaid` or `PostPaid`. Default is `PostPaid`.
* `db_engine_version` - (Optional, ForceNew) The version of db engine, valid value contains `MongoDB_4_0`, `MongoDB_4_2`, `MongoDB_4_4`, `MongoDB_5_0`, `MongoDB_6_0`.
* `instance_name` - (Optional) The instance name.
* `instance_type` - (Optional, ForceNew) The type of instance, the valid value contains `ReplicaSet` or `ShardedCluster`. Default is `ReplicaSet`.
* `mongos_node_number` - (Optional) The mongos node number of shard cluster, value range is `2~23`, this parameter is required when the `InstanceType` is `ShardedCluster`.
* `mongos_node_spec` - (Optional) The mongos node spec of shard cluster, this parameter is required when the `InstanceType` is `ShardedCluster`.
* `node_availability_zone` - (Optional, ForceNew) The readonly node of the instance. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `period_unit` - (Optional) The period unit, valid value contains `Year` or `Month`. This parameter is required when the `ChargeType` is `Prepaid`.
* `period` - (Optional) The instance purchase duration, the value range is `1~3` when `PeriodUtil` is `Year`, the value range is `1~9` when `PeriodUtil` is `Month`. This parameter is required when the `ChargeType` is `Prepaid`.
* `project_name` - (Optional) The project name to which the instance belongs.
* `shard_number` - (Optional) The number of shards in shard cluster, value range is `2~32`, this parameter is required when the `InstanceType` is `ShardedCluster`.
* `super_account_password` - (Optional) The password of database account. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `tags` - (Optional) Tags.
* `vpc_id` - (Optional, ForceNew) The vpc ID.
* `zone_id` - (Optional, ForceNew, **Deprecated**) This field has been deprecated after version-0.0.156. Please use `zone_ids` to deploy multiple availability zones. The zone ID of instance.
* `zone_ids` - (Optional, ForceNew) The list of zone ids. If you need to deploy multiple availability zones for a newly created instance, you can specify three availability zone IDs at the same time. By default, the first available zone passed in is the primary available zone, and the two available zones passed in afterwards are the backup available zones.

The `node_availability_zone` object supports the following:

* `node_number` - (Required, ForceNew) The number of readonly nodes in current zone. Currently, only ReplicaSet instances and Shard in ShardedCluster instances support adding readonly nodes.
When the instance_type is ReplicaSet, this value represents the total number of readonly nodes in a single replica set instance. Each instance of the replica set supports adding up to 5 readonly nodes.
When the instance_type is ShardedCluster, this value represents the number of readonly nodes in each shard. Each shard can add up to 5 readonly nodes.
* `zone_id` - (Required, ForceNew) The zone id of readonly nodes.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `config_servers_id` - The config servers id of the ShardedCluster instance.
* `mongos_id` - The mongos id of the ShardedCluster instance.
* `mongos` - The mongos information of the ShardedCluster instance.
    * `mongos_node_id` - The mongos node ID.
    * `node_spec` - The node spec.
    * `node_status` - The node status.
* `private_endpoint` - The private endpoint address of instance.
* `read_only_node_number` - The number of readonly node in instance.
* `shards` - The shards information of the ShardedCluster instance.
    * `shard_id` - The shard id.


## Import
mongodb instance can be imported using the id, e.g.
```
$ terraform import volcengine_mongodb_instance.default mongo-replica-e405f8e2****
```

