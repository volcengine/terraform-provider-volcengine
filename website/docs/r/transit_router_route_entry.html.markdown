---
subcategory: "TRANSIT_ROUTER"
layout: "volcengine"
page_title: "Volcengine: volcengine_transit_router_route_entry"
sidebar_current: "docs-volcengine-resource-transit_router_route_entry"
description: |-
  Provides a resource to manage transit router route entry
---
# volcengine_transit_router_route_entry
Provides a resource to manage transit router route entry
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

resource "volcengine_vpn_gateway" "foo" {
  vpc_id           = volcengine_vpc.foo.id
  subnet_id        = volcengine_subnet.foo.id
  bandwidth        = 20
  vpn_gateway_name = "acc-test"
  description      = "acc-test"
  period           = 2
}

resource "volcengine_customer_gateway" "foo" {
  ip_address            = "192.0.1.3"
  customer_gateway_name = "acc-test"
  description           = "acc-test"
}

resource "volcengine_vpn_connection" "foo" {
  vpn_connection_name   = "acc-tf-test"
  description           = "acc-tf-test"
  attach_type           = "TransitRouter"
  vpn_gateway_id        = volcengine_vpn_gateway.foo.id
  customer_gateway_id   = volcengine_customer_gateway.foo.id
  local_subnet          = ["192.168.0.0/22"]
  remote_subnet         = ["192.161.0.0/20"]
  dpd_action            = "none"
  nat_traversal         = true
  ike_config_psk        = "acctest@!3"
  ike_config_version    = "ikev1"
  ike_config_mode       = "main"
  ike_config_enc_alg    = "aes"
  ike_config_auth_alg   = "md5"
  ike_config_dh_group   = "group2"
  ike_config_lifetime   = 9000
  ike_config_local_id   = "acc_test"
  ike_config_remote_id  = "acc_test"
  ipsec_config_enc_alg  = "aes"
  ipsec_config_auth_alg = "sha256"
  ipsec_config_dh_group = "group2"
  ipsec_config_lifetime = 9000
  log_enabled           = false
}

resource "volcengine_transit_router_vpn_attachment" "foo" {
  zone_id                        = "cn-beijing-a"
  transit_router_attachment_name = "tf-test-acc"
  description                    = "tf-test-acc-desc"
  transit_router_id              = volcengine_transit_router.foo.id
  vpn_connection_id              = volcengine_vpn_connection.foo.id
}

resource "volcengine_transit_router" "foo" {
  transit_router_name = "test-tf-acc"
  description         = "test-tf-acc"
}

resource "volcengine_transit_router_route_table" "foo" {
  description                     = "tf-test-acc-description-route-route-table"
  transit_router_route_table_name = "tf-table-test-acc"
  transit_router_id               = volcengine_transit_router.foo.id
}

resource "volcengine_transit_router_route_entry" "foo" {
  description                              = "tf-test-acc-description-entry"
  transit_router_route_entry_name          = "tf-acc-test-entry"
  destination_cidr_block                   = "192.168.0.0/24"
  transit_router_route_entry_next_hop_type = "Attachment"
  transit_router_route_table_id            = volcengine_transit_router_route_table.foo.transit_router_route_table_id
  transit_router_route_entry_next_hop_id   = volcengine_transit_router_vpn_attachment.foo.transit_router_attachment_id
}
```
## Argument Reference
The following arguments are supported:
* `destination_cidr_block` - (Required, ForceNew) The target network segment of the route entry.
* `transit_router_route_entry_next_hop_type` - (Required, ForceNew) The next hop type of the routing entry. The value can be Attachment or BlackHole.
* `transit_router_route_table_id` - (Required, ForceNew) The id of the route table.
* `description` - (Optional) Description of the transit router route entry.
* `transit_router_route_entry_name` - (Optional) The name of the route entry.
* `transit_router_route_entry_next_hop_id` - (Optional, ForceNew) The next hot id of the routing entry. When the parameter TransitRouterRouteEntryNextHopType is Attachment, this parameter must be filled.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `creation_time` - The creation time of the route entry.
* `status` - The status of the route entry.
* `transit_router_route_entry_id` - The id of the route entry.
* `transit_router_route_entry_type` - The type of the route entry.
* `update_time` - The update time of the route entry.


## Import
transit router route entry can be imported using the table and entry id, e.g.
```
$ terraform import volcengine_transit_router_route_entry.default tr-rtb-12b7qd3fmzf2817q7y2jkbd55:tr-rte-1i5i8khf9m58gae5kcx6***
```

