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
data "volcengine_redis_accounts" "default" {
  instance_id = "redis-cn0398aizj8cwmopx"
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


