---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_group_users"
sidebar_current: "docs-volcengine-datasource-iam_group_users"
description: |-
  Use this data source to query detailed information of iam group users
---
# volcengine_iam_group_users
Use this data source to query detailed information of iam group users
## Example Usage
```hcl
data "volcengine_iam_group_users" "default" {
  user_name = "jonny"
}
```
## Argument Reference
The following arguments are supported:
* `user_name` - (Required) The name of user.
* `output_file` - (Optional) File name where to save data source results.
* `query` - (Optional) Fuzzy search, supports searching for user group names, display names, and remarks.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of query.
* `user_groups` - The collection of user group.
    * `description` - The description of the user group.
    * `display_name` - The display name of the user group.
    * `join_date` - The join date of the user group.
    * `user_group_id` - The id of the user group.
    * `user_group_name` - The name of the user group.


