---
subcategory: "REDIS"
layout: "volcengine"
page_title: "Volcengine: volcengine_redis_backup_restore"
sidebar_current: "docs-volcengine-resource-redis_backup_restore"
description: |-
  Provides a resource to manage redis backup restore
---
# volcengine_redis_backup_restore
Provides a resource to manage redis backup restore
## Example Usage
```hcl
resource "volcengine_redis_backup_restore" "default" {
  instance_id = "redis-cnlfvrv4qye6u4lpa"
  time_point  = "2023-04-14T02:51:51Z"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required, ForceNew) Id of instance.
* `time_point` - (Required) Time point of backup, for example: 2021-11-09T06:07:26Z. Use lifecycle and ignore_changes in import.
* `backup_type` - (Optional) The type of backup. The value can be Full or Inc.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Redis Backup Restore can be imported using the restore:instanceId, e.g.
```
$ terraform import volcengine_redis_backup_restore.default restore:redis-asdljioeixxxx
```

