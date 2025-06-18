---
subcategory: "WAF"
layout: "volcengine"
page_title: "Volcengine: volcengine_waf_custom_bots"
sidebar_current: "docs-volcengine-datasource-waf_custom_bots"
description: |-
  Use this data source to query detailed information of waf custom bots
---
# volcengine_waf_custom_bots
Use this data source to query detailed information of waf custom bots
## Example Usage
```hcl
data "volcengine_waf_custom_bots" "foo" {
  host = "www.tf-test.com"
}
```
## Argument Reference
The following arguments are supported:
* `host` - (Optional) The domain names that need to be viewed.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `data` - The Details of Custom bot.
    * `accurate` - Advanced conditions.
        * `accurate_rules` - Details of advanced conditions.
            * `http_obj` - The HTTP object to be added to the advanced conditions.
            * `obj_type` - The matching field for HTTP objects.
            * `opretar` - The logical operator for the condition.
            * `property` - Operate the properties of the http object.
            * `value_string` - The value to be matched.
        * `logic` - The logical relationship of advanced conditions.
    * `action` - The execution action of the Bot.
    * `advanced` - Whether to set advanced conditions.
    * `bot_type` - bot name.
    * `description` - The description of bot.
    * `enable` - Whether to enable bot.
    * `id` - The actual count bits of the rule unique identifier (corresponding to the RuleTag).
    * `rule_tag` - Rule unique identifier.
    * `update_time` - The update time.
* `total_count` - The total count of query.


