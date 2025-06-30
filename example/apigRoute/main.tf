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
