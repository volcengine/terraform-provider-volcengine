---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_mysql_databases"
sidebar_current: "docs-volcengine-datasource-rds_mysql_databases"
description: |-
  Use this data source to query detailed information of rds mysql databases
---
# volcengine_rds_mysql_databases
Use this data source to query detailed information of rds mysql databases
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

resource "volcengine_rds_mysql_database" "foo" {
  db_name     = "acc-test"
  instance_id = volcengine_rds_mysql_instance.foo.id
}
data "volcengine_rds_mysql_databases" "foo" {
  db_name     = "acc-test"
  instance_id = volcengine_rds_mysql_instance.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The id of the RDS instance.
* `db_name` - (Optional) The name of the RDS database.
* `name_regex` - (Optional) A Name Regex of RDS database.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `databases` - The collection of RDS instance account query.
    * `character_set_name` - The character set of the RDS database.
    * `database_privileges` - The privilege detail list of RDS mysql instance database.
        * `account_name` - The name of account.
        * `account_privilege_detail` - The privilege detail of the account.
        * `account_privilege` - The privilege type of the account.
    * `db_name` - The name of the RDS database. This field supports fuzzy queries.
* `total_count` - The total count of RDS database query.


