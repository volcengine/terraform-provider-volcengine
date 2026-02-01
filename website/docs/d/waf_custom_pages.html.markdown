---
subcategory: "WAF"
layout: "volcengine"
page_title: "Volcengine: volcengine_waf_custom_pages"
sidebar_current: "docs-volcengine-datasource-waf_custom_pages"
description: |-
  Use this data source to query detailed information of waf custom pages
---
# volcengine_waf_custom_pages
Use this data source to query detailed information of waf custom pages
## Example Usage
```hcl
data "volcengine_waf_custom_pages" "foo" {
  host = "www.tf-test.com"
}
```
## Argument Reference
The following arguments are supported:
* `host` - (Required) The domain names that need to be viewed.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The name of the project to which your domain names belong.
* `rule_tag` - (Optional) Unique identification of the rules.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `data` - Details of the rules.
    * `accurate` - Advanced conditions.
        * `accurate_rules` - Details of advanced conditions.
            * `http_obj` - The HTTP object to be added to the advanced conditions.
            * `obj_type` - The matching field for HTTP objects.
            * `opretar` - The logical operator for the condition.
            * `property` - Operate the properties of the http object.
            * `value_string` - The value to be matched.
        * `logic` - The logical relationship of advanced conditions.
    * `advanced` - Whether to configure advanced conditions.
    * `body` - The layout content of the response page.
    * `client_ip` - Fill in ALL, which means this rule will take effect on all IP addresses.
    * `code` - Custom HTTP code returned when the request is blocked. Required if PageMode=0 or 1.
    * `content_type` - The layout template of the response page. Required if PageMode=0 or 1.
    * `description` - Rule description.
    * `enable` - Whether to enable the rule.
    * `group_id` - The ID of the advanced conditional rule group.
    * `header` - Request header information.
    * `host` - Domain name to be protected.
    * `id` - The ID of rule.
    * `isolation_id` - The ID of Region.
    * `name` - Rule name.
    * `page_mode` - The layout template of the response page.
    * `policy` - Action to be taken on requests that match the rule.
    * `redirect_url` - The path where users should be redirected.
    * `rule_tag` - Unique identification of the rules.
    * `update_time` - Rule update time.
    * `url` - Match the path.
* `total_count` - The total count of query.


