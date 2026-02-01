---
subcategory: "RDS_MSSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_mssql_backups"
sidebar_current: "docs-volcengine-datasource-rds_mssql_backups"
description: |-
  Use this data source to query detailed information of rds mssql backups
---
# volcengine_rds_mssql_backups
Use this data source to query detailed information of rds mssql backups
## Example Usage
```hcl
data "volcengine_rds_mssql_backups" "foo" {
  instance_id = "mssql-40914121fd22"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The id of the instance.
* `backup_end_time` - (Optional) The end time of the backup.
* `backup_id` - (Optional) The id of the backup.
* `backup_start_time` - (Optional) The start time of the backup.
* `backup_type` - (Optional) The type of the backup.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `backups` - The collection of query.
    * `backup_database_detail` - The detail of the database.
        * `backup_download_link_eip` - External backup download link.
        * `backup_download_link_inner` - Intranet backup download link.
        * `backup_end_time` - The end time of the backup.
        * `backup_file_name` - The name of the backup file.
        * `backup_file_size` - The size of the backup file.
        * `backup_start_time` - The start time of the backup.
        * `backup_type` - The type of the backup.
        * `database_name` - The name of the database.
        * `download_progress` - Backup file preparation progress, unit: %.
        * `download_status` - Download status.
        * `link_expired_time` - Download link expiration time.
    * `backup_end_time` - The end time of the backup.
    * `backup_file_size` - The size of the backup file.
    * `backup_id` - The id of the backup.
    * `backup_method` - The name of the backup method.
    * `backup_start_time` - The start time of the backup.
    * `backup_status` - The status of the backup.
    * `backup_type` - The type of the backup.
    * `create_type` - The type of the backup create.
    * `id` - The id of the backup.
* `total_count` - The total count of query.


