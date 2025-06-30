---
subcategory: "REDIS"
layout: "volcengine"
page_title: "Volcengine: volcengine_redis_slow_logs"
sidebar_current: "docs-volcengine-datasource-redis_slow_logs"
description: |-
  Use this data source to query detailed information of redis slow logs
---
# volcengine_redis_slow_logs
Use this data source to query detailed information of redis slow logs
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

data "volcengine_redis_slow_logs" "default" {
  instance_id = volcengine_redis_instance.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The ID of Instance.
* `context` - (Optional) The context of the query results for slow log recording is used when more slow log records need to be loaded.
* `db_name` - (Optional) The database where the slow log is located.
* `name_regex` - (Optional) A Name Regex of Resource.
* `node_ids` - (Optional) The node ID of the slow log needs to be queried.
* `output_file` - (Optional) File name where to save data source results.
* `query_end_time` - (Optional) Query the end time in the format of yyyy-MM-ddTHH:mm:ssZ (UTC).
* `query_start_time` - (Optional) Query the start time in the format of yyyy-MM-ddTHH:mm:ssZ (UTC).
* `slow_log_type` - (Optional) The types of slow logs.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `slow_query` - The Details of the slow log.
    * `db_name` - The name of the database.
    * `execution_start_time` - The execution start time of the slow query statement is in the format of yyyy-MM-ddTHH:mm: ssZ (UTC).
    * `host_address` - The address of the client that issues the slow query request.
    * `instance_id` - The ID of Instance.
    * `node_id` - The node ID to which the slow log belongs.
    * `query_text` - Slow query statement.
    * `query_times` - Slow query statement execution time, unit: microseconds (us).
    * `user_name` - The name of the account.
* `total_count` - The total count of query.


