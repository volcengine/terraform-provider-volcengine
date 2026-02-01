---
subcategory: "EBS"
layout: "volcengine"
page_title: "Volcengine: volcengine_ebs_auto_snapshot_policies"
sidebar_current: "docs-volcengine-datasource-ebs_auto_snapshot_policies"
description: |-
  Use this data source to query detailed information of ebs auto snapshot policies
---
# volcengine_ebs_auto_snapshot_policies
Use this data source to query detailed information of ebs auto snapshot policies
## Example Usage
```hcl
resource "volcengine_ebs_auto_snapshot_policy" "foo" {
  auto_snapshot_policy_name = "acc-test-auto-snapshot-policy"
  time_points               = [1, 5, 9]
  retention_days            = -1
  repeat_weekdays           = [2, 6]
  project_name              = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
  count = 2
}

data "volcengine_ebs_auto_snapshot_policies" "foo" {
  ids = volcengine_ebs_auto_snapshot_policy.foo[*].id
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of auto snapshot policy IDs.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name of auto snapshot policy.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `auto_snapshot_policies` - The collection of query.
    * `auto_snapshot_policy_id` - The id of the auto snapshot policy.
    * `auto_snapshot_policy_name` - The name of the auto snapshot policy.
    * `created_at` - The creation time of the auto snapshot policy.
    * `id` - The id of the auto snapshot policy.
    * `project_name` - The project name of the auto snapshot policy.
    * `repeat_days` - Create snapshots repeatedly on a daily basis, with intervals of a certain number of days between each snapshot.
    * `repeat_weekdays` - The date of creating snapshot repeatedly by week. The value range is `1-7`, for example, 1 represents Monday.
    * `retention_days` - The retention days of the auto snapshot. `-1` means permanently preserving the snapshot.
    * `status` - The status of the auto snapshot policy.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `time_points` - The creation time points of the auto snapshot policy. The value range is `0~23`, representing a total of 24 time points from 00:00 to 23:00, for example, 1 represents 01:00.
    * `updated_at` - The updated time of the auto snapshot policy.
    * `volume_nums` - The number of volumes associated with the auto snapshot policy.
* `total_count` - The total count of query.


