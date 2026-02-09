---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_entities_policies"
sidebar_current: "docs-volcengine-datasource-iam_entities_policies"
description: |-
  Use this data source to query detailed information of iam entities policies
---
# volcengine_iam_entities_policies
Use this data source to query detailed information of iam entities policies
## Example Usage
```hcl
data "volcengine_iam_entities_policies" "default" {
  policy_name = "AdministratorAccess"
  policy_type = "System"
}
```
## Argument Reference
The following arguments are supported:
* `policy_name` - (Required) The name of the policy.
* `policy_type` - (Required) The type of the policy.
* `entity_filter` - (Optional) The entity filter.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `roles` - The collection of roles.
    * `attach_date` - The attach date of the role.
    * `description` - The description of the role.
    * `display_name` - The display name of the role.
    * `id` - The id of the role.
    * `policy_scope` - The scope of the policy.
        * `attach_date` - The attach date of the policy scope.
        * `policy_scope_type` - The type of the policy scope.
        * `project_display_name` - The display name of the project.
        * `project_name` - The name of the project.
    * `role_name` - The name of the role.
* `total_count` - The total count of query.
* `user_groups` - The collection of user groups.
    * `attach_date` - The attach date of the user group.
    * `description` - The description of the user group.
    * `display_name` - The display name of the user group.
    * `id` - The id of the user group.
    * `policy_scope` - The scope of the policy.
        * `attach_date` - The attach date of the policy scope.
        * `policy_scope_type` - The type of the policy scope.
        * `project_display_name` - The display name of the project.
        * `project_name` - The name of the project.
    * `user_group_name` - The name of the user group.
* `users` - The collection of users.
    * `attach_date` - The attach date of the user.
    * `description` - The description of the user.
    * `display_name` - The display name of the user.
    * `id` - The id of the user.
    * `policy_scope` - The scope of the policy.
        * `attach_date` - The attach date of the policy scope.
        * `policy_scope_type` - The type of the policy scope.
        * `project_display_name` - The display name of the project.
        * `project_name` - The name of the project.
    * `user_name` - The name of the user.


