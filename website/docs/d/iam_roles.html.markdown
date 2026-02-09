---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_roles"
sidebar_current: "docs-volcengine-datasource-iam_roles"
description: |-
  Use this data source to query detailed information of iam roles
---
# volcengine_iam_roles
Use this data source to query detailed information of iam roles
## Example Usage
```hcl
data "volcengine_iam_roles" "default" {
  query = "CustomRoleForOOS"
}
```
## Argument Reference
The following arguments are supported:
* `name_regex` - (Optional) A Name Regex of Role.
* `output_file` - (Optional) File name where to save data source results.
* `query` - (Optional) Fuzzy query. Can query by role name, display name or description.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `roles` - The collection of Role query.
    * `create_date` - The create time of the Role.
    * `description` - The description of the Role.
    * `display_name` - The display name of the Role.
    * `is_service_linked_role` - Whether the Role is a service linked role.
    * `max_session_duration` - The max session duration of the Role.
    * `role_id` - The id of the Role.
    * `role_name` - The name of the Role.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `trn` - The resource name of the Role.
    * `trust_policy_document` - The trust policy document of the Role.
    * `update_date` - The update time of the Role.
* `total_count` - The total count of Role query.


