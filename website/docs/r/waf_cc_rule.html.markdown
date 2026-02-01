---
subcategory: "WAF"
layout: "volcengine"
page_title: "Volcengine: volcengine_waf_cc_rule"
sidebar_current: "docs-volcengine-resource-waf_cc_rule"
description: |-
  Provides a resource to manage waf cc rule
---
# volcengine_waf_cc_rule
Provides a resource to manage waf cc rule
## Example Usage
```hcl
resource "volcengine_waf_cc_rule" "foo" {
  name             = "tf-test"
  url              = "/"
  field            = "HEADER:User-Agemnt"
  single_threshold = "100"
  path_threshold   = 101
  count_time       = 102
  cc_type          = 1
  effect_time      = 200
  rule_priority    = 2
  enable           = 1
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
  host           = "www.tf-test.com"
  exemption_time = 0
  cron_enable    = 1
  cron_confs {
    crontab          = "* 0 * * 1,2,3,4,5,6,0"
    path_threshold   = 123
    single_threshold = 234
  }
  cron_confs {
    crontab          = "* 3-8 * * 1,2,3,4,5,6,0"
    path_threshold   = 345
    single_threshold = 456
  }
}
```
## Argument Reference
The following arguments are supported:
* `cc_type` - (Required) The actions performed on subsequent requests after meeting the statistical conditions.
* `count_time` - (Required) The statistical period of the strategy.
* `effect_time` - (Required) Limit the duration, that is, the effective duration of the action.
* `enable` - (Required) Whether to enable the rules.
* `field` - (Required) statistical object.
* `host` - (Required, ForceNew) Website domain names that require the setting of protection rules.
* `name` - (Required) The name of cc rule.
* `path_threshold` - (Required) The threshold of the total number of times the request path is accessed.
* `rule_priority` - (Required) Rule execution priority.
* `single_threshold` - (Required) The threshold of the number of times each statistical object accesses the request path.
* `url` - (Required) The website request path that needs protection.
* `accurate_group` - (Optional) Advanced conditions.
* `advanced_enable` - (Optional) Whether to enable advanced conditions.
* `cron_confs` - (Optional, ForceNew) Details of the periodic loop configuration.
* `cron_enable` - (Optional, ForceNew) Whether to set the cycle to take effect.
* `exemption_time` - (Optional, ForceNew) Strategy exemption time.
* `project_name` - (Optional) The Name of the affiliated project resource.

The `accurate_group` object supports the following:

* `accurate_rules` - (Required) Details of advanced conditions.
* `logic` - (Required) The logical relationship of advanced conditions.

The `accurate_rules` object supports the following:

* `http_obj` - (Required) The HTTP object to be added to the advanced conditions.
* `obj_type` - (Required) The matching field for HTTP objects.
* `opretar` - (Required) The logical operator for the condition.
* `property` - (Required) Operate the properties of the http object.
* `value_string` - (Required) The value to be matched.

The `cron_confs` object supports the following:

* `crontab` - (Required) The weekly cycle days and cycle time periods.
* `path_threshold` - (Required) The threshold of the number of requests for path access.
* `single_threshold` - (Required) The threshold of the number of visits to each statistical object.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
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


## Import
WafCcRule can be imported using the id, e.g.
```
$ terraform import volcengine_waf_cc_rule.default resource_id:Host
```

