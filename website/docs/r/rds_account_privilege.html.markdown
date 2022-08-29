---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_account_privilege"
sidebar_current: "docs-volcengine-resource-rds_account_privilege"
description: |-
  Provides a resource to manage rds account privilege
---
# volcengine_rds_account_privilege
Provides a resource to manage rds account privilege
## Example Usage
```hcl
resource "volcengine_rds_account" "app_name" {
  instance_id      = "mysql-0fdd3bab2e7c"
  account_name     = "terraform-test-app"
  account_password = "Aatest123"
  account_type     = "Normal"
}

resource "volcengine_rds_account_privilege" "foo" {
  instance_id  = "mysql-0fdd3bab2e7c"
  account_name = volcengine_rds_account.app_name.account_name

  db_privileges {
    db_name               = "foo"
    account_privilege     = "Custom"
    account_privilege_str = "ALTER,ALTER ROUTINE,CREATE,CREATE ROUTINE,CREATE TEMPORARY TABLES"
  }

  db_privileges {
    db_name           = "bar"
    account_privilege = "DDLOnly"
  }

  db_privileges {
    db_name           = "demo"
    account_privilege = "ReadWrite"
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
* `db_privileges` - (Required) The privileges of the account.
* `instance_id` - (Required, ForceNew) The ID of the RDS instance.

The `db_privileges` object supports the following:

* `account_privilege` - (Required) The privilege type of the account.
* `db_name` - (Required) The name of database.
* `account_privilege_str` - (Optional) The privilege string of the account.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
RDS account privilege can be imported using the id, e.g.
```
$ terraform import volcengine_rds_account_privilege.default mysql-42b38c769c4b:account_name
```

