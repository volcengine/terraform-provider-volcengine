---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_mysql_instances"
sidebar_current: "docs-volcengine-datasource-rds_mysql_instances"
description: |-
  Use this data source to query detailed information of rds mysql instances
---
# volcengine_rds_mysql_instances
Use this data source to query detailed information of rds mysql instances
## Example Usage
```hcl
data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-project1"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-subnet-test-2"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_rds_mysql_instance" "foo" {
  db_engine_version      = "MySQL_5_7"
  node_spec              = "rds.mysql.1c2g"
  primary_zone_id        = data.volcengine_zones.foo.zones[0].id
  secondary_zone_id      = data.volcengine_zones.foo.zones[0].id
  storage_space          = 80
  subnet_id              = volcengine_subnet.foo.id
  instance_name          = "acc-test"
  lower_case_table_names = "1"

  charge_info {
    charge_type = "PostPaid"
  }

  parameters {
    parameter_name  = "auto_increment_increment"
    parameter_value = "2"
  }
  parameters {
    parameter_name  = "auto_increment_offset"
    parameter_value = "4"
  }
}

data "volcengine_rds_mysql_instances" "foo" {
  instance_id = volcengine_rds_mysql_instance.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `charge_type` - (Optional) The charge type of the RDS instance.
* `create_time_end` - (Optional) The end time of creating RDS instance.
* `create_time_start` - (Optional) The start time of creating RDS instance.
* `db_engine_version` - (Optional) The version of the RDS instance.
* `instance_id` - (Optional) The id of the RDS instance.
* `instance_name` - (Optional) The name of the RDS instance.
* `instance_status` - (Optional) The status of the RDS instance.
* `name_regex` - (Optional) A Name Regex of RDS instance.
* `output_file` - (Optional) File name where to save data source results.
* `tags` - (Optional) Tags.
* `zone_id` - (Optional) The available zone of the RDS instance.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `rds_mysql_instances` - The collection of RDS instance query.
    * `allow_list_version` - The version of allow list.
    * `backup_use` - The instance has used backup space. Unit: GB.
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
        * `temp_modify_end_time` - Restore time of temporary upgrade.
        * `temp_modify_start_time` - Temporary upgrade start time.
    * `create_time` - The create time of the RDS instance.
    * `data_sync_mode` - Data synchronization mode.
    * `db_engine_version` - The engine version of the RDS instance.
    * `endpoints` - The endpoint info of the RDS instance.
        * `addresses` - Address list.
            * `dns_visibility` - DNS Visibility.
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
        * `idle_connection_reclaim` - Whether the idle connection reclaim function is enabled. true: Enabled. false: Disabled.
        * `node_weight` - The list of nodes configured by the connection terminal and the corresponding read-only weights.
            * `node_id` - The ID of the node.
            * `node_type` - The type of the node.
            * `weight` - The weight of the node.
        * `read_write_mode` - Read and write mode:
ReadWrite: read and write
ReadOnly: read only (default).
    * `id` - The ID of the RDS instance.
    * `instance_id` - The ID of the RDS instance.
    * `instance_name` - The name of the RDS instance.
    * `instance_status` - The status of the RDS instance.
    * `lower_case_table_names` - Whether the table name is case sensitive, the default value is 1.
Ranges:
0: Table names are stored as fixed and table names are case-sensitive.
1: Table names will be stored in lowercase and table names are not case sensitive.
    * `maintenance_window` - Maintenance Window.
        * `day_kind` - DayKind of maintainable window. Value: Week. Month.
        * `day_of_month` - Days of maintainable window of the month.
        * `day_of_week` - Days of maintainable window of the week.
        * `maintenance_time` - The maintainable time of the RDS instance.
    * `memory` - Memory size.
    * `node_cpu_used_percentage` - Average CPU usage of the instance master node in nearly one minute.
    * `node_memory_used_percentage` - Average memory usage of the instance master node in nearly one minute.
    * `node_number` - The number of nodes.
    * `node_space_used_percentage` - Average disk usage of the instance master node in nearly one minute.
    * `node_spec` - The specification of primary node.
    * `nodes` - Instance node information.
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
    * `project_name` - The project name of the RDS instance.
    * `region_id` - The region of the RDS instance.
    * `storage_space` - Total instance storage space. Unit: GB.
    * `storage_type` - Instance storage type.
    * `storage_use` - The instance has used storage space. Unit: GB.
    * `subnet_id` - The subnet ID of the RDS instance.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `time_zone` - Time zone.
    * `update_time` - The update time of the RDS instance.
    * `v_cpu` - CPU size.
    * `vpc_id` - The vpc ID of the RDS instance.
    * `zone_id` - The available zone of the RDS instance.
    * `zone_ids` - List of availability zones where each node of the instance is located.
* `total_count` - The total count of RDS instance query.


