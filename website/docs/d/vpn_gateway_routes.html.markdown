---
subcategory: "VPN"
layout: "volcengine"
page_title: "Volcengine: volcengine_vpn_gateway_routes"
sidebar_current: "docs-volcengine-datasource-vpn_gateway_routes"
description: |-
  Use this data source to query detailed information of vpn gateway routes
---
# volcengine_vpn_gateway_routes
Use this data source to query detailed information of vpn gateway routes
## Example Usage
```hcl
resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = "cn-beijig-a"
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

data "volcengine_vpn_gateway_routes" "foo" {
  ids = [volcengine_vpn_gateway_route.foo.id]
}
```
## Argument Reference
The following arguments are supported:
* `destination_cidr_block` - (Optional) A destination cidr block.
* `ids` - (Optional) A list of VPN gateway route ids.
* `next_hop_id` - (Optional) An ID of next hop.
* `output_file` - (Optional) File name where to save data source results.
* `route_type` - (Optional) The type of the VPN gateway route. Valid values: `Static`, `BGP`, `Cloud`.
* `status` - (Optional) The status of the VPN gateway route.
* `vpn_gateway_id` - (Optional) An ID of VPN gateway.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of VPN gateway route query.
* `vpn_gateway_routes` - The collection of VPN gateway route query.
    * `creation_time` - The create time of VPN gateway route.
    * `destination_cidr_block` - The destination cidr block of the VPN gateway route.
    * `id` - The ID of the VPN gateway route.
    * `next_hop_id` - The next hop id of the VPN gateway route.
    * `status` - The status of the VPN gateway route.
    * `update_time` - The update time of VPN gateway route.
    * `vpn_gateway_id` - The ID of the VPN gateway of the VPN gateway route.
    * `vpn_gateway_route_id` - The ID of the VPN gateway route.


