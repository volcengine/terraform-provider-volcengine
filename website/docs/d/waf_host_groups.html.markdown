---
subcategory: "WAF"
layout: "volcengine"
page_title: "Volcengine: volcengine_waf_host_groups"
sidebar_current: "docs-volcengine-datasource-waf_host_groups"
description: |-
  Use this data source to query detailed information of waf host groups
---
# volcengine_waf_host_groups
Use this data source to query detailed information of waf host groups
## Example Usage
```hcl
data "volcengine_waf_host_groups" "foo" {
  host_fix      = "www.tf-test.com"
  time_order_by = "DESC"
}
```
## Argument Reference
The following arguments are supported:
* `time_order_by` - (Required) The list of rule ids associated with the domain name group shows the timing sequence.
* `host_fix` - (Optional) The domain name information queried.
* `host_group_id` - (Optional) The ID of the domain name group.
* `ids` - (Optional) A list of IDs.
* `list_all` - (Optional) Whether to return all domain name groups and their name information, it returns by default.
* `name_fix` - (Optional) The name of the domain name group being queried.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `rule_tag` - (Optional) The rule ID associated with domain name groups.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `host_group_list` - Details of the domain name group list.
    * `description` - Domain name group description.
    * `host_count` - The number of domain names contained in the domain name group.
    * `host_group_id` - The ID of the domain name group.
    * `name` - The name of the domain name group.
    * `related_rules` - The list of associated rules.
        * `rule_name` - The name of the rule.
        * `rule_tag` - The ID of the rule.
        * `rule_type` - The type of the rule.
    * `update_time` - Domain name group update time.
* `total_count` - The total count of query.


