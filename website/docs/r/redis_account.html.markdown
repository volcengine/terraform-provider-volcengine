---
subcategory: "REDIS"
layout: "volcengine"
page_title: "Volcengine: volcengine_redis_account"
sidebar_current: "docs-volcengine-resource-redis_account"
description: |-
  Provides a resource to manage redis account
---
# volcengine_redis_account
Provides a resource to manage redis account
## Example Usage
```hcl
resource "volcengine_redis_account" "foo" {
  instance_id  = "redis-cn0398aizj8cwmopx"
  account_name = "test"
  password     = "1qaz!QAZ"
  role_name    = "ReadOnly"
  description  = "test12345"
}
```
## Argument Reference
The following arguments are supported:
* `account_name` - (Required, ForceNew) Redis account name.
* `instance_id` - (Required, ForceNew) The ID of the Redis instance.
* `password` - (Required) The password of the redis account. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `role_name` - (Required) Role type, the valid value can be `Administrator`, `ReadWrite`, `ReadOnly`, `NotDangerous`.
* `description` - (Optional) The description of the redis account.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Redis account can be imported using the instanceId:accountName, e.g.
```
$ terraform import volcengine_redis_account.default redis-42b38c769c4b:test
```

