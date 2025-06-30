---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_mysql_account_table_column_infos"
sidebar_current: "docs-volcengine-datasource-rds_mysql_account_table_column_infos"
description: |-
  Use this data source to query detailed information of rds mysql account table column infos
---
# volcengine_rds_mysql_account_table_column_infos
Use this data source to query detailed information of rds mysql account table column infos
## Example Usage
```hcl
data "volcengine_rds_mysql_account_table_column_infos" "foo" {
  db_name     = "ddd"
  instance_id = "mysql-b51d37110dd1"
}
```
## Argument Reference
The following arguments are supported:
* `db_name` - (Required) The name of the database.
* `instance_id` - (Required) The id of the mysql instance.
* `account_name` - (Optional) The name of the account.
* `column_name` - (Optional) The name of the column.
* `host` - (Optional) Specify the IP address for the account to access the database. The default value is %.
* `output_file` - (Optional) File name where to save data source results.
* `table_limit` - (Optional) Specify the number of tables in the table column permission information to be returned. If it exceeds the setting, it will be truncated.
* `table_name` - (Optional) The name of the table.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `table_infos` - The collection of query.
    * `account_privileges` - The table permissions of the account.
    * `column_infos` - The column permission information of the account.
        * `account_privileges` - The column privileges of the account.
        * `column_name` - The name of the column.
    * `table_name` - The name of the table.
* `total_count` - The total count of query.


