---
subcategory: "WAF"
layout: "volcengine"
page_title: "Volcengine: volcengine_waf_system_bot"
sidebar_current: "docs-volcengine-resource-waf_system_bot"
description: |-
  Provides a resource to manage waf system bot
---
**❗Notice:**
The current provider is no longer being maintained. We recommend that you use the [volcenginecc](https://registry.terraform.io/providers/volcengine/volcenginecc/latest/docs) instead.
# volcengine_waf_system_bot
Provides a resource to manage waf system bot
## Example Usage
```hcl
resource "volcengine_waf_system_bot" "foo" {
  bot_type     = "feed_fetcher"
  project_name = "default"
  action       = "observe"
  enable       = 0
  host         = "www.tf-test.com"
}
```
## Argument Reference
The following arguments are supported:
* `bot_type` - (Required) The name of bot.
* `host` - (Required) Domain name information.
* `action` - (Optional) The execution action of the Bot.
* `enable` - (Optional) Whether to enable bot.
* `project_name` - (Optional) The Name of the affiliated project resource.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `description` - The description of the Bot.
* `rule_tag` - The ID of the Bot rule.


## Import
WafSystemBot can be imported using the id, e.g.
```
$ terraform import volcengine_waf_system_bot.default BotType:Host
```

