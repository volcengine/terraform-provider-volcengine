---
subcategory: "CLOUD_IDENTITY"
layout: "volcengine"
page_title: "Volcengine: volcengine_cloud_identity_user"
sidebar_current: "docs-volcengine-resource-cloud_identity_user"
description: |-
  Provides a resource to manage cloud identity user
---
# volcengine_cloud_identity_user
Provides a resource to manage cloud identity user
## Example Usage
```hcl
resource "volcengine_cloud_identity_user" "foo" {
  user_name    = "acc-test-user"
  display_name = "tf-test-user"
  description  = "tf"
  email        = "88@qq.com"
  phone        = "1810000****"
}
```
## Argument Reference
The following arguments are supported:
* `user_name` - (Required) The name of the cloud identity user.
* `description` - (Optional) The description of the cloud identity user.
* `display_name` - (Optional) The display name of the cloud identity user.
* `email` - (Optional) The email of the cloud identity user.
* `phone` - (Optional) The phone of the cloud identity user. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `identity_type` - The identity type of the cloud identity user.
* `source` - The source of the cloud identity user.


## Import
CloudIdentityUser can be imported using the id, e.g.
```
$ terraform import volcengine_cloud_identity_user.default resource_id
```

