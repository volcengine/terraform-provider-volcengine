---
subcategory: "VPN"
layout: "volcengine"
page_title: "Volcengine: volcengine_vpn_gateway_route"
sidebar_current: "docs-volcengine-resource-vpn_gateway_route"
description: |-
  Provides a resource to manage vpn gateway route
---
# volcengine_vpn_gateway_route
Provides a resource to manage vpn gateway route
## Example Usage
```hcl
resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = "cn-beijing-a"
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_vpn_gateway" "foo" {
  vpc_id           = volcengine_vpc.foo.id
  subnet_id        = volcengine_subnet.foo.id
  bandwidth        = 20
  vpn_gateway_name = "acc-test"
  description      = "acc-test"
  period           = 2
  project_name     = "default"
}

resource "volcengine_customer_gateway" "foo" {
  ip_address            = "192.0.1.3"
  customer_gateway_name = "acc-test"
  description           = "acc-test"
  project_name          = "default"
}

resource "volcengine_vpn_connection" "foo" {
  vpn_connection_name   = "acc-tf-test"
  description           = "acc-tf-test"
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
  project_name          = "default"
  log_enabled           = false
}

resource "volcengine_vpn_gateway_route" "foo" {
  vpn_gateway_id         = volcengine_vpn_gateway.foo.id
  destination_cidr_block = "192.168.0.0/20"
  next_hop_id            = volcengine_vpn_connection.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `destination_cidr_block` - (Required, ForceNew) The destination cidr block of the VPN gateway route.
* `next_hop_id` - (Required, ForceNew) The next hop id of the VPN gateway route.
* `vpn_gateway_id` - (Required, ForceNew) The ID of the VPN gateway of the VPN gateway route.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `creation_time` - The create time of VPN gateway route.
* `status` - The status of the VPN gateway route.
* `update_time` - The update time of VPN gateway route.
* `vpn_gateway_route_id` - The ID of the VPN gateway route.


## Import
VpnGatewayRoute can be imported using the id, e.g.
```
$ terraform import volcengine_vpn_gateway_route.default vgr-3tex2c6c0v844c****
```

