---
subcategory: "PRIVATELINK"
layout: "volcengine"
page_title: "Volcengine: volcengine_privatelink_vpc_endpoint_service_resource"
sidebar_current: "docs-volcengine-resource-privatelink_vpc_endpoint_service_resource"
description: |-
  Provides a resource to manage privatelink vpc endpoint service resource
---
# volcengine_privatelink_vpc_endpoint_service_resource
Provides a resource to manage privatelink vpc endpoint service resource
## Example Usage
```hcl
data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_clb" "foo" {
  type                       = "public"
  subnet_id                  = volcengine_subnet.foo.id
  load_balancer_spec         = "small_1"
  description                = "acc-test-demo"
  load_balancer_name         = "acc-test-clb"
  load_balancer_billing_type = "PostPaid"
  eip_billing_config {
    isp              = "BGP"
    eip_billing_type = "PostPaidByBandwidth"
    bandwidth        = 1
  }
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_clb" "foo1" {
  type                       = "public"
  subnet_id                  = volcengine_subnet.foo.id
  load_balancer_spec         = "small_1"
  description                = "acc-test-demo"
  load_balancer_name         = "acc-test-clb-new"
  load_balancer_billing_type = "PostPaid"
  eip_billing_config {
    isp              = "BGP"
    eip_billing_type = "PostPaidByBandwidth"
    bandwidth        = 1
  }
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_privatelink_vpc_endpoint_service" "foo" {
  resources {
    resource_id   = volcengine_clb.foo.id
    resource_type = "CLB"
  }
  description         = "acc-test"
  auto_accept_enabled = true
}

resource "volcengine_privatelink_vpc_endpoint_service_resource" "foo" {
  service_id  = volcengine_privatelink_vpc_endpoint_service.foo.id
  resource_id = volcengine_clb.foo1.id
}
```
## Argument Reference
The following arguments are supported:
* `resource_id` - (Required, ForceNew) The id of resource. It is not recommended to use this resource for binding resources, it is recommended to use the resources field of vpc_endpoint_service for binding. If using this resource and vpc_endpoint_service jointly for operations, use lifecycle ignore_changes to suppress changes to the resources field in vpc_endpoint_service.
* `service_id` - (Required, ForceNew) The id of service.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
VpcEndpointServiceResource can be imported using the serviceId:resourceId, e.g.
```
$ terraform import volcengine_privatelink_vpc_endpoint_service_resource.default epsvc-2fe630gurkl37k5gfuy33****:clb-bp1o94dp5i6ea****
```
It is not recommended to use this resource for binding resources, it is recommended to use the resources field of vpc_endpoint_service for binding.
If using this resource and vpc_endpoint_service jointly for operations, use lifecycle ignore_changes to suppress changes to the resources field in vpc_endpoint_service.

