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
resource "volcengine_rds_mysql_instance" "foo" {
  db_engine_version      = "MySQL_5_7"
  node_spec              = "rds.mysql.1c2g"
  primary_zone_id        = "cn-guilin-a"
  secondary_zone_id      = "cn-guilin-b"
  storage_space          = 80
  subnet_id              = "subnet-2d72yi377stts58ozfdrlk9f6"
  instance_name          = "tf-test"
  lower_case_table_names = "1"

  charge_info {
    charge_type = "PostPaid"
  }

  allow_list_ids = ["acl-2dd8f8317e4d4159b21630d13ae2e6ec", "acl-2eaa2a053b2a4a58b988e38ae975e81c"]

  parameters {
    parameter_name  = "auto_increment_increment"
    parameter_value = "2"
  }
  parameters {
    parameter_name  = "auto_increment_offset"
    parameter_value = "4"
  }
}

resource "volcengine_rds_mysql_instance_readonly_node" "readonly" {
  instance_id = volcengine_rds_mysql_instance.foo.id
  node_spec   = "rds.mysql.2c4g"
  zone_id     = "cn-guilin-a"
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
* `db_time_zone` - (Optional, ForceNew) Time zone. Support UTC -12:00 ~ +13:00. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `instance_name` - (Optional) Instance name. Cannot start with a number or a dash
Can only contain Chinese characters, letters, numbers, underscores and dashes
The length is limited between 1 ~ 128.
* `lower_case_table_names` - (Optional, ForceNew) Whether the table name is case sensitive, the default value is 1.
Ranges:
0: Table names are stored as fixed and table names are case-sensitive.
1: Table names will be stored in lowercase and table names are not case sensitive.
* `parameters` - (Optional) Parameter of the RDS instance. This field can only be added or modified. Deleting this field is invalid.
* `storage_space` - (Optional) Instance storage space. Value range: [20, 3000], unit: GB, increments every 100GB. Default value: 100.

The `charge_info` object supports the following:

* `charge_type` - (Required, ForceNew) Payment type. Value:
PostPaid - Pay-As-You-Go
PrePaid - Yearly and monthly (default).
* `auto_renew` - (Optional, ForceNew) Whether to automatically renew in prepaid scenarios.
* `period_unit` - (Optional, ForceNew) The purchase cycle in the prepaid scenario.
Month - monthly subscription (default)
Year - Package year.
* `period` - (Optional, ForceNew) Purchase duration in prepaid scenarios. Default: 1.

The `parameters` object supports the following:

* `parameter_name` - (Required) Parameter name.
* `parameter_value` - (Required) Parameter value.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
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
* `create_time` - The create time of the RDS instance.
* `data_sync_mode` - Data synchronization mode.
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
    * `node_weight` - The list of nodes configured by the connection terminal and the corresponding read-only weights.
        * `node_id` - The ID of the node.
        * `node_type` - The type of the node.
        * `weight` - The weight of the node.
    * `read_write_mode` - Read and write mode:
ReadWrite: read and write
ReadOnly: read only (default).
* `instance_id` - The ID of the RDS instance.
* `instance_status` - The status of the RDS instance.
* `maintenance_window` - Maintenance Window.
    * `day_kind` - DayKind of maintainable window. Value: Week. Month.
    * `day_of_month` - Days of maintainable window of the month.
    * `day_of_week` - Days of maintainable window of the week.
    * `maintenance_time` - The maintainable time of the RDS instance.
* `memory` - Memory size.
* `node_number` - The number of nodes.
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


## Import
Rds Mysql Instance can be imported using the id, e.g.
```
$ terraform import volcengine_rds_mysql_instance.default mysql-72da4258c2c7
```

