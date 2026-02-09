---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_security_config"
sidebar_current: "docs-volcengine-resource-iam_security_config"
description: |-
  Provides a resource to manage iam security config
---
# volcengine_iam_security_config
Provides a resource to manage iam security config
## Example Usage
```hcl
resource "volcengine_iam_security_config" "foo" {
  user_name                 = "jonny"
  safe_auth_type            = "email"
  safe_auth_exempt_duration = 11
}
```
## Argument Reference
The following arguments are supported:
* `safe_auth_type` - (Required, ForceNew) The type of safe auth, Ensure the setting scope is for a single sub-account only.
* `user_name` - (Required, ForceNew) The user name.
* `safe_auth_exempt_duration` - (Optional, ForceNew) The exempt duration of safe auth, Ensure the setting scope is for a single sub-account only.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `safe_auth_close` - The status of safe auth, Ensure the setting scope is for a single sub-account only.
* `user_id` - The user id.


## Import
Iam SecurityConfig key don't support import

