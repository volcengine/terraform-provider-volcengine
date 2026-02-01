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
# Basic edition
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

# Standard edition
resource "volcengine_alb_rule" "example" {
  listener_id = "lsn-bddjp5fcof0g8dv40naga1yd"
  rule_action = ""
  description = "standard edition alb rule"
  url         = ""
  priority    = 1
  # Matching condition: Host + Path
  rule_conditions {
    type = "Host"
    host_config {
      values = ["www.example.com"]
    }
  }
  rule_conditions {
    type = "Path"
    path_config {
      values = ["/app/*"]
    }
  }
  rule_actions {
    type = "ForwardGroup"
    forward_group_config {
      server_group_tuples {
        server_group_id = "rsp-bdd1lpcbvv288dv40ov1sye0"
        weight          = 50
      }
      server_group_sticky_session {
        enabled = "off"
      }
    }
  }
}
```
## Argument Reference
The following arguments are supported:
* `listener_id` - (Required, ForceNew) The ID of listener.
* `rule_action` - (Required) The forwarding rule action, if this parameter is empty(`""`), forward to server group, if value is `Redirect`, will redirect.
* `description` - (Optional) The description of the Rule.
* `domain` - (Optional, ForceNew) The domain of Rule.
* `priority` - (Optional) The priority of the Rule.Only the standard version is supported.
* `redirect_config` - (Optional) The redirect related configuration.
* `rewrite_config` - (Optional) The list of rewrite configurations.
* `rewrite_enabled` - (Optional) Rewrite configuration switch for forwarding rules, only allows configuration and takes effect when RuleAction is empty (i.e., forwarding to server group). Only available for whitelist users, please submit an application to experience. Supported values are as follows:
on: enable.
off: disable.
* `rule_actions` - (Optional) The rule actions for standard edition forwarding rules.
* `rule_conditions` - (Optional) The rule conditions for standard edition forwarding rules.
* `server_group_id` - (Optional) Server group ID, this parameter is required if `rule_action` is empty.
* `server_group_tuples` - (Optional) Weight forwarded to the corresponding backend server group.
* `sticky_session_enabled` - (Optional) Whether to enable group session stickiness. Valid values are 'on' and 'off'.
* `sticky_session_timeout` - (Optional) The group session stickiness timeout, in seconds.
* `traffic_limit_enabled` - (Optional) Forwarding rule QPS rate limiting switch:
 on: enable.
 off: disable (default).
* `traffic_limit_qps` - (Optional) When Rules.N.TrafficLimitEnabled is turned on, this field is required. Requests per second. Valid values are between 100 and 100000.
* `url` - (Optional, ForceNew) The Url of Rule.

The `fixed_response_config` object supports the following:


The `forward_group_config` object supports the following:

* `server_group_sticky_session` - (Optional) The config of group session stickiness.
* `server_group_tuples` - (Optional) The server group tuples.

The `header_config` object supports the following:

* `key` - (Required) The header key.
* `values` - (Required) The list of header values.

The `host_config` object supports the following:

* `values` - (Required) The list of domain names.

The `method_config` object supports the following:

* `values` - (Required) The values of the method. Vaild values: HEAD,GET,POST,OPTIONS,PUT,PATCH,DELETE.

The `path_config` object supports the following:

* `values` - (Optional) The list of absolute paths.

The `query_string_config` object supports the following:


The `redirect_config` object supports the following:

* `http_code` - (Optional) The redirect HTTP code.
* `path` - (Optional) The path to which the request was redirected.
* `protocol` - (Optional) The redirect protocol.

The `redirect_config` object supports the following:

* `redirect_domain` - (Optional) The redirect domain, only support exact domain name.
* `redirect_http_code` - (Optional) The redirect http code, support 301(default), 302, 307, 308.
* `redirect_port` - (Optional) The redirect port.
* `redirect_protocol` - (Optional) The redirect protocol, support HTTP, HTTPS(default).
* `redirect_uri` - (Optional) The redirect URI.

The `rewrite_config` object supports the following:


The `rewrite_config` object supports the following:

* `rewrite_path` - (Required) Rewrite path.

The `rule_actions` object supports the following:

* `fixed_response_config` - (Optional) Fixed response configuration for fixed response type rule.
* `forward_group_config` - (Optional) Forward group configuration for ForwardGroup type action.
* `redirect_config` - (Optional) Redirect configuration for Redirect type action.
* `rewrite_config` - (Optional) Rewrite configuration for Rewrite type action.
* `traffic_limit_config` - (Optional) Traffic limit configuration for TrafficLimit type action.
* `type` - (Optional) The type of rule action. Valid values: ForwardGroup, Redirect, Rewrite, TrafficLimit.

The `rule_conditions` object supports the following:

* `header_config` - (Optional) Header configuration for Header type condition.
* `host_config` - (Optional) Host configuration for Host type condition.
* `method_config` - (Optional) Method configuration for Method type condition.
* `path_config` - (Optional) Path configuration for Path type condition.
* `query_string_config` - (Optional) Query string configuration for QueryString type condition.
* `type` - (Optional) The type of rule condition. Valid values: Host, Path, Header, Method, QueryString.

The `server_group_sticky_session` object supports the following:

* `enabled` - (Optional) Whether to enable sticky session stickiness. Valid values are 'on' and 'off'.
* `timeout` - (Optional) The sticky session timeout, in seconds.

The `server_group_tuples` object supports the following:

* `server_group_id` - (Optional) The server group ID.
* `weight` - (Optional) The weight of the server group.

The `server_group_tuples` object supports the following:

* `server_group_id` - (Required) The server group ID. The priority of this parameter is higher than that of `server_group_id`.
* `weight` - (Optional) The weight of the server group.

The `traffic_limit_config` object supports the following:

* `qps` - (Optional) The QPS limit.

The `values` object supports the following:


## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `rule_id` - The ID of rule.


## Import
AlbRule can be imported using the listener id and rule id, e.g.
```
$ terraform import volcengine_alb_rule.default lsn-273yv0mhs5xj47fap8sehiiso:rule-****
```

