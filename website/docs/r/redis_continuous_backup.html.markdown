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
resource "volcengine_redis_continuous_backup" "foo" {
  instance_id = "redis-cnlficlt4974s****"
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

