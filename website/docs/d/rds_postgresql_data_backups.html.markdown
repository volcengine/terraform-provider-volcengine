---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_data_backups"
sidebar_current: "docs-volcengine-datasource-rds_postgresql_data_backups"
description: |-
  Use this data source to query detailed information of rds postgresql data backups
---
# volcengine_rds_postgresql_data_backups
Use this data source to query detailed information of rds postgresql data backups
## Example Usage
```hcl
data "volcengine_rds_postgresql_data_backups" "example" {
  instance_id       = "postgres-72715e0d9f58"
  backup_id         = "20251214-172343F"
  backup_start_time = "2025-12-01T00:00:00.000Z"
  backup_end_time   = "2025-12-15T23:59:59.999Z"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The ID of the PostgreSQL instance.
* `backup_database_name` - (Optional) The name of the database included in the backup set. Only effective when the value of backup_method is Logical.
* `backup_description` - (Optional) The description of the backup set.
* `backup_end_time` - (Optional) The latest time when the backup is created, in the format of yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).
* `backup_id` - (Optional) The ID of the backup.
* `backup_method` - (Optional) The method of the backup: Physical, Logical.
* `backup_scope` - (Optional) The scope of the backup: Instance, Database.
* `backup_start_time` - (Optional) The earliest time when the backup is created, in the format of yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).
* `backup_status` - (Optional) The status of the backup: Success, Failed, Running.
* `backup_type` - (Optional) The type of the backup: Full, Increment.
* `create_type` - (Optional) The creation type of the backup: System, User.
* `download_status` - (Optional) The downloadable status of the backup set. NotAllowed: download is not supported. NeedToPrepare: the backup set is in place and needs background preparation for backup. LinkReady: the backup set is ready for download.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `backups` - The collection of the query.
    * `backup_data_size` - The original size of the data contained in the backup, in Bytes.
    * `backup_description` - The description of the backup set.
    * `backup_end_time` - The end time of the backup. The time format is yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).
    * `backup_file_name` - The name of the backup file.
    * `backup_file_size` - The size of the backup file, in Byte.
    * `backup_id` - The ID of the backup.
    * `backup_meta` - The information about the databases included in the backup.
        * `db_name` - The name of the database.
    * `backup_method` - The method of the backup: Physical, Logical.
    * `backup_progress` - The progress of the backup. The unit is percentage.
    * `backup_scope` - The scope of the backup: Instance, Database.
    * `backup_start_time` - The start time of the backup. The time format is yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).
    * `backup_status` - The status of the backup: Success, Failed, Running.
    * `backup_type` - The type of the backup: Full, Increment.
    * `create_type` - The creation type of the backup: System, User.
    * `download_status` - The downloadable status of the backup set.
* `total_count` - The total count of query.


