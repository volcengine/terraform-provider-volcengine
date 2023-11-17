---
subcategory: "ALB"
layout: "volcengine"
page_title: "Volcengine: volcengine_alb_rule"
sidebar_current: "docs-volcengine-resource-alb_rule"
description: |-
  Provides a resource to manage alb rule
---
# volcengine_alb_rule
Provides a resource to manage alb rule
## Example Usage
```hcl
resource "volcengine_alb_rule" "foo" {
  listener_id           = "lsn-1iidd19u4oni874adhezjkyj3"
  domain                = "www.test.com"
  url                   = "/test"
  rule_action           = "Redirect"
  server_group_id       = "rsp-1g72w74y4umf42zbhq4k4hnln"
  description           = "test"
  traffic_limit_enabled = "off"
  traffic_limit_qps     = 100
  rewrite_enabled       = "off"
  redirect_config {
    redirect_domain    = "www.testtest.com"
    redirect_uri       = "/testtest"
    redirect_port      = "555"
    redirect_http_code = "302"
    //redirect_http_protocol = ""
  }
  rewrite_config {
    rewrite_path = "/test"
  }
}
```
## Argument Reference
The following arguments are supported:
* `listener_id` - (Required, ForceNew) The ID of listener.
* `rule_action` - (Required) The forwarding rule action, if this parameter is empty(`""`), forward to server group, if value is `Redirect`, will redirect.
* `description` - (Optional) The description of the Rule.
* `domain` - (Optional, ForceNew) The domain of Rule.
* `redirect_config` - (Optional) The redirect related configuration.
* `rewrite_config` - (Optional) The list of rewrite configurations.
* `rewrite_enabled` - (Optional) Rewrite configuration switch for forwarding rules, only allows configuration and takes effect when RuleAction is empty (i.e., forwarding to server group). Only available for whitelist users, please submit an application to experience. Supported values are as follows:
on: enable.
off: disable.
* `server_group_id` - (Optional) Server group ID, this parameter is required if `rule_action` is empty.
* `traffic_limit_enabled` - (Optional) Forwarding rule QPS rate limiting switch:
 on: enable.
 off: disable (default).
* `traffic_limit_qps` - (Optional) When Rules.N.TrafficLimitEnabled is turned on, this field is required. Requests per second. Valid values are between 100 and 100000.
* `url` - (Optional, ForceNew) The Url of Rule.

The `redirect_config` object supports the following:

* `redirect_domain` - (Optional) The redirect domain, only support exact domain name.
* `redirect_http_code` - (Optional) The redirect http code, support 301(default), 302, 307, 308.
* `redirect_port` - (Optional) The redirect port.
* `redirect_protocol` - (Optional) The redirect protocol, support HTTP, HTTPS(default).
* `redirect_uri` - (Optional) The redirect URI.

The `rewrite_config` object supports the following:

* `rewrite_path` - (Required) Rewrite path.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `rule_id` - The ID of rule.


## Import
AlbRule can be imported using the listener id and rule id, e.g.
```
$ terraform import volcengine_alb_rule.default lsn-273yv0mhs5xj47fap8sehiiso:rule-****
```

