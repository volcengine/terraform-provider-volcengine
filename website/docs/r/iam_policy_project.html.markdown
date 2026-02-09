---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_policy_project"
sidebar_current: "docs-volcengine-resource-iam_policy_project"
description: |-
  Provides a resource to manage iam policy project
---
# volcengine_iam_policy_project
Provides a resource to manage iam policy project
## Example Usage
```hcl
resource "volcengine_iam_policy_project" "foo" {
  principal_type = "User"
  principal_name = "jonny"
  policy_type    = "Custom"
  policy_name    = "restart-oas-ecs"
  project_names  = ["default"]
}
```
## Argument Reference
The following arguments are supported:
* `policy_name` - (Required, ForceNew) The name of the policy.
* `policy_type` - (Required, ForceNew) The type of the policy. Valid values: System, Custom.
* `principal_name` - (Required, ForceNew) The name of the principal.
* `principal_type` - (Required, ForceNew) The type of the principal. Valid values: User, Role, UserGroup.
* `project_names` - (Required, ForceNew) The list of project names, which is the scope of the policy.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
IamPolicyProject can be imported using the id, e.g.
```
$ terraform import volcengine_iam_policy_project.default PrincipalType:PrincipalName:PolicyType:PolicyName:ProjectName
```

