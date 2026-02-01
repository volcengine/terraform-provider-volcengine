---
subcategory: "CLOUD_IDENTITY"
layout: "volcengine"
page_title: "Volcengine: volcengine_cloud_identity_permission_set_provisioning"
sidebar_current: "docs-volcengine-resource-cloud_identity_permission_set_provisioning"
description: |-
  Provides a resource to manage cloud identity permission set provisioning
---
# volcengine_cloud_identity_permission_set_provisioning
Provides a resource to manage cloud identity permission set provisioning
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
  permission_set_id    = volcengine_cloud_identity_permission_set.foo.id
  target_id            = "210005****"
  principal_type       = "User"
  principal_id         = volcengine_cloud_identity_user.foo.id
  deprovision_strategy = "None"
}

# It is not recommended to use this resource to provision the permission_set.
# When the `volcengine_cloud_identity_permission_set` is updated, you can use this resource to provision the permission set.
# When deleting this resource, resource `volcengine_cloud_identity_permission_set_assignment` must be deleted first, and the `deprovision_strategy` of `volcengine_cloud_identity_permission_set_assignment` should be set as `None`.
resource "volcengine_cloud_identity_permission_set_provisioning" "foo" {
  permission_set_id   = volcengine_cloud_identity_permission_set.foo.id
  target_id           = "210005****"
  provisioning_status = "Provisioned"
}
```
## Argument Reference
The following arguments are supported:
* `permission_set_id` - (Required, ForceNew) The id of the cloud identity permission set.
* `provisioning_status` - (Required) The target provisioning status of the cloud identity permission set. This field must be specified as `Provisioned` in order to provision the updated permission set. 
When deleting this resource, resource `volcengine_cloud_identity_permission_set_assignment` must be deleted first.
* `target_id` - (Required, ForceNew) The target account id of the cloud identity permission set provisioning.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
CloudIdentityPermissionSetProvisioning can be imported using the permission_set_id:target_id, e.g.
```
$ terraform import volcengine_cloud_identity_permission_set_provisioning.default resource_id
```

