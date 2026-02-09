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
  query = "jonny"
}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.
* `query` - (Optional) Fuzzy query. Can query by user name, display name or description.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of user query.
* `users` - The collection of user.
    * `account_id` - Main account ID to which the sub-user belongs.
    * `create_date` - The create date of the user.
    * `description` - The description of the user.
    * `display_name` - The display name of the user.
    * `email_is_verify` - Whether the email has been verified.
    * `email` - The email of the user.
    * `mobile_phone_is_verify` - Whether the phone number has been verified.
    * `mobile_phone` - The mobile phone of the user.
    * `trn` - The trn of the user.
    * `update_date` - The update date of the user.
    * `user_id` - The id of the user.
    * `user_name` - The name of the user.


