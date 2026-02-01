---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_instance_backup_wal_logs"
sidebar_current: "docs-volcengine-datasource-rds_postgresql_instance_backup_wal_logs"
description: |-
  Use this data source to query detailed information of rds postgresql instance backup wal logs
---
# volcengine_rds_postgresql_instance_backup_wal_logs
Use this data source to query detailed information of rds postgresql instance backup wal logs
## Example Usage
```hcl
data "volcengine_rds_postgresql_instance_backup_wal_logs" "example" {
  instance_id = "postgres-ac541555dd74"
  backup_id   = "000000030000000E00000006"
  start_time  = "2025-12-10T00:00:00Z"
  end_time    = "2025-12-15T23:59:59Z"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The id of the PostgreSQL instance.
* `backup_id` - (Optional) The id of the backup.
* `end_time` - (Optional) The end time of the query. The format is yyyy-MM-ddTHH:mm:ssZ (UTC time). Note: The maximum interval between start_time and end_time cannot exceed 7 days.
* `output_file` - (Optional) File name where to save data source results.
* `start_time` - (Optional) The start time of the query. The format is yyyy-MM-ddTHH:mm:ssZ (UTC time).

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of query.
* `wal_log_backups` - List of WAL log backups.
    * `backup_file_size` - The size of the WAL log backup file. The unit is bytes (Byte).
    * `backup_id` - The ID of the WAL log backup.
    * `backup_status` - The status of the WAL log backup.
    * `check_sum` - The checksum in the ETag format using the crc64 algorithm.
    * `download_status` - The downloadable status of the WAL log backup.
    * `project_name` - The project to which the instance of the WAL log backup belongs.
    * `wal_log_backup_end_time` - The end time of the WAL log backup, in the format of yyyy-MM-ddTHH:mm:ssZ (UTC time).


