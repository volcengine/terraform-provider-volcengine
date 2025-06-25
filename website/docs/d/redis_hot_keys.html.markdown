---
subcategory: "REDIS"
layout: "volcengine"
page_title: "Volcengine: volcengine_redis_hot_keys"
sidebar_current: "docs-volcengine-datasource-redis_hot_keys"
description: |-
  Use this data source to query detailed information of redis hot keys
---
# volcengine_redis_hot_keys
Use this data source to query detailed information of redis hot keys
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

data "volcengine_redis_hot_keys" "foo" {
  instance_id = volcengine_redis_instance.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The ID of Instance.
* `key_type` - (Optional) Specify the data type used to filter the query results of hot keys.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `query_end_time` - (Optional) Query the end time in the format of yyyy-MM-ddTHH:mm:ssZ (UTC).
* `query_start_time` - (Optional) Query the start time in the format of yyyy-MM-ddTHH:mm:ssZ (UTC).
* `shard_ids` - (Optional) Specify the list of shard ids used to filter the query results of hot keys.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `hot_key` - The List of hot Key details.
    * `create_time` - The collection time of the hot Key, in the format of yyyy-MM-ddTHH:mm:ssZ (UTC).
    * `db_name` - The name of the database to which the hot Key belongs.
    * `key_info` - The name of the hot Key.
    * `key_type` - The type of hot Key.
    * `node_id` - The node ID to which the hot Key belongs.
    * `query_count` - The number of accesses to the hot Key.
    * `shard_id` - The shard ID to which the hot Key belongs.
* `total_count` - The total count of query.


