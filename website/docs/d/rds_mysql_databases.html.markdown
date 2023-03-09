---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_mysql_databases"
sidebar_current: "docs-volcengine-datasource-rds_mysql_databases"
description: |-
  Use this data source to query detailed information of rds mysql databases
---
# volcengine_rds_mysql_databases
Use this data source to query detailed information of rds mysql databases
## Example Usage
```hcl
data "volcengine_rds_mysql_databases" "default" {
  instance_id = ""
  db_name     = ""
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The id of the RDS instance.
* `db_name` - (Optional) The name of the RDS database.
* `name_regex` - (Optional) A Name Regex of RDS database.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `databases` - The collection of RDS instance account query.
    * `character_set_name` - The character set of the RDS database.
    * `database_privileges` - The privilege detail list of RDS mysql instance database.
        * `account_name` - The name of account.
        * `account_privilege_detail` - The privilege detail of the account.
        * `account_privilege` - The privilege type of the account.
    * `db_name` - The name of the RDS database.
    * `id` - The ID of the RDS database.
* `total_count` - The total count of RDS database query.


