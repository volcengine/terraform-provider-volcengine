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
    * `forward_group_config` - The list of forward group configurations.
        * `server_group_tuples` - The list of destination server groups to forward to.
            * `server_group_id` - The destination server group ID to forward to.
            * `weight` - Server group weight.
        * `sticky_session_enabled` - Whether to enable inter-group session hold.
        * `sticky_session_timeout` - The group session stickiness timeout, in seconds.
    * `id` - The Id of Rule.
    * `priority` - The priority of the Rule. Only the standard version is supported.
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
    * `rule_actions` - The rule actions for standard edition forwarding rules.
        * `fixed_response_config` - Fixed response configuration for fixed response type rule.
            * `content_type` - The content type of the fixed response.
            * `response_body` - The response body of the fixed response.
            * `response_code` - The fixed response HTTP status code.
            * `response_message` - The fixed response message.
        * `forward_group_config` - Forward group configuration for ForwardGroup type action.
            * `server_group_sticky_session` - The config of group session stickiness.
                * `enabled` - Whether to enable sticky session stickiness. Valid values are 'on' and 'off'.
                * `timeout` - The sticky session timeout, in seconds.
            * `server_group_tuples` - The server group tuples.
                * `server_group_id` - The server group ID.
                * `weight` - The weight of the server group.
        * `redirect_config` - Redirect configuration for Redirect type action.
            * `redirect_domain` - The redirect domain.
            * `redirect_http_code` - The redirect HTTP code.
            * `redirect_port` - The redirect port.
            * `redirect_protocol` - The redirect protocol.
            * `redirect_uri` - The redirect URI.
        * `rewrite_config` - Rewrite configuration for Rewrite type action.
            * `path` - The rewrite path.
        * `traffic_limit_config` - Traffic limit configuration for TrafficLimit type action.
            * `qps` - The QPS limit.
        * `type` - The type of rule action. Valid values: ForwardGroup, Redirect, Rewrite, TrafficLimit.
    * `rule_conditions` - The rule conditions for standard edition forwarding rules.
        * `header_config` - Header configuration for Header type condition.
            * `key` - The header key.
            * `values` - The list of header values.
        * `host_config` - Host configuration for host type condition.
            * `values` - The list of domain names.
        * `method_config` - Method configuration for Method type condition.
            * `values` - The list of HTTP methods.
        * `path_config` - Path configuration for Path type condition.
            * `values` - The list of absolute paths.
        * `query_string_config` - Query string configuration.
            * `values` - The list of query string values.
                * `key` - The query string key.
                * `value` - The query string value.
        * `type` - The type of rule condition. Valid values: Host, Path, Header.
    * `rule_id` - The Id of Rule.
    * `server_group_id` - The Id of Server Group.
    * `traffic_limit_enabled` - Forwarding rule QPS rate limiting switch:
 on: enable.
off: disable (default).
    * `traffic_limit_qps` - When Rules.N.TrafficLimitEnabled is turned on, this field is required. Requests per second. Valid values are between 100 and 100000.
    * `url` - The Url of Rule.
* `total_count` - The total count of Rule query.


