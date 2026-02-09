---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_user_policy_attachments"
sidebar_current: "docs-volcengine-datasource-iam_user_policy_attachments"
description: |-
  Use this data source to query detailed information of iam user policy attachments
---
# volcengine_iam_user_policy_attachments
Use this data source to query detailed information of iam user policy attachments
## Example Usage
```hcl
data "volcengine_iam_user_policy_attachments" "default" {
  user_name = "jonny"
}
```
## Argument Reference
The following arguments are supported:
* `user_name` - (Required) The name of the user.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `policies` - The collection of policies.
    * `attach_date` - The attach date of the policy.
    * `description` - The description of the policy.
    * `policy_name` - The name of the policy.
    * `policy_scope` - The scope of the policy.
        * `attach_date` - The attach date of the policy scope.
        * `policy_scope_type` - The type of the policy scope.
        * `project_display_name` - The display name of the project.
        * `project_name` - The name of the project.
    * `policy_trn` - The trn of the policy.
    * `policy_type` - The type of the policy.


