---
subcategory: "NAS"
layout: "volcengine"
page_title: "Volcengine: volcengine_nas_mount_points"
sidebar_current: "docs-volcengine-datasource-nas_mount_points"
description: |-
  Use this data source to query detailed information of nas mount points
---
# volcengine_nas_mount_points
Use this data source to query detailed information of nas mount points
## Example Usage
```hcl
data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-project1"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-subnet-test-2"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

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

data "volcengine_nas_zones" "foo" {

}

resource "volcengine_nas_file_system" "foo" {
  file_system_name = "acc-test-fs"
  description      = "acc-test"
  zone_id          = data.volcengine_nas_zones.foo.zones[0].id
  capacity         = 103
  project_name     = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_nas_mount_point" "foo" {
  file_system_id      = volcengine_nas_file_system.foo.id
  mount_point_name    = "acc-test"
  permission_group_id = volcengine_nas_permission_group.foo.id
  subnet_id           = volcengine_subnet.foo.id
}

data "volcengine_nas_mount_points" "foo" {
  file_system_id = volcengine_nas_file_system.foo.id
  mount_point_id = volcengine_nas_mount_point.foo.mount_point_id
}
```
## Argument Reference
The following arguments are supported:
* `file_system_id` - (Required) The id of the file system.
* `mount_point_id` - (Optional) The id of the mount point.
* `mount_point_name` - (Optional) The name of the mount point.
* `output_file` - (Optional) File name where to save data source results.
* `vpcs_id` - (Optional) The id of the vpc.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `mount_points` - List of mount points.
    * `create_time` - The creation time of the mount point.
    * `domain` - The dns address.
    * `ip` - The address of the mount point.
    * `mount_point_id` - The id of the mount point.
    * `mount_point_name` - The name of the mount point.
    * `permission_group` - The struct of the permission group.
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
    * `status` - The status of the mount point.
    * `subnet_id` - The id of the subnet.
    * `subnet_name` - The name of the subnet.
    * `update_time` - The update time of the mount point.
    * `vpc_id` - The id of the vpc.
    * `vpc_name` - The name of the vpc.
* `total_count` - The total count of nas mount points query.


