---
subcategory: "CLOUDFS"
layout: "volcengine"
page_title: "Volcengine: volcengine_cloudfs_file_systems"
sidebar_current: "docs-volcengine-datasource-cloudfs_file_systems"
description: |-
  Use this data source to query detailed information of cloudfs file systems
---
# volcengine_cloudfs_file_systems
Use this data source to query detailed information of cloudfs file systems
## Example Usage
```hcl
data "volcengine_cloudfs_file_systems" "default" {
  fs_name = "tftest2"
}
```
## Argument Reference
The following arguments are supported:
* `fs_name` - (Optional) The name of file system.
* `meta_status` - (Optional) The status of file system.
* `name_regex` - (Optional) A Name Regex of cloudfs.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `file_systems` - The collection of query.
    * `cache_capacity_tib` - The capacity of cache.
    * `cache_plan` - The plan of cache.
    * `created_time` - The creation time.
    * `id` - The ID of file system.
    * `mode` - The mode of file system.
    * `mount_point` - The point mount.
    * `name` - The name of file system.
    * `region_id` - The id of region.
    * `security_group_id` - The id of security group.
    * `status` - The status of file system.
    * `subnet_id` - The id of subnet.
    * `tos_bucket` - The tos bucket.
    * `tos_prefix` - The tos prefix.
    * `vpc_id` - The id of vpc.
    * `zone_id` - The id of zone.
* `total_count` - The total count of query.


