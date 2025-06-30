---
subcategory: "REDIS"
layout: "volcengine"
page_title: "Volcengine: volcengine_redis_planned_events"
sidebar_current: "docs-volcengine-datasource-redis_planned_events"
description: |-
  Use this data source to query detailed information of redis planned events
---
# volcengine_redis_planned_events
Use this data source to query detailed information of redis planned events
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

data "volcengine_redis_planned_events" "foo" {
  instance_id = volcengine_redis_instance.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Optional) The ID of instance.
* `max_start_time` - (Optional) The latest execution time of the planned events that need to be queried. The format is yyyy-MM-ddTHH:mm:ssZ (UTC).
* `min_start_time` - (Optional) The earliest execution time of the planned event that needs to be queried. The format is yyyy-MM-ddTHH:mm:ssZ (UTC).
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `planned_events` - The List of planned event information.
    * `action_name` - Event operation name.
    * `can_cancel` - Whether the current event is allowed to be cancelled for execution.
    * `can_modify_time` - Whether the execution time of the current event can be changed.
    * `event_id` - The ID of Event.
    * `instance_id` - The ID of instance.
    * `instance_name` - The name of instance.
    * `max_end_time` - The latest execution time at which changes are allowed for the current event.
    * `plan_end_time` - The latest execution time of the event plan. The format is yyyy-MM-ddTHH:mm:ssZ (UTC).
    * `plan_start_time` - The earliest planned execution time of the event. The format is yyyy-MM-ddTHH:mm:ssZ (UTC).
    * `status` - The status of event.
    * `type` - The type of event.
* `total_count` - The total count of query.


