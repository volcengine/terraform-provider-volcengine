---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_data_backup"
sidebar_current: "docs-volcengine-resource-rds_postgresql_data_backup"
description: |-
  Provides a resource to manage rds postgresql data backup
---
# volcengine_rds_postgresql_data_backup
Provides a resource to manage rds postgresql data backup
## Example Usage
```hcl
resource "volcengine_rds_postgresql_data_backup" "example" {
  instance_id        = "postgres-72715e0d9f58"
  backup_scope       = "Instance"
  backup_method      = "Physical"
  backup_type        = "Full"
  backup_description = "tf demo full backup2"
}

resource "volcengine_rds_postgresql_data_backup" "example1" {
  instance_id        = "postgres-72715e0d9f58"
  backup_scope       = "Instance"
  backup_method      = "Logical"
  backup_description = "tf demo logical backup"
}

resource "volcengine_rds_postgresql_data_backup" "example2" {
  instance_id        = "postgres-72715e0d9f58"
  backup_scope       = "Database"
  backup_method      = "Logical"
  backup_description = "tf demo database full backup"
  backup_meta {
    db_name = "test"
  }
  backup_meta {
    db_name = "test-1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required, ForceNew) The id of the PostgreSQL instance.
* `backup_description` - (Optional, ForceNew) The description of the backup set.
* `backup_meta` - (Optional, ForceNew) Specify the database that needs to be backed up. This parameter can only be set when the value of backup_scope is Database.
* `backup_method` - (Optional, ForceNew) The method of the backup: Physical, Logical.When the value of backup_scope is Database, the value of backup_method can only be Logical.
* `backup_scope` - (Optional, ForceNew) The scope of the backup: Instance, Database.
* `backup_type` - (Optional, ForceNew) The backup type of the backup: Full(default), Increment. Do not set this parameter when backup_method is Logical; otherwise, the creation will fail.

The `backup_meta` object supports the following:

* `db_name` - (Required) The name of the database.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `backup_id` - The id of the backup.


## Import
RdsPostgresqlDataBackup can be imported using the id, e.g.
```
$ terraform import volcengine_rds_postgresql_data_backup.default resource_id
```

