---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_instance_v2"
sidebar_current: "docs-volcengine-resource-rds_instance_v2"
description: |-
  Provides a resource to manage rds instance v2
---
# volcengine_rds_instance_v2
(Deprecated! Recommend use volcengine_rds_mysql_*** replace) Provides a resource to manage rds instance v2
## Example Usage
```hcl
data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_rds_instance_v2" "foo" {
  db_engine_version = "MySQL_5_7"
  node_info {
    node_type = "Primary"
    node_spec = "rds.mysql.2c4g"
    zone_id   = data.volcengine_zones.foo.zones[0].id
  }
  node_info {
    node_type = "Secondary"
    node_spec = "rds.mysql.2c4g"
    zone_id   = data.volcengine_zones.foo.zones[0].id
  }
  storage_type           = "LocalSSD"
  storage_space          = 100
  vpc_id                 = volcengine_vpc.foo.id
  subnet_id              = volcengine_subnet.foo.id
  instance_name          = "tf-test-v2"
  lower_case_table_names = "1"
  charge_info {
    charge_type = "PostPaid"
  }
}
```
## Argument Reference
The following arguments are supported:
* `charge_info` - (Required, ForceNew) Payment methods.
* `db_engine_version` - (Required, ForceNew) Instance type. Value:
MySQL_5_7
MySQL_8_0.
* `node_info` - (Required) Instance specification configuration. This parameter is required for RDS for MySQL, RDS for PostgreSQL and MySQL Sharding. There is one and only one Primary node, one and only one Secondary node, and 0-10 Read-Only nodes.
* `storage_type` - (Required) Instance storage type. When the database type is MySQL/PostgreSQL/SQL_Server/MySQL Sharding, the value is:
LocalSSD - local SSD disk
When the database type is veDB_MySQL/veDB_PostgreSQL, the value is:
DistributedStorage - Distributed Storage.
* `subnet_id` - (Required, ForceNew) Subnet ID.
* `vpc_id` - (Required, ForceNew) Private network (VPC) ID. You can call the DescribeVpcs query and use this parameter to specify the VPC where the instance is to be created.
* `db_param_group_id` - (Optional, ForceNew) Parameter template ID. It only takes effect when the database type is MySQL/PostgreSQL/SQL_Server. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `db_time_zone` - (Optional, ForceNew) Time zone. Support UTC -12:00 ~ +13:00. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `instance_name` - (Optional, ForceNew) Instance name. Cannot start with a number or a dash
Can only contain Chinese characters, letters, numbers, underscores and dashes
The length is limited between 1 ~ 128.
* `instance_type` - (Optional, **Deprecated**) The field instance_type is no longer support. Instance type. Value:
HA: High availability version.
* `lower_case_table_names` - (Optional, ForceNew) Whether the table name is case sensitive, the default value is 1.
Ranges:
0: Table names are stored as fixed and table names are case-sensitive.
1: Table names will be stored in lowercase and table names are not case sensitive.
* `project_name` - (Optional) Subordinate to the project.
* `storage_space` - (Optional) Instance storage space.
When the database type is MySQL/PostgreSQL/SQL_Server/MySQL Sharding, value range: [20, 3000], unit: GB, increments every 100GB.
When the database type is veDB_MySQL/veDB_PostgreSQL, this parameter does not need to be passed.

The `charge_info` object supports the following:

* `charge_type` - (Required, ForceNew) Payment type. Value:
PostPaid - Pay-As-You-Go
PrePaid - Yearly and monthly (default).
* `auto_renew` - (Optional, ForceNew) Whether to automatically renew in prepaid scenarios.
* `period_unit` - (Optional, ForceNew) The purchase cycle in the prepaid scenario.
Month - monthly subscription (default)
Year - Package year.
* `period` - (Optional, ForceNew) Purchase duration in prepaid scenarios. Default: 1.

The `node_info` object supports the following:

* `node_spec` - (Required) Masternode specs. Pass
DescribeDBInstanceSpecs Query the instance specifications that can be sold.
* `node_type` - (Required) Node type, the value is "Primary", "Secondary", "ReadOnly".
* `zone_id` - (Required) Zone ID.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
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


## Import
RDS Instance can be imported using the id, e.g.
```
$ terraform import volcengine_rds_instance_v2.default mysql-42b38c769c4b
```

