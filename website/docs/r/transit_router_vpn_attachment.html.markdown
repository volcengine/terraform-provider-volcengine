---
subcategory: "TRANSIT_ROUTER"
layout: "volcengine"
page_title: "Volcengine: volcengine_transit_router_vpn_attachment"
sidebar_current: "docs-volcengine-resource-transit_router_vpn_attachment"
description: |-
  Provides a resource to manage transit router vpn attachment
---
# volcengine_transit_router_vpn_attachment
Provides a resource to manage transit router vpn attachment
## Example Usage
```hcl
resource "volcengine_transit_router" "foo" {
  transit_router_name = "test-tf-acc"
  description         = "test-tf-acc"
}

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
```
## Argument Reference
The following arguments are supported:
* `transit_router_id` - (Required, ForceNew) The id of the transit router.
* `vpn_connection_id` - (Required, ForceNew) The ID of the IPSec connection.
* `zone_id` - (Required, ForceNew) The ID of the availability zone.
* `description` - (Optional) The description of the transit router vpn attachment.
* `transit_router_attachment_name` - (Optional) The name of the transit router vpn attachment.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `creation_time` - The create time.
* `status` - The status of the transit router.
* `transit_router_attachment_id` - The id of the transit router vpn attachment.
* `update_time` - The update time.


## Import
TransitRouterVpnAttachment can be imported using the transitRouterId:attachmentId, e.g.
```
$ terraform import volcengine_transit_router_vpn_attachment.default tr-2d6fr7mzya2gw58ozfes5g2oh:tr-attach-7qthudw0ll6jmc****
```

