---
subcategory: "NAS"
layout: "volcengine"
page_title: "Volcengine: volcengine_nas_auto_snapshot_policy"
sidebar_current: "docs-volcengine-resource-nas_auto_snapshot_policy"
description: |-
  Provides a resource to manage nas auto snapshot policy
---
# volcengine_nas_auto_snapshot_policy
Provides a resource to manage nas auto snapshot policy
## Example Usage
```hcl
resource "volcengine_nas_auto_snapshot_policy" "foo" {
  auto_snapshot_policy_name = "acc-test-auto_snapshot_policy"
  repeat_weekdays           = "1,3,5,7"
  time_points               = "0,7,17"
  retention_days            = 20
}
```
## Argument Reference
The following arguments are supported:
* `auto_snapshot_policy_name` - (Required) The name of the auto snapshot policy.
* `repeat_weekdays` - (Required) The repeat weekdays of the auto snapshot policy. Support setting multiple dates, separated by English commas. Valid values: `1` ~ `7`.
* `time_points` - (Required) The time points of the auto snapshot policy. Support setting multiple dates, separated by English commas. Valid values: `0` ~ `23`.
* `retention_days` - (Optional) The retention days of the auto snapshot policy. Valid values: -1(permanent) or 1 ~ 65536. Default is 30.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `create_time` - The create time of auto snapshot policy.
* `file_system_count` - The count of file system which auto snapshot policy bind.
* `status` - The status of auto snapshot policy.


## Import
NasAutoSnapshotPolicy can be imported using the id, e.g.
```
$ terraform import volcengine_nas_auto_snapshot_policy.default resource_id
```

