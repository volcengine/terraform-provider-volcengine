---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_user_group"
sidebar_current: "docs-volcengine-resource-iam_user_group"
description: |-
  Provides a resource to manage iam user group
---
**❗Notice:**
The current provider is no longer being maintained. We recommend that you use the [volcenginecc](https://registry.terraform.io/providers/volcengine/volcenginecc/latest/docs) instead.
# volcengine_iam_user_group
Provides a resource to manage iam user group
## Example Usage
```hcl
resource "volcengine_iam_user_group" "foo" {
  user_group_name = "acc-test-modify-ggg"
  description     = "acc-modify-gg"
  display_name    = "modify-gg"
}
```
## Argument Reference
The following arguments are supported:
* `user_group_name` - (Required) The name of the user group.
* `description` - (Optional) The description of the user group.
* `display_name` - (Optional) The display name of the user group.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
IamUserGroup can be imported using the id, e.g.
```
$ terraform import volcengine_iam_user_group.default user_group_name
```

