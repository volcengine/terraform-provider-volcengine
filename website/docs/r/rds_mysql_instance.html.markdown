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
# query available zones in current region
data "volcengine_zones" "foo" {
}

# create vpc
resource "volcengine_vpc" "foo" {
  vpc_name     = "acc-test-vpc"
  cidr_block   = "172.16.0.0/16"
  dns_servers  = ["8.8.8.8", "114.114.114.114"]
  project_name = "default"
}

# create subnet
resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

# create mysql instance
resource "volcengine_rds_mysql_instance" "foo" {
  db_engine_version      = "MySQL_5_7"
  node_spec              = "rds.mysql.2c4g"
  primary_zone_id        = data.volcengine_zones.foo.zones[0].id
  secondary_zone_id      = data.volcengine_zones.foo.zones[0].id
  storage_space          = 80
  subnet_id              = volcengine_subnet.foo.id
  instance_name          = "acc-test-mysql-instance"
  lower_case_table_names = "1"
  project_name           = "default"
  tags {
    key   = "k1"
    value = "v1"
  }

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
  deletion_protection = "Disabled"
  data_sync_mode      = "SemiSync"
  auto_storage_scaling_config {
    enable_storage_auto_scale = true
    storage_threshold         = 40
    storage_upper_bound       = 110
  }
}

# create mysql instance readonly node
resource "volcengine_rds_mysql_instance_readonly_node" "foo" {
  instance_id = volcengine_rds_mysql_instance.foo.id
  node_spec   = "rds.mysql.2c4g"
  zone_id     = data.volcengine_zones.foo.zones[0].id
}

# create mysql allow list
resource "volcengine_rds_mysql_allowlist" "foo" {
  allow_list_name = "acc-test-allowlist"
  allow_list_desc = "acc-test"
  allow_list_type = "IPv4"
  allow_list      = ["192.168.0.0/24", "192.168.1.0/24"]
}

# associate mysql allow list to mysql instance
resource "volcengine_rds_mysql_allowlist_associate" "foo" {
  allow_list_id = volcengine_rds_mysql_allowlist.foo.id
  instance_id   = volcengine_rds_mysql_instance.foo.id
}

# create mysql database
resource "volcengine_rds_mysql_database" "foo" {
  db_name     = "acc-test-database"
  instance_id = volcengine_rds_mysql_instance.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `charge_info` - (Required) Payment methods.
* `db_engine_version` - (Required, ForceNew) Instance type. Value:
MySQL_5_7
MySQL_8_0.
* `node_spec` - (Required) The specification of primary node and secondary node.
* `primary_zone_id` - (Required, ForceNew) The available zone of primary node.
* `secondary_zone_id` - (Required, ForceNew) The available zone of secondary node.
* `subnet_id` - (Required, ForceNew) Subnet ID of the RDS instance.
* `allow_list_ids` - (Optional) Allow list Ids of the RDS instance.
* `auto_storage_scaling_config` - (Optional) Auto - storage scaling configuration.
* `connection_pool_type` - (Optional) Connection pool type. Value range:
Direct: Direct connection mode.
Transaction: Transaction-level connection pool (default).
* `data_sync_mode` - (Optional) Data synchronization methods:
SemiSync: Semi - synchronous(Default).
Async: Asynchronous.
* `db_time_zone` - (Optional, ForceNew) Time zone. Support UTC -12:00 ~ +13:00. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `deletion_protection` - (Optional) Whether to enable the deletion protection function. Values:
Enabled: Yes.
Disabled: No.
* `global_read_only` - (Optional) Whether to enable global read-only for the instance.
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

The `auto_storage_scaling_config` object supports the following:

* `enable_storage_auto_scale` - (Required) Whether to enable the instance's auto - scaling function. Values:
true: Yes.
false: No. Description: When StorageConfig is used as a request parameter, if the value of EnableStorageAutoScale is false, the StorageThreshold and StorageUpperBound parameters do not need to be passed in.
* `storage_threshold` - (Optional) The proportion of available storage space that triggers automatic expansion. The value range is 10 to 50, and the default value is 10, with the unit being %.
* `storage_upper_bound` - (Optional) The upper limit of the storage space that can be automatically expanded. The lower limit of the value of this field is the instance storage space + 20GB; the upper limit of the value is the upper limit of the storage space value range corresponding to the instance master node specification, with the unit being GB. For detailed information on the selectable storage space value range of different specifications, please refer to Product Specifications.

The `charge_info` object supports the following:

* `charge_type` - (Required) Payment type. Value:
PostPaid - Pay-As-You-Go
PrePaid - Yearly and monthly (default).
* `auto_renew` - (Optional) Whether to automatically renew in prepaid scenarios.
* `period_unit` - (Optional) The purchase cycle in the prepaid scenario.
Month - monthly subscription (default)
Year - Package year.
* `period` - (Optional) Purchase duration in prepaid scenarios. Default: 1.

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
* `auto_upgrade_minor_version` - The upgrade strategy for the minor version of the instance kernel. Values:
Auto: Auto upgrade.
Manual: Manual upgrade.
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
* `db_proxy_status` - The running status of the proxy instance. This parameter is returned only when the database proxy is enabled. Values:
Creating: The proxy is being started.
Running: The proxy is running.
Shutdown: The proxy is closed.
Deleting: The proxy is being closed.
* `dr_dts_task_id` - The ID of the data synchronization task in DTS for the data synchronization link between the primary instance and the disaster recovery instance.
* `dr_dts_task_name` - The name of the DTS data synchronization task for the data synchronization link between the primary instance and the disaster recovery instance.
* `dr_dts_task_status` - The status of the DTS data synchronization task for the data synchronization link between the primary instance and the disaster recovery instance.
* `dr_seconds_behind_master` - The number of seconds that the disaster recovery instance is behind the primary instance.
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
* `instance_id` - The ID of the RDS instance.
* `instance_status` - The status of the RDS instance.
* `kernel_version` - The current kernel version of the RDS instance.
* `master_instance_id` - The ID of the primary instance of the disaster recovery instance.
* `master_instance_name` - The name of the primary instance of the disaster recovery instance.
* `master_region` - The region where the primary instance of the disaster recovery instance is located.
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
* `storage_max_capacity` - The upper limit of the storage space that can be set for automatic expansion. The value is the upper limit of the storage space value range corresponding to the instance master node specification, with the unit being GB. For detailed information on the selectable storage space value ranges of different specifications, please refer to Product Specifications.
* `storage_max_trigger_threshold` - The upper limit of the proportion of available storage space that triggers automatic expansion. When supported, the value is 50%.
* `storage_min_capacity` - The lower limit of the storage space that can be set for automatic expansion. The value is the lower limit of the storage space value range corresponding to the instance master node specification, with the unit being GB. For detailed information on the selectable storage space value ranges of different specifications, please refer to Product Specifications.
* `storage_min_trigger_threshold` - The lower limit of the proportion of available storage space that triggers automatic expansion. When supported, the value is 10%.
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

