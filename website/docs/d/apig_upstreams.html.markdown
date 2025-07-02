---
subcategory: "APIG"
layout: "volcengine"
page_title: "Volcengine: volcengine_apig_upstreams"
sidebar_current: "docs-volcengine-datasource-apig_upstreams"
description: |-
  Use this data source to query detailed information of apig upstreams
---
# volcengine_apig_upstreams
Use this data source to query detailed information of apig upstreams
## Example Usage
```hcl
data "volcengine_apig_upstreams" "foo" {
  gateway_id = "gd13d8c6eq1emkiunq6p0"
  ids        = ["ud18p5krj5ce3htvrd0v0", "ud18ouitrjp6fhvu61n7g"]
}
```
## Argument Reference
The following arguments are supported:
* `gateway_id` - (Optional) The id of api gateway.
* `ids` - (Optional) A list of apig upstream IDs.
* `name_regex` - (Optional) A Name Regex of Resource.
* `name` - (Optional) The name of apig upstream. This field support fuzzy query.
* `output_file` - (Optional) File name where to save data source results.
* `resource_type` - (Optional) The resource type of apig upstream. Valid values: `Console`, `Ingress`.
* `source_type` - (Optional) The source type of apig upstream. Valid values: `VeFaas`, `ECS`, `FixedIP`, `K8S`, `Nacos`, `Domain`, `AIProvider`, `VeMLP`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of query.
* `upstreams` - The collection of query.
    * `backend_target_list` - The backend target list of apig upstream.
        * `health_status` - The health status of apig upstream backend.
        * `ip` - The ip of apig upstream backend.
        * `port` - The port of apig upstream backend.
    * `circuit_breaking_settings` - The circuit breaking settings of apig upstream.
        * `base_ejection_time` - The base ejection time of circuit breaking. Unit: ms.
        * `consecutive_errors` - The consecutive errors of circuit breaking.
        * `enable` - Whether the circuit breaking is enabled.
        * `interval` - The interval of circuit breaking. Unit: ms.
        * `max_ejection_percent` - The max ejection percent of circuit breaking.
        * `min_health_percent` - The min health percent of circuit breaking.
    * `comments` - The comments of apig upstream.
    * `create_time` - The create time of apig upstream.
    * `gateway_id` - The id of api gateway.
    * `id` - The id of apig upstream.
    * `load_balancer_settings` - The load balancer settings of apig upstream.
        * `consistent_hash_lb` - The consistent hash lb of apig upstream.
            * `hash_key` - The hash key of apig upstream consistent hash lb.
            * `http_cookie` - The http cookie of apig upstream consistent hash lb.
                * `name` - The name of apig upstream consistent hash lb http cookie.
                * `path` - The path of apig upstream consistent hash lb http cookie.
                * `ttl` - The ttl of apig upstream consistent hash lb http cookie.
            * `http_header_name` - The http header name of apig upstream consistent hash lb.
            * `http_query_parameter_name` - The http query parameter name of apig upstream consistent hash lb.
            * `use_source_ip` - The use source ip of apig upstream consistent hash lb.
        * `lb_policy` - The load balancer policy of apig upstream.
        * `simple_lb` - The simple load balancer of apig upstream.
        * `warmup_duration` - The warmup duration of apig upstream lb.
    * `name` - The name of apig upstream.
    * `protocol` - The protocol of apig upstream.
    * `resource_type` - The resource type of apig upstream.
    * `source_type` - The source type of apig upstream.
    * `tls_settings` - The tls settings of apig upstream.
        * `sni` - The sni of apig upstream tls setting.
        * `tls_mode` - The tls mode of apig upstream tls setting.
    * `update_time` - The update time of apig upstream.
    * `upstream_spec` - The upstream spec of apig upstream.
        * `ai_provider` - The ai provider of apig upstream.
            * `base_url` - The base url of ai provider.
            * `custom_body_params` - The custom body params of ai provider.
            * `custom_header_params` - The custom header params of ai provider.
            * `custom_model_service` - The custom model service of ai provider.
                * `name` - The name of custom model service.
                * `namespace` - The namespace of custom model service.
                * `port` - The port of custom model service.
            * `name` - The name of ai provider.
            * `token` - The token of ai provider.
        * `domain` - The domain of apig upstream.
            * `domain_list` - The domain list of apig upstream.
                * `domain` - The domain of apig upstream.
                * `port` - The port of domain.
            * `protocol` - The protocol of apig upstream.
        * `ecs_list` - The ecs list of apig upstream.
            * `ecs_id` - The instance id of ecs.
            * `ip` - The ip of ecs.
            * `port` - The port of ecs.
        * `fixed_ip_list` - The fixed ip list of apig upstream.
            * `ip` - The ip of apig upstream.
            * `port` - The port of apig upstream.
        * `k8s_service` - The k8s service of apig upstream.
            * `name` - The name of k8s service.
            * `namespace` - The namespace of k8s service.
            * `port` - The port of k8s service.
        * `nacos_service` - The nacos service of apig upstream.
            * `group` - The group of nacos service.
            * `namespace_id` - The namespace id of nacos service.
            * `namespace` - The namespace of nacos service.
            * `service` - The service of nacos service.
            * `upstream_source_id` - The upstream source id.
        * `ve_faas` - The vefaas of apig upstream.
            * `function_id` - The function id of vefaas.
        * `ve_mlp` - The mlp of apig upstream.
            * `k8s_service` - The k8s service of mlp.
                * `cluster_info` - The cluster info of k8s service.
                    * `account_id` - The account id of k8s service.
                    * `cluster_name` - The cluster name of k8s service.
                * `name` - The name of k8s service.
                * `namespace` - The namespace of k8s service.
                * `port` - The port of k8s service.
            * `service_discover_type` - The service discover type of mlp.
            * `service_id` - The service id of mlp.
            * `service_name` - The service name of mlp.
            * `service_url` - The service url of mlp.
            * `upstream_source_id` - The upstream source id.
    * `version_details` - The version details of apig upstream.
        * `labels` - The labels of apig upstream version.
            * `key` - The key of apig upstream version label.
            * `value` - The value of apig upstream version label.
        * `name` - The name of apig upstream version.
        * `update_time` - The update time of apig upstream version.


