---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_instance"
sidebar_current: "docs-volcengine-resource-rds_postgresql_instance"
description: |-
  Provides a resource to manage rds postgresql instance
---
# volcengine_rds_postgresql_instance
Provides a resource to manage rds postgresql instance
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

# create postgresql instance
resource "volcengine_rds_postgresql_instance" "foo" {
  db_engine_version = "PostgreSQL_12"
  node_spec         = "rds.postgres.1c2g"
  primary_zone_id   = data.volcengine_zones.foo.zones[0].id
  secondary_zone_id = data.volcengine_zones.foo.zones[0].id
  storage_space     = 40
  subnet_id         = volcengine_subnet.foo.id
  instance_name     = "acc-test-postgresql-instance"
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

# create postgresql instance readonly node
resource "volcengine_rds_postgresql_instance_readonly_node" "foo" {
  instance_id = volcengine_rds_postgresql_instance.foo.id
  node_spec   = "rds.postgres.1c2g"
  zone_id     = data.volcengine_zones.foo.zones[0].id
}

# create postgresql allow list
resource "volcengine_rds_postgresql_allowlist" "foo" {
  allow_list_name = "acc-test-allowlist"
  allow_list_desc = "acc-test"
  allow_list_type = "IPv4"
  allow_list      = ["192.168.0.0/24", "192.168.1.0/24"]
}

# associate postgresql allow list to postgresql instance
resource "volcengine_rds_postgresql_allowlist_associate" "foo" {
  instance_id   = volcengine_rds_postgresql_instance.foo.id
  allow_list_id = volcengine_rds_postgresql_allowlist.foo.id
}

# create postgresql database
resource "volcengine_rds_postgresql_database" "foo" {
  db_name     = "acc-test-database"
  instance_id = volcengine_rds_postgresql_instance.foo.id
  c_type      = "C"
  collate     = "zh_CN.utf8"
}

# create postgresql account
resource "volcengine_rds_postgresql_account" "foo" {
  account_name       = "acc-test-account"
  account_password   = "9wc@********12"
  account_type       = "Normal"
  instance_id        = volcengine_rds_postgresql_instance.foo.id
  account_privileges = "Inherit,Login,CreateRole,CreateDB"
}

# create postgresql schema
resource "volcengine_rds_postgresql_schema" "foo" {
  db_name     = volcengine_rds_postgresql_database.foo.db_name
  instance_id = volcengine_rds_postgresql_instance.foo.id
  owner       = volcengine_rds_postgresql_account.foo.account_name
  schema_name = "acc-test-schema"
}
```
## Argument Reference
The following arguments are supported:
* `charge_info` - (Required, ForceNew) Payment methods.
* `db_engine_version` - (Required, ForceNew) Instance type. Value: PostgreSQL_11, PostgreSQL_12, PostgreSQL_13.
* `node_spec` - (Required) The specification of primary node and secondary node.
* `primary_zone_id` - (Required, ForceNew) The available zone of primary node.
* `secondary_zone_id` - (Required, ForceNew) The available zone of secondary node.
* `subnet_id` - (Required, ForceNew) Subnet ID of the RDS PostgreSQL instance.
* `instance_name` - (Optional) Instance name. Cannot start with a number or a dash. Can only contain Chinese characters, letters, numbers, underscores and dashes. The length is limited between 1 ~ 128.
* `parameters` - (Optional) Parameter of the RDS PostgreSQL instance. This field can only be added or modified. Deleting this field is invalid.
* `project_name` - (Optional) The project name of the RDS instance.
* `storage_space` - (Optional) Instance storage space. Value range: [20, 3000], unit: GB, increments every 100GB. Default value: 100.
* `tags` - (Optional) Tags.

The `charge_info` object supports the following:

* `charge_type` - (Required, ForceNew) Payment type. Value:
PostPaid - Pay-As-You-Go
PrePaid - Yearly and monthly (default). 
When the value of this field is `PrePaid`, the postgresql instance cannot be deleted through terraform. Please unsubscribe the instance from the Volcengine console first, and then use `terraform state rm volcengine_rds_postgresql_instance.resource_name` command to remove it from terraform state file and management.
* `auto_renew` - (Optional, ForceNew) Whether to automatically renew in prepaid scenarios.
* `period_unit` - (Optional, ForceNew) The purchase cycle in the prepaid scenario.
Month - monthly subscription (default)
Year - Package year.
* `period` - (Optional, ForceNew) Purchase duration in prepaid scenarios. Default: 1.

The `parameters` object supports the following:

* `name` - (Required) Parameter name.
* `value` - (Required) Parameter value.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
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
* `instance_id` - The ID of the RDS PostgreSQL instance.
* `instance_status` - The status of the RDS PostgreSQL instance.
* `instance_type` - The instance type of the RDS PostgreSQL instance.
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
* `region_id` - The region of the RDS PostgreSQL instance.
* `storage_type` - Instance storage type.
* `update_time` - The update time of the RDS PostgreSQL instance.
* `v_cpu` - CPU size.
* `vpc_id` - The vpc ID of the RDS PostgreSQL instance.
* `zone_id` - The available zone of the RDS PostgreSQL instance.
* `zone_ids` - ID of the availability zone where each instance is located.


## Import
RdsPostgresqlInstance can be imported using the id, e.g.
```
$ terraform import volcengine_rds_postgresql_instance.default postgres-21a3333b****
```

