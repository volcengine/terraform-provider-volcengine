---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_user_group_policy_attachment"
sidebar_current: "docs-volcengine-resource-iam_user_group_policy_attachment"
description: |-
  Provides a resource to manage iam user group policy attachment
---
**❗Notice:**
The current provider is no longer being maintained. We recommend that you use the [volcenginecc](https://registry.terraform.io/providers/volcengine/volcenginecc/latest/docs) instead.
# volcengine_iam_user_group_policy_attachment
Provides a resource to manage iam user group policy attachment
## Example Usage
```hcl
resource "volcengine_iam_user_group_policy_attachment" "foo" {
  policy_name     = "test"
  policy_type     = "Custom"
  user_group_name = "xRqElT"
}
```
## Argument Reference
The following arguments are supported:
* `policy_name` - (Required, ForceNew) The policy name.
* `policy_type` - (Required, ForceNew) Strategy types, System strategy, Custom strategy.
* `user_group_name` - (Required, ForceNew) The user group name.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
IamUserGroupPolicyAttachment can be imported using the user group name and policy name, e.g.
```
$ terraform import volcengine_iam_user_group_policy_attachment.default userGroupName:policyName
```

