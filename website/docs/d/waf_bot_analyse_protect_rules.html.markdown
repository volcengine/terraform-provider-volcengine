---
subcategory: "WAF"
layout: "volcengine"
page_title: "Volcengine: volcengine_waf_bot_analyse_protect_rules"
sidebar_current: "docs-volcengine-datasource-waf_bot_analyse_protect_rules"
description: |-
  Use this data source to query detailed information of waf bot analyse protect rules
---
# volcengine_waf_bot_analyse_protect_rules
Use this data source to query detailed information of waf bot analyse protect rules
## Example Usage
```hcl
data "volcengine_waf_bot_analyse_protect_rules" "foo" {
  host      = "www.tf-test.com"
  bot_space = "BotRepeat"
}
```
## Argument Reference
The following arguments are supported:
* `bot_space` - (Required) Bot protection rule type.
* `host` - (Required) Website domain names that require the setting of protection rules.
* `name_regex` - (Optional) A Name Regex of Resource.
* `name` - (Optional) The name of the rule.
* `output_file` - (Optional) File name where to save data source results.
* `path` - (Optional) Protective path.
* `project_name` - (Optional) The name of the project to which your domain names belong.
* `rule_tag` - (Optional) Unique identification of rules.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `data` - The details of the Bot rules.
    * `enable_count` - The number of statistical protection rules enabled under the current domain name.
    * `path` - The requested path.
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
            * `accurate_group_priority` - Advanced condition priority.
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
            * `action_after_verification` - Perform actions after human-machine verification /JS challenges.
            * `action_type` - perform the action.
            * `effect_time` - Limit the duration.
            * `enable` - Whether to enable the rules.
            * `exemption_time` - Exemption time.
            * `field` - statistical object.
            * `host` - The domain name where the protection rule is located.
            * `id` - Rule unique identifier.
            * `name` - The name of rule.
            * `pass_ratio` - JS challenge/human-machine verification pass rate.
            * `path_threshold` - Threshold of path access times.
            * `path` - Request path.
            * `rule_priority` - Rule execution priority.
            * `rule_tag` - Rule label, that is, the complete rule ID.
            * `single_proportion` - The IP proportion of the same statistical object.
            * `single_threshold` - The maximum number of ips for the same statistical object.
            * `statistical_duration` - The duration of the statistics.
            * `statistical_type` - Statistical content method.
            * `update_time` - Rule update time.
    * `total_count` - The total number of statistical protection rules under the current domain name.
* `total_count` - The total count of query.


