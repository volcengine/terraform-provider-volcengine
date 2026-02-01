---
subcategory: "CLOUD_FIREWALL"
layout: "volcengine"
page_title: "Volcengine: volcengine_cfw_nat_firewall_control_policies"
sidebar_current: "docs-volcengine-datasource-cfw_nat_firewall_control_policies"
description: |-
  Use this data source to query detailed information of cfw nat firewall control policies
---
# volcengine_cfw_nat_firewall_control_policies
Use this data source to query detailed information of cfw nat firewall control policies
## Example Usage
```hcl
data "volcengine_cfw_nat_firewall_control_policies" "foo" {
  direction       = "in"
  nat_firewall_id = "nfw-ydmkayvjsw2vsavx****"
}
```
## Argument Reference
The following arguments are supported:
* `direction` - (Required) The direction of nat firewall control policy. Valid values: `in`, `out`.
* `nat_firewall_id` - (Required) The nat firewall id of the nat firewall control policy.
* `action` - (Optional) The action list of the nat firewall control policy. Valid values: `accept`, `deny`, `monitor`.
* `description` - (Optional) The description of the nat firewall control policy. This field support fuzzy query.
* `dest_port` - (Optional) The dest port of the nat firewall control policy.
* `destination` - (Optional) The destination of the nat firewall control policy. This field support fuzzy query.
* `output_file` - (Optional) File name where to save data source results.
* `proto` - (Optional) The proto list of the nat firewall control policy. Valid values: `TCP`, `ICMP`, `UDP`, `ANY`. When the destination_type is `domain`, The proto must be `TCP`.
* `repeat_type` - (Optional) The repeat type of the nat firewall control policy. Valid values: `Permanent`, `Once`, `Daily`, `Weekly`, `Monthly`.
* `rule_id` - (Optional) The rule id of the nat firewall control policy. This field support fuzzy query.
* `source` - (Optional) The source of the nat firewall control policy. This field support fuzzy query.
* `status` - (Optional) The enable status list of the nat firewall control policy.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `nat_firewall_control_policies` - The collection of query.
    * `account_id` - The account id of the nat firewall control policy.
    * `action` - The action of the nat firewall control policy.
    * `description` - The description of the nat firewall control policy.
    * `dest_port_group_list` - The dest port group list of the nat firewall control policy.
    * `dest_port_group_type` - The dest port group type of the nat firewall control policy.
    * `dest_port_list` - The dest port list of the nat firewall control policy.
    * `dest_port_type` - The dest port type of the nat firewall control policy.
    * `dest_port` - The dest port of the nat firewall control policy.
    * `destination_cidr_list` - The destination cidr list of the nat firewall control policy.
    * `destination_group_list` - The destination group list of the nat firewall control policy.
    * `destination_group_type` - The destination group type of the nat firewall control policy.
    * `destination_type` - The destination type of the nat firewall control policy.
    * `destination` - The destination of the nat firewall control policy.
    * `direction` - The direction of the nat firewall control policy.
    * `effect_status` - The effect status of the nat firewall control policy. 1: Not yet effective, 2: Issued in progress, 3: Effective.
    * `end_time` - The end time of the nat firewall control policy. Unix timestamp.
    * `hit_cnt` - The hit count of the nat firewall control policy.
    * `id` - The id of the nat firewall control policy.
    * `is_effected` - Whether the nat firewall control policy is effected.
    * `nat_firewall_id` - The id of the nat firewall.
    * `nat_firewall_name` - The name of the nat firewall.
    * `prio` - The priority of the nat firewall control policy.
    * `proto` - The proto of the nat firewall control policy.
    * `repeat_days` - The repeat days of the nat firewall control policy.
    * `repeat_end_time` - The repeat end time of the nat firewall control policy.
    * `repeat_start_time` - The repeat start time of the nat firewall control policy.
    * `repeat_type` - The repeat type of the nat firewall control policy.
    * `rule_id` - The id of the nat firewall control policy.
    * `source_cidr_list` - The source cidr list of the nat firewall control policy.
    * `source_group_list` - The source group list of the nat firewall control policy.
    * `source_group_type` - The source group type of the nat firewall control policy.
    * `source_type` - The source type of the nat firewall control policy.
    * `source` - The source of the nat firewall control policy.
    * `start_time` - The start time of the nat firewall control policy. Unix timestamp.
    * `status` - Whether to enable the nat firewall control policy.
    * `update_time` - The update time of the nat firewall control policy.
    * `use_count` - The use count of the nat firewall control policy.
* `total_count` - The total count of query.


