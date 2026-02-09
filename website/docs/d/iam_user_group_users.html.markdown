---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_user_group_users"
sidebar_current: "docs-volcengine-datasource-iam_user_group_users"
description: |-
  Use this data source to query detailed information of iam user group users
---
# volcengine_iam_user_group_users
Use this data source to query detailed information of iam user group users
## Example Usage
```hcl
data "volcengine_iam_user_group_users" "default" {
  user_group_name = "KR23A6ItftestA37rerA"
}
```
## Argument Reference
The following arguments are supported:
* `user_group_name` - (Required) The name of user group.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of query.
* `users` - The collection of user.
    * `description` - The description of the user.
    * `display_name` - The display name of the user.
    * `join_date` - The join date of the user.
    * `user_id` - The id of the user.
    * `user_name` - The name of the user.


