---
subcategory: "RDS_MSSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_mssql_backup"
sidebar_current: "docs-volcengine-resource-rds_mssql_backup"
description: |-
  Provides a resource to manage rds mssql backup
---
# volcengine_rds_mssql_backup
Provides a resource to manage rds mssql backup
## Example Usage
```hcl
resource "volcengine_rds_mssql_backup" "foo" {
  instance_id = "mssql-40914121fd22"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required, ForceNew) The id of the instance.
* `backup_meta` - (Optional, ForceNew) Backup repository information. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `backup_type` - (Optional, ForceNew) Backup type. Currently only supports full backup, with a value of Full (default).

The `backup_meta` object supports the following:

* `db_name` - (Required, ForceNew) The name of the database.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `backup_id` - The ID of the backup.


## Import
Rds Mssql Backup can be imported using the id, e.g.
```
$ terraform import volcengine_rds_mssql_backup.default instanceId:backupId
```

