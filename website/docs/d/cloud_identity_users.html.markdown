---
subcategory: "CLOUD_IDENTITY"
layout: "volcengine"
page_title: "Volcengine: volcengine_cloud_identity_users"
sidebar_current: "docs-volcengine-datasource-cloud_identity_users"
description: |-
  Use this data source to query detailed information of cloud identity users
---
# volcengine_cloud_identity_users
Use this data source to query detailed information of cloud identity users
## Example Usage
```hcl
resource "volcengine_cloud_identity_user" "foo" {
  user_name    = "acc-test-user-${count.index}"
  display_name = "tf-test-user-${count.index}"
  description  = "tf"
  email        = "88@qq.com"
  phone        = "181"

  count = 2
}

data "volcengine_cloud_identity_users" "foo" {
  user_name = "acc-test-user"
  source    = "Manual"
}
```
## Argument Reference
The following arguments are supported:
* `department_id` - (Optional) The department id.
* `display_name` - (Optional) The display name of cloud identity user.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `source` - (Optional) The source of cloud identity user. Valid values: `Sync`, `Manual`.
* `user_name` - (Optional) The name of cloud identity user.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of query.
* `users` - The collection of query.
    * `created_time` - The created time of the cloud identity user.
    * `description` - The description of the cloud identity user.
    * `display_name` - The display name of the cloud identity user.
    * `email` - The email of the cloud identity user.
    * `id` - The id of the cloud identity user.
    * `identity_type` - The identity type of the cloud identity user.
    * `phone` - The phone of the cloud identity user.
    * `source` - The source of the cloud identity user.
    * `updated_time` - The updated time of the cloud identity user.
    * `user_id` - The id of the cloud identity user.
    * `user_name` - The name of the cloud identity user.


