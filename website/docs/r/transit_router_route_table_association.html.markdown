---
subcategory: "TRANSIT_ROUTER"
layout: "volcengine"
page_title: "Volcengine: volcengine_transit_router_route_table_association"
sidebar_current: "docs-volcengine-resource-transit_router_route_table_association"
description: |-
  Provides a resource to manage transit router route table association
---
# volcengine_transit_router_route_table_association
Provides a resource to manage transit router route table association
## Example Usage
```hcl
resource "volcengine_transit_router" "foo" {
  transit_router_name = "test-tf-acc"
  description         = "test-tf-acc"
}

resource "volcengine_transit_router_route_table" "foo" {
  description                     = "tf-test-acc-description"
  transit_router_route_table_name = "tf-table-test-acc"
  transit_router_id               = volcengine_transit_router.foo.id
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

resource "volcengine_transit_router_route_table_association" "foo" {
  transit_router_attachment_id  = volcengine_transit_router_vpn_attachment.foo.transit_router_attachment_id
  transit_router_route_table_id = volcengine_transit_router_route_table.foo.transit_router_route_table_id
}
```
## Argument Reference
The following arguments are supported:
* `transit_router_attachment_id` - (Required, ForceNew) The ID of the network instance connection.
* `transit_router_route_table_id` - (Required, ForceNew) The ID of the routing table associated with the transit router instance.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
TransitRouterRouteTableAssociation can be imported using the TransitRouterAttachmentId:TransitRouterRouteTableId, e.g.
```
$ terraform import volcengine_transit_router_route_table_association.default tr-attach-13n2l4c****:tr-rt-1i5i8khf9m58gae5kcx6****
```

