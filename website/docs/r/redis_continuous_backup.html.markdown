---
subcategory: "REDIS"
layout: "volcengine"
page_title: "Volcengine: volcengine_redis_continuous_backup"
sidebar_current: "docs-volcengine-resource-redis_continuous_backup"
description: |-
  Provides a resource to manage redis continuous backup
---
# volcengine_redis_continuous_backup
Provides a resource to manage redis continuous backup
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

resource "volcengine_redis_continuous_backup" "foo" {
  instance_id = volcengine_redis_instance.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required, ForceNew) The Id of redis instance.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Redis Continuous Backup can be imported using the continuous:instanceId, e.g.
```
$ terraform import volcengine_redis_continuous_backup.default continuous:redis-asdljioeixxxx
```

