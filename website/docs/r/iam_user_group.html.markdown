---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_user_group"
sidebar_current: "docs-volcengine-resource-iam_user_group"
description: |-
  Provides a resource to manage iam user group
---
# volcengine_iam_user_group
Provides a resource to manage iam user group
## Example Usage
```hcl
resource "volcengine_iam_user_group" "foo" {
  user_group_name = "acc-test1"
  description     = "acc"
  display_name    = "modify-xx"
}
```
## Argument Reference
The following arguments are supported:
* `user_group_name` - (Required, ForceNew) The name of the user group.
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

