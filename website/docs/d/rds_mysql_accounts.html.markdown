---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_mysql_accounts"
sidebar_current: "docs-volcengine-datasource-rds_mysql_accounts"
description: |-
  Use this data source to query detailed information of rds mysql accounts
---
# volcengine_rds_mysql_accounts
Use this data source to query detailed information of rds mysql accounts
## Example Usage
```hcl
data "volcengine_rds_mysql_accounts" "default" {
  instance_id  = ""
  account_name = ""
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The id of the RDS instance.
* `account_name` - (Optional) The name of the database account.
* `name_regex` - (Optional) A Name Regex of database account.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `accounts` - The collection of RDS instance account query.
    * `account_name` - The name of the database account.
    * `account_privileges` - The privilege detail list of RDS mysql instance account.
        * `account_privilege_detail` - The privilege detail of the account.
        * `account_privilege` - The privilege type of the account.
        * `db_name` - The name of database.
    * `account_status` - The status of the database account.
    * `account_type` - The type of the database account.
    * `id` - The ID of the RDS instance account.
* `total_count` - The total count of database account query.


