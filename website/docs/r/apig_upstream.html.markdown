---
subcategory: "APIG"
layout: "volcengine"
page_title: "Volcengine: volcengine_apig_upstream"
sidebar_current: "docs-volcengine-resource-apig_upstream"
description: |-
  Provides a resource to manage apig upstream
---
# volcengine_apig_upstream
Provides a resource to manage apig upstream
## Example Usage
```hcl
data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo1" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_subnet" "foo2" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.1.0/24"
  zone_id     = data.volcengine_zones.foo.zones[1].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_apig_gateway" "foo" {
  name         = "acc-test-apig"
  type         = "standard"
  comments     = "acc-test"
  project_name = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
  network_spec {
    vpc_id     = volcengine_vpc.foo.id
    subnet_ids = [volcengine_subnet.foo1.id, volcengine_subnet.foo2.id]
  }
  resource_spec {
    replicas                    = 2
    instance_spec_code          = "1c2g"
    clb_spec_code               = "small_1"
    public_network_billing_type = "bandwidth"
    public_network_bandwidth    = 1
    network_type {
      enable_public_network  = true
      enable_private_network = true
    }
  }
  log_spec {
    enable     = true
    project_id = "d3cb87c0-faeb-4074-b1ee-9bd747865a76"
    topic_id   = "d339482e-d86d-4bd8-a9bb-f270417f00a1"
  }
  monitor_spec {
    enable       = true
    workspace_id = "4ed1caf3-279d-4c5f-8301-87ea38e92ffc"
  }
}

resource "volcengine_apig_gateway_service" "foo" {
  gateway_id   = volcengine_apig_gateway.foo.id
  service_name = "acc-test-apig-service"
  comments     = "acc-test"
  protocol     = ["HTTP", "HTTPS"]
  auth_spec {
    enable = false
  }
}

resource "volcengine_apig_upstream" "foo-fixed-ip" {
  gateway_id    = volcengine_apig_gateway.foo.id
  name          = "acc-test-apig-upstream-ip"
  comments      = "acc-test"
  resource_type = "Console"
  protocol      = "HTTP"
  load_balancer_settings {
    lb_policy = "ConsistentHashLB"
    consistent_hash_lb {
      hash_key = "HTTPCookie"
      http_cookie {
        name = "test"
        path = "/"
        ttl  = 300
      }
    }
  }
  tls_settings {
    tls_mode = "SIMPLE"
    sni      = "test"
  }
  circuit_breaking_settings {
    enable               = false
    consecutive_errors   = 5
    interval             = 10000
    base_ejection_time   = 30000
    max_ejection_percent = 20
    min_health_percent   = 60
  }

  source_type = "FixedIP"
  upstream_spec {
    fixed_ip_list {
      ip   = "172.16.0.10"
      port = 8080
    }
    fixed_ip_list {
      ip   = "172.16.0.20"
      port = 8090
    }
  }
}

resource "volcengine_apig_upstream" "foo-vefaas" {
  gateway_id    = volcengine_apig_gateway.foo.id
  name          = "acc-test-apig-upstream-vefaas"
  comments      = "acc-test"
  resource_type = "Console"
  protocol      = "HTTP"
  source_type   = "VeFaas"
  upstream_spec {
    ve_faas {
      function_id = "crnrfajj"
    }
  }
}

resource "volcengine_apig_upstream_source" "foo-k8s" {
  gateway_id  = volcengine_apig_gateway.foo.id
  comments    = "acc-test-k8s"
  source_type = "K8S"
  source_spec {
    k8s_source {
      cluster_id = "cd197sac4mpmnruh7um80"
    }
  }
  ingress_settings {
    enable_ingress   = true
    update_status    = true
    ingress_classes  = ["test"]
    watch_namespaces = ["default"]
  }
}

resource "volcengine_apig_upstream" "foo-k8s" {
  gateway_id    = volcengine_apig_gateway.foo.id
  name          = "acc-test-apig-upstream-k8s"
  comments      = "acc-test"
  resource_type = "Console"
  protocol      = "HTTP"
  load_balancer_settings {
    lb_policy = "ConsistentHashLB"
    consistent_hash_lb {
      hash_key = "HTTPCookie"
      http_cookie {
        name = "test"
        path = "/"
        ttl  = 300
      }
    }
  }
  tls_settings {
    tls_mode = "SIMPLE"
    sni      = "test"
  }
  circuit_breaking_settings {
    enable               = false
    consecutive_errors   = 5
    interval             = 10000
    base_ejection_time   = 30000
    max_ejection_percent = 20
    min_health_percent   = 60
  }
  source_type = "K8S"
  upstream_spec {
    k8s_service {
      namespace = "default"
      name      = "kubernetes"
      port      = 443
    }
  }
  depends_on = [volcengine_apig_upstream_source.foo-k8s]
}
```
## Argument Reference
The following arguments are supported:
* `gateway_id` - (Required, ForceNew) The gateway id of the apig upstream.
* `name` - (Required, ForceNew) The name of the apig upstream.
* `protocol` - (Required) The protocol of the apig upstream. Valid values: `HTTP`, `HTTP2`, `GRPC`.
* `source_type` - (Required, ForceNew) The source type of the apig upstream. Valid values: `VeFaas`, `ECS`, `FixedIP`, `K8S`, `Nacos`, `Domain`, `AIProvider`, `VeMLP`.
* `upstream_spec` - (Required) The upstream spec of apig upstream.
* `circuit_breaking_settings` - (Optional) The circuit breaking settings of apig upstream.
* `comments` - (Optional) The comments of the apig upstream.
* `load_balancer_settings` - (Optional) The load balancer settings of apig upstream.
* `resource_type` - (Optional, ForceNew) The resource type of the apig upstream. Valid values: `Console`, `Ingress`.
* `tls_settings` - (Optional) The tls settings of apig upstream.

