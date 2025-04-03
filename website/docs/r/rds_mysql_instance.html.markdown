---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_mysql_instance"
sidebar_current: "docs-volcengine-resource-rds_mysql_instance"
description: |-
  Provides a resource to manage rds mysql instance
---
# volcengine_rds_mysql_instance
Provides a resource to manage rds mysql instance
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
    parameter_value = "5"
  }

  #  maintenance_window {
  #    maintenance_time = "18:00Z-21:59Z"
  #    day_kind = "Week"
  #    day_of_week = ["Monday","Tuesday","Wednesday","Thursday","Friday","Saturday","Sunday"]
  #  }

  project_name = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `charge_info` - (Required, ForceNew) Payment methods.
* `db_engine_version` - (Required, ForceNew) Instance type. Value:
MySQL_5_7
MySQL_8_0.
* `node_spec` - (Required) The specification of primary node and secondary node.
* `primary_zone_id` - (Required, ForceNew) The available zone of primary node.
* `secondary_zone_id` - (Required, ForceNew) The available zone of secondary node.
* `subnet_id` - (Required, ForceNew) Subnet ID of the RDS instance.
* `allow_list_ids` - (Optional) Allow list Ids of the RDS instance.
* `connection_pool_type` - (Optional) Connection pool type. Value range:
Direct: Direct connection mode.
Transaction: Transaction-level connection pool (default).
* `db_time_zone` - (Optional, ForceNew) Time zone. Support UTC -12:00 ~ +13:00. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `instance_name` - (Optional) Instance name. Cannot start with a number or a dash
Can only contain Chinese characters, letters, numbers, underscores and dashes
The length is limited between 1 ~ 128.
* `lower_case_table_names` - (Optional, ForceNew) Whether the table name is case sensitive, the default value is 1.
Ranges:
0: Table names are stored as fixed and table names are case-sensitive.
1: Table names will be stored in lowercase and table names are not case sensitive.
* `maintenance_window` - (Optional) Specify the maintainable time period of the instance when creating the instance. This field is optional. If not set, it defaults to 18:00Z - 21:59Z of every day within a week (that is, 02:00 - 05:59 Beijing time).
* `parameters` - (Optional) Parameter of the RDS instance. This field can only be added or modified. Deleting this field is invalid.
* `project_name` - (Optional) The project name of the RDS instance.
* `storage_space` - (Optional) Instance storage space. Value range: [20, 3000], unit: GB, increments every 100GB. Default value: 100.
* `tags` - (Optional) Tags.

The `charge_info` object supports the following:

* `charge_type` - (Required, ForceNew) Payment type. Value:
PostPaid - Pay-As-You-Go
PrePaid - Yearly and monthly (default). 
When the value of this field is `PrePaid`, the mysql instance cannot be deleted through terraform. Please unsubscribe the instance from the Volcengine console first, and then use `terraform state rm volcengine_rds_mysql_instance.resource_name` command to remove it from terraform state file and management.
* `auto_renew` - (Optional, ForceNew) Whether to automatically renew in prepaid scenarios.
* `period_unit` - (Optional, ForceNew) The purchase cycle in the prepaid scenario.
Month - monthly subscription (default)
Year - Package year.
* `period` - (Optional, ForceNew) Purchase duration in prepaid scenarios. Default: 1.

The `maintenance_window` object supports the following:

* `day_kind` - (Optional) Maintenance cycle granularity, values: Week: Week. Month: Month.
* `day_of_week` - (Optional) Specify the maintainable time period of a certain day of the week. The values are: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday. Multiple selections are allowed. If this value is not specified or is empty, it defaults to specifying all seven days of the week.
* `maintenance_time` - (Optional) Maintenance period of an instance. Format: HH:mmZ-HH:mmZ (UTC time).

The `parameters` object supports the following:

* `parameter_name` - (Required) Parameter name.
* `parameter_value` - (Required) Parameter value.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `allow_list_version` - The version of allow list.
* `backup_use` - The instance has used backup space. Unit: GB.
* `binlog_dump` - Does it support the binlog capability? This parameter is returned only when the database proxy is enabled. Values:
true: Yes.
false: No.
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
* `db_proxy_status` - The running status of the proxy instance. This parameter is returned only when the database proxy is enabled. Values:
Creating: The proxy is being started.
Running: The proxy is running.
Shutdown: The proxy is closed.
Deleting: The proxy is being closed.
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
* `feature_states` - Feature status.
    * `enable` - Whether it is enabled. Values:
true: Enabled.
false: Disabled.
    * `feature_name` - Feature name.
    * `support` - Whether it support this function. Value:
true: Supported.
false: Not supported.
* `global_read_only` - Whether to enable global read-only.
true: Yes.
false: No.
* `instance_id` - The ID of the RDS instance.
* `instance_status` - The status of the RDS instance.
* `memory` - Memory size.
* `node_cpu_used_percentage` - Average CPU usage of the instance master node in nearly one minute.
* `node_memory_used_percentage` - Average memory usage of the instance master node in nearly one minute.
* `node_number` - The number of nodes.
* `node_space_used_percentage` - Average disk usage of the instance master node in nearly one minute.
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
* `region_id` - The region of the RDS instance.
* `storage_type` - Instance storage type.
* `storage_use` - The instance has used storage space. Unit: GB.
* `time_zone` - Time zone.
* `update_time` - The update time of the RDS instance.
* `v_cpu` - CPU size.
* `vpc_id` - The vpc ID of the RDS instance.
* `zone_id` - The available zone of the RDS instance.
* `zone_ids` - List of availability zones where each node of the instance is located.


## Import
Rds Mysql Instance can be imported using the id, e.g.
```
$ terraform import volcengine_rds_mysql_instance.default mysql-72da4258c2c7
```

