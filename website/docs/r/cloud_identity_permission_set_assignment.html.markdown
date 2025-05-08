---
subcategory: "CLOUD_IDENTITY"
layout: "volcengine"
page_title: "Volcengine: volcengine_cloud_identity_permission_set_assignment"
sidebar_current: "docs-volcengine-resource-cloud_identity_permission_set_assignment"
description: |-
  Provides a resource to manage cloud identity permission set assignment
---
# volcengine_cloud_identity_permission_set_assignment
Provides a resource to manage cloud identity permission set assignment
## Example Usage
```hcl
resource "volcengine_cloud_identity_permission_set" "foo" {
  name             = "acc-test-permission_set"
  description      = "tf"
  session_duration = 5000
  permission_policies {
    permission_policy_type = "System"
    permission_policy_name = "AdministratorAccess"
    inline_policy_document = ""
  }
  permission_policies {
    permission_policy_type = "System"
    permission_policy_name = "ReadOnlyAccess"
    inline_policy_document = ""
  }
  permission_policies {
    permission_policy_type = "Inline"
    inline_policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"auto_scaling:DescribeScalingGroups\"],\"Resource\":[\"*\"]}]}"
  }
}

resource "volcengine_cloud_identity_user" "foo" {
  user_name    = "acc-test-user"
  display_name = "tf-test-user"
  description  = "tf"
  email        = "88@qq.com"
  phone        = "181"
}

resource "volcengine_cloud_identity_permission_set_assignment" "foo" {
  permission_set_id = volcengine_cloud_identity_permission_set.foo.id
  target_id         = "210026****"
  principal_type    = "User"
  principal_id      = volcengine_cloud_identity_user.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `permission_set_id` - (Required, ForceNew) The id of the cloud identity permission set.
* `principal_id` - (Required, ForceNew) The principal id of the cloud identity permission set. When the `principal_type` is `User`, this field is specified to `UserId`. When the `principal_type` is `Group`, this field is specified to `GroupId`.
* `principal_type` - (Required, ForceNew) The principal type of the cloud identity permission set. Valid values: `User`, `Group`.
* `target_id` - (Required, ForceNew) The target account id of the cloud identity permission set assignment.
* `deprovision_strategy` - (Optional) The deprovision strategy when deleting the cloud identity permission set assignment. Valid values: `DeprovisionForLastPermissionSetOnAccount`, `None`. Default is `DeprovisionForLastPermissionSetOnAccount`. 
When the `deprovision_strategy` is `DeprovisionForLastPermissionSetOnAccount`, and the permission set assignment to be deleted is the last assignment for the same account, this option is used for the DeprovisionPermissionSet operation.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
CloudIdentityPermissionSetAssignment can be imported using the permission_set_id:target_id:principal_type:principal_id, e.g.
```
$ terraform import volcengine_cloud_identity_permission_set_assignment.default resource_id
```

