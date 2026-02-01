---
subcategory: "VEPFS"
layout: "volcengine"
page_title: "Volcengine: volcengine_vepfs_file_systems"
sidebar_current: "docs-volcengine-datasource-vepfs_file_systems"
description: |-
  Use this data source to query detailed information of vepfs file systems
---
# volcengine_vepfs_file_systems
Use this data source to query detailed information of vepfs file systems
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

data "volcengine_vepfs_file_systems" "foo" {
  ids = [volcengine_vepfs_file_system.foo.id]
}
```
## Argument Reference
The following arguments are supported:
* `file_system_name` - (Optional) The Name of Vepfs File System. This field support fuzzy query.
* `ids` - (Optional) A list of Vepfs File System IDs.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `project` - (Optional) The project of Vepfs File System.
* `status` - (Optional) The query status list of Vepfs File System.
* `store_type` - (Optional) The Store Type of Vepfs File System.
* `zone_id` - (Optional) The zone id of File System.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `file_systems` - The collection of query.
    * `account_id` - The id of the account.
    * `bandwidth` - The bandwidth info of the vepfs file system.
    * `capacity_info` - The capacity info of the vepfs file system.
        * `total_tib` - The total size. Unit: TiB.
        * `used_gib` - The used size. Unit: GiB.
    * `charge_status` - The charge status of the vepfs file system.
    * `charge_type` - The charge type of the vepfs file system.
    * `create_time` - The create time of the vepfs file system.
    * `description` - The description of the vepfs file system.
    * `expire_time` - The expire time of the vepfs file system.
    * `file_system_id` - The id of the vepfs file system.
    * `file_system_name` - The name of the vepfs file system.
    * `file_system_type` - The type of the vepfs file system.
    * `free_time` - The free time of the vepfs file system.
    * `id` - The id of the vepfs file system.
    * `last_modify_time` - The last modify time of the vepfs file system.
    * `project` - The project name of the vepfs file system.
    * `protocol_type` - The protocol type of the vepfs file system.
    * `region_id` - The id of the region.
    * `status` - The status of the vepfs file system.
    * `stop_service_time` - The stop service time of the vepfs file system.
    * `store_type_cn` - The store type cn name of the vepfs file system.
    * `store_type` - The store type of the vepfs file system.
    * `tags` - The tags of the vepfs file system.
        * `key` - The Key of Tags.
        * `type` - The Type of Tags.
        * `value` - The Value of Tags.
    * `version` - The version info of the vepfs file system.
    * `zone_id` - The id of the zone.
    * `zone_name` - The name of the zone.
* `total_count` - The total count of query.


