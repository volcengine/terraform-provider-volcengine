---
subcategory: "CLOUD_FIREWALL"
layout: "volcengine"
page_title: "Volcengine: volcengine_cfw_control_policies"
sidebar_current: "docs-volcengine-datasource-cfw_control_policies"
description: |-
  Use this data source to query detailed information of cfw control policies
---
# volcengine_cfw_control_policies
Use this data source to query detailed information of cfw control policies
## Example Usage
```hcl
data "volcengine_cfw_control_policies" "foo" {
  direction = "in"
  action    = ["deny"]
}
```
## Argument Reference
The following arguments are supported:
* `direction` - (Required) The direction of control policy. Valid values: `in`, `out`.
* `action` - (Optional) The action list of the control policy. Valid values: `accept`, `deny`, `monitor`.
* `description` - (Optional) The description of the control policy. This field support fuzzy query.
* `destination` - (Optional) The destination of the control policy. This field support fuzzy query.
* `output_file` - (Optional) File name where to save data source results.
* `proto` - (Optional) The proto list of the control policy. Valid values: `TCP`, `ICMP`, `UDP`, `ANY`. When the destination_type is `domain`, The proto must be `TCP`.
* `repeat_type` - (Optional) The repeat type of the control policy. Valid values: `Permanent`, `Once`, `Daily`, `Weekly`, `Monthly`.
* `rule_id` - (Optional) The rule id of the control policy. This field support fuzzy query.
* `source` - (Optional) The source of the control policy. This field support fuzzy query.
* `status` - (Optional) The enable status list of the control policy.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `control_policies` - The collection of query.
    * `account_id` - The account id of the control policy.
    * `action` - The action of the control policy.
    * `description` - The description of the control policy.
    * `dest_port_group_type` - The dest port group type of the control policy.
    * `dest_port_list` - The dest port list of the control policy.
    * `dest_port_type` - The dest port type of the control policy.
    * `dest_port` - The dest port of the control policy.
    * `destination_cidr_list` - The destination cidr list of the control policy.
    * `destination_group_type` - The destination group type of the control policy.
    * `destination_type` - The destination type of the control policy.
    * `destination` - The destination of the control policy.
    * `direction` - The direction of the control policy.
    * `effect_status` - The effect status of the control policy. 1: Not yet effective, 2: Issued in progress, 3: Effective.
    * `end_time` - The end time of the control policy. Unix timestamp.
    * `hit_cnt` - The hit count of the control policy.
    * `id` - The id of the control policy.
    * `is_effected` - Whether the control policy is effected.
    * `prio` - The priority of the control policy.
    * `proto` - The proto of the control policy.
    * `repeat_days` - The repeat days of the control policy.
    * `repeat_end_time` - The repeat end time of the control policy.
    * `repeat_start_time` - The repeat start time of the control policy.
    * `repeat_type` - The repeat type of the control policy.
    * `rule_id` - The id of the control policy.
    * `source_cidr_list` - The source cidr list of the control policy.
    * `source_group_type` - The source group type of the control policy.
    * `source_type` - The source type of the control policy.
    * `source` - The source of the control policy.
    * `start_time` - The start time of the control policy. Unix timestamp.
    * `status` - Whether to enable the control policy.
    * `update_time` - The update time of the control policy.
    * `use_count` - The use count of the control policy.
* `total_count` - The total count of query.