The `ai_provider` object supports the following:

* `base_url` - (Required) The base url of ai provider.
* `name` - (Required) The name of ai provider.
* `token` - (Required) The token of ai provider.
* `custom_body_params` - (Optional) The custom body params of ai provider.
* `custom_header_params` - (Optional) The custom header params of ai provider.
* `custom_model_service` - (Optional) The custom model service of ai provider.

The `circuit_breaking_settings` object supports the following:

* `enable` - (Required) Whether the circuit breaking is enabled.
* `base_ejection_time` - (Optional) The base ejection time of circuit breaking. Unit: ms. Default is 10s.
* `consecutive_errors` - (Optional) The consecutive errors of circuit breaking. Default is 5.
* `interval` - (Optional) The interval of circuit breaking. Unit: ms. Default is 10s.
* `max_ejection_percent` - (Optional) The max ejection percent of circuit breaking. Default is 20%.
* `min_health_percent` - (Optional) The min health percent of circuit breaking. Default is 60%.

The `cluster_info` object supports the following:

* `account_id` - (Required) The account id of k8s service.
* `cluster_name` - (Required) The cluster name of k8s service.

The `consistent_hash_lb` object supports the following:

* `hash_key` - (Required) The hash key of apig upstream consistent hash lb. Valid values: `HTTPCookie`, `HttpHeaderName`, `HttpQueryParameterName`, `UseSourceIp`.
* `http_cookie` - (Optional) The http cookie of apig upstream consistent hash lb.
* `http_header_name` - (Optional) The http header name of apig upstream consistent hash lb.
* `http_query_parameter_name` - (Optional) The http query parameter name of apig upstream consistent hash lb.
* `use_source_ip` - (Optional) The use source ip of apig upstream consistent hash lb.

The `custom_model_service` object supports the following:

* `name` - (Required) The name of custom model service.
* `namespace` - (Required) The namespace of custom model service.
* `port` - (Required) The port of custom model service.

The `domain_list` object supports the following:

* `domain` - (Required) The domain of apig upstream.
* `port` - (Optional) The port of domain. Default is 80 for HTTP, 443 for HTTPS.

