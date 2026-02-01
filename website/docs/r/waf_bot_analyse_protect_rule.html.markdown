---
subcategory: "WAF"
layout: "volcengine"
page_title: "Volcengine: volcengine_waf_bot_analyse_protect_rule"
sidebar_current: "docs-volcengine-resource-waf_bot_analyse_protect_rule"
description: |-
  Provides a resource to manage waf bot analyse protect rule
---
# volcengine_waf_bot_analyse_protect_rule
Provides a resource to manage waf bot analyse protect rule
## Example Usage
```hcl
resource "volcengine_waf_bot_analyse_protect_rule" "foo" {
  project_name              = "default"
  statistical_type          = 2
  statistical_duration      = 50
  single_threshold          = 100
  single_proportion         = 0.25
  rule_priority             = 3
  path_threshold            = 1000
  path                      = "/mod"
  name                      = "tf-test-mod"
  host                      = "www.tf-test.com"
  field                     = "HEADER:User-Agent"
  exemption_time            = 60
  enable                    = 1
  effect_time               = 1000
  action_type               = 1
  action_after_verification = 1
  accurate_group {
    accurate_rules {
      http_obj     = "request.uri"
      obj_type     = 1
      opretar      = 2
      property     = 0
      value_string = "tf"
    }
    accurate_rules {
      http_obj     = "request.schema"
      obj_type     = 0
      opretar      = 2
      property     = 0
      value_string = "tf-2"
    }
    logic = 2
  }
}
```
## Argument Reference
The following arguments are supported:
* `action_type` - (Required) perform the action.
* `effect_time` - (Required) Limit the duration.
* `enable` - (Required) Whether to enable the rules.
* `field` - (Required) Statistical objects, with multiple objects separated by commas.
* `host` - (Required, ForceNew) Website domain names that require the setting of protection rules.
* `name` - (Required) The name of rule.
* `path` - (Required) The requested path.
* `rule_priority` - (Required) Priority of rule effectiveness.
* `single_threshold` - (Required) The maximum number of ips of the same statistical object is enabled when StatisticalType=2.
* `statistical_duration` - (Required) The duration of statistics.
* `statistical_type` - (Required) Statistical content and methods.
* `accurate_group` - (Optional) Advanced conditions.
* `action_after_verification` - (Optional) Perform the action after verification/challenge.
* `exemption_time` - (Optional) Exemption time takes effect when the execution action is human-machine challenge /JS/ Proof of work.
* `path_threshold` - (Optional) The path access frequency threshold is enabled when StatisticalType=1.
* `project_name` - (Optional) The Name of the affiliated project resource.
* `single_proportion` - (Optional) The IP proportion of the same statistical object needs to be configured when StatisticalType=3.

The `accurate_group` object supports the following:

* `accurate_rules` - (Optional) Request characteristic information of the rule group.
* `logic` - (Optional) In the rule group, the high-level conditional operation relationships corresponding to each rule.

The `accurate_rules` object supports the following:

* `http_obj` - (Optional) Custom object.
* `obj_type` - (Optional) matching field.
* `opretar` - (Optional) The logical operator for the condition.
* `property` - (Optional) Operate the properties of the http object.
* `value_string` - (Optional) The value to be matched.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `enable_count` - The number of statistical protection rules enabled under the current domain name.
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


## Import
WafBotAnalyseProtectRule can be imported using the id, e.g.
```
$ terraform import volcengine_waf_bot_analyse_protect_rule.default resource_id:bot_space:host
```

