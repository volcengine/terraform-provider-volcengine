---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_role"
sidebar_current: "docs-volcengine-resource-iam_role"
description: |-
  Provides a resource to manage iam role
---
# volcengine_iam_role
Provides a resource to manage iam role
## Example Usage
```hcl
resource "volcengine_iam_role" "foo" {
  role_name             = "TerraformTestRole"
  display_name          = "terraform role"
  trust_policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"sts:AssumeRole\"],\"Principal\":{\"Service\":[\"auto_scaling\"]}}]}"
  description           = "created by terraform"
  max_session_duration  = 43200
}
```
## Argument Reference
The following arguments are supported:
* `display_name` - (Required) The display name of the Role.
* `role_name` - (Required, ForceNew) The name of the Role.
* `trust_policy_document` - (Required) The trust policy document of the Role.
* `description` - (Optional) The description of the Role.
* `max_session_duration` - (Optional) The max session duration of the Role.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `trn` - The resource name of the Role.


## Import
Iam role can be imported using the id, e.g.
```
$ terraform import volcengine_iam_role.default TerraformTestRole
```

