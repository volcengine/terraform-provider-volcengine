---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_restore_backup"
sidebar_current: "docs-volcengine-resource-rds_postgresql_restore_backup"
description: |-
  Provides a resource to manage rds postgresql restore backup
---
# volcengine_rds_postgresql_restore_backup
Provides a resource to manage rds postgresql restore backup
## Example Usage
```hcl
resource "volcengine_rds_postgresql_restore_backup" "example" {
  backup_id                  = "20251214-200431-0698LD"
  source_db_instance_id      = "postgres-72715e0d9f58"
  target_db_instance_id      = "postgres-72715e0d9f58"
  target_db_instance_account = "super"

  databases {
    db_name     = "test"
    new_db_name = "test_restored"
  }
}
```
## Argument Reference
The following arguments are supported:
* `backup_id` - (Required) The backup ID used for restore.Only supports restoring data to an existing instance through logical backup.
* `databases` - (Required) Information of the database to be restored.
* `source_db_instance_id` - (Required) The ID of the instance to which the backup belongs.
* `target_db_instance_account` - (Required) The account used as the Owner of the newly restored database in the target instance.
* `target_db_instance_id` - (Required) The ID of the target instance for restore.

The `databases` object supports the following:

* `db_name` - (Required) Original database name.
* `new_db_name` - (Required) New database name.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
RdsPostgresqlRestoreBackup can be imported using the id, e.g.
```
$ terraform import volcengine_rds_postgresql_restore_backup.default resource_id
```

