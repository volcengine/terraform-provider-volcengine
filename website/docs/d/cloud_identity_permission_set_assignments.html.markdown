---
subcategory: "CLOUD_IDENTITY"
layout: "volcengine"
page_title: "Volcengine: volcengine_cloud_identity_permission_set_assignments"
sidebar_current: "docs-volcengine-datasource-cloud_identity_permission_set_assignments"
description: |-
  Use this data source to query detailed information of cloud identity permission set assignments
---
# volcengine_cloud_identity_permission_set_assignments
Use this data source to query detailed information of cloud identity permission set assignments
## Example Usage
```hcl
resource "volcengine_cloud_identity_permission_set" "foo" {
  name             = "acc-test-permission_set"
  description      = "tf"
  session_duration = 5000
  permission_policies {
    permission_policy_type = "System"
    permission_policy_name = "AdministratorAccess"
    inline_policy_document = ""
  }
  permission_policies {
    permission_policy_type = "System"
    permission_policy_name = "ReadOnlyAccess"
    inline_policy_document = ""
  }
  permission_policies {
    permission_policy_type = "Inline"
    inline_policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"auto_scaling:DescribeScalingGroups\"],\"Resource\":[\"*\"]}]}"
  }
}

resource "volcengine_cloud_identity_user" "foo" {
  user_name    = "acc-test-user"
  display_name = "tf-test-user"
  description  = "tf"
  email        = "88@qq.com"
  phone        = "181"
}

resource "volcengine_cloud_identity_permission_set_assignment" "foo" {
  permission_set_id = volcengine_cloud_identity_permission_set.foo.id
  target_id         = "210026****"
  principal_type    = "User"
  principal_id      = volcengine_cloud_identity_user.foo.id
}

data "volcengine_cloud_identity_permission_set_assignments" "foo" {
  permission_set_id = volcengine_cloud_identity_permission_set_assignment.foo.permission_set_id
}
```
## Argument Reference
The following arguments are supported:
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `permission_set_id` - (Optional) The id of cloud identity permission set.
* `principal_id` - (Optional) The principal id of cloud identity permission set. When the `principal_type` is `User`, this field is specified to `UserId`. When the `principal_type` is `Group`, this field is specified to `GroupId`.
* `principal_type` - (Optional) The principal type of cloud identity permission set. Valid values: `User`, `Group`.
* `target_id` - (Optional) The target account id of cloud identity permission set assignment.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `assignments` - The collection of query.
    * `create_time` - The create time of the cloud identity permission set assignment.
    * `id` - The id of the cloud identity permission set.
    * `permission_set_id` - The id of the cloud identity permission set.
    * `permission_set_name` - The name of the cloud identity permission set.
    * `principal_id` - The principal id of the cloud identity permission set assignment.
    * `principal_type` - The principal type of the cloud identity permission set assignment.
    * `target_id` - The target account id of the cloud identity permission set assignment.
* `total_count` - The total count of query.


