---
subcategory: "ALB"
layout: "volcengine"
page_title: "Volcengine: volcengine_alb_rules"
sidebar_current: "docs-volcengine-datasource-alb_rules"
description: |-
  Use this data source to query detailed information of alb rules
---
# volcengine_alb_rules
Use this data source to query detailed information of alb rules
## Example Usage
```hcl
data "volcengine_alb_rules" "foo" {
  listener_id = "lsn-1iidd19u4oni874adhezjkyj3"
}
```
## Argument Reference
The following arguments are supported:
* `listener_id` - (Required) The Id of listener.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `rules` - The collection of Rule query.
    * `description` - The Description of Rule.
    * `domain` - The Domain of Rule.
    * `id` - The Id of Rule.
    * `redirect_config` - Redirect related configuration.
        * `redirect_domain` - The redirect domain.
        * `redirect_http_code` - The redirect HTTP code,support 301(default), 302, 307, 308.
        * `redirect_port` - The redirect port.
        * `redirect_protocol` - The redirect protocol,support HTTP,HTTPS(default).
        * `redirect_uri` - The redirect URI.
    * `rewrite_config` - The list of rewrite configurations.
        * `rewrite_path` - Rewrite path.
    * `rewrite_enabled` - Rewrite configuration switch for forwarding rules, only allows configuration and takes effect when RuleAction is empty (i.e., forwarding to server group). Only available for whitelist users, please submit an application to experience. Supported values are as follows:
on: enable.
off: disable.
    * `rule_action` - The forwarding rule action, if this parameter is empty, forward to server group, if value is `Redirect`, will redirect.
    * `rule_id` - The Id of Rule.
    * `server_group_id` - The Id of Server Group.
    * `traffic_limit_enabled` - Forwarding rule QPS rate limiting switch:
 on: enable.
off: disable (default).
    * `traffic_limit_qps` - When Rules.N.TrafficLimitEnabled is turned on, this field is required. Requests per second. Valid values are between 100 and 100000.
    * `url` - The Url of Rule.
* `total_count` - The total count of Rule query.


