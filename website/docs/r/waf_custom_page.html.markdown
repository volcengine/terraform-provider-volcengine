---
subcategory: "WAF"
layout: "volcengine"
page_title: "Volcengine: volcengine_waf_custom_page"
sidebar_current: "docs-volcengine-resource-waf_custom_page"
description: |-
  Provides a resource to manage waf custom page
---
# volcengine_waf_custom_page
Provides a resource to manage waf custom page
## Example Usage
```hcl
resource "volcengine_waf_custom_page" "foo" {
  host         = "www.123.com"
  policy       = 1
  client_ip    = "ALL"
  name         = "tf-test"
  description  = "tf-test"
  url          = "/tf-test"
  enable       = 1
  code         = 403
  page_mode    = 1
  content_type = "text/html"
  body         = "tf-test-body"
  advanced     = 1
  redirect_url = "/test/tf/path"
  project_name = "default"
  accurate {
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
* `client_ip` - (Required) Fill in ALL, which means this rule will take effect on all IP addresses.
* `code` - (Required) Custom HTTP code returned when the request is blocked. Required if PageMode=0 or 1.
* `enable` - (Required) Whether to enable the rule.
* `host` - (Required, ForceNew) Domain name to be protected.
* `name` - (Required) Rule name.
* `page_mode` - (Required) The layout template of the response page.
* `policy` - (Required) Action to be taken on requests that match the rule.
* `url` - (Required) Match the path.
* `accurate` - (Optional) Advanced conditions.
* `advanced` - (Optional) Whether to configure advanced conditions.
* `body` - (Optional) The layout content of the response page.
* `content_type` - (Optional) The layout template of the response page. Required if PageMode=0 or 1.
* `description` - (Optional) Rule description.
* `project_name` - (Optional) The name of the project to which your domain names belong.
* `redirect_url` - (Optional) The path where users should be redirected.

The `accurate_rules` object supports the following:

* `http_obj` - (Required) The HTTP object to be added to the advanced conditions.
* `obj_type` - (Required) The matching field for HTTP objects.
* `opretar` - (Required) The logical operator for the condition.
* `property` - (Required) Operate the properties of the http object.
* `value_string` - (Required) The value to be matched.

The `accurate` object supports the following:

* `accurate_rules` - (Required) Details of advanced conditions.
* `logic` - (Required) The logical relationship of advanced conditions.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `group_id` - The ID of the advanced conditional rule group.
* `header` - Request header information.
* `isolation_id` - The ID of Region.
* `rule_tag` - Unique identification of the rules.
* `update_time` - Rule update time.


## Import
WafCustomPage can be imported using the id, e.g.
```
$ terraform import volcengine_waf_custom_page.default resource_id:Host
```

