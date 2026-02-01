---
subcategory: "VEPFS"
layout: "volcengine"
page_title: "Volcengine: volcengine_vepfs_filesets"
sidebar_current: "docs-volcengine-datasource-vepfs_filesets"
description: |-
  Use this data source to query detailed information of vepfs filesets
---
# volcengine_vepfs_filesets
Use this data source to query detailed information of vepfs filesets
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

data "volcengine_vepfs_filesets" "foo" {
  file_system_id = volcengine_vepfs_file_system.foo.id
  fileset_id     = volcengine_vepfs_fileset.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `file_system_id` - (Required) The id of Vepfs File System.
* `fileset_id` - (Optional) The id of Vepfs Fileset.
* `fileset_name` - (Optional) The name of Vepfs Fileset. This field support fuzzy query.
* `fileset_path` - (Optional) The path of Vepfs Fileset. This field support fuzzy query.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `status` - (Optional) The query status list of Vepfs Fileset.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `filesets` - The collection of query.
    * `bandwidth_qos` - The bandwidth Qos of the vepfs fileset.
    * `capacity_limit` - The capacity limit of the vepfs fileset. Unit: GiB.
    * `capacity_used` - The used capacity of the vepfs fileset. Unit: GiB.
    * `create_time` - The create time of the vepfs fileset.
    * `file_limit` - Quota for the number of files or directories. A return of 0 indicates that there is no quota limit set for the number of directories after the file.
    * `file_used` - The used file number of the vepfs fileset.
    * `fileset_id` - The id of the vepfs fileset.
    * `fileset_name` - The name of the vepfs fileset.
    * `fileset_path` - The path of the vepfs fileset.
    * `id` - The id of the vepfs fileset.
    * `iops_qos` - The IOPS Qos of the vepfs fileset.
    * `max_inode_num` - The max number of inode in the vepfs fileset.
    * `status` - The status of the vepfs fileset.
* `total_count` - The total count of query.


