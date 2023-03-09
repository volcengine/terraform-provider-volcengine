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
resource "volcengine_rds_mysql_account" "default" {
  instance_id      = "mysql-xxx"
  account_name     = "xxx"
  account_password = "xxx"
  account_type     = "Normal"
  account_privileges {
    db_name                  = "xxx"
    account_privilege        = "Custom"
    account_privilege_detail = "SELECT,UPDATE,INSERT"
  }
  account_privileges {
    db_name                  = "xx"
    account_privilege        = "Custom"
    account_privilege_detail = "SELECT,UPDATE,INSERT"
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
illustrate
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
$ terraform import volcengine_rds_account.default mysql-42b38c769c4b:test
```

