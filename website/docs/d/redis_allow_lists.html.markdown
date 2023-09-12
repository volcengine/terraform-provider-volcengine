---
subcategory: "REDIS"
layout: "volcengine"
page_title: "Volcengine: volcengine_redis_allow_lists"
sidebar_current: "docs-volcengine-datasource-redis_allow_lists"
description: |-
  Use this data source to query detailed information of redis allow lists
---
# volcengine_redis_allow_lists
Use this data source to query detailed information of redis allow lists
## Example Usage
```hcl
resource "volcengine_redis_allow_list" "foo" {
  allow_list      = ["192.168.0.0/24"]
  allow_list_name = "acc-test-allowlist"
}

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

resource "volcengine_redis_allow_list_associate" "foo" {
  allow_list_id = volcengine_redis_allow_list.foo.id
  instance_id   = volcengine_redis_instance.foo.id
}

data "volcengine_redis_allow_lists" "foo" {
  instance_id = volcengine_redis_allow_list_associate.foo.instance_id
  region_id   = "cn-beijing"
  name_regex  = volcengine_redis_allow_list.foo.allow_list_name
}
```
## Argument Reference
The following arguments are supported:
* `region_id` - (Required) The Id of region.
* `instance_id` - (Optional) The Id of instance.
* `name_regex` - (Optional) A Name Regex of Allow List.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `allow_lists` - Information of list of allow list.
    * `allow_list_desc` - Description of allow list.
    * `allow_list_id` - Id of allow list.
    * `allow_list_ip_num` - The IP number of allow list.
    * `allow_list_name` - Name of allow list.
    * `allow_list_type` - Type of allow list.
    * `allow_list` - Ip list of allow list.
    * `associated_instance_num` - The number of instance that associated to allow list.
    * `associated_instances` - Instances associated by this allow list.
        * `instance_id` - Id of instance.
        * `instance_name` - Name of instance.
        * `vpc` - Id of virtual private cloud.
* `total_count` - The total count of allow list query.


