---
subcategory: "APIG"
layout: "volcengine"
page_title: "Volcengine: volcengine_apig_upstream_source"
sidebar_current: "docs-volcengine-resource-apig_upstream_source"
description: |-
  Provides a resource to manage apig upstream source
---
# volcengine_apig_upstream_source
Provides a resource to manage apig upstream source
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

resource "volcengine_apig_upstream_source" "foo-nacos" {
  gateway_id  = volcengine_apig_gateway.foo.id
  comments    = "acc-test-nacos"
  source_type = "Nacos"
  source_spec {
    nacos_source {
      nacos_id = "nd197ls631meck48imm7g"
      auth_config {
        basic {
          username = "nacos"
          password = "******"
        }
      }
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
```
## Argument Reference
The following arguments are supported:
* `gateway_id` - (Required, ForceNew) The gateway id of the apig upstream source.
* `source_spec` - (Required, ForceNew) The source spec of apig upstream source.
* `source_type` - (Required, ForceNew) The source type of the apig upstream. Valid values: `K8S`, `Nacos`.
* `comments` - (Optional) The comments of the apig upstream source.
* `ingress_settings` - (Optional) The ingress settings of apig upstream source.

The `auth_config` object supports the following:

* `basic` - (Optional, ForceNew) The basic auth config of nacos source.

The `basic` object supports the following:

* `password` - (Required, ForceNew) The password of basic auth config of nacos source.
* `username` - (Required, ForceNew) The username of basic auth config of nacos source.

The `ingress_settings` object supports the following:

* `enable_all_ingress_classes` - (Optional) Whether to enable all ingress classes.
* `enable_all_namespaces` - (Optional) Whether to enable all namespaces.
* `enable_ingress_without_ingress_class` - (Optional) Whether to enable ingress without ingress class.
* `enable_ingress` - (Optional) Whether to enable ingress.
* `ingress_classes` - (Optional) The ingress classes of ingress settings.
* `update_status` - (Optional) The update status of ingress settings.
* `watch_namespaces` - (Optional) The watch namespaces of ingress settings.

The `k8s_source` object supports the following:

* `cluster_id` - (Required, ForceNew) The cluster id of k8s source.
* `cluster_type` - (Optional, ForceNew) The cluster type of k8s source.

The `nacos_source` object supports the following:

* `nacos_id` - (Required, ForceNew) The nacos id of nacos source.
* `address` - (Optional, ForceNew) The address of nacos source.
* `auth_config` - (Optional, ForceNew) The auth config of nacos source.
* `context_path` - (Optional, ForceNew) The context path of nacos source.
* `grpc_port` - (Optional, ForceNew) The grpc port of nacos source.
* `http_port` - (Optional, ForceNew) The http port of nacos source.
* `nacos_name` - (Optional, ForceNew) The nacos name of nacos source.

The `source_spec` object supports the following:

* `k8s_source` - (Optional, ForceNew) The k8s source of apig upstream source.
* `nacos_source` - (Optional, ForceNew) The nacos source of apig upstream source.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
ApigUpstreamSource can be imported using the id, e.g.
```
$ terraform import volcengine_apig_upstream_source.default resource_id
```

