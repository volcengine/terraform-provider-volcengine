---
subcategory: "NAS"
layout: "volcengine"
page_title: "Volcengine: volcengine_nas_file_system"
sidebar_current: "docs-volcengine-resource-nas_file_system"
description: |-
  Provides a resource to manage nas file system
---
# volcengine_nas_file_system
Provides a resource to manage nas file system
## Example Usage
```hcl
// query available zones in current region
data "volcengine_nas_zones" "foo" {

}

// create nas file system
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

// create vpc
resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

// create subnet
resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_nas_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

// create nas permission group
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

// create nas mount point
resource "volcengine_nas_mount_point" "foo" {
  file_system_id      = volcengine_nas_file_system.foo.id
  mount_point_name    = "acc-test"
  permission_group_id = volcengine_nas_permission_group.foo.id
  subnet_id           = volcengine_subnet.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `capacity` - (Required) The capacity of the nas file system. Unit: GiB.
* `file_system_name` - (Required) The name of the nas file system.
* `zone_id` - (Required, ForceNew) The zone id of the nas file system.
* `description` - (Optional) The description of the nas file system.
* `project_name` - (Optional) The project name of the nas file system.
* `snapshot_id` - (Optional, ForceNew) The snapshot id when creating the nas file system. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `charge_type` - The charge type of the nas file system.
* `create_time` - The create time of the nas file system.
* `file_system_type` - The type of the nas file system.
* `protocol_type` - The protocol type of the nas file system.
* `region_id` - The region id of the nas file system.
* `snapshot_count` - The snapshot count of the nas file system.
* `status` - The status of the nas file system.
* `storage_type` - The storage type of the nas file system.
* `update_time` - The update time of the nas file system.
* `version` - The version of the nas file system.
* `zone_name` - The zone name of the nas file system.


## Import
NasFileSystem can be imported using the id, e.g.
```
$ terraform import volcengine_nas_file_system.default enas-cnbjd3879745****
```

