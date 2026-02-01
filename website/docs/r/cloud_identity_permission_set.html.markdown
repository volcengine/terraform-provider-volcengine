---
subcategory: "CLOUD_IDENTITY"
layout: "volcengine"
page_title: "Volcengine: volcengine_cloud_identity_permission_set"
sidebar_current: "docs-volcengine-resource-cloud_identity_permission_set"
description: |-
  Provides a resource to manage cloud identity permission set
---
# volcengine_cloud_identity_permission_set
Provides a resource to manage cloud identity permission set
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
```
## Argument Reference
The following arguments are supported:
* `name` - (Required, ForceNew) The name of the cloud identity permission set.
* `description` - (Optional) The description of the cloud identity permission set.
* `permission_policies` - (Optional) The policies of the cloud identity permission set.
* `relay_state` - (Optional) The relay state of the cloud identity permission set.
* `session_duration` - (Optional) The session duration of the cloud identity permission set. Unit: second. Valid value range in 3600~43200.

The `permission_policies` object supports the following:

* `permission_policy_type` - (Required) The type of the cloud identity permission set policy. Valid values: `System`, `Inline`.
* `inline_policy_document` - (Optional) The document of the cloud identity permission set inline policy. When the `permission_policy_type` is `Inline`, this field must be specified.
* `permission_policy_name` - (Optional) The name of the cloud identity permission set system policy. When the `permission_policy_type` is `System`, this field must be specified.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
CloudIdentityPermissionSet can be imported using the id, e.g.
```
$ terraform import volcengine_cloud_identity_permission_set.default resource_id
```

