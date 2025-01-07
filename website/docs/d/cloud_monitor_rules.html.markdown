---
subcategory: "CLOUD_MONITOR"
layout: "volcengine"
page_title: "Volcengine: volcengine_cloud_monitor_rules"
sidebar_current: "docs-volcengine-datasource-cloud_monitor_rules"
description: |-
  Use this data source to query detailed information of cloud monitor rules
---
# volcengine_cloud_monitor_rules
Use this data source to query detailed information of cloud monitor rules
## Example Usage
```hcl
data "volcengine_cloud_monitor_rules" "foo" {
  ids = ["174402785374661****"]
}
```
## Argument Reference
The following arguments are supported:
* `alert_state` - (Optional) The alert state of the cloud monitor rule. Valid values: `altering`, `normal`.
* `enable_state` - (Optional) The enable state of the cloud monitor rule. Valid values: `enable`, `disable`.
* `ids` - (Optional) A list of cloud monitor ids.
* `level` - (Optional) The level of the cloud monitor rule. Valid values: `critical`, `warning`, `notice`.
* `name_regex` - (Optional) A Name Regex of Resource.
* `namespace` - (Optional) The namespace of the cloud monitor rule.
* `output_file` - (Optional) File name where to save data source results.
* `rule_name` - (Optional) The name of the cloud monitor rule. This field support fuzzy query.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `rules` - The collection of query.
    * `alert_methods` - The alert methods of the cloud monitor rule.
    * `alert_state` - The alert state of the cloud monitor rule.
    * `condition_operator` - The condition operator of the cloud monitor rule. Valid values: `&&`, `||`.
    * `conditions` - The conditions of the cloud monitor rule.
        * `comparison_operator` - The comparison operation of the cloud monitor rule.
        * `metric_name` - The metric name of the cloud monitor rule.
        * `metric_unit` - The metric unit of the cloud monitor rule.
        * `period` - The period of the cloud monitor rule.
        * `statistics` - The statistics of the cloud monitor rule.
        * `threshold` - The threshold of the cloud monitor rule.
    * `contact_group_ids` - The contact group ids of the cloud monitor rule.
    * `created_at` - The created time of the cloud monitor rule.
    * `description` - The description of the cloud monitor rule.
    * `effect_end_at` - The effect end time of the cloud monitor rule.
    * `effect_start_at` - The effect start time of the cloud monitor rule.
    * `enable_state` - The enable state of the cloud monitor rule.
    * `evaluation_count` - The evaluation count of the cloud monitor rule.
    * `id` - The id of the cloud monitor rule.
    * `level` - The level of the cloud monitor rule.
    * `multiple_conditions` - Whether to enable the multiple conditions function of the cloud monitor rule.
    * `namespace` - The namespace of the cloud monitor rule.
    * `original_dimensions` - The original dimensions of the cloud monitor rule.
        * `key` - The key of the dimension.
        * `value` - The value of the dimension.
    * `regions` - The region id of the cloud monitor rule.
    * `rule_name` - The name of the cloud monitor rule.
    * `silence_time` - The silence time of the cloud monitor rule. Unit in minutes.
    * `sub_namespace` - The sub namespace of the cloud monitor rule.
    * `updated_at` - The updated time of the cloud monitor rule.
    * `web_hook` - The web hook of the cloud monitor rule.
    * `webhook_ids` - The webhook id list of the cloud monitor rule.
* `total_count` - The total count of query.


