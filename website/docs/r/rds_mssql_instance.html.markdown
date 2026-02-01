---
subcategory: "RDS_MSSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_mssql_instance"
sidebar_current: "docs-volcengine-resource-rds_mssql_instance"
description: |-
  Provides a resource to manage rds mssql instance
---
# volcengine_rds_mssql_instance
Provides a resource to manage rds mssql instance
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

resource "volcengine_rds_mssql_instance" "foo" {
  db_engine_version      = "SQLServer_2019_Std"
  instance_type          = "HA"
  node_spec              = "rds.mssql.se.ha.d2.2c4g"
  storage_space          = 20
  subnet_id              = [volcengine_subnet.foo.id]
  super_account_password = "Tftest110"
  instance_name          = "acc-test-mssql"
  project_name           = "default"
  charge_info {
    charge_type = "PostPaid"
  }
  tags {
    key   = "k1"
    value = "v1"
  }

  backup_time             = "18:00Z-19:00Z"
  full_backup_period      = ["Monday", "Tuesday"]
  backup_retention_period = 14
}
```
## Argument Reference
The following arguments are supported:
* `charge_info` - (Required, ForceNew) The charge info.
* `db_engine_version` - (Required, ForceNew) The Compatible version. Valid values: `SQLServer_2019_Std`, `SQLServer_2019_Web`, `SQLServer_2019_Ent`.
* `instance_type` - (Required, ForceNew) The Instance type. When the value of the `db_engine_version` is `SQLServer_2019_Std`, the value of this field can be `HA` or `Basic`.When the value of the `db_engine_version` is `SQLServer_2019_Ent`, the value of this field can be `Cluster` or `Basic`.When the value of the `db_engine_version` is `SQLServer_2019_Web`, the value of this field can be `Basic`.
* `node_spec` - (Required, ForceNew) The node specification.
* `storage_space` - (Required, ForceNew) Storage space size, measured in GiB. The range of values is 20GiB to 4000GiB, with a step size of 10GiB.
* `subnet_id` - (Required, ForceNew) The subnet id of the instance node. When creating an instance that includes primary and backup nodes and needs to deploy primary and backup nodes across availability zones, you can specify two subnet_id. By default, the first is the primary node availability zone, and the second is the backup node availability zone.
* `super_account_password` - (Required, ForceNew) The super account password. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `backup_retention_period` - (Optional) Data backup retention days, value range: 7~30. 
This field is valid and required when updating the backup plan of instance.
* `backup_time` - (Optional) The time window for starting the backup task is one hour interval. 
This field is valid and required when updating the backup plan of instance.
* `full_backup_period` - (Optional) Full backup cycle. Multiple values separated by commas. The values are as follows: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday. 
This field is valid and required when updating the backup plan of instance.
* `instance_name` - (Optional, ForceNew) Name of the instance.
* `project_name` - (Optional) The project name.
* `tags` - (Optional) Tags.

The `charge_info` object supports the following:

* `charge_type` - (Required, ForceNew) The charge type. Valid values: `PostPaid`, `PrePaid`.
* `auto_renew` - (Optional, ForceNew) Whether to enable automatic renewal in the prepaid scenario. This parameter can be set when the ChargeType is `Prepaid`.
* `period` - (Optional, ForceNew) Purchase duration in a prepaid scenario. This parameter is required when the ChargeType is `Prepaid`.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Rds Mssql Instance can be imported using the id, e.g.
```
$ terraform import volcengine_rds_mssql_instance.default resource_id
```

