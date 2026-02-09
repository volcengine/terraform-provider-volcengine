---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_user_group_attachment"
sidebar_current: "docs-volcengine-resource-iam_user_group_attachment"
description: |-
  Provides a resource to manage iam user group attachment
---
# volcengine_iam_user_group_attachment
Provides a resource to manage iam user group attachment
## Example Usage
```hcl
resource "volcengine_iam_user_group_attachment" "foo" {
  user_group_name = "xRqElT"
  user_name       = "jonny-tt"
}
```
## Argument Reference
The following arguments are supported:
* `user_group_name` - (Required, ForceNew) The name of the user group.
* `user_name` - (Required, ForceNew) The name of the user.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
IamUserGroupAttachment can be imported using the id, e.g.
```
$ terraform import volcengine_iam_user_group_attachment.default user_group_id:user_id
```

