---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_role_policy_attachment"
sidebar_current: "docs-volcengine-resource-iam_role_policy_attachment"
description: |-
  Provides a resource to manage iam role policy attachment
---
# volcengine_iam_role_policy_attachment
Provides a resource to manage iam role policy attachment
## Example Usage
```hcl
resource "volcengine_iam_role" "role" {
  role_name             = "TerraformTestRole"
  display_name          = "terraform role"
  trust_policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"sts:AssumeRole\"],\"Principal\":{\"Service\":[\"auto_scaling\"]}}]}"
  description           = "created by terraform"
  max_session_duration  = 43200
}

resource "volcengine_iam_policy" "policy" {
  policy_name     = "TerraformResourceTest1"
  description     = "created by terraform 1"
  policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"auto_scaling:DescribeScalingGroups\"],\"Resource\":[\"*\"]}]}"
}

resource "volcengine_iam_role_policy_attachment" "foo" {
  role_name   = volcengine_iam_role.role.id
  policy_name = volcengine_iam_policy.policy.id
  policy_type = volcengine_iam_policy.policy.policy_type
}
```
## Argument Reference
The following arguments are supported:
* `policy_name` - (Required, ForceNew) The name of the Policy.
* `policy_type` - (Required, ForceNew) The type of the Policy.
* `role_name` - (Required, ForceNew) The name of the Role.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Iam role policy attachment can be imported using the id, e.g.
```
$ terraform import volcengine_iam_role_policy_attachment.default TerraformTestRole:TerraformTestPolicy:Custom
```

