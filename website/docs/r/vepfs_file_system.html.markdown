---
subcategory: "VEPFS"
layout: "volcengine"
page_title: "Volcengine: volcengine_vepfs_file_system"
sidebar_current: "docs-volcengine-resource-vepfs_file_system"
description: |-
  Provides a resource to manage vepfs file system
---
# volcengine_vepfs_file_system
Provides a resource to manage vepfs file system
## Example Usage
```hcl
resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = "cn-beijing-a"
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_vepfs_file_system" "foo" {
  file_system_name = "acc-test-file-system"
  subnet_id        = volcengine_subnet.foo.id
  store_type       = "Advance_100"
  description      = "tf-test"
  capacity         = 12
  project          = "default"
  enable_restripe  = false
  tags {
    key   = "k1"
    value = "v1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `capacity` - (Required) The capacity of the vepfs file system.
* `file_system_name` - (Required) The name of the vepfs file system.
* `store_type` - (Required, ForceNew) The store type of the vepfs file system. Valid values: `Advance_100`, `Performance`, `Intelligent_Computing`.
* `subnet_id` - (Required, ForceNew) The subnet id of the vepfs file system.
* `description` - (Optional) The description info of the vepfs file system.
* `enable_restripe` - (Optional) Whether to enable data balance after capacity expansion. This filed is valid only when expanding capacity.
* `project` - (Optional, ForceNew) The project of the vepfs file system.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `account_id` - The id of the account.
* `bandwidth` - The bandwidth info of the vepfs file system.
* `charge_status` - The charge status of the vepfs file system.
* `charge_type` - The charge type of the vepfs file system.
* `create_time` - The create time of the vepfs file system.
* `expire_time` - The expire time of the vepfs file system.
* `file_system_type` - The type of the vepfs file system.
* `free_time` - The free time of the vepfs file system.
* `last_modify_time` - The last modify time of the vepfs file system.
* `protocol_type` - The protocol type of the vepfs file system.
* `region_id` - The id of the region.
* `status` - The status of the vepfs file system.
* `stop_service_time` - The stop service time of the vepfs file system.
* `store_type_cn` - The store type cn name of the vepfs file system.
* `version` - The version info of the vepfs file system.
* `zone_id` - The id of the zone.
* `zone_name` - The name of the zone.


## Import
VepfsFileSystem can be imported using the id, e.g.
```
$ terraform import volcengine_vepfs_file_system.default resource_id
```

