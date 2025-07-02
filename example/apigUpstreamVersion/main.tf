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

resource "volcengine_apig_upstream_version" "foo" {
  upstream_id = volcengine_apig_upstream.foo-k8s.id
  upstream_version {
    name = "acc-test-version"
    labels {
      key   = "k1"
      value = "v2"
    }
  }
}
