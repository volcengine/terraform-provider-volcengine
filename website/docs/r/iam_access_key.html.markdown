---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_access_key"
sidebar_current: "docs-volcengine-resource-iam_access_key"
description: |-
  Provides a resource to manage iam access key
---
# volcengine_iam_access_key
Provides a resource to manage iam access key
## Example Usage
```hcl
resource "volcengine_iam_access_key" "foo" {
  user_name = "jonny"
  status    = "active"
}
```
## Argument Reference
The following arguments are supported:
* `status` - (Optional) The status of the access key, Optional choice contains `active` or `inactive`.
* `user_name` - (Optional, ForceNew) The user name. If not specified, the current user is used.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `access_key_id` - The access key id.
* `create_date` - The create date of the access key.
* `secret_access_key` - The secret access key.
* `update_date` - The update date of the access key.


## Import
Iam access key don't support import

