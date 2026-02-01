---
subcategory: "REDIS"
layout: "volcengine"
page_title: "Volcengine: volcengine_redis_instance_specs"
sidebar_current: "docs-volcengine-datasource-redis_instance_specs"
description: |-
  Use this data source to query detailed information of redis instance specs
---
# volcengine_redis_instance_specs
Use this data source to query detailed information of redis instance specs
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

data "volcengine_redis_instance_specs" "foo" {
  instance_id = volcengine_redis_instance.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `arch_type` - (Optional) The architecture type of the Redis instance.
* `instance_class` - (Optional) The type of Redis instance.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `instance_specs` - The List of Redis instance specifications.
    * `arch_type` - The architecture type of the Redis instance.
    * `node_numbers` - The list of the number of nodes allowed to be used per shard. The number of nodes allowed for different instance types varies.
    * `shard_capacity_specs` - The List of capacity specifications for a single shard.
        * `default_bandwidth_per_shard` - The default bandwidth of the instance under the current memory capacity.
        * `max_additional_bandwidth_per_shard` - The upper limit of bandwidth that an instance is allowed to modify under the current memory capacity.
        * `max_connections_per_shard` - The default maximum number of connections for a single shard.
        * `shard_capacity` - Single-shard memory capacity.
    * `shard_numbers` - The list of shards that the instance is allowed to use. The number of shards allowed for use varies among different instance architecture types.
* `total_count` - The total count of query.


