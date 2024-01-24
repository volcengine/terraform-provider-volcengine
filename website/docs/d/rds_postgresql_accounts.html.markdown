---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_accounts"
sidebar_current: "docs-volcengine-datasource-rds_postgresql_accounts"
description: |-
  Use this data source to query detailed information of rds postgresql accounts
---
# volcengine_rds_postgresql_accounts
Use this data source to query detailed information of rds postgresql accounts
## Example Usage
```hcl
data "volcengine_rds_postgresql_accounts" "foo" {
  instance_id = "postgres-954****f7233"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The id of the RDS instance.
* `account_name` - (Optional) The name of the database account. This field supports fuzzy query.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `accounts` - The collection of RDS instance account query.
    * `account_name` - The name of the database account.
    * `account_privileges` - The privileges of the database account.
    * `account_status` - The status of the database account.
    * `account_type` - The type of the database account.
* `total_count` - The total count of database account query.


