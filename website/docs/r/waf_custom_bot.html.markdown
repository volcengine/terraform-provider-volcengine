---
subcategory: "WAF"
layout: "volcengine"
page_title: "Volcengine: volcengine_waf_custom_bot"
sidebar_current: "docs-volcengine-resource-waf_custom_bot"
description: |-
  Provides a resource to manage waf custom bot
---
# volcengine_waf_custom_bot
Provides a resource to manage waf custom bot
## Example Usage
```hcl
resource "volcengine_waf_custom_bot" "foo" {
  host         = "www.tf-test.com"
  bot_type     = "tf-test"
  description  = "tf-test"
  project_name = "default"
  action       = "observe"
  enable       = 1
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
* `accurate` - (Required) Advanced conditions.
* `action` - (Required) The execution action of the Bot.
* `bot_type` - (Required) bot name.
* `enable` - (Required) Whether to enable bot.
* `host` - (Required, ForceNew) Domain name information.
* `description` - (Optional) The description of bot.
* `project_name` - (Optional) The Name of the affiliated project resource.

The `accurate_rules` object supports the following:

* `http_obj` - (Optional) The HTTP object to be added to the advanced conditions.
* `obj_type` - (Optional) The matching field for HTTP objects.
* `opretar` - (Optional) The logical operator for the condition.
* `property` - (Optional) Operate the properties of the http object.
* `value_string` - (Optional) The value to be matched.

The `accurate` object supports the following:

* `accurate_rules` - (Optional) Details of advanced conditions.
* `logic` - (Optional) The logical relationship of advanced conditions.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `advanced` - Whether to set advanced conditions.
* `rule_tag` - Rule unique identifier.
* `update_time` - The update time.


## Import
WafCustomBot can be imported using the id, e.g.
```
$ terraform import volcengine_waf_custom_bot.default resource_id
```

