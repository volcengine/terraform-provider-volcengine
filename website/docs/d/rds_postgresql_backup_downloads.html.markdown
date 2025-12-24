---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_backup_downloads"
sidebar_current: "docs-volcengine-datasource-rds_postgresql_backup_downloads"
description: |-
  Use this data source to query detailed information of rds postgresql backup downloads
---
# volcengine_rds_postgresql_backup_downloads
Use this data source to query detailed information of rds postgresql backup downloads
## Example Usage
```hcl
data "volcengine_rds_postgresql_backup_downloads" "example" {
  instance_id = "postgres-72715e0d9f58"
  backup_id   = "20251214-200431-0698LD"
}
```
## Argument Reference
The following arguments are supported:
* `backup_id` - (Required) The ID of the logical backup to be downloaded.
* `instance_id` - (Required) The id of the PostgreSQL instance.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `downloads` - Download link information (if needed, please trigger the download task first).
    * `backup_description` - The description of the backup set.
    * `backup_download_link` - The public network download address of the backup.
    * `backup_file_name` - The name of the backup file.
    * `backup_file_size` - The size of the backup file, in Byte.
    * `backup_id` - The id of the backup.
    * `backup_method` - The type of the backup.
    * `inner_backup_download_link` - The inner network download address of the backup.
    * `instance_id` - The id of the PostgreSQL instance.
    * `link_expired_time` - Expiration time of the download link, format:yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).
    * `prepare_progress` - The prepare progress of the backup.
* `total_count` - The total count of query.


