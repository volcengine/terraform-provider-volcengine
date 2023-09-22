---
subcategory: "NAS"
layout: "volcengine"
page_title: "Volcengine: volcengine_nas_snapshots"
sidebar_current: "docs-volcengine-datasource-nas_snapshots"
description: |-
  Use this data source to query detailed information of nas snapshots
---
# volcengine_nas_snapshots
Use this data source to query detailed information of nas snapshots
## Example Usage
```hcl
data "volcengine_nas_snapshots" "default" {
  file_system_id = "enas-cnbj5c18f02afe0e"
  ids            = ["snap-022c648fed8b", "snap-e53591b05fbd"]
}
```
## Argument Reference
The following arguments are supported:
* `file_system_id` - (Optional) The ID of file system.
* `ids` - (Optional) A list of Snapshot IDs.
* `output_file` - (Optional) File name where to save data source results.
* `snapshot_name` - (Optional) The name of snapshot.
* `snapshot_type` - (Optional) The type of snapshot.
* `status` - (Optional) The status of snapshot.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `snapshots` - The collection of query.
    * `create_time` - The create time of snapshot.
    * `description` - The description of snapshot.
    * `file_system_id` - The id of file system.
    * `file_system_name` - The name of file system.
    * `id` - The ID of snapshot.
    * `is_encrypt` - Whether is encrypt.
    * `progress` - The progress of snapshot.
    * `retention_days` - The retention days of snapshot.
    * `snapshot_id` - The ID of snapshot.
    * `snapshot_name` - The name of snapshot.
    * `snapshot_type` - The type of snapshot.
    * `source_size` - The size of source.
    * `source_version` - The source version info.
    * `status` - The status of snapshot.
    * `zone_id` - The ID of zone.
* `total_count` - The total count of query.


