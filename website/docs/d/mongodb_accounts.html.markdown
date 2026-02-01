---
subcategory: "MONGODB"
layout: "volcengine"
page_title: "Volcengine: volcengine_mongodb_accounts"
sidebar_current: "docs-volcengine-datasource-mongodb_accounts"
description: |-
  Use this data source to query detailed information of mongodb accounts
---
# volcengine_mongodb_accounts
Use this data source to query detailed information of mongodb accounts
## Example Usage
```hcl
data "volcengine_mongodb_accounts" "default" {
  instance_id = "mongo-replica-xxx"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) Target query mongodb instance id.
* `account_name` - (Optional) The name of account. This field support fuzzy query.
* `auth_db` - (Optional) The database of account. This field support fuzzy query.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `accounts` - The collection of accounts query.
    * `account_desc` - The description of account.
    * `account_name` - The name of account.
    * `account_privileges` - The privilege info of mongo instance.
        * `db_name` - The Name of DB.
        * `role_name` - The Name of role.
    * `account_type` - The type of account.
    * `auth_db` - The database of account.
    * `create_time` - The create time of account.
    * `modify_time` - The modify time of account.
* `total_count` - The total count of accounts query.