The `domain` object supports the following:

* `domain_list` - (Required) The domain list of apig upstream.
* `protocol` - (Optional) The protocol of apig upstream. Valid values: `HTTP`, `HTTPS`.

The `ecs_list` object supports the following:

* `ecs_id` - (Required) The instance id of ecs.
* `ip` - (Required) The ip of ecs.
* `port` - (Required) The port of ecs.

The `fixed_ip_list` object supports the following:

* `ip` - (Required) The ip of apig upstream.
* `port` - (Required) The port of apig upstream.

The `http_cookie` object supports the following:

* `name` - (Required) The name of apig upstream consistent hash lb http cookie.
* `path` - (Required) The path of apig upstream consistent hash lb http cookie.
* `ttl` - (Required) The ttl of apig upstream consistent hash lb http cookie.

The `k8s_service` object supports the following:

* `cluster_info` - (Required) The cluster info of k8s service.
* `name` - (Required) The name of k8s service.
* `namespace` - (Required) The namespace of k8s service.
* `port` - (Required) The port of k8s service.

The `k8s_service` object supports the following:

* `name` - (Required) The name of k8s service.
* `namespace` - (Required) The namespace of k8s service.
* `port` - (Required) The port of k8s service.

The `load_balancer_settings` object supports the following:

* `lb_policy` - (Required) The load balancer policy of apig upstream. Valid values: `SimpleLB`, `ConsistentHashLB`.
* `consistent_hash_lb` - (Optional) The consistent hash lb of apig upstream.
* `simple_lb` - (Optional) The simple load balancer of apig upstream. Valid values: `ROUND_ROBIN`, `LEAST_CONN`, `RANDOM`.
* `warmup_duration` - (Optional) The warmup duration of apig upstream lb. This field is valid when the simple_lb is `ROUND_ROBIN` or `LEAST_CONN`.

The `nacos_service` object supports the following:

* `group` - (Required) The group of nacos service.
* `namespace` - (Required) The namespace of nacos service.
* `service` - (Required) The service of nacos service.
* `upstream_source_id` - (Required) The upstream source id.
* `namespace_id` - (Optional) The namespace id of nacos service.

The `tls_settings` object supports the following:

* `tls_mode` - (Required) The tls mode of apig upstream tls setting. Valid values: `DISABLE`, `SIMPLE`.
* `sni` - (Optional) The sni of apig upstream tls setting.

The `upstream_spec` object supports the following:

* `ai_provider` - (Optional) The ai provider of apig upstream.
* `domain` - (Optional) The domain of apig upstream.
* `ecs_list` - (Optional) The ecs list of apig upstream.
* `fixed_ip_list` - (Optional) The fixed ip list of apig upstream.
* `k8s_service` - (Optional) The k8s service of apig upstream.
* `nacos_service` - (Optional) The nacos service of apig upstream.
* `ve_faas` - (Optional) The vefaas of apig upstream.
* `ve_mlp` - (Optional) The mlp of apig upstream.

The `ve_faas` object supports the following:

* `function_id` - (Required) The function id of vefaas.

The `ve_mlp` object supports the following:

* `k8s_service` - (Required) The k8s service of mlp.
* `service_discover_type` - (Required) The service discover type of mlp.
* `service_id` - (Required) The service id of mlp.
* `service_name` - (Optional) The service name of mlp.
* `service_url` - (Optional) The service url of mlp.
* `upstream_source_id` - (Optional) The upstream source id.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `create_time` - The create time of apig upstream.
* `update_time` - The update time of apig upstream.
* `version_details` - The version details of apig upstream.
    * `labels` - The labels of apig upstream version.
        * `key` - The key of apig upstream version label.
        * `value` - The value of apig upstream version label.
    * `name` - The name of apig upstream version.
    * `update_time` - The update time of apig upstream version.


## Import
ApigUpstream can be imported using the id, e.g.
```
$ terraform import volcengine_apig_upstream.default resource_id
```

