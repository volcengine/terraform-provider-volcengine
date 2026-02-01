---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_databases"
sidebar_current: "docs-volcengine-datasource-rds_postgresql_databases"
description: |-
  Use this data source to query detailed information of rds postgresql databases
---
# volcengine_rds_postgresql_databases
Use this data source to query detailed information of rds postgresql databases
## Example Usage
```hcl
data "volcengine_rds_postgresql_databases" "foo" {
  instance_id = "postgres-95******8233"
  db_name     = "test001"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The id of the RDS instance.
* `db_name` - (Optional) The name of the RDS database.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `databases` - The collection of RDS instance account query.
    * `c_type` - Character classification.
    * `character_set_name` - The character set of the RDS database.
    * `collate` - The collate of database.
    * `db_name` - The name of the RDS database.
    * `db_status` - The status of the RDS database.
    * `owner` - The owner of database.
* `total_count` - The total count of RDS database query.


