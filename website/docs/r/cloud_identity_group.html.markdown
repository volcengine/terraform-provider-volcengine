---
subcategory: "CLOUD_IDENTITY"
layout: "volcengine"
page_title: "Volcengine: volcengine_cloud_identity_group"
sidebar_current: "docs-volcengine-resource-cloud_identity_group"
description: |-
  Provides a resource to manage cloud identity group
---
# volcengine_cloud_identity_group
Provides a resource to manage cloud identity group
## Example Usage
```hcl
resource "volcengine_cloud_identity_group" "foo" {
  group_name   = "acc-test-group"
  display_name = "tf-test-group"
  join_type    = "Manual"
  description  = "tf"
}
```
## Argument Reference
The following arguments are supported:
* `group_name` - (Required, ForceNew) The name of the cloud identity group.
* `join_type` - (Required, ForceNew) The user join type of the cloud identity group.
* `description` - (Optional) The description of the cloud identity group.
* `display_name` - (Optional) The display name of the cloud identity group.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `members` - The member user info of the cloud identity group.
    * `description` - The description of the cloud identity user.
    * `display_name` - The display name of the cloud identity user.
    * `email` - The email of the cloud identity user.
    * `identity_type` - The identity type of the cloud identity user.
    * `join_time` - The join time of the cloud identity user.
    * `phone` - The phone of the cloud identity user.
    * `source` - The source of the cloud identity user.
    * `user_id` - The id of the cloud identity user.
    * `user_name` - The name of the cloud identity user.
* `source` - The source of the cloud identity group.


## Import
CloudIdentityGroup can be imported using the id, e.g.
```
$ terraform import volcengine_cloud_identity_group.default resource_id
```

