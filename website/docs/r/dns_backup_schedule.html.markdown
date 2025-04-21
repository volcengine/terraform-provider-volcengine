---
subcategory: "DNS"
layout: "volcengine"
page_title: "Volcengine: volcengine_dns_backup_schedule"
sidebar_current: "docs-volcengine-resource-dns_backup_schedule"
description: |-
  Provides a resource to manage dns backup schedule
---
# volcengine_dns_backup_schedule
Provides a resource to manage dns backup schedule
## Example Usage
```hcl
resource "volcengine_dns_backup_schedule" "foo" {
  zid      = 58846
  schedule = 1
}
```
## Argument Reference
The following arguments are supported:
* `schedule` - (Required) The backup schedule. 0: Turn off automatic backup. 1: Automatic backup once per hour. 2: Automatic backup once per day. 3: Automatic backup once per month.
* `zid` - (Required, ForceNew) The ID of the domain for which you want to update the backup schedule.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `count_limit` - Maximum number of backups per domain.


## Import
DnsBackupSchedule can be imported using the id, e.g.
```
$ terraform import volcengine_dns_backup_schedule.default resource_id
```

