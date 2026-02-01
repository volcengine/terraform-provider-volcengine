---
subcategory: "NAS"
layout: "volcengine"
page_title: "Volcengine: volcengine_nas_snapshot"
sidebar_current: "docs-volcengine-resource-nas_snapshot"
description: |-
  Provides a resource to manage nas snapshot
---
# volcengine_nas_snapshot
Provides a resource to manage nas snapshot
## Example Usage
```hcl
resource "volcengine_nas_snapshot" "foo" {
  file_system_id = "enas-cnbj5c18f02afe0e"
  snapshot_name  = "tfsnap3"
  description    = "desc2"
}
```
## Argument Reference
The following arguments are supported:
* `file_system_id` - (Required, ForceNew) The ID of the file system.
* `snapshot_name` - (Required) The name of snapshot.
* `description` - (Optional) The description of snapshot.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `create_time` - The create time of snapshot.
* `file_system_name` - The name of file system.
* `is_encrypt` - Whether is encrypt.
* `progress` - The progress of snapshot.
* `snapshot_type` - The type of snapshot.
* `source_size` - The size of source.
* `source_version` - The source version info.
* `status` - The status of snapshot.
* `zone_id` - The ID of zone.


## Import
Nas Snapshot can be imported using the id, e.g.
```
$ terraform import volcengine_nas_snapshot.default snap-472a716f****
```

