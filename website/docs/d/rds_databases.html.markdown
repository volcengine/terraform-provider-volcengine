---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_databases"
sidebar_current: "docs-volcengine-datasource-rds_databases"
description: |-
  Use this data source to query detailed information of rds databases
---
# volcengine_rds_databases
Use this data source to query detailed information of rds databases
## Example Usage
```hcl
data "volcengine_rds_databases" "default" {
  instance_id = "mysql-0fdd3bab2e7c"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The id of the RDS instance.
* `db_status` - (Optional) The status of the RDS database.
* `name_regex` - (Optional) A Name Regex of RDS database.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `rds_databases` - The collection of RDS instance account query.
    * `account_names` - The account names of the RDS database.
    * `character_set_name` - The character set of the RDS database.
    * `db_name` - The name of the RDS database.
    * `db_status` - The status of the RDS database.
    * `id` - The ID of the RDS database.
* `total_count` - The total count of RDS database query.


