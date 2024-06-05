---
subcategory: "CLOUD_IDENTITY"
layout: "volcengine"
page_title: "Volcengine: volcengine_cloud_identity_user_attachment"
sidebar_current: "docs-volcengine-resource-cloud_identity_user_attachment"
description: |-
  Provides a resource to manage cloud identity user attachment
---
# volcengine_cloud_identity_user_attachment
Provides a resource to manage cloud identity user attachment
## Example Usage
```hcl
resource "volcengine_cloud_identity_group" "foo" {
  group_name   = "acc-test-group"
  display_name = "tf-test-group"
  join_type    = "Manual"
  description  = "tf"
}

resource "volcengine_cloud_identity_user" "foo" {
  user_name    = "acc-test-user"
  display_name = "tf-test-user"
  description  = "tf"
  email        = "88@qq.com"
  phone        = "181"
}

resource "volcengine_cloud_identity_user_attachment" "foo" {
  user_id  = volcengine_cloud_identity_user.foo.id
  group_id = volcengine_cloud_identity_group.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `group_id` - (Required, ForceNew) The id of the cloud identity group.
* `user_id` - (Required, ForceNew) The id of the cloud identity user.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
CloudIdentityUserAttachment can be imported using the group_id:user_id, e.g.
```
$ terraform import volcengine_cloud_identity_user_attachment.default resource_id
```

