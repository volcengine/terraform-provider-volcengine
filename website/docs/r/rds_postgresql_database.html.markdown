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
  owner       = "super"
}
resource "volcengine_rds_postgresql_database" "clone_example" {
  db_name        = "clone-test"
  source_db_name = "acc-test"
  instance_id    = "postgres-95*******233"
  data_option    = "Metadata"
}
```
## Argument Reference
The following arguments are supported:
* `db_name` - (Required, ForceNew) The name of database.
* `instance_id` - (Required, ForceNew) The ID of the RDS instance.
* `c_type` - (Optional, ForceNew) Character classification. Value range: C (default), C.UTF-8, en_US.utf8, zh_CN.utf8, and POSIX.
* `character_set_name` - (Optional, ForceNew) Database character set. Currently supported character sets include: utf8, latin1, ascii. Default is utf8.
* `collate` - (Optional, ForceNew) The collate of database. Sorting rules. Value range: C (default), C.UTF-8, en_US.utf8, zh_CN.utf8 and POSIX.
* `data_option` - (Optional, ForceNew) The data option of the new database. Currently only Metadata is supported. This parameter is optional when clone an existing database.
* `owner` - (Optional) The owner of database.
* `plpgsql_option` - (Optional, ForceNew) The pl_pgsql option of the new database. Value range: View, Procedure, Function, Trigger. This parameter is optional when clone an existing database.
* `source_db_name` - (Optional, ForceNew) The name of the source database. This parameter is required when clone an existing database.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `db_status` - The status of the RDS database.


## Import
Database can be imported using the instanceId:dbName, e.g.
```
$ terraform import volcengine_rds_postgresql_database.default postgres-ca7b7019****:dbname
```

