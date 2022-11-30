---
subcategory: "AUTOSCALING"
layout: "volcengine"
page_title: "Volcengine: volcengine_scaling_activities"
sidebar_current: "docs-volcengine-datasource-scaling_activities"
description: |-
  Use this data source to query detailed information of scaling activities
---
# volcengine_scaling_activities
Use this data source to query detailed information of scaling activities
## Example Usage
```hcl
data "volcengine_scaling_activities" "default" {
  scaling_group_id = "scg-ybqm0b6kcigh9zu9ce6t"
}
```
## Argument Reference
The following arguments are supported:
* `scaling_group_id` - (Required) A Id of Scaling Group.
* `end_time` - (Optional) An end time to start a Scaling Activity.
* `ids` - (Optional) A list of Scaling Activity IDs.
* `output_file` - (Optional) File name where to save data source results.
* `start_time` - (Optional) A start time to start a Scaling Activity.
* `status_code` - (Optional) A status code of Scaling Activity. Valid values: Init, Running, Success, PartialSuccess, Error, Rejected, Exception.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `activities` - The collection of Scaling Activity query.
    * `activity_type` - The Actual Type.
    * `actual_adjust_instance_number` - The Actual Adjustment Instance Number.
    * `cooldown` - The Cooldown time.
    * `created_at` - The create time of Scaling Activity.
    * `current_instance_number` - The Current Instance Number.
    * `expected_run_time` - The expected run time of Scaling Activity.
    * `id` - The ID of Scaling Activity.
    * `max_instance_number` - The Max Instance Number.
    * `min_instance_number` - The Min Instance Number.
    * `related_instances` - The related instances.
        * `instance_id` - The Instance ID.
        * `message` - The message of Instance.
        * `operate_type` - The Operation Type.
        * `status` - The Status.
    * `result_msg` - The Result of Scaling Activity.
    * `scaling_activity_id` - The ID of Scaling Activity.
    * `scaling_group_id` - The scaling group Id.
    * `status_code` - The Status Code of Scaling Activity.
    * `stopped_at` - The stopped time of Scaling Activity.
    * `task_category` - The task category of Scaling Activity.
* `total_count` - The total count of Scaling Activity query.


