---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_database"
sidebar_current: "docs-volcengine-resource-rds_postgresql_database"
description: |-
  Provides a resource to manage rds postgresql database
---
# volcengine_rds_postgresql_database
Provides a resource to manage rds postgresql database
## Example Usage
```hcl
resource "volcengine_rds_postgresql_database" "foo" {
  db_name     = "acc-test"
  instance_id = "postgres-95*******233"
  c_type      = "C"
  collate     = "zh_CN.utf8"
}
```
## Argument Reference
The following arguments are supported:
* `db_name` - (Required, ForceNew) The name of database.
* `instance_id` - (Required, ForceNew) The ID of the RDS instance.
* `c_type` - (Optional, ForceNew) Character classification. Value range: C (default), C.UTF-8, en_US.utf8, zh_CN.utf8, and POSIX.
* `character_set_name` - (Optional, ForceNew) Database character set. Currently supported character sets include: utf8, latin1, ascii. Default is utf8.
* `collate` - (Optional, ForceNew) The collate of database. Sorting rules. Value range: C (default), C.UTF-8, en_US.utf8, zh_CN.utf8 and POSIX.
* `owner` - (Optional, ForceNew) The owner of database.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `db_status` - The status of the RDS database.


## Import
Database can be imported using the instanceId:dbName, e.g.
```
$ terraform import volcengine_rds_postgresql_database.default postgres-ca7b7019****:dbname
```

