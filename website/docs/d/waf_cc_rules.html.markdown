---
subcategory: "WAF"
layout: "volcengine"
page_title: "Volcengine: volcengine_waf_cc_rules"
sidebar_current: "docs-volcengine-datasource-waf_cc_rules"
description: |-
  Use this data source to query detailed information of waf cc rules
---
# volcengine_waf_cc_rules
Use this data source to query detailed information of waf cc rules
## Example Usage
```hcl
data "volcengine_waf_cc_rules" "foo" {
  cc_type       = [1]
  host          = "www.tf-test.com"
  rule_name     = "tf"
  path_order_by = "ASC"
}
```
## Argument Reference
The following arguments are supported:
* `host` - (Required) Website domain names that require the setting of protection rules.
* `path_order_by` - (Required) The list shows the order.
* `cc_type` - (Optional) The actions performed on subsequent requests after meeting the statistical conditions.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `rule_name` - (Optional) Search by rule name in a fuzzy manner.
* `rule_tag` - (Optional) Search precisely according to the rule ID.
* `url` - (Optional) Fuzzy search by the requested path.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `data` - The collection of query.
    * `enable_count` - The total number of enabled rules within the rule group.
    * `insert_time` - The creation time of the rule group.
    * `rule_group` - Details of the rule group.
        * `group` - Rule group information.
            * `accurate_group_priority` - After the rule creation is completed, the priority of the automatically generated rule group.
            * `accurate_rules` - Request characteristic information of the rule group.
                * `http_obj` - Custom object.
                * `obj_type` - matching field.
                * `opretar` - The logical operator for the condition.
                * `property` - Operate the properties of the http object.
                * `value_string` - The value to be matched.
            * `id` - The ID of Rule group.
            * `logic` - In the rule group, the high-level conditional operation relationships corresponding to each rule.
        * `rules` - Specific rule information within the rule group.
            * `accurate_group_priority` - After the rule creation is completed, the priority of the automatically generated rule group.
            * `accurate_group` - Advanced conditions.
                * `accurate_group_priority` - After the rule creation is completed, the priority of the automatically generated rule group.
                * `accurate_rules` - Request characteristic information of the rule group.
                    * `http_obj` - Custom object.
                    * `obj_type` - matching field.
                    * `opretar` - The logical operator for the condition.
                    * `property` - Operate the properties of the http object.
                    * `value_string` - The value to be matched.
                * `id` - The ID of Rule group.
                * `logic` - In the rule group, the high-level conditional operation relationships corresponding to each rule.
            * `cc_type` - The actions performed on subsequent requests after meeting the statistical conditions.
            * `count_time` - The statistical period of the strategy.
            * `cron_confs` - Details of the periodic loop configuration.
                * `crontab` - The weekly cycle days and cycle time periods.
                * `path_threshold` - The threshold of the number of requests for path access.
                * `single_threshold` - The threshold of the number of visits to each statistical object.
            * `cron_enable` - Whether to set the cycle to take effect.
            * `effect_time` - Limit the duration, that is, the effective duration of the action.
            * `enable` - Whether the rule is enabled.
            * `exemption_time` - Strategy exemption time.
            * `field` - statistical object.
            * `host` - Protected website domain names.
            * `id` - The ID of Rule group.
            * `name` - The Name of Rule group.
            * `path_threshold` - The threshold of the number of requests for path access.
            * `rule_priority` - Rule execution priority.
            * `rule_tag` - Rule label, that is, the complete rule ID.
            * `single_threshold` - The threshold of the number of visits to each statistical object.
            * `update_time` - Rule update time.
            * `url` - Request path.
    * `total_count` - The total number of rules within the rule group.
    * `url` - The requested path.
* `total_count` - The total count of query.


