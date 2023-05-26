---
subcategory: "REDIS"
layout: "volcengine"
page_title: "Volcengine: volcengine_redis_backups"
sidebar_current: "docs-volcengine-datasource-redis_backups"
description: |-
  Use this data source to query detailed information of redis backups
---
# volcengine_redis_backups
Use this data source to query detailed information of redis backups
## Example Usage
```hcl
data "volcengine_redis_backups" "default" {
  instance_id          = "redis-cnlfvrv4qye6u4lpa"
  backup_strategy_list = ["ManualBackup"]
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) Id of instance.
* `backup_strategy_list` - (Optional) The list of backup strategy, support AutomatedBackup and ManualBackup.
* `end_time` - (Optional) Query end time.
* `output_file` - (Optional) File name where to save data source results.
* `start_time` - (Optional) Query start time.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `backups` - Information of backups.
    * `backup_point_id` - The id of backup point.
    * `backup_strategy` - Backup strategy.
    * `backup_type` - Backup type.
    * `end_time` - End time of backup.
    * `instance_detail` - Information of instance.
        * `account_id` - Id of account.
        * `arch_type` - Arch type of instance(Standard/Cluster).
        * `charge_type` - Charge type of instance(Postpaid/Prepaid).
        * `engine_version` - Engine version of instance.
        * `expired_time` - Expired time of instance.
        * `instance_id` - Id of instance.
        * `instance_name` - Name of instance.
        * `maintenance_time` - The maintainable period (in UTC) of the instance.
        * `network_type` - Network type of instance.
        * `project_name` - Project name of instance.
        * `region_id` - Id of region.
        * `replicas` - Count of replica in which shard.
        * `server_cpu` - Count of cpu cores of instance.
        * `shard_capacity` - Capacity of shard.
        * `shard_count` - Count of shard.
        * `total_capacity` - Total capacity of instance.
        * `used_capacity` - Capacity used of this instance.
        * `vpc_info` - Information of vpc.
            * `id` - Id of vpc.
            * `name` - Name of vpc.
        * `zone_ids` - List of id of zone.
    * `instance_id` - Id of instance.
    * `size` - Size in MiB.
    * `start_time` - Start time of backup.
    * `status` - Status of backup (Creating/Available/Unavailable/Deleting).
* `total_count` - The total count of backup query.


