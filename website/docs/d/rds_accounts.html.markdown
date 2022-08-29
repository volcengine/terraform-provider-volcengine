---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_accounts"
sidebar_current: "docs-volcengine-datasource-rds_accounts"
description: |-
  Use this data source to query detailed information of rds accounts
---
# volcengine_rds_accounts
Use this data source to query detailed information of rds accounts
## Example Usage
```hcl
data "volcengine_rds_accounts" "default" {
  instance_id = "mysql-0fdd3bab2e7c"
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
* `rds_accounts` - The collection of RDS instance account query.
    * `account_name` - The name of the database account.
    * `account_status` - The status of the database account.
    * `account_type` - The type of the database account.
    * `db_privileges` - The privilege detail list of RDS instance account.
        * `account_privilege_str` - The privilege string of the account.
        * `account_privilege` - The privilege type of the account.
        * `db_name` - The name of database.
    * `id` - The ID of the RDS instance account.
* `total_count` - The total count of database account query.


