---
subcategory: "CLOUD_FIREWALL"
layout: "volcengine"
page_title: "Volcengine: volcengine_cfw_vpc_firewall_acl_rules"
sidebar_current: "docs-volcengine-datasource-cfw_vpc_firewall_acl_rules"
description: |-
  Use this data source to query detailed information of cfw vpc firewall acl rules
---
# volcengine_cfw_vpc_firewall_acl_rules
Use this data source to query detailed information of cfw vpc firewall acl rules
## Example Usage
```hcl
data "volcengine_cfw_vpc_firewall_acl_rules" "foo" {
  vpc_firewall_id = "vfw-ydmjakzksgf7u99j6sby"
  action          = ["accept", "deny"]
}
```
## Argument Reference
The following arguments are supported:
* `vpc_firewall_id` - (Required) The vpc firewall id of the vpc firewall acl rule.
* `action` - (Optional) The action list of the vpc firewall acl rule. Valid values: `accept`, `deny`, `monitor`.
* `description` - (Optional) The description of the vpc firewall acl rule. This field support fuzzy query.
* `destination` - (Optional) The destination of the vpc firewall acl rule. This field support fuzzy query.
* `output_file` - (Optional) File name where to save data source results.
* `proto` - (Optional) The proto list of the vpc firewall acl rule. Valid values: `TCP`, `ICMP`, `UDP`, `ANY`. When the destination_type is `domain`, The proto must be `TCP`.
* `repeat_type` - (Optional) The repeat type of the vpc firewall acl rule. Valid values: `Permanent`, `Once`, `Daily`, `Weekly`, `Monthly`.
* `rule_id` - (Optional) The rule id of the vpc firewall acl rule. This field support fuzzy query.
* `source` - (Optional) The source of the vpc firewall acl rule. This field support fuzzy query.
* `status` - (Optional) The enable status list of the vpc firewall acl rule.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of query.
* `vpc_firewall_acl_rules` - The collection of query.
    * `account_id` - The account id of the vpc firewall acl rule.
    * `action` - The action of the vpc firewall acl rule.
    * `description` - The description of the vpc firewall acl rule.
    * `dest_port_group_type` - The dest port group type of the vpc firewall acl rule.
    * `dest_port_list` - The dest port list of the vpc firewall acl rule.
    * `dest_port_type` - The dest port type of the vpc firewall acl rule.
    * `dest_port` - The dest port of the vpc firewall acl rule.
    * `destination_cidr_list` - The destination cidr list of the vpc firewall acl rule.
    * `destination_group_type` - The destination group type of the vpc firewall acl rule.
    * `destination_type` - The destination type of the vpc firewall acl rule.
    * `destination` - The destination of the vpc firewall acl rule.
    * `effect_status` - The effect status of the vpc firewall acl rule. 1: Not yet effective, 2: Issued in progress, 3: Effective.
    * `end_time` - The end time of the vpc firewall acl rule. Unix timestamp.
    * `hit_cnt` - The hit count of the vpc firewall acl rule.
    * `id` - The id of the vpc firewall acl rule.
    * `is_effected` - Whether the vpc firewall acl rule is effected.
    * `prio` - The priority of the vpc firewall acl rule.
    * `proto` - The proto of the vpc firewall acl rule.
    * `repeat_days` - The repeat days of the vpc firewall acl rule.
    * `repeat_end_time` - The repeat end time of the vpc firewall acl rule.
    * `repeat_start_time` - The repeat start time of the vpc firewall acl rule.
    * `repeat_type` - The repeat type of the vpc firewall acl rule.
    * `rule_id` - The id of the vpc firewall acl rule.
    * `source_cidr_list` - The source cidr list of the vpc firewall acl rule.
    * `source_group_type` - The source group type of the vpc firewall acl rule.
    * `source_type` - The source type of the vpc firewall acl rule.
    * `source` - The source of the vpc firewall acl rule.
    * `start_time` - The start time of the vpc firewall acl rule. Unix timestamp.
    * `status` - Whether to enable the vpc firewall acl rule.
    * `update_time` - The update time of the vpc firewall acl rule.
    * `use_count` - The use count of the vpc firewall acl rule.
    * `vpc_firewall_id` - The id of the vpc firewall.
    * `vpc_firewall_name` - The name of the vpc firewall.


