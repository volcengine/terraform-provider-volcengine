---
subcategory: "NAS"
layout: "volcengine"
page_title: "Volcengine: volcengine_nas_auto_snapshot_policies"
sidebar_current: "docs-volcengine-datasource-nas_auto_snapshot_policies"
description: |-
  Use this data source to query detailed information of nas auto snapshot policies
---
# volcengine_nas_auto_snapshot_policies
Use this data source to query detailed information of nas auto snapshot policies
## Example Usage
```hcl
resource "volcengine_nas_auto_snapshot_policy" "foo" {
  auto_snapshot_policy_name = "acc-test-auto_snapshot_policy"
  repeat_weekdays           = "1,3,5,7"
  time_points               = "0,7,17"
  retention_days            = 20
}

data "volcengine_nas_auto_snapshot_policies" "foo" {
  auto_snapshot_policy_id = volcengine_nas_auto_snapshot_policy.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `auto_snapshot_policy_id` - (Optional) The id of auto snapshot policy.
* `auto_snapshot_policy_name` - (Optional) The name of auto snapshot policy.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `auto_snapshot_polices` - The collection of query.
    * `auto_snapshot_policy_id` - The ID of auto snapshot policy.
    * `auto_snapshot_policy_name` - The name of auto snapshot policy.
    * `create_time` - The create time of auto snapshot policy.
    * `file_system_count` - The count of file system which auto snapshot policy bind.
    * `id` - The ID of auto snapshot policy.
    * `repeat_weekdays` - The repeat weekdays of auto snapshot policy. Unit: day.
    * `retention_days` - The retention days of auto snapshot policy. Unit: day.
    * `status` - The status of auto snapshot policy.
    * `time_points` - The time points of auto snapshot policy. Unit: hour.
* `total_count` - The total count of query.


