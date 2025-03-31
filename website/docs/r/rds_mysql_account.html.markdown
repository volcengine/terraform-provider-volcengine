---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_mysql_account"
sidebar_current: "docs-volcengine-resource-rds_mysql_account"
description: |-
  Provides a resource to manage rds mysql account
---
# volcengine_rds_mysql_account
Provides a resource to manage rds mysql account
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

resource "volcengine_rds_mysql_instance" "foo" {
  instance_name          = "acc-test-rds-mysql"
  db_engine_version      = "MySQL_5_7"
  node_spec              = "rds.mysql.1c2g"
  primary_zone_id        = data.volcengine_zones.foo.zones[0].id
  secondary_zone_id      = data.volcengine_zones.foo.zones[0].id
  storage_space          = 80
  subnet_id              = volcengine_subnet.foo.id
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

resource "volcengine_rds_mysql_database" "foo1" {
  db_name     = "acc-test-db1"
  instance_id = volcengine_rds_mysql_instance.foo.id
}

resource "volcengine_rds_mysql_database" "foo" {
  db_name     = "acc-test-db"
  instance_id = volcengine_rds_mysql_instance.foo.id
}

resource "volcengine_rds_mysql_account" "foo" {
  account_name     = "acc-test-account"
  account_password = "93f0cb0614Aab12"
  account_type     = "Normal"
  instance_id      = volcengine_rds_mysql_instance.foo.id
  account_privileges {
    db_name                  = volcengine_rds_mysql_database.foo.db_name
    account_privilege        = "Custom"
    account_privilege_detail = "SELECT,INSERT,UPDATE"
  }
  account_privileges {
    db_name           = volcengine_rds_mysql_database.foo1.db_name
    account_privilege = "DDLOnly"
  }
}
```
## Argument Reference
The following arguments are supported:
* `account_name` - (Required, ForceNew) Database account name. The rules are as follows:
Unique name.
Start with a letter and end with a letter or number.
Consists of lowercase letters, numbers, or underscores (_).
The length is 2~32 characters.
The [keyword list](https://www.volcengine.com/docs/6313/66162) is disabled for database accounts, and certain reserved words, including root, admin, etc., cannot be used.
* `account_password` - (Required) The password of the database account.
Illustrate:
Cannot start with `!` or `@`.
The length is 8~32 characters.
It consists of any three of uppercase letters, lowercase letters, numbers, and special characters.
The special characters are `!@#$%^*()_+-=`. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `account_type` - (Required, ForceNew) Database account type, value:
Super: A high-privilege account. Only one database account can be created for an instance.
Normal: An account with ordinary privileges.
* `instance_id` - (Required, ForceNew) The ID of the RDS instance.
* `account_privileges` - (Optional) The privilege information of account.

The `account_privileges` object supports the following:

* `account_privilege` - (Required) The privilege type of the account.
* `db_name` - (Required) The name of database.
* `account_privilege_detail` - (Optional) The privilege detail of the account.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
RDS mysql account can be imported using the instance_id:account_name, e.g.
```
$ terraform import volcengine_rds_mysql_account.default mysql-42b38c769c4b:test
```

