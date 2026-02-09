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
  user_name                 = "jonny"
  password                  = ""
  login_allowed             = true
  password_reset_required   = true
  safe_auth_flag            = true
  safe_auth_type            = "phone"
  safe_auth_exempt_required = 1
  safe_auth_exempt_unit     = 1
  safe_auth_exempt_duration = 1
}
```
## Argument Reference
The following arguments are supported:
* `password` - (Required) The password.
* `user_name` - (Required, ForceNew) The user name.
* `login_allowed` - (Optional) The flag of login allowed.
* `password_reset_required` - (Optional) Is required reset password when next time login in.
* `safe_auth_exempt_duration` - (Optional) The duration of safe auth exempt.
* `safe_auth_exempt_required` - (Optional) The flag of safe auth exempt required.
* `safe_auth_exempt_unit` - (Optional) The unit of safe auth exempt.
* `safe_auth_flag` - (Optional) The flag of safe auth.
* `safe_auth_type` - (Optional) The type of safe auth.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `create_date` - The create date.
* `last_login_date` - The last login date.
* `last_login_ip` - The last login ip.
* `last_reset_password_time` - The last reset password time.
* `login_locked` - The flag of login locked.
* `password_expire_at` - The password expire at.
* `update_date` - The update date.
* `user_id` - The user id.


## Import
Login profile can be imported using the UserName, e.g.
```
$ terraform import volcengine_iam_login_profile.default user_name
```

