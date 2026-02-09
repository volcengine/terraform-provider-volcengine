---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_user_group_policy_attachments"
sidebar_current: "docs-volcengine-datasource-iam_user_group_policy_attachments"
description: |-
  Use this data source to query detailed information of iam user group policy attachments
---
# volcengine_iam_user_group_policy_attachments
Use this data source to query detailed information of iam user group policy attachments
## Example Usage
```hcl
data "volcengine_iam_user_group_policy_attachments" "default" {
  user_group_name = "xRqElT"
}
```
## Argument Reference
The following arguments are supported:
* `user_group_name` - (Required) A name of user group.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `policies` - The collection of query.
    * `attach_date` - Attached time.
    * `description` - The description.
    * `policy_name` - Name of the policy.
    * `policy_scope` - The scope of the policy.
        * `attach_date` - The attach date of the policy scope.
        * `policy_scope_type` - The type of the policy scope.
        * `project_display_name` - The display name of the project.
        * `project_name` - The name of the project.
    * `policy_trn` - Resource name of the strategy.
    * `policy_type` - The type of the policy.
* `total_count` - The total count of query.


