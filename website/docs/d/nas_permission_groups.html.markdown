---
subcategory: "NAS"
layout: "volcengine"
page_title: "Volcengine: volcengine_nas_permission_groups"
sidebar_current: "docs-volcengine-datasource-nas_permission_groups"
description: |-
  Use this data source to query detailed information of nas permission groups
---
# volcengine_nas_permission_groups
Use this data source to query detailed information of nas permission groups
## Example Usage
```hcl
resource "volcengine_nas_permission_group" "foo" {
  permission_group_name = "acc-test"
  description           = "acctest"
  permission_rules {
    cidr_ip  = "*"
    rw_mode  = "RW"
    use_mode = "All_squash"
  }
  permission_rules {
    cidr_ip  = "192.168.0.0"
    rw_mode  = "RO"
    use_mode = "All_squash"
  }
}

data "volcengine_nas_permission_groups" "default" {
  filters {
    key   = "PermissionGroupId"
    value = volcengine_nas_permission_group.foo.id
  }
}
```
## Argument Reference
The following arguments are supported:
* `filters` - (Optional) Filter permission groups for specified characteristics.
* `output_file` - (Optional) File name where to save data source results.

The `filters` object supports the following:

* `key` - (Required) Filters permission groups for specified characteristics based on attributes. The parameters that support filtering are as follows: `PermissionGroupName`, `PermissionGroupId`.
* `value` - (Required) The value of the filter item.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `permission_groups` - The list of permissions groups.
    * `create_time` - The creation time of the permission group.
    * `description` - The description of the permission group.
    * `file_system_count` - The number of the file system.
    * `file_system_type` - The file system type of the permission group.
    * `mount_points` - The list of the mount point.
        * `file_system_id` - The id of the file system.
        * `mount_point_id` - The id of the mount point.
        * `mount_point_name` - The name of the mount point.
    * `permission_group_id` - The id of the permission group.
    * `permission_group_name` - The name of the permission group.
    * `permission_rule_count` - The number of the permission rule.
    * `permission_rules` - The list of permissions rules.
        * `cidr_ip` - Client IP addresses that are allowed access.
        * `permission_rule_id` - The id of the permission rule.
        * `rw_mode` - Permission group read and write rules.
        * `user_mode` - Permission group user permissions.
* `total_count` - The total count of nas permission groups query.


