---
subcategory: "APIG"
layout: "volcengine"
page_title: "Volcengine: volcengine_apig_gateway_service"
sidebar_current: "docs-volcengine-resource-apig_gateway_service"
description: |-
  Provides a resource to manage apig gateway service
---
# volcengine_apig_gateway_service
Provides a resource to manage apig gateway service
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
```
## Argument Reference
The following arguments are supported:
* `auth_spec` - (Required) The auth spec of the api gateway service.
* `gateway_id` - (Required, ForceNew) The gateway id of api gateway service.
* `protocol` - (Required) The protocol of api gateway service.
* `service_name` - (Required, ForceNew) The name of api gateway service.
* `comments` - (Optional) The comments of api gateway service.

The `auth_spec` object supports the following:

* `enable` - (Required) Whether the api gateway service enable auth.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `create_time` - The create time of the api gateway service.
* `message` - The error message of the api gateway service.
* `status` - The status of the api gateway service.


## Import
ApigGatewayService can be imported using the id, e.g.
```
$ terraform import volcengine_apig_gateway_service.default resource_id
```

