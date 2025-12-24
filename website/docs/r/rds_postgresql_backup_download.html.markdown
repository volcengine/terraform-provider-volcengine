---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_backup_download"
sidebar_current: "docs-volcengine-resource-rds_postgresql_backup_download"
description: |-
  Provides a resource to manage rds postgresql backup download
---
# volcengine_rds_postgresql_backup_download
Provides a resource to manage rds postgresql backup download
## Example Usage
```hcl
resource "volcengine_rds_postgresql_backup_download" "example" {
  instance_id = "postgres-72715e0d9f58"
  backup_id   = "20251214-200431-0698LD"
}
```
## Argument Reference
The following arguments are supported:
* `backup_id` - (Required, ForceNew) The ID of the logical backup to be downloaded.
* `instance_id` - (Required, ForceNew) The id of the PostgreSQL instance.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
RdsPostgresqlBackupDownload can be imported using the id, e.g.
```
$ terraform import volcengine_rds_postgresql_backup_download.default resource_id
```

