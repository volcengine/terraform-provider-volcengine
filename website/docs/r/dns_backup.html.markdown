---
subcategory: "DNS"
layout: "volcengine"
page_title: "Volcengine: volcengine_dns_backup"
sidebar_current: "docs-volcengine-resource-dns_backup"
description: |-
  Provides a resource to manage dns backup
---
# volcengine_dns_backup
Provides a resource to manage dns backup
## Example Usage
```hcl
resource "volcengine_dns_backup" "foo" {
  zid = 58846
}
```
## Argument Reference
The following arguments are supported:
* `zid` - (Required, ForceNew) The ID of the domain for which you want to get the backup schedule.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `backup_id` - The ID of backup.
* `backup_time` - Time when the backup was created. Timezone is UTC.


## Import
DnsBackup can be imported using the id, e.g.
```
$ terraform import volcengine_dns_backup.default ZID:BackupID
```

