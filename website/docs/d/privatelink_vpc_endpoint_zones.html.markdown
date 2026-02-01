---
subcategory: "PRIVATELINK"
layout: "volcengine"
page_title: "Volcengine: volcengine_privatelink_vpc_endpoint_zones"
sidebar_current: "docs-volcengine-datasource-privatelink_vpc_endpoint_zones"
description: |-
  Use this data source to query detailed information of privatelink vpc endpoint zones
---
# volcengine_privatelink_vpc_endpoint_zones
Use this data source to query detailed information of privatelink vpc endpoint zones
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

data "volcengine_privatelink_vpc_endpoint_zones" "foo" {
  endpoint_id = volcengine_privatelink_vpc_endpoint_zone.foo.endpoint_id
}
```
## Argument Reference
The following arguments are supported:
* `endpoint_id` - (Optional) The endpoint id of query.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - Returns the total amount of the data list.
* `vpc_endpoint_zones` - The collection of query.
    * `id` - The Id of vpc endpoint zone.
    * `network_interface_id` - The network interface id of vpc endpoint.
    * `network_interface_ip` - The network interface ip of vpc endpoint.
    * `service_status` - The status of vpc endpoint service.
    * `subnet_id` - The subnet id of vpc endpoint.
    * `zone_domain` - The domain of vpc endpoint zone.
    * `zone_id` - The Id of vpc endpoint zone.
    * `zone_status` - The status of vpc endpoint zone.


