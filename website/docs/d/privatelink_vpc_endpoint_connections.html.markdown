---
subcategory: "PRIVATELINK"
layout: "volcengine"
page_title: "Volcengine: volcengine_privatelink_vpc_endpoint_connections"
sidebar_current: "docs-volcengine-datasource-privatelink_vpc_endpoint_connections"
description: |-
  Use this data source to query detailed information of privatelink vpc endpoint connections
---
# volcengine_privatelink_vpc_endpoint_connections
Use this data source to query detailed information of privatelink vpc endpoint connections
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
  description = "acc-test"
}

resource "volcengine_privatelink_vpc_endpoint" "foo" {
  security_group_ids = [volcengine_security_group.foo.id]
  service_id         = volcengine_privatelink_vpc_endpoint_service.foo.id
  endpoint_name      = "acc-test-ep"
  description        = "acc-test"
}

resource "volcengine_privatelink_vpc_endpoint_connection" "foo" {
  endpoint_id = volcengine_privatelink_vpc_endpoint.foo.id
  service_id  = volcengine_privatelink_vpc_endpoint_service.foo.id
}

data "volcengine_privatelink_vpc_endpoint_connections" "foo" {
  endpoint_id = volcengine_privatelink_vpc_endpoint_connection.foo.endpoint_id
  service_id  = volcengine_privatelink_vpc_endpoint_connection.foo.service_id
}
```
## Argument Reference
The following arguments are supported:
* `service_id` - (Required) The id of the vpc endpoint service.
* `endpoint_id` - (Optional) The id of the vpc endpoint.
* `endpoint_owner_account_id` - (Optional) The account id of the vpc endpoint.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `connections` - The list of query.
    * `connection_status` - The status of the connection.
    * `creation_time` - The create time of the connection.
    * `endpoint_id` - The id of the vpc endpoint.
    * `endpoint_owner_account_id` - The account id of the vpc endpoint.
    * `endpoint_vpc_id` - The vpc id of the vpc endpoint.
    * `service_id` - The id of the vpc endpoint service.
    * `update_time` - The update time of the connection.
    * `zones` - The available zones.
        * `network_interface_id` - The id of the network interface.
        * `network_interface_ip` - The ip address of the network interface.
        * `resource_id` - The id of the resource.
        * `subnet_id` - The id of the subnet.
        * `zone_domain` - The domain of the zone.
        * `zone_id` - The id of the zone.
        * `zone_status` - The status of the zone.
* `total_count` - Returns the total amount of the data list.


