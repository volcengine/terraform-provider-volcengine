---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_mysql_backup"
sidebar_current: "docs-volcengine-resource-rds_mysql_backup"
description: |-
  Provides a resource to manage rds mysql backup
---
# volcengine_rds_mysql_backup
Provides a resource to manage rds mysql backup
## Example Usage
```hcl
resource "volcengine_rds_mysql_backup" "foo" {
  instance_id = "mysql-c8c3f45c4b07"
  #backup_type = "Full"
  backup_method = "Logical"
  backup_meta {
    db_name = "order"
  }
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required, ForceNew) The id of the instance.
* `backup_meta` - (Optional, ForceNew) When creating a library table backup of logical backup type, it is used to specify the library table information to be backed up.
Prerequisite: When the value of BackupMethod is Logical, and the BackupType field is not passed.
Mutual exclusion situation: When the value of the BackupType field is DumpAll, this field is not effective.
Quantity limit: When creating a specified library table backup, the upper limit of the number of libraries is 5000, and the upper limit of the number of tables in each library is 5000.
* `backup_method` - (Optional, ForceNew) Backup method. Value range: Full, full backup under physical backup type. Default value. DumpAll: full database backup under logical backup type. Prerequisite: If you need to create a full database backup of logical backup type, that is, when the value of BackupType is DumpAll, the backup type should be set to logical backup, that is, the value of BackupMethod should be Logical. If you need to create a database table backup of logical backup type, you do not need to pass in this field. You only need to specify the database and table to be backed up in the BackupMeta field.
* `backup_type` - (Optional, ForceNew) Backup type. Currently, only full backup is supported. The value is Full.

The `backup_meta` object supports the following:

* `db_name` - (Required, ForceNew) Specify the database that needs to be backed up.
* `table_names` - (Optional, ForceNew) Specify the tables to be backed up in the specified database. When this field is empty, it defaults to full database backup.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `backup_id` - The id of the backup.


## Import
RdsMysqlBackup can be imported using the id, e.g.
```
$ terraform import volcengine_rds_mysql_backup.default instanceId:backupId
```

