---
subcategory: "VEDB_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_vedb_mysql_account"
sidebar_current: "docs-volcengine-resource-vedb_mysql_account"
description: |-
  Provides a resource to manage vedb mysql account
---
# volcengine_vedb_mysql_account
Provides a resource to manage vedb mysql account
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
  zone_id     = data.volcengine_zones.foo.zones[2].id
  vpc_id      = volcengine_vpc.foo.id
}


resource "volcengine_vedb_mysql_instance" "foo" {
  charge_type         = "PostPaid"
  storage_charge_type = "PostPaid"
  db_engine_version   = "MySQL_8_0"
  db_minor_version    = "3.0"
  node_number         = 2
  node_spec           = "vedb.mysql.x4.large"
  subnet_id           = volcengine_subnet.foo.id
  instance_name       = "tf-test"
  project_name        = "testA"
  tags {
    key   = "tftest"
    value = "tftest"
  }
  tags {
    key   = "tftest2"
    value = "tftest2"
  }
}

resource "volcengine_vedb_mysql_database" "foo" {
  db_name     = "tf-table"
  instance_id = volcengine_vedb_mysql_instance.foo.id
}

resource "volcengine_vedb_mysql_account" "foo" {
  account_name     = "tftest"
  account_password = "93f0cb0614Aab12"
  account_type     = "Normal"
  instance_id      = volcengine_vedb_mysql_instance.foo.id
  account_privileges {
    db_name                  = volcengine_vedb_mysql_database.foo.db_name
    account_privilege        = "Custom"
    account_privilege_detail = "SELECT,INSERT,DELETE"
  }
}
```
## Argument Reference
The following arguments are supported:
* `account_name` - (Required, ForceNew) Database account name. The account name must meet the following requirements:
 The name is unique and within 2 to 32 characters in length.
 Consists of lowercase letters, numbers, or underscores (_).
 Starts with a lowercase letter and ends with a letter or number.
 The name cannot contain certain prohibited words. For detailed information, please refer to prohibited keywords. And certain reserved words such as root, admin, etc. cannot be used.
* `account_password` - (Required) Password of database account. The account password must meet the following requirements:
 It can only contain upper and lower case letters, numbers and the following special characters _#!@$%^&*()+=-. 
It must be within 8 to 32 characters in length.
 It must contain at least three of upper case letters, lower case letters, numbers or special characters.
* `account_type` - (Required, ForceNew) Database account type. Values: 
Super: High-privilege account. Only one high-privilege account can be created for an instance. It has all permissions for all databases under this instance and can manage all ordinary accounts and databases. 
Normal: Multiple ordinary accounts can be created for an instance. Specific database permissions need to be manually granted to ordinary accounts.
* `instance_id` - (Required, ForceNew) The id of the instance.
* `account_privileges` - (Optional) Database permission information. When the value of AccountType is Super, this parameter does not need to be passed. High-privilege accounts by default have all permissions for all databases under this instance. When the value of AccountType is Normal, it is recommended to pass this parameter to grant specified permissions for specified databases to ordinary accounts. If not set, this account does not have any permissions for any database.

The `account_privileges` object supports the following:

* `account_privilege` - (Required) Authorization database privilege types: 
ReadWrite: Read and write privilege.
 ReadOnly: Read-only privilege.
 DDLOnly: Only DDL privilege.
 DMLOnly: Only DML privilege.
 Custom: Custom privilege.
* `db_name` - (Required) Database name requiring authorization.
* `account_privilege_detail` - (Optional) The specific SQL operation permissions contained in the permission type are separated by English commas (,) between multiple strings.
 When used as a request parameter in the CreateDatabase interface, when the AccountPrivilege value is Custom, this parameter is required. Value range (multiple selections allowed): SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, REFERENCES, INDEX, ALTER, CREATE TEMPORARY TABLES, LOCK TABLES, EXECUTE, CREATE VIEW, SHOW VIEW, CREATE ROUTINE, ALTER ROUTINE, EVENT, TRIGGER. When used as a return parameter in the DescribeDatabases interface, regardless of the value of AccountPrivilege, the details of the SQL operation permissions contained in this permission type are returned. For the specific SQL operation permissions contained in each permission type, please refer to the account permission list.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
VedbMysqlAccount can be imported using the instance id and account name, e.g.
```
$ terraform import volcengine_vedb_mysql_account.default vedbm-r3xq0zdl****:testuser

```

