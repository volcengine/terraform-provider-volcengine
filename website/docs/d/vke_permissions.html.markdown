---
subcategory: "VKE"
layout: "volcengine"
page_title: "Volcengine: volcengine_vke_permissions"
sidebar_current: "docs-volcengine-datasource-vke_permissions"
description: |-
  Use this data source to query detailed information of vke permissions
---
# volcengine_vke_permissions
Use this data source to query detailed information of vke permissions
## Example Usage
```hcl
data "volcengine_vke_permissions" "foo" {
  ids          = ["apd10o9jhqqno0ba25****"]
  grantee_type = "User"
}
```
## Argument Reference
The following arguments are supported:
* `cluster_ids` - (Optional) A list of Cluster IDs.
* `grantee_ids` - (Optional) A list of Grantee IDs.
* `grantee_type` - (Optional) The type of Grantee. Valid values: `User`, `Role`.
* `ids` - (Optional) A list of RBAC Permission IDs.
* `namespaces` - (Optional) A list of Namespaces.
* `output_file` - (Optional) File name where to save data source results.
* `role_names` - (Optional) A list of RBAC Role Names.
* `status` - (Optional) The status of RBAC Permission.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `access_policies` - The collection of query.
    * `authorized_at` - The authorized time of the RBAC Permission.
    * `authorizer_id` - The ID of the Authorizer.
    * `authorizer_name` - The name of the Authorizer.
    * `authorizer_type` - The type of the Authorizer.
    * `cluster_id` - The ID of the Cluster.
    * `granted_at` - The granted time of the RBAC Permission.
    * `grantee_id` - The ID of the Grantee.
    * `grantee_type` - The type of the Grantee.
    * `id` - The id of the RBAC Permission.
    * `is_custom_role` - Whether the RBAC Role is custom role.
    * `kube_role_binding_name` - The name of the Kube Role Binding.
    * `message` - The message of the RBAC Permission.
    * `namespace` - The Namespace of the RBAC Permission.
    * `revoked_at` - The revoked time of the RBAC Permission.
    * `role_name` - The name of the RBAC Role.
    * `status` - The status of the RBAC Permission.
* `total_count` - The total count of query.


