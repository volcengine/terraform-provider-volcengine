---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_instances"
sidebar_current: "docs-volcengine-datasource-rds_postgresql_instances"
description: |-
  Use this data source to query detailed information of rds postgresql instances
---
# volcengine_rds_postgresql_instances
Use this data source to query detailed information of rds postgresql instances
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


resource "volcengine_rds_postgresql_instance" "foo" {
  db_engine_version = "PostgreSQL_12"
  node_spec         = "rds.postgres.1c2g"
  primary_zone_id   = data.volcengine_zones.foo.zones[0].id
  secondary_zone_id = data.volcengine_zones.foo.zones[0].id
  storage_space     = 40
  subnet_id         = volcengine_subnet.foo.id
  instance_name     = "acc-test-1"
  charge_info {
    charge_type = "PostPaid"
  }
  project_name = "default"
  tags {
    key   = "tfk1"
    value = "tfv1"
  }
  parameters {
    name  = "auto_explain.log_analyze"
    value = "off"
  }
  parameters {
    name  = "auto_explain.log_format"
    value = "text"
  }
}

data "volcengine_rds_postgresql_instances" "foo" {
  instance_id = volcengine_rds_postgresql_instance.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `charge_type` - (Optional) The charge type of the RDS instance.
* `create_time_end` - (Optional) The end time of creating RDS PostgreSQL instance.
* `create_time_start` - (Optional) The start time of creating RDS PostgreSQL instance.
* `db_engine_version` - (Optional) The version of the RDS PostgreSQL instance.
* `instance_id` - (Optional) The id of the RDS PostgreSQL instance.
* `instance_name` - (Optional) The name of the RDS PostgreSQL instance.
* `instance_status` - (Optional) The status of the RDS PostgreSQL instance.
* `name_regex` - (Optional) A Name Regex of RDS instance.
* `output_file` - (Optional) File name where to save data source results.
* `tags` - (Optional) Tags.
* `zone_id` - (Optional) The available zone of the RDS PostgreSQL instance.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `instances` - The collection of query.
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
        * `temp_modify_end_time` - Temporary upgrade of restoration time.
        * `temp_modify_start_time` - Start time of temporary upgrade.
    * `create_time` - The create time of the RDS PostgreSQL instance.
    * `data_sync_mode` - Data synchronization mode.
    * `db_engine_version` - The engine version of the RDS PostgreSQL instance.
    * `endpoints` - The endpoint info of the RDS instance.
        * `address` - Address list.
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
        * `read_only_node_weight` - The list of nodes configured by the connection terminal and the corresponding read-only weights.
            * `node_id` - The ID of the node.
            * `node_type` - The type of the node.
            * `weight` - The weight of the node.
        * `read_write_mode` - Read and write mode:
ReadWrite: read and write
ReadOnly: read only (default).
    * `id` - The ID of the RDS PostgreSQL instance.
    * `instance_id` - The ID of the RDS PostgreSQL instance.
    * `instance_name` - The name of the RDS PostgreSQL instance.
    * `instance_status` - The status of the RDS PostgreSQL instance.
    * `instance_type` - The instance type of the RDS PostgreSQL instance.
    * `memory` - Memory size.
    * `node_number` - The number of nodes.
    * `node_spec` - Master node specifications.
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
    * `project_name` - The project name of the RDS PostgreSQL instance.
    * `region_id` - The region of the RDS PostgreSQL instance.
    * `storage_space` - Total instance storage space. Unit: GB.
    * `storage_type` - Instance storage type.
    * `subnet_id` - The subnet ID of the RDS PostgreSQL instance.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `update_time` - The update time of the RDS PostgreSQL instance.
    * `v_cpu` - CPU size.
    * `vpc_id` - The vpc ID of the RDS PostgreSQL instance.
    * `zone_id` - The available zone of the RDS PostgreSQL instance.
    * `zone_ids` - ID of the availability zone where each instance is located.
* `total_count` - The total count of RDS instance query.


