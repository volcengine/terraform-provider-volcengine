---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_user_policy_attachment"
sidebar_current: "docs-volcengine-resource-iam_user_policy_attachment"
description: |-
  Provides a resource to manage iam user policy attachment
---
# volcengine_iam_user_policy_attachment
Provides a resource to manage iam user policy attachment
## Example Usage
```hcl
resource "volcengine_iam_user" "user" {
  user_name   = "TfTest"
  description = "test"
}

resource "volcengine_iam_policy" "policy" {
  policy_name     = "TerraformResourceTest1"
  description     = "created by terraform 1"
  policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"auto_scaling:DescribeScalingGroups\"],\"Resource\":[\"*\"]}]}"
}

resource "volcengine_iam_user_policy_attachment" "foo" {
  user_name   = volcengine_iam_user.user.user_name
  policy_name = volcengine_iam_policy.policy.policy_name
  policy_type = volcengine_iam_policy.policy.policy_type
}
```
## Argument Reference
The following arguments are supported:
* `policy_name` - (Required, ForceNew) The name of the Policy.
* `policy_type` - (Required, ForceNew) The type of the Policy.
* `user_name` - (Required, ForceNew) The name of the user.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Iam user policy attachment can be imported using the UserName:PolicyName:PolicyType, e.g.
```
$ terraform import volcengine_iam_user_policy_attachment.default TerraformTestUser:TerraformTestPolicy:Custom
```

