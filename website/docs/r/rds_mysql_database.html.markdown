---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_mysql_database"
sidebar_current: "docs-volcengine-resource-rds_mysql_database"
description: |-
  Provides a resource to manage rds mysql database
---
# volcengine_rds_mysql_database
Provides a resource to manage rds mysql database
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
```
## Argument Reference
The following arguments are supported:
* `db_name` - (Required, ForceNew) Name database.
illustrate:
Unique name.
The length is 2~64 characters.
Start with a letter and end with a letter or number.
Consists of lowercase letters, numbers, and underscores (_) or dashes (-).
Database names are disabled [keywords](https://www.volcengine.com/docs/6313/66162).
* `instance_id` - (Required, ForceNew) The ID of the RDS instance.
* `character_set_name` - (Optional, ForceNew) Database character set. Currently supported character sets include: utf8, utf8mb4, latin1, ascii.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Database can be imported using the instanceId:dbName, e.g.
```
$ terraform import volcengine_rds_mysql_database.default mysql-42b38c769c4b:dbname
```

