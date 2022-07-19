---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_login_profile"
sidebar_current: "docs-volcengine-resource-iam_login_profile"
description: |-
  Provides a resource to manage iam login profile
---
# volcengine_iam_login_profile
Provides a resource to manage iam login profile
## Example Usage
```hcl
resource "volcengine_iam_login_profile" "foo" {
  user_name               = "tf-test"
  password                = "******"
  login_allowed           = true
  password_reset_required = false
}
```
## Argument Reference
The following arguments are supported:
* `password` - (Required) The password.
* `user_name` - (Required, ForceNew) The user name.
* `login_allowed` - (Optional) The flag of login allowed.
* `password_reset_required` - (Optional) Is required reset password when next time login in.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Login profile can be imported using the UserName, e.g.
```
$ terraform import volcengine_iam_login_profile.default user_name
```

