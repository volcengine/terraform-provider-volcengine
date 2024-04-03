---
subcategory: "PRIVATELINK"
layout: "volcengine"
page_title: "Volcengine: volcengine_privatelink_vpc_endpoint_zone"
sidebar_current: "docs-volcengine-resource-privatelink_vpc_endpoint_zone"
description: |-
  Provides a resource to manage privatelink vpc endpoint zone
---
# volcengine_privatelink_vpc_endpoint_zone
Provides a resource to manage privatelink vpc endpoint zone
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

resource "volcengine_security_group" "foo" {
  security_group_name = "acc-test-security-group"
  vpc_id              = volcengine_vpc.foo.id
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

resource "volcengine_privatelink_vpc_endpoint_service" "foo" {
  resources {
    resource_id   = volcengine_clb.foo.id
    resource_type = "CLB"
  }
  description         = "acc-test"
  auto_accept_enabled = true
}

resource "volcengine_privatelink_vpc_endpoint" "foo" {
  security_group_ids = [volcengine_security_group.foo.id]
  service_id         = volcengine_privatelink_vpc_endpoint_service.foo.id
  endpoint_name      = "acc-test-ep"
  description        = "acc-test"
}

resource "volcengine_privatelink_vpc_endpoint_zone" "foo" {
  endpoint_id        = volcengine_privatelink_vpc_endpoint.foo.id
  subnet_id          = volcengine_subnet.foo.id
  private_ip_address = "172.16.0.251"
}
```
## Argument Reference
The following arguments are supported:
* `endpoint_id` - (Required, ForceNew) The endpoint id of vpc endpoint zone.
* `subnet_id` - (Required, ForceNew) The subnet id of vpc endpoint zone.
* `private_ip_address` - (Optional, ForceNew) The private ip address of vpc endpoint zone.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `network_interface_id` - The network interface id of vpc endpoint.
* `zone_domain` - The domain of vpc endpoint zone.
* `zone_id` - The Id of vpc endpoint zone.
* `zone_status` - The status of vpc endpoint zone.


## Import
VpcEndpointZone can be imported using the endpointId:subnetId, e.g.
```
$ terraform import volcengine_privatelink_vpc_endpoint_zone.default ep-3rel75r081l345zsk2i59****:subnet-2bz47q19zhx4w2dx0eevn****
```

