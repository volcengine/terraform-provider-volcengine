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
resource "volcengine_rds_mysql_database" "default" {
  instance_id        = "mysql-xxx"
  db_name            = "xxx"
  character_set_name = "utf8"
}
```
## Argument Reference
The following arguments are supported:
* `character_set_name` - (Required, ForceNew) Database character set. Currently supported character sets include: utf8, utf8mb4, latin1, ascii.
* `db_name` - (Required, ForceNew) Name database.
illustrate:
Unique name.
The length is 2~64 characters.
Start with a letter and end with a letter or number.
Consists of lowercase letters, numbers, and underscores (_) or dashes (-).
Database names are disabled [keywords](https://www.volcengine.com/docs/6313/66162).
* `instance_id` - (Required, ForceNew) The ID of the RDS instance.
* `database_privileges` - (Optional) The privilege detail list of RDS mysql instance database.

The `database_privileges` object supports the following:

* `account_name` - (Required) The name of account.
* `account_privilege` - (Required) The privilege type of the account.
* `account_privilege_detail` - (Optional) The privilege detail of the account.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Database can be imported using the instanceId:dbName, e.g.
```
$ terraform import volcengine_rds_database.default mysql-42b38c769c4b:dbname
```

