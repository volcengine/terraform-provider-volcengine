---
subcategory: "APIG"
layout: "volcengine"
page_title: "Volcengine: volcengine_apig_route"
sidebar_current: "docs-volcengine-resource-apig_route"
description: |-
  Provides a resource to manage apig route
---
# volcengine_apig_route
Provides a resource to manage apig route
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
    fixed_ip_list {
      ip   = "172.16.0.30"
      port = 8099
    }
  }
}

resource "volcengine_apig_route" "foo" {
  service_id    = volcengine_apig_gateway_service.foo.id
  name          = "acc-test-route"
  resource_type = "Console"
  priority      = 2
  enable        = true
  upstream_list {
    upstream_id = volcengine_apig_upstream.foo-fixed-ip.id
    weight      = 100
  }
  match_rule {
    path {
      match_type    = "Prefix"
      match_content = "/test"
    }
    method = ["GET", "POST"]
    query_string {
      key = "test-key"
      value {
        match_type    = "Exact"
        match_content = "test-value"
      }
    }
    header {
      key = "test-header"
      value {
        match_type    = "Regex"
        match_content = "test-value"
      }
    }
  }
  advanced_setting {
    timeout_setting {
      enable  = false
      timeout = 10
    }
    cors_policy_setting {
      enable = false
    }
    url_rewrite_setting {
      enable      = true
      url_rewrite = "/test"
    }
    retry_policy_setting {
      enable          = true
      attempts        = 5
      per_try_timeout = 1000
      retry_on        = ["5xx", "reset"]
      http_codes      = ["500", "502", "503", "504"]
    }
    header_operations {
      operation      = "add"
      key            = "test-header-req"
      value          = "test-value"
      direction_type = "request"
    }
    header_operations {
      operation      = "set"
      key            = "test-header-resp"
      value          = "test-value"
      direction_type = "response"
    }
    mirror_policies {
      upstream {
        upstream_id = volcengine_apig_upstream.foo-fixed-ip.id
        type        = "fixed_ip"
      }
      percent {
        value = 50
      }
    }
  }
}
```
## Argument Reference
The following arguments are supported:
* `match_rule` - (Required) The match rule of the api gateway route.
* `name` - (Required, ForceNew) The name of the apig route.
* `service_id` - (Required, ForceNew) The service id of the apig route.
* `upstream_list` - (Required) The upstream list of the api gateway route.
* `advanced_setting` - (Optional) The advanced setting of the api gateway route.
* `enable` - (Optional) Whether the apig route is enabled. Default is `false`.
* `priority` - (Optional) The priority of the apig route. Valid values: 0~100.
* `resource_type` - (Optional, ForceNew) The resource type of the apig route. Valid values: `Console`, `Ingress` Default is `Console`.

The `advanced_setting` object supports the following:

* `cors_policy_setting` - (Optional) The cors policy setting of the api gateway route.
* `header_operations` - (Optional) The header operations of the api gateway route.
* `mirror_policies` - (Optional) The mirror policies of the api gateway route.
* `retry_policy_setting` - (Optional) The retry policy setting of the api gateway route.
* `timeout_setting` - (Optional) The timeout setting of the api gateway route.
* `url_rewrite_setting` - (Optional) The url rewrite setting of the api gateway route.

The `ai_provider_settings` object supports the following:

* `model` - (Required) The model of the ai provider.
* `target_path` - (Required) The target path of the ai provider.

The `cors_policy_setting` object supports the following:

* `enable` - (Optional) Whether the cors policy setting is enabled.

The `header_operations` object supports the following:

* `key` - (Required) The key of the header.
* `operation` - (Required) The operation of the header. Valid values: `set`, `add`, `remove`.
* `direction_type` - (Optional) The direction type of the header. Valid values: `request`, `response`.
* `value` - (Optional) The value of the header.

The `header` object supports the following:

* `key` - (Required) The key of the header.
* `value` - (Required) The path of the api gateway route.

The `match_rule` object supports the following:

* `path` - (Required) The path of the api gateway route.
* `header` - (Optional) The header of the api gateway route.
* `method` - (Optional) The method of the api gateway route. Valid values: `GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `CONNECT`.
* `query_string` - (Optional) The query string of the api gateway route.

The `mirror_policies` object supports the following:

* `upstream` - (Required) The upstream of the mirror policy.
* `percent` - (Optional) The percent of the mirror policy.

The `path` object supports the following:

* `match_content` - (Required) The match content of the api gateway route.
* `match_type` - (Required) The match type of the api gateway route. Valid values: `Prefix`, `Exact`, `Regex`.

The `percent` object supports the following:

* `value` - (Required) The percent value of the mirror policy.

The `query_string` object supports the following:

* `key` - (Required) The key of the query string.
* `value` - (Required) The path of the api gateway route.

The `retry_policy_setting` object supports the following:

* `attempts` - (Optional) The attempts of the api gateway route.
* `enable` - (Optional) Whether the retry policy setting is enabled.
* `http_codes` - (Optional) The http codes of the api gateway route.
* `per_try_timeout` - (Optional) The per try timeout of the api gateway route.
* `retry_on` - (Optional) The retry on of the api gateway route. Valid values: `5xx`, `reset`, `connect-failure`, `refused-stream`, `cancelled`, `deadline-exceeded`, `internal`, `resource-exhausted`, `unavailable`.

The `timeout_setting` object supports the following:

* `enable` - (Optional) Whether the timeout setting is enabled.
* `timeout` - (Optional) The timeout of the api gateway route. Unit: s.

The `upstream_list` object supports the following:

* `upstream_id` - (Required) The id of the api gateway upstream.
* `weight` - (Required) The weight of the api gateway upstream. Valid values: 0~10000.
* `ai_provider_settings` - (Optional) The ai provider settings of the api gateway route.
* `version` - (Optional) The version of the api gateway upstream.

The `upstream` object supports the following:

* `type` - (Required) The type of the api gateway upstream.
* `upstream_id` - (Required) The id of the api gateway upstream.
* `version` - (Optional) The version of the api gateway upstream.

The `url_rewrite_setting` object supports the following:

* `enable` - (Optional) Whether the url rewrite setting is enabled.
* `url_rewrite` - (Optional) The url rewrite path of the api gateway route.

The `value` object supports the following:

* `match_content` - (Required) The match content of the api gateway route.
* `match_type` - (Required) The match type of the api gateway route. Valid values: `Prefix`, `Exact`, `Regex`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `create_time` - The create time of the api gateway route.
* `custom_domains` - The custom domains of the api gateway route.
    * `domain` - The custom domain of the api gateway route.
    * `id` - The id of the custom domain.
* `domains` - The domains of the api gateway route.
    * `domain` - The domain of the api gateway route.
    * `type` - The type of the domain.
* `reason` - The reason of the api gateway route.
* `status` - The status of the api gateway route.
* `update_time` - The update time of the api gateway route.


## Import
ApigRoute can be imported using the id, e.g.
```
$ terraform import volcengine_apig_route.default resource_id
```

