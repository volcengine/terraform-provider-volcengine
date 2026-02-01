---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_instance_backup_detacheds"
sidebar_current: "docs-volcengine-datasource-rds_postgresql_instance_backup_detacheds"
description: |-
  Use this data source to query detailed information of rds postgresql instance backup detacheds
---
# volcengine_rds_postgresql_instance_backup_detacheds
Use this data source to query detailed information of rds postgresql instance backup detacheds
## Example Usage
```hcl
data "volcengine_rds_postgresql_instance_backup_detacheds" "example" {
  project_name      = "default"
  backup_status     = "Success"
  backup_type       = "Full"
  backup_start_time = "2025-12-01T00:00:00.000Z"
  backup_end_time   = "2025-12-15T23:59:59.999Z"
}
```
## Argument Reference
The following arguments are supported:
* `backup_end_time` - (Optional) The latest time when the backup is created, in the format of yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).
* `backup_id` - (Optional) The ID of the backup.
* `backup_start_time` - (Optional) The earliest time when the backup is created, in the format of yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).
* `backup_status` - (Optional) The status of the backup.
* `backup_type` - (Optional) The type of the backup.
* `instance_id` - (Optional) The ID of the PostgreSQL instance.
* `instance_name` - (Optional) The name of the PostgreSQL instance.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project to which the instance belongs.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `backups` - List of deleted instance backups.
    * `backup_end_time` - The end time of the backup. The time format is yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).
    * `backup_file_name` - The name of the backup file.
    * `backup_file_size` - The size of the backup file, in Byte.
    * `backup_id` - The ID of the backup.
    * `backup_progress` - The progress of the backup. The unit is percentage.
    * `backup_start_time` - The start time of the backup. The time format is yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).
    * `backup_status` - The status of the backup: Success, Failed, Running.
    * `backup_type` - The type of the backup: Full, Increment.
    * `create_type` - The creation type of the backup: System, User.
    * `instance_info` - Information about the PostgreSQL instance associated with this backup.
        * `db_engine_version` - The version of the database engine.
        * `instance_id` - The ID of the instance.
        * `instance_name` - The name of the instance.
        * `instance_status` - The status of the instance.
* `total_count` - The total count of query.


