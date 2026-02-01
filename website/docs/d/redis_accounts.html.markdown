---
subcategory: "REDIS"
layout: "volcengine"
page_title: "Volcengine: volcengine_redis_accounts"
sidebar_current: "docs-volcengine-datasource-redis_accounts"
description: |-
  Use this data source to query detailed information of redis accounts
---
# volcengine_redis_accounts
Use this data source to query detailed information of redis accounts
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

resource "volcengine_redis_account" "foo" {
  account_name = "acc_test_account"
  instance_id  = volcengine_redis_instance.foo.id
  password     = "Password@@"
  role_name    = "ReadOnly"
}

data "volcengine_redis_accounts" "foo" {
  account_name = volcengine_redis_account.foo.account_name
  instance_id  = volcengine_redis_instance.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The id of the Redis instance.
* `account_name` - (Optional) The name of the redis account.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `accounts` - The collection of redis instance account query.
    * `account_name` - The name of the redis account.
    * `description` - The description of the redis account.
    * `instance_id` - The id of instance.
    * `role_name` - The role info.
* `total_count` - The total count of redis accounts query.


