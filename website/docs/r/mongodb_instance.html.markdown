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
## Example Usage
```hcl
resource "volcengine_mongodb_instance" "foo" {
  zone_id       = "cn-beijing-a"
  instance_type = "ReplicaSet"
  node_spec     = "mongo.2c4g"
  #    mongos_node_spec="mongo.2c4g"
  #    shard_number=3
  storage_space_gb       = 20
  subnet_id              = "subnet-rrx4ns6abw1sv0x57wq6h47"
  instance_name          = "mongo-replica-be9995d32e4a"
  charge_type            = "PostPaid"
  super_account_password = "******"
  # period_unit="Month"
  # period=1
  # auto_renew=false
  # ssl_action="Close"
  #    lifecycle {
  #        ignore_changes = [
  #            super_account_password,
  #        ]
  #    }
}
```
## Argument Reference
The following arguments are supported:
* `node_spec` - (Required) The spec of node.
* `storage_space_gb` - (Required) The total storage space of a replica set instance, or the storage space of a single shard in a sharded cluster, in GiB.
* `subnet_id` - (Required, ForceNew) The subnet id of instance.
* `super_account_password` - (Required) The password of database account.
* `auto_renew` - (Optional) Whether to enable automatic renewal.
* `charge_type` - (Optional) The charge type of instance,valid value contains `Prepaid` or `PostPaid`.
* `instance_name` - (Optional) The instance name.
* `instance_type` - (Optional) The type of instance,the valid value contains `ReplicaSet` or `ShardedCluster`.
* `mongos_node_number` - (Optional) The mongos node number of shard cluster,value range is `2~23`,this parameter is required when `InstanceType` is `ShardedCluster`.
* `mongos_node_spec` - (Optional) The mongos node spec of shard cluster,this parameter is required when `InstanceType` is `ShardedCluster`.
* `period_unit` - (Optional) The period unit,valid value contains `Year` or `Month`,this parameter is required when `ChargeType` is `Prepaid`.
* `period` - (Optional) The instance purchase duration,the value range is `1~3` when `PeriodUtil` is `Year`,the value range is `1~9` when `PeriodUtil` is `Month`,this parameter is required when `ChargeType` is `Prepaid`.
* `project_name` - (Optional) The project name to which the instance belongs.
* `shard_number` - (Optional) The number of shards in shard cluster,value range is `2~23`,this parameter is required when `InstanceType` is `ShardedCluster`.
* `vpc_id` - (Optional) The vpc ID.
* `zone_id` - (Optional) The zone ID of instance.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
mongodb instance can be imported using the id, e.g.
```
$ terraform import volcengine_mongodb_instance.default mongo-replica-e405f8e2****
```

