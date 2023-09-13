---
subcategory: "REDIS"
layout: "volcengine"
page_title: "Volcengine: volcengine_redis_backup"
sidebar_current: "docs-volcengine-resource-redis_backup"
description: |-
  Provides a resource to manage redis backup
---
# volcengine_redis_backup
Provides a resource to manage redis backup
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

resource "volcengine_redis_backup" "foo" {
  instance_id = volcengine_redis_instance.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required, ForceNew) Id of instance to create backup.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `backup_point_id` - The id of backup point.
* `backup_strategy` - Backup strategy.
* `backup_type` - Backup type.
* `end_time` - End time of backup.
* `instance_detail` - Information of instance.
    * `account_id` - Id of account.
    * `arch_type` - Arch type of instance(Standard/Cluster).
    * `charge_type` - Charge type of instance(Postpaid/Prepaid).
    * `engine_version` - Engine version of instance.
    * `expired_time` - Expired time of instance.
    * `instance_id` - Id of instance.
    * `instance_name` - Name of instance.
    * `maintenance_time` - The maintainable period (in UTC) of the instance.
    * `network_type` - Network type of instance.
    * `project_name` - Project name of instance.
    * `region_id` - Id of region.
    * `replicas` - Count of replica in which shard.
    * `server_cpu` - Count of cpu cores of instance.
    * `shard_capacity` - Capacity of shard.
    * `shard_count` - Count of shard.
    * `total_capacity` - Total capacity of instance.
    * `used_capacity` - Capacity used of this instance.
    * `vpc_info` - Information of vpc.
        * `id` - Id of vpc.
        * `name` - Name of vpc.
    * `zone_ids` - List of id of zone.
* `size` - Size in MiB.
* `start_time` - Start time of backup.
* `status` - Status of backup (Creating/Available/Unavailable/Deleting).


## Import
Redis Backup can be imported using the instanceId:backupId, e.g.
```
$ terraform import volcengine_redis_backup.default redis-cn02aqusft7ws****:b-cn02xmmrp751i9cdzcphjmk4****
```

