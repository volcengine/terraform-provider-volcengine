---
subcategory: "NAS"
layout: "volcengine"
page_title: "Volcengine: volcengine_nas_mount_point"
sidebar_current: "docs-volcengine-resource-nas_mount_point"
description: |-
  Provides a resource to manage nas mount point
---
# volcengine_nas_mount_point
Provides a resource to manage nas mount point
## Example Usage
```hcl
data "volcengine_nas_zones" "foo" {

}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-project1"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-subnet-test-2"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_nas_zones.foo.zones[0].id
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
```
## Argument Reference
The following arguments are supported:
* `file_system_id` - (Required, ForceNew) The file system id.
* `mount_point_name` - (Required) The mount point name.
* `permission_group_id` - (Required) The permission group id.
* `subnet_id` - (Required, ForceNew) The subnet id.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `mount_point_id` - The mount point id.


## Import
Nas Mount Point can be imported using the file system id and mount point id, e.g.
```
$ terraform import volcengine_nas_mount_point.default enas-cnbj18bcb923****:mount-a6ee****
```

