---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_database"
sidebar_current: "docs-volcengine-resource-rds_database"
description: |-
  Provides a resource to manage rds database
---
# volcengine_rds_database
(Deprecated! Recommend use volcengine_rds_mysql_*** replace) Provides a resource to manage rds database
## Example Usage
```hcl
resource "volcengine_rds_database" "foo" {
  instance_id        = "mysql-0fdd3bab2e7c"
  db_name            = "foo"
  character_set_name = "utf8mb4"
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

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Database can be imported using the id, e.g.
```
$ terraform import volcengine_rds_database.default mysql-42b38c769c4b:dbname
```

