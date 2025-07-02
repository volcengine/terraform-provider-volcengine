---
subcategory: "WAF"
layout: "volcengine"
page_title: "Volcengine: volcengine_waf_ip_groups"
sidebar_current: "docs-volcengine-datasource-waf_ip_groups"
description: |-
  Use this data source to query detailed information of waf ip groups
---
# volcengine_waf_ip_groups
Use this data source to query detailed information of waf ip groups
## Example Usage
```hcl
data "volcengine_waf_ip_groups" "foo" {
  time_order_by = "DESC"
}
```
## Argument Reference
The following arguments are supported:
* `time_order_by` - (Required) The arrangement order of the address group.
* `ip` - (Optional) The address or address segment of the query.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `rule_tag` - (Optional) Query the association rule ID.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `ip_group_list` - Address group list information.
    * `ip_count` - The number of IP addresses within the address group.
    * `ip_group_id` - The ID of the ip group.
    * `ip_list` - The IP address to be added.
    * `name` - The name of the ip group.
    * `related_rules` - The list of associated rules.
        * `host` - The information of the protected domain names associated with the rules.
        * `rule_name` - The name of the rule.
        * `rule_tag` - The ID of the rule.
        * `rule_type` - The type of the rule.
    * `update_time` - ip group update time.
* `total_count` - The total count of query.


