---
subcategory: "CLOUD_IDENTITY"
layout: "volcengine"
page_title: "Volcengine: volcengine_cloud_identity_user_provisioning"
sidebar_current: "docs-volcengine-resource-cloud_identity_user_provisioning"
description: |-
  Provides a resource to manage cloud identity user provisioning
---
# volcengine_cloud_identity_user_provisioning
Provides a resource to manage cloud identity user provisioning
## Example Usage
```hcl
resource "volcengine_cloud_identity_user" "foo" {
  user_name    = "acc-test-user"
  display_name = "tf-test-user"
  description  = "tf"
  email        = "88@qq.com"
  phone        = "181"
}

resource "volcengine_cloud_identity_user_provisioning" "foo" {
  principal_type           = "User"
  principal_id             = volcengine_cloud_identity_user.foo.id
  target_id                = "210026****"
  description              = "tf"
  identity_source_strategy = "Ignore"
  duplication_strategy     = "KeepBoth"
  duplication_suffix       = "-tf"
  deletion_strategy        = "Delete"
  policy_name              = ["AdministratorAccess"]
}
```
## Argument Reference
The following arguments are supported:
* `deletion_strategy` - (Required, ForceNew) The deletion strategy of the cloud identity user provisioning. Valid values: `Keep`, `Delete`.
* `duplication_strategy` - (Required, ForceNew) The duplication strategy of the cloud identity user provisioning. Valid values: `KeepBoth`, `Takeover`.
* `identity_source_strategy` - (Required, ForceNew) The identity source strategy of the cloud identity user provisioning. Valid values: `Create`, `Ignore`.
* `principal_id` - (Required, ForceNew) The principal id of the cloud identity user provisioning. When the `principal_type` is `User`, this field is specified to `UserId`. When the `principal_type` is `Group`, this field is specified to `GroupId`.
* `principal_type` - (Required, ForceNew) The principal type of the cloud identity user provisioning. Valid values: `User`, `Group`.
* `target_id` - (Required, ForceNew) The target account id of the cloud identity user provisioning.
* `description` - (Optional, ForceNew) The description of the cloud identity user provisioning.
* `duplication_suffix` - (Optional, ForceNew) The duplication suffix of the cloud identity user provisioning. When the `duplication_strategy` is `KeepBoth`, this field must be specified.
* `policy_name` - (Optional) A list of policy name. Valid values: `AdministratorAccess`. This field is valid when the `principal_type` is `User`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `principal_name` - The principal name of the cloud identity user provisioning. When the `principal_type` is `User`, this field is specified to `UserName`. When the `principal_type` is `Group`, this field is specified to `GroupName`.
* `provision_status` - The status of the cloud identity user provisioning.


## Import
CloudIdentityUserProvisioning can be imported using the id, e.g.
```
$ terraform import volcengine_cloud_identity_user_provisioning.default resource_id
```

