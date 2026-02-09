---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_security_configs"
sidebar_current: "docs-volcengine-datasource-iam_security_configs"
description: |-
  Use this data source to query detailed information of iam security configs
---
# volcengine_iam_security_configs
Use this data source to query detailed information of iam security configs
## Example Usage
```hcl
data "volcengine_iam_security_configs" "default" {
  user_name = "jonny"
}
```
## Argument Reference
The following arguments are supported:
* `user_name` - (Required) The user name.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `security_configs` - The collection of security configs.
    * `safe_auth_close` - The status of safe auth.
    * `safe_auth_exempt_duration` - The exempt duration of safe auth.
    * `safe_auth_type` - The type of safe auth.
    * `user_id` - The user id.
    * `user_name` - The user name.


