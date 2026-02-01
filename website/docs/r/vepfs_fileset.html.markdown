---
subcategory: "VEPFS"
layout: "volcengine"
page_title: "Volcengine: volcengine_vepfs_fileset"
sidebar_current: "docs-volcengine-resource-vepfs_fileset"
description: |-
  Provides a resource to manage vepfs fileset
---
# volcengine_vepfs_fileset
Provides a resource to manage vepfs fileset
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

resource "volcengine_vepfs_fileset" "foo" {
  file_system_id = volcengine_vepfs_file_system.foo.id
  fileset_name   = "acc-test-fileset"
  fileset_path   = "/tf-test/"
  max_iops       = 100
  max_bandwidth  = 10
  file_limit     = 20
  capacity_limit = 30
}
```
## Argument Reference
The following arguments are supported:
* `file_system_id` - (Required, ForceNew) The id of the vepfs file system.
* `fileset_name` - (Required) The name of the vepfs fileset.
* `fileset_path` - (Required, ForceNew) The path of the vepfs fileset.
* `capacity_limit` - (Optional) The capacity limit of the vepfs fileset. Unit: Gib.
* `file_limit` - (Optional) The file number limit of the vepfs fileset.
* `max_bandwidth` - (Optional) The max bandwidth qos limit of the vepfs fileset. Unit: MB/s.
* `max_iops` - (Optional) The max IOPS qos limit of the vepfs fileset.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `capacity_used` - The used capacity of the vepfs fileset. Unit: GiB.
* `create_time` - The create time of the vepfs fileset.
* `file_used` - The used file number of the vepfs fileset.
* `max_inode_num` - The max number of inode in the vepfs fileset.
* `status` - The status of the vepfs fileset.


## Import
VepfsFileset can be imported using the file_system_id:fileset_id, e.g.
```
$ terraform import volcengine_vepfs_fileset.default file_system_id:fileset_id
```

