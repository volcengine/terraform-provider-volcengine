---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_login_profiles"
sidebar_current: "docs-volcengine-datasource-iam_login_profiles"
description: |-
  Use this data source to query detailed information of iam login profiles
---
# volcengine_iam_login_profiles
Use this data source to query detailed information of iam login profiles
## Example Usage
```hcl
data "volcengine_iam_login_profiles" "default" {
  user_name = "xx"
}
```
## Argument Reference
The following arguments are supported:
* `user_name` - (Required) The user name.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `login_profiles` - The collection of login profiles.
    * `create_date` - The create date.
    * `last_login_date` - The last login date.
    * `last_login_ip` - The last login ip.
    * `last_reset_password_time` - The last reset password time.
    * `login_allowed` - The flag of login allowed.
    * `login_locked` - The flag of login locked.
    * `password_expire_at` - The password expire at.
    * `password_reset_required` - Is required reset password when next time login in.
    * `safe_auth_exempt_duration` - The duration of safe auth exempt.
    * `safe_auth_exempt_required` - The flag of safe auth exempt required.
    * `safe_auth_exempt_unit` - The unit of safe auth exempt.
    * `safe_auth_flag` - The flag of safe auth.
    * `safe_auth_type` - The type of safe auth.
    * `update_date` - The update date.
    * `user_id` - The user id.
    * `user_name` - The user name.


