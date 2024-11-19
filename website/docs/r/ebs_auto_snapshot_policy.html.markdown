---
subcategory: "EBS"
layout: "volcengine"
page_title: "Volcengine: volcengine_ebs_auto_snapshot_policy"
sidebar_current: "docs-volcengine-resource-ebs_auto_snapshot_policy"
description: |-
  Provides a resource to manage ebs auto snapshot policy
---
# volcengine_ebs_auto_snapshot_policy
Provides a resource to manage ebs auto snapshot policy
## Example Usage
```hcl
resource "volcengine_ebs_auto_snapshot_policy" "foo" {
  auto_snapshot_policy_name = "acc-test-auto-snapshot-policy"
  time_points               = [1, 5, 9]
  retention_days            = -1
  repeat_weekdays           = [2, 6]
  #  repeat_days               = 5
  project_name = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `auto_snapshot_policy_name` - (Required) The name of the auto snapshot policy.
* `retention_days` - (Required) The retention days of the auto snapshot. Valid values: -1 and 1~65536. `-1` means permanently preserving the snapshot.
* `time_points` - (Required) The creation time points of the auto snapshot policy. The value range is `0~23`, representing a total of 24 time points from 00:00 to 23:00, for example, 1 represents 01:00.
* `project_name` - (Optional) The project name of the auto snapshot policy.
* `repeat_days` - (Optional) Create snapshots repeatedly on a daily basis, with intervals of a certain number of days between each snapshot. The value range is `1-30`. Only one of `repeat_weekdays, repeat_days` can be specified.
* `repeat_weekdays` - (Optional) The date of creating snapshot repeatedly by week. The value range is `1-7`, for example, 1 represents Monday. Only one of `repeat_weekdays, repeat_days` can be specified.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `created_at` - The creation time of the auto snapshot policy.
* `status` - The status of the auto snapshot policy.
* `updated_at` - The updated time of the auto snapshot policy.
* `volume_nums` - The number of volumes associated with the auto snapshot policy.


## Import
EbsAutoSnapshotPolicy can be imported using the id, e.g.
```
$ terraform import volcengine_ebs_auto_snapshot_policy.default resource_id
```

