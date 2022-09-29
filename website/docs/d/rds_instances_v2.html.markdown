---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_instances_v2"
sidebar_current: "docs-volcengine-datasource-rds_instances_v2"
description: |-
  Use this data source to query detailed information of rds instances v2
---
# volcengine_rds_instances_v2
Use this data source to query detailed information of rds instances v2
## Example Usage
```hcl
data "volcengine_rds_instances_v2" "default" {

}
```
## Argument Reference
The following arguments are supported:
* `charge_type` - (Optional) The charge type of the RDS instance.
* `create_time_end` - (Optional) The end time of creating RDS instance.
* `create_time_start` - (Optional) The start time of creating RDS instance.
* `db_engine_version` - (Optional) The version of the RDS instance, Value:
MySQL Community:
    MySQL_5.7 - MySQL 5.7
    MySQL_8_0 - MySQL 8.0
PostgreSQL Community:
    PostgreSQL_11 - PostgreSQL 11
    PostgreSQL_12 - PostgreSQL 12
Microsoft SQL Server: Not available at this time
    SQLServer_2019 - SQL Server 2019
veDB for MySQL:
    MySQL_8_0 - MySQL 8.0
veDB for PostgreSQL:
    PostgreSQL_13 - PostgreSQL 13.
* `instance_id` - (Optional) The id of the RDS instance.
* `instance_name` - (Optional) The name of the RDS instance.
* `instance_status` - (Optional) The status of the RDS instance, Value:
Running - running
Creating - Creating
Deleting - Deleting
Restarting - Restarting
Restoring - Restoring
Updating - changing
Upgrading - Upgrading
Error - the error.
* `instance_type` - (Optional) The type of the RDS instance, Value:
Value:
RDS for MySQL:
    HA - high availability version;
RDS for PostgreSQL:
    HA - high availability version;
Microsoft SQL Server: Not available at this time
    Enterprise - Enterprise Edition
    Standard - Standard Edition
    Web - Web version
veDB for MySQL:
    Cluster - Cluster Edition
veDB for PostgreSQL:
    Cluster - Cluster Edition
MySQL Sharding:
    HA - high availability version;.
* `name_regex` - (Optional) A Name Regex of RDS instance.
* `output_file` - (Optional) File name where to save data source results.
* `zone_id` - (Optional) The available zone of the RDS instance.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `rds_instances` - The collection of RDS instance query.
    * `charge_detail` - Payment methods.
        * `auto_renew` - Whether to automatically renew in prepaid scenarios.
Autorenew_Enable
Autorenew_Disable (default).
        * `charge_end_time` - Billing expiry time (yearly and monthly only).
        * `charge_start_time` - Billing start time (pay-as-you-go & monthly subscription).
        * `charge_status` - Pay status. Value:
normal - normal
overdue - overdue
.
        * `charge_type` - Payment type. Value:
PostPaid - Pay-As-You-Go
PrePaid - Yearly and monthly (default).
        * `overdue_reclaim_time` - Estimated release time when arrears are closed (pay-as-you-go & monthly subscription).
        * `overdue_time` - Shutdown time in arrears (pay-as-you-go & monthly subscription).
        * `period_unit` - The purchase cycle in the prepaid scenario.
Month - monthly subscription (default)
Year - Package year.
        * `period` - Purchase duration in prepaid scenarios. Default: 1.
    * `connection_info` - The connection info ot the RDS instance.
        * `address` - Address list.
            * `domain` - Connect domain name.
            * `eip_id` - The ID of the EIP, only valid for Public addresses.
            * `ip_address` - The IP Address.
            * `network_type` - Network address type, temporarily Private, Public, PublicService.
            * `port` - The Port.
            * `subnet_id` - Subnet ID, valid only for private addresses.
        * `auto_add_new_nodes` - When the terminal type is read-write terminal or read-only terminal, it supports setting whether new nodes are automatically added.
        * `description` - Address description.
        * `enable_read_only` - Whether global read-only is enabled, value: Enable: Enable. Disable: Disabled.
        * `enable_read_write_splitting` - Whether read-write separation is enabled, value: Enable: Enable. Disable: Disabled.
        * `endpoint_id` - Instance connection terminal ID.
        * `endpoint_name` - The instance connection terminal name.
        * `endpoint_type` - Terminal type:
Cluster: The default terminal. (created by default)
Primary: Primary node terminal.
Custom: Custom terminal.
Direct: Direct connection to the terminal. (Only the operation and maintenance side)
AllNode: All node terminals. (Only the operation and maintenance side).
        * `read_only_node_weight` - The list of nodes configured by the connection terminal and the corresponding read-only weights.
            * `node_id` - The ID of the node.
            * `node_type` - The type of the node.
            * `weight` - The weight of the node.
        * `read_write_mode` - Read and write mode:
ReadWrite: read and write
ReadOnly: read only (default).
    * `create_time` - The create time of the RDS instance.
    * `db_engine_version` - The engine version of the RDS instance.
    * `db_engine` - The engine of the RDS instance.
    * `id` - The ID of the RDS instance.
    * `instance_id` - The ID of the RDS instance.
    * `instance_name` - The name of the RDS instance.
    * `instance_status` - The status of the RDS instance.
    * `instance_type` - The type of the RDS instance.
    * `node_detail_info` - Instance node information.
        * `create_time` - Node creation local time.
        * `instance_id` - Instance ID.
        * `memory` - Memory size in GB.
        * `node_id` - Node ID.
        * `node_spec` - General instance type, different from Custom instance type.
        * `node_status` - Node state, value: aligned with instance state.
        * `node_type` - Node type. Value: Primary: Primary node.
Secondary: Standby node.
ReadOnly: Read-only node.
        * `region_id` - Region ID, you can call the DescribeRegions query and use this parameter to specify the region where the instance is to be created.
        * `update_time` - Node updates local time.
        * `v_cpu` - CPU size. For example: 1 means 1U.
        * `zone_id` - Availability zone ID. Subsequent support for multi-availability zones can be separated and displayed by an English colon.
    * `node_number` - The number of nodes.
    * `node_spec` - General instance type, different from Custom instance type.
    * `port` - Instance intranet port.
    * `project_name` - Subordinate to the project.
    * `region_id` - The region of the RDS instance.
    * `shard_number` - The number of shards.
    * `storage_space` - Total instance storage space. Unit: GB.
    * `storage_type` - Instance storage type. When the database type is MySQL/PostgreSQL/SQL_Server/MySQL Sharding, the value is:
LocalSSD - local SSD disk
When the database type is veDB_MySQL/veDB_PostgreSQL, the value is:
DistributedStorage - Distributed Storage.
    * `storage_use` - The instance has used storage space. Unit: GB.
    * `subnet_id` - The subnet ID of the RDS instance.
    * `time_zone` - Time zone.
    * `vpc_id` - The vpc ID of the RDS instance.
    * `zone_id` - The available zone of the RDS instance.
* `total_count` - The total count of RDS instance query.


