---
subcategory: "NAS"
layout: "volcengine"
page_title: "Volcengine: volcengine_nas_permission_group"
sidebar_current: "docs-volcengine-resource-nas_permission_group"
description: |-
  Provides a resource to manage nas permission group
---
# volcengine_nas_permission_group
Provides a resource to manage nas permission group
## Example Usage
```hcl
resource "volcengine_nas_permission_group" "foo" {
  permission_group_name = "acc-test1"
  description           = "acctest1"
  permission_rules {
    cidr_ip  = "*"
    rw_mode  = "RW"
    use_mode = "All_squash"
  }
  permission_rules {
    cidr_ip  = "192.168.0.0"
    rw_mode  = "RO"
    use_mode = "No_all_squash"
  }
}
```
## Argument Reference
The following arguments are supported:
* `permission_group_name` - (Required) The name of the permission group.
* `description` - (Optional) The description of the permission group.
* `permission_rules` - (Optional) The list of permissions rules.

The `permission_rules` object supports the following:

* `cidr_ip` - (Required) Client IP addresses that are allowed access.
* `rw_mode` - (Required) Permission group read and write rules. The value description is as follows:
`RW`: Allows reading and writing.
`RO`: read-only mode.
* `use_mode` - (Required) Permission group user permissions. The value description is as follows:
`All_squash`: All access users are mapped to anonymous users or user groups.
`No_all_squash`: The access user is first matched with the local user, and then mapped to an anonymous user or user group after the match fails.
`Root_squash`: Map the Root user as an anonymous user or user group.
`No_root_squash`: The Root user maintains the Root account authority.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `permission_group_id` - The id of the permission group.


## Import
Nas Permission Group can be imported using the id, e.g.
```
$ terraform import volcengine_nas_permission_group.default pgroup-1f85db2c****
```

