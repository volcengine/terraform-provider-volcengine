---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_backup_policy"
sidebar_current: "docs-volcengine-resource-rds_postgresql_backup_policy"
description: |-
  Provides a resource to manage rds postgresql backup policy
---
# volcengine_rds_postgresql_backup_policy
Provides a resource to manage rds postgresql backup policy
## Example Usage
```hcl
resource "volcengine_rds_postgresql_backup_policy" "example" {
  instance_id                = "postgres-72715e0d9f58"
  backup_retention_period    = 7
  full_backup_period         = "Monday,Wednesday,Friday"
  full_backup_time           = "18:00Z-19:00Z"
  data_incr_backup_periods   = "Tuesday,Sunday"
  hourly_incr_backup_enable  = true
  increment_backup_frequency = 12
  wal_log_space_limit_enable = false
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The id of the PostgreSQL instance.
* `backup_retention_period` - (Optional) The number of days to retain backups, with a value range of 7 to 365.
* `data_incr_backup_periods` - (Optional) The incremental backup method follows the backup frequency for normal increments, with multiple values separated by English commas (,). The selected values must not overlap with the full backup cycle. Can select at most six days a week for incremental backup.
* `full_backup_period` - (Optional) Full backup period. Separate multiple values with an English comma (,).Select at least one day per week for a full backup.
* `full_backup_time` - (Optional) The time when the backup task is executed. Format: HH:mmZ-HH:mmZ (UTC time).
* `hourly_incr_backup_enable` - (Optional) Whether to enable the high-frequency backup function. To disable incremental backup, need to pass an empty string for the parameter data_incr_backup_periods and pass false for the parameter hourly_incr_backup_enable.
* `increment_backup_frequency` - (Optional) The method of incremental backup is the backup frequency for high-frequency increments. The Unit: hours. The valid values are 1, 2, 4, 6, and 12.
* `wal_log_space_limit_enable` - (Optional) Status of the local remaining available space protection function. When enabled, it will automatically start clearing the earliest local WAL logs when the total storage space usage rate of the instance exceeds 80% or the remaining space is less than 5GB, until the total space usage rate is below 80% and the remaining space is greater than 5GB.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
RdsPostgresqlBackupPolicy can be imported using the id, e.g.
```
$ terraform import volcengine_rds_postgresql_backup_policy.default resource_id
```

