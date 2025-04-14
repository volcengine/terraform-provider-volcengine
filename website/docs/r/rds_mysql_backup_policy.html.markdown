---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_mysql_backup_policy"
sidebar_current: "docs-volcengine-resource-rds_mysql_backup_policy"
description: |-
  Provides a resource to manage rds mysql backup policy
---
# volcengine_rds_mysql_backup_policy
Provides a resource to manage rds mysql backup policy
## Example Usage
```hcl
resource "volcengine_rds_mysql_backup_policy" "foo" {
  instance_id               = "mysql-c8c3f45c4b07"
  data_full_backup_periods  = ["Monday", "Sunday"]
  binlog_file_counts_enable = true
  binlog_space_limit_enable = true
  lock_ddl_time             = 80
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required, ForceNew) The ID of the RDS instance.
* `binlog_backup_all_retention` - (Optional) Whether to retain all log backups before releasing an instance. Values:
true: Yes.
false: No. Description: BinlogBackupAllRetention is ineffective when the value of RetentionPolicySynced is true.
* `binlog_backup_enabled` - (Optional) Whether to enable log backup function. Values:
true: Yes.
false: No.
* `binlog_backup_encryption_enabled` - (Optional) Is encryption enabled for log backups? Values:
true: Yes.
false: No.
* `binlog_file_counts_enable` - (Optional) Whether to enable the upper limit of local Binlog retention. Values: true: Enabled. false: Disabled. Description:When modifying the log backup policy, this parameter needs to be passed in.
* `binlog_limit_count` - (Optional) Number of local Binlog retained, ranging from 6 to 1000, in units of pieces. Automatically delete local logs that exceed the retained number. Explanation: When modifying the log backup policy, this parameter needs to be passed in.
* `binlog_local_retention_hour` - (Optional) Local Binlog retention duration, with a value ranging from 0 to 168, in hours. Local logs exceeding the retention duration will be automatically deleted. When set to 0, local logs will not be automatically deleted. Note: When modifying the log backup policy, this parameter needs to be passed.
* `binlog_space_limit_enable` - (Optional) Whether to enable automatic cleanup of Binlog when space is too large. When the total storage space occupancy rate of the instance exceeds 80% or the remaining space is less than 5GB, the system will automatically start cleaning up the earliest local Binlog until the total space occupancy rate is lower than 80% and the remaining space is greater than 5GB. true: Enabled. false: Disabled. Description: This parameter needs to be passed in when modifying the log backup policy.
* `binlog_storage_percentage` - (Optional) Maximum storage space usage rate can be set to 20% - 50%. After exceeding this limit, the earliest Binlog file will be automatically deleted until the space usage rate is lower than this ratio. Local Binlog space usage rate = Local Binlog size / Total available (purchased) instance space size. When modifying the log backup policy, this parameter needs to be passed in. Explanation: When modifying the log backup policy, this parameter needs to be passed in.
* `data_backup_all_retention` - (Optional) Whether to retain all data backups before releasing the instance. Values:
true: Yes.
false: No.
* `data_backup_encryption_enabled` - (Optional) Whether to enable encryption for data backup. Values:
true: Yes.
false: No.
* `data_backup_retention_day` - (Optional) Data backup retention days, value range: 7 to 365 days. Default retention is 7 days.
* `data_full_backup_periods` - (Optional) Full backup period. It is recommended to select at least 2 days for full backup every week. Multiple values are separated by English commas (,). Values: Monday. Tuesday. Wednesday. Thursday. Friday. Saturday. Sunday. When modifying the data backup policy, this parameter needs to be passed in.
* `data_full_backup_start_utc_hour` - (Optional) The start point (UTC time) of the time window for starting the full backup task. The time window length is 1 hour. Explanation: Both DataFullBackupStartUTCHour and DataFullBackupTime can be used to indicate the full backup time period of an instance. DataFullBackupStartUTCHour has higher priority. If both fields are returned at the same time, DataFullBackupStartUTCHour shall prevail.
* `data_full_backup_time` - (Optional) Time window for executing backup tasks is one hour. Format: HH:mmZ-HH:mmZ (UTC time). Explanation: This parameter needs to be passed in when modifying the data backup policy.
* `data_incr_backup_periods` - (Optional) Incremental backup period. Multiple values are separated by commas (,). Values: Monday. Tuesday. Wednesday. Thursday. Friday. Saturday. Sunday.Description: The incremental backup period cannot conflict with the full backup. When modifying the data backup policy, this parameter needs to be passed in.
* `data_keep_days_after_released` - (Optional) Backup retention days when an instance is released. Currently, only a value of 7 is supported.
* `data_keep_policy_after_released` - (Optional) Policy for retaining a backup of an instance after it is released. The values are: Last: Keep the last backup. Default value. All: Keep all backups of the instance.
* `hourly_incr_backup_enable` - (Optional) Whether to enable high-frequency backup function. Values:
true: Yes.
false: No.
* `incr_backup_hour_period` - (Optional) Frequency of performing high-frequency incremental backups. Values: 2: Perform an incremental backup every 2 hours. 4: Perform an incremental backup every 4 hours. 6: Perform an incremental backup every 6 hours. 12: Perform an incremental backup every 12 hours. Description: This parameter takes effect only when HourlyIncrBackupEnable is set to true.
* `lock_ddl_time` - (Optional) Maximum waiting time for DDL. The default value is 30. The minimum value is 10. The maximum value is 1440. The unit is minutes. Description: Only instances of MySQL 8.0 version support this setting.
* `log_backup_retention_day` - (Optional) Binlog backup retention period. The value range is 7 to 365, in days. Explanation: When modifying the log backup policy, this parameter needs to be passed in.
* `retention_policy_synced` - (Optional) Is the retention policy for log backups the same as that for data backups?
Explanation: When the value is true, LogBackupRetentionDay and BinlogBackupAllRetention are ignored.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
RdsMysqlBackupPolicy can be imported using the id, e.g.
```
$ terraform import volcengine_rds_mysql_backup_policy.default instanceId:backupPolicy
```
Warning:The resource cannot be deleted, and the destroy operation will not perform any actions.

