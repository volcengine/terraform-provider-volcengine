---
subcategory: "APIG"
layout: "volcengine"
page_title: "Volcengine: volcengine_apig_routes"
sidebar_current: "docs-volcengine-datasource-apig_routes"
description: |-
  Use this data source to query detailed information of apig routes
---
# volcengine_apig_routes
Use this data source to query detailed information of apig routes
## Example Usage
```hcl
data "volcengine_apig_routes" "foo" {
  gateway_id = "gd1ek1ki9optek6ooabh0"
}
```
## Argument Reference
The following arguments are supported:
* `gateway_id` - (Optional) The id of api gateway.
* `name_regex` - (Optional) A Name Regex of Resource.
* `name` - (Optional) The name of api gateway route. This field support fuzzy query.
* `output_file` - (Optional) File name where to save data source results.
* `path` - (Optional) The path of api gateway route.
* `resource_type` - (Optional) The resource type of route. Valid values: `Console`, `Ingress`.
* `service_id` - (Optional) The id of api gateway service.
* `upstream_id` - (Optional) The id of api gateway upstream.
* `upstream_version` - (Optional) The version of api gateway upstream.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `routes` - The collection of query.
    * `advanced_setting` - The advanced setting of the api gateway route.
        * `cors_policy_setting` - The cors policy setting of the api gateway route.
            * `enable` - Whether the cors policy setting is enabled.
        * `header_operations` - The header operations of the api gateway route.
            * `direction_type` - The direction type of the header.
            * `key` - The key of the header.
            * `operation` - The operation of the header.
            * `value` - The value of the header.
        * `mirror_policies` - The mirror policies of the api gateway route.
            * `percent` - The percent of the mirror policy.
                * `value` - The percent value of the mirror policy.
            * `upstream` - The upstream of the mirror policy.
                * `type` - The type of the api gateway upstream.
                * `upstream_id` - The id of the api gateway upstream.
                * `version` - The version of the api gateway upstream.
        * `retry_policy_setting` - The retry policy setting of the api gateway route.
            * `attempts` - The attempts of the api gateway route.
            * `enable` - Whether the retry policy setting is enabled.
            * `http_codes` - The http codes of the api gateway route.
            * `per_try_timeout` - The per try timeout of the api gateway route.
            * `retry_on` - The retry on of the api gateway route.
        * `timeout_setting` - The timeout setting of the api gateway route.
            * `enable` - Whether the timeout setting is enabled.
            * `timeout` - The timeout of the api gateway route.
        * `url_rewrite_setting` - The url rewrite setting of the api gateway route.
            * `enable` - Whether the url rewrite setting is enabled.
            * `url_rewrite` - The url rewrite path of the api gateway route.
    * `create_time` - The create time of the api gateway route.
    * `custom_domains` - The custom domains of the api gateway route.
        * `domain` - The custom domain of the api gateway route.
        * `id` - The id of the custom domain.
    * `domains` - The domains of the api gateway route.
        * `domain` - The domain of the api gateway route.
        * `type` - The type of the domain.
    * `enable` - Whether the api gateway route is enabled.
    * `id` - The id of the api gateway route.
    * `match_rule` - The match rule of the api gateway route.
        * `header` - The header of the api gateway route.
            * `key` - The key of the header.
            * `value` - The path of the api gateway route.
                * `match_content` - The match content of the api gateway route.
                * `match_type` - The match type of the api gateway route.
        * `method` - The method of the api gateway route.
        * `path` - The path of the api gateway route.
            * `match_content` - The match content of the api gateway route.
            * `match_type` - The match type of the api gateway route.
        * `query_string` - The query string of the api gateway route.
            * `key` - The key of the query string.
            * `value` - The path of the api gateway route.
                * `match_content` - The match content of the api gateway route.
                * `match_type` - The match type of the api gateway route.
    * `name` - The name of the api gateway route.
    * `priority` - The priority of the api gateway route.
    * `reason` - The reason of the api gateway route.
    * `resource_type` - The resource type of route. Valid values: `Console`, `Ingress`.
    * `service_id` - The id of the api gateway service.
    * `service_name` - The name of the api gateway service.
    * `status` - The status of the api gateway route.
    * `update_time` - The update time of the api gateway route.
    * `upstream_list` - The upstream list of the api gateway route.
        * `ai_provider_settings` - The ai provider settings of the api gateway route.
            * `model` - The model of the ai provider.
            * `target_path` - The target path of the ai provider.
        * `upstream_id` - The id of the api gateway upstream.
        * `version` - The version of the api gateway upstream.
        * `weight` - The weight of the api gateway upstream.
* `total_count` - The total count of query.


