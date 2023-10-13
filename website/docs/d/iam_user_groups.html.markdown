---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_user_groups"
sidebar_current: "docs-volcengine-datasource-iam_user_groups"
description: |-
  Use this data source to query detailed information of iam user groups
---
# volcengine_iam_user_groups
Use this data source to query detailed information of iam user groups
## Example Usage
```hcl
resource "volcengine_iam_user_group" "foo" {
  user_group_name = "acc-test-group"
  description     = "acc-test"
  display_name    = "acc-test"
}

data "volcengine_iam_user_groups" "foo" {
  query = "acc-test"
}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.
* `query` - (Optional) Fuzzy search, supports searching for user group names, display names, and remarks.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of query.
* `user_groups` - The collection of query.
    * `account_id` - The id of the account.
    * `create_date` - The creation date of the user group.
    * `description` - The description of the user group.
    * `display_name` - The display name of the user group.
    * `update_date` - The update date of the user group.
    * `user_group_name` - The name of the user group.


