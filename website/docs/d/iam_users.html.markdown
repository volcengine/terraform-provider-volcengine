---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_users"
sidebar_current: "docs-volcengine-datasource-iam_users"
description: |-
  Use this data source to query detailed information of iam users
---
# volcengine_iam_users
Use this data source to query detailed information of iam users
## Example Usage
```hcl
data "volcengine_iam_users" "default" {
  #  user_names = ["tf-test"]
}
```
## Argument Reference
The following arguments are supported:
* `name_regex` - (Optional) A Name Regex of IAM.
* `output_file` - (Optional) File name where to save data source results.
* `user_names` - (Optional) A list of user names.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of user query.
* `users` - The collection of user.
  * `account_id` - The account id of the user.
  * `create_date` - The create date of the user.
  * `trn` - The trn of the user.
  * `update_date` - The update date of the user.
  * `user_name` - The name of the user.


