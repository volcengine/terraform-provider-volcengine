---
subcategory: "WAF"
layout: "volcengine"
page_title: "Volcengine: volcengine_waf_acl_rules"
sidebar_current: "docs-volcengine-datasource-waf_acl_rules"
description: |-
  Use this data source to query detailed information of waf acl rules
---
# volcengine_waf_acl_rules
Use this data source to query detailed information of waf acl rules
## Example Usage
```hcl
data "volcengine_waf_acl_rules" "foo" {
  acl_type      = "Block"
  action        = ["observe"]
  defence_host  = ["www.tf-test.com"]
  enable        = [1]
  rule_name     = "tf-test"
  time_order_by = "ASC"
  project_name  = "default"
}
```
## Argument Reference
The following arguments are supported:
* `acl_type` - (Required) The types of access control rules.
* `action` - (Optional) Action to be taken on requests that match the rule.
* `defence_host` - (Optional) The list of queried domain names.
* `enable` - (Optional) The enabled status of the rule.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The name of the project to which your domain names belong.
* `rule_name` - (Optional) Rule name, fuzzy search.
* `rule_tag` - (Optional) Rule unique identifier, precise search.
* `time_order_by` - (Optional) The list shows the timing sequence.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `rules` - Details of the rules.
    * `accurate_group` - Advanced conditions.
        * `accurate_rules` - Details of advanced conditions.
            * `http_obj` - The HTTP object to be added to the advanced conditions.
            * `obj_type` - The matching field for HTTP objects.
            * `opretar` - The logical operator for the condition.
            * `property` - Operate the properties of the http object.
            * `value_string` - The value to be matched.
        * `logic` - The logical relationship of advanced conditions.
    * `action` - Action to be taken on requests that match the rule.
    * `advanced` - Whether to set advanced conditions.
    * `client_ip` - IP address.
    * `description` - Rule description.
    * `enable` - Whether to enable the rule.
    * `host_add_type` - Type of domain name addition.
    * `host_group_id` - The ID of the domain group.
    * `host_groups` - The list of domain name groups.
        * `host_group_id` - The ID of host group.
        * `name` - The Name of host group.
    * `host_list` - Single or multiple domain names are supported.
    * `id` - Rule ID.
    * `ip_add_type` - Type of IP address addition.
    * `ip_group_id` - Add the list of address group ids in the address group mode.
    * `ip_groups` - The list of domain name groups.
        * `ip_group_id` - The ID of the IP address group.
        * `name` - The Name of the IP address group.
    * `ip_list` - Single or multiple IP addresses are supported.
    * `ip_location_country` - Country or region code.
    * `ip_location_subregion` - Domestic region code.
    * `name` - Rule name.
    * `rule_tag` - Rule unique identifier.
    * `update_time` - Update time of the rule.
    * `url` - The path of Matching.
* `total_count` - The total count of query.


