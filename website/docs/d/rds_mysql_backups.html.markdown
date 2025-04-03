---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_mysql_backups"
sidebar_current: "docs-volcengine-datasource-rds_mysql_backups"
description: |-
  Use this data source to query detailed information of rds mysql backups
---
# volcengine_rds_mysql_backups
Use this data source to query detailed information of rds mysql backups
## Example Usage
```hcl
data "volcengine_rds_mysql_backups" "foo" {
  backup_end_time   = ""
  backup_id         = ""
  backup_method     = ""
  backup_start_time = ""
  backup_status     = ""
  backup_type       = ""
  create_type       = ""
  instance_id       = ""
}
```
## Argument Reference
The following arguments are supported:
* `backup_end_time` - (Optional) The end time of the backup.
* `backup_id` - (Optional) The id of the backup.
* `backup_method` - (Optional) Backup type, value: Physical: Physical backup. Default value. Logical: Logical backup. Description: There is no default value. When this field is not passed, backups of all states under the query conditions limited by other fields are returned.
* `backup_start_time` - (Optional) The start time of the backup.
* `backup_status` - (Optional) Backup status, values: Success: Success. Failed: Failed. Running: In progress. Description: There is no default value. When this field is not passed, all backups in all states under the query conditions limited by other fields are returned.
* `backup_type` - (Optional) Backup method, value: Full: Full backup under physical backup type or library table backup under logical backup type. Increment: Incremental backup under physical backup type. DumpAll: Full database backup under logical backup type. Description: There is no default value. When this field is not passed, all backups of all methods under the query conditions limited by other fields are returned.
* `create_type` - (Optional) Creator of backup. Values: System: System. User: User. Description: There is no default value. When this field is not passed, all types of backups under the query conditions limited by other fields are returned.
* `instance_id` - (Optional) The id of the instance.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `backups` - The collection of query.
    * `backup_end_time` - The end time of backup, in the format of yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).
    * `backup_file_name` - Backup file name.
    * `backup_file_size` - Backup file size, in bytes.
    * `backup_id` - The id of the backup.
    * `backup_method` - Backup type, value: Physical: Physical backup. Logical: Logical backup.
    * `backup_region` - The region where the backup is located.
    * `backup_start_time` - The start time of backup, in the format of yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).
    * `backup_status` - Backup status, values: Success. Failed. Running.
    * `backup_type` - Backup method, values:
Full: Full backup under physical backup type or library table backup under logical backup type.
Increment: Incremental backup under physical backup type (created by the system).
DumpAll: Full database backup under logical backup type.
Description:
There is no default value. When this field is not passed, all types of backups under the query conditions limited by other fields are returned.
    * `consistent_time` - The time point of a consistent snapshot is in the format of yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).
    * `create_type` - Creator of backup. Values: System. User.
    * `db_table_infos` - The database table information contained in the backup set can include up to 10,000 tables.
Explanation:
When the database is empty, this field is not returned.
        * `database` - Database name.
        * `tables` - Table names.
    * `download_status` - Download status. Values:
NotDownload: Not downloaded.
Success: Downloaded.
Failed: Download failed.
Running: Downloading.
    * `error_message` - Error message.
    * `expired_time` - Expired time of backup, in the format of yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).
    * `id` - The id of the backup.
    * `is_encrypted` - Is the data backup encrypted? Value:
true: Encrypted.
false: Not encrypted.
    * `is_expired` - Whether the backup has expired. Value:
true: Expired.
false: Not expired.
* `total_count` - The total count of query.


