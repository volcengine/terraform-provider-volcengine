---
subcategory: "NAS"
layout: "volcengine"
page_title: "Volcengine: volcengine_nas_file_systems"
sidebar_current: "docs-volcengine-datasource-nas_file_systems"
description: |-
  Use this data source to query detailed information of nas file systems
---
# volcengine_nas_file_systems
Use this data source to query detailed information of nas file systems
## Example Usage
```hcl
data "volcengine_nas_zones" "foo" {

}

resource "volcengine_nas_file_system" "foo" {
  file_system_name = "acc-test-fs-${count.index}"
  description      = "acc-test"
  zone_id          = data.volcengine_nas_zones.foo.zones[0].id
  capacity         = 103
  project_name     = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
  count = 3
}

data "volcengine_nas_file_systems" "foo" {
  ids = volcengine_nas_file_system.foo[*].id
}
```
## Argument Reference
The following arguments are supported:
* `charge_type` - (Optional) The charge type of nas file system.
* `file_system_name` - (Optional) The name of nas file system. This field supports fuzzy queries.
* `ids` - (Optional) A list of nas file system ids.
* `mount_point_id` - (Optional) The mount point id of nas file system.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `permission_group_id` - (Optional) The permission group id of nas file system.
* `project_name` - (Optional) The project name of nas file system.
* `protocol_type` - (Optional) The protocol type of nas file system.
* `status` - (Optional) The status of nas file system.
* `storage_type` - (Optional) The storage type of nas file system.
* `tags` - (Optional) Tags.
* `zone_id` - (Optional) The zone id of nas file system.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `file_systems` - The collection of query.
    * `capacity` - The capacity of the nas file system.
        * `total` - The total capacity of the nas file system. Unit: GiB.
        * `used` - The used capacity of the nas file system. Unit: MiB.
    * `charge_type` - The charge type of the nas file system.
    * `create_time` - The create time of the nas file system.
    * `description` - The description of the nas file system.
    * `file_system_id` - The id of the nas file system.
    * `file_system_name` - The name of the nas file system.
    * `file_system_type` - The type of the nas file system.
    * `id` - The id of the nas file system.
    * `project_name` - The project name of the nas file system.
    * `protocol_type` - The protocol type of the nas file system.
    * `region_id` - The region id of the nas file system.
    * `snapshot_count` - The snapshot count of the nas file system.
    * `status` - The status of the nas file system.
    * `storage_type` - The storage type of the nas file system.
    * `tags` - Tags of the nas file system.
        * `key` - The Key of Tags.
        * `type` - The Type of Tags.
        * `value` - The Value of Tags.
    * `update_time` - The update time of the nas file system.
    * `version` - The version of the nas file system.
    * `zone_id` - The zone id of the nas file system.
    * `zone_name` - The zone name of the nas file system.
* `total_count` - The total count of query.


