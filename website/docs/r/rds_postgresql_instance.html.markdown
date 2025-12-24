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
data "volcengine_rds_postgresql_zones" "foo" {
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

# Restore the backup to a new instance

resource "volcengine_rds_postgresql_instance" "example" {
  src_instance_id   = "postgres-faa4921fdde4"
  backup_id         = "20251215-215628F"
  db_engine_version = "PostgreSQL_12"
  node_spec         = "rds.postgres.1c2g"
  subnet_id         = volcengine_subnet.foo.id
  instance_name     = "acc-test-postgresql-instance-restore"
  charge_info {
    charge_type = "PostPaid"
    number      = 1
  }
  primary_zone_id   = data.volcengine_zones.foo.zones[0].id
  secondary_zone_id = data.volcengine_zones.foo.zones[0].id
}
```
## Argument Reference
The following arguments are supported:
* `charge_info` - (Required) Payment methods.
* `db_engine_version` - (Required, ForceNew) Instance type. Value: PostgreSQL_11, PostgreSQL_12, PostgreSQL_13, PostgreSQL_14, PostgreSQL_15, PostgreSQL_16, PostgreSQL_17.
* `node_spec` - (Required) The specification of primary node and secondary node.
* `primary_zone_id` - (Required, ForceNew) The available zone of primary node.
* `secondary_zone_id` - (Required) The available zone of secondary node.
* `subnet_id` - (Required, ForceNew) Subnet ID of the RDS PostgreSQL instance.
* `allow_list_ids` - (Optional) Allow list IDs to bind at creation.
* `backup_id` - (Optional, ForceNew) Backup ID (choose either this or restore_time; if both are set, backup_id shall prevail).
* `estimate_only` - (Optional) Whether to initiate a configuration change assessment. Only estimate spec change impact without executing. Default value: false.
* `instance_name` - (Optional) Instance name. Cannot start with a number or a dash. Can only contain Chinese characters, letters, numbers, underscores and dashes. The length is limited between 1 ~ 128.
* `modify_type` - (Optional) Spec change type. Usually(default) or Temporary.
* `parameters` - (Optional) Parameter of the RDS PostgreSQL instance. This field can only be added or modified. Deleting this field is invalid.
* `project_name` - (Optional) The project name of the RDS instance.
* `restore_time` - (Optional, ForceNew) The point in time to restore to, in UTC format yyyy-MM-ddTHH:mm:ssZ (choose either this or backup_id).
* `rollback_time` - (Optional) Rollback time for Temporary change, UTC format yyyy-MM-ddTHH:mm:ss.sssZ.
* `src_instance_id` - (Optional, ForceNew) Source instance ID. After setting it, a new instance will be created by restoring from the backup/time point.
* `storage_space` - (Optional) Instance storage space. Value range: [20, 3000], unit: GB, step 10GB. Default value: 100.
* `tags` - (Optional) Tags.
* `zone_migrations` - (Optional) Nodes to migrate AZ. Only Secondary or ReadOnly nodes are allowed. If you want to migrate the availability zone of the secondary node, you need to add the zone_migrations field. Modifying the secondary_zone_id directly will not work. Cross-AZ instance migration is not supported.

The `charge_info` object supports the following:

* `charge_type` - (Required) Payment type. Value:
PostPaid - Pay-As-You-Go
PrePaid - Yearly and monthly (default). 
When the value of this field is `PrePaid`, the postgresql instance cannot be deleted through terraform. Please unsubscribe the instance from the Volcengine console first, and then use `terraform state rm volcengine_rds_postgresql_instance.resource_name` command to remove it from terraform state file and management.
* `auto_renew` - (Optional) Whether to automatically renew in prepaid scenarios.
* `number` - (Optional) Purchase number of the RDS PostgreSQL instance. Range: [1, 20]. Default: 1.
* `period_unit` - (Optional) The purchase cycle in the prepaid scenario.
Month - monthly subscription (default)
Year - Package year.
* `period` - (Optional) Purchase duration in prepaid scenarios. Default: 1.

The `parameters` object supports the following:

* `name` - (Required) Parameter name.
* `value` - (Required) Parameter value.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

The `zone_migrations` object supports the following:

* `node_id` - (Required) Node ID to migrate.
* `zone_id` - (Required) Target zone ID.
* `node_type` - (Optional) Node type: Secondary or ReadOnly.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `allow_list_version` - The allow list version of the RDS PostgreSQL instance.
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
unpaid - unpaid.
    * `charge_type` - Payment type. Value:
PostPaid - Pay-As-You-Go
PrePaid - Yearly and monthly (default).
    * `number` - The number of the RDS PostgreSQL instance.
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
        * `cross_region_domain` - Address that can be accessed across regions.
        * `dns_visibility` - Whether to enable public network resolution. Values: false: Default value. PrivateZone of Volcano Engine. true: Private and public network resolution of Volcano Engine.
        * `domain_visibility_setting` - The type of private network address. Values: LocalDomain: Local domain name. CrossRegionDomain: Domains accessible across regions.
        * `domain` - Connect domain name.
        * `eip_id` - The ID of the EIP, only valid for Public addresses.
        * `internet_protocol` - Address IP protocol, IPv4 or IPv6.
        * `ip_address` - The IP Address.
        * `ipv6_address` - The IPv6 Address.
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
    * `read_only_node_distribution_type` - The distribution type of the read-only nodes, value:
Default: Default distribution.
Custom: Custom distribution.
    * `read_only_node_max_delay_time` - Maximum latency threshold of read-only node. If the latency of a read-only node exceeds this value, reading traffic won't be routed to this node. Unit: seconds.Values: 0~3600.Default value: 30.
    * `read_only_node_weight` - The list of nodes configured by the connection terminal and the corresponding read-only weights.
        * `node_id` - The ID of the node.
        * `node_type` - The type of the node.
        * `weight` - The weight of the node.
    * `read_write_mode` - Read and write mode:
ReadWrite: read and write
ReadOnly: read only (default).
    * `read_write_proxy_connection` - After the terminal enables read-write separation, the number of proxy connections set for the terminal. The lower limit of the number of proxy connections is 20. The upper limit of the number of proxy connections depends on the specifications of the instance master node.
    * `write_node_halt_writing` - Whether the endpoint sends write requests to the write node (currently only the master node is a write node). Values: true: Yes(Default). false: No.
* `estimation_result` - The estimated impact on the instance after the current configuration changes.
    * `effects` - After changing according to the current configuration, the estimated impact on the read and write connections of the instance.
    * `plans` - Estimated impact on the instance after the current configuration changes.
* `instance_id` - The ID of the RDS PostgreSQL instance.
* `instance_status` - The status of the RDS PostgreSQL instance.
* `instance_type` - The instance type of the RDS PostgreSQL instance.
* `memory` - Memory size. Unit: GB.
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
* `storage_data_use` - The instance's primary node has used storage space. Unit: Byte.
* `storage_log_use` - The instance's primary node has used log storage space. Unit: Byte.
* `storage_temp_use` - The instance's primary node has used temporary storage space. Unit: Byte.
* `storage_type` - Instance storage type.
* `storage_use` - The instance has used storage space. Unit: Byte.
* `storage_wal_use` - The instance's primary node has used WAL storage space. Unit: Byte.
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

