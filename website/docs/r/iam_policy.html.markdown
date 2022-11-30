---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_policy"
sidebar_current: "docs-volcengine-resource-iam_policy"
description: |-
  Provides a resource to manage iam policy
---
# volcengine_iam_policy
Provides a resource to manage iam policy
## Example Usage
```hcl
resource "volcengine_iam_policy" "foo" {
  policy_name     = "TerraformResourceTest1"
  description     = "created by terraform 1"
  policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"auto_scaling:DescribeScalingGroups\"],\"Resource\":[\"*\"]}]}"
}
```
## Argument Reference
The following arguments are supported:
* `policy_document` - (Required) The document of the Policy.
* `policy_name` - (Required) The name of the Policy.
* `description` - (Optional) The description of the Policy.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `create_date` - The create time of the Policy.
* `policy_trn` - The resource name of the Policy.
* `policy_type` - The type of the Policy.
* `update_date` - The update time of the Policy.


## Import
Iam policy can be imported using the id, e.g.
```
$ terraform import volcengine_iam_policy.default TerraformTestPolicy
```

