---
subcategory: "VPN(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_vpn_connection"
sidebar_current: "docs-volcengine-resource-vpn_connection"
description: |-
  Provides a resource to manage vpn connection
---
# volcengine_vpn_connection
Provides a resource to manage vpn connection
## Example Usage
```hcl
resource "volcengine_vpn_connection" "foo" {
  vpn_connection_name   = "tf-test"
  description           = "tf-test"
  vpn_gateway_id        = "vgw-2feq19gnyc9hc59gp68914u6o"
  customer_gateway_id   = "cgw-12ayj1s157gn417q7y29bixqy"
  local_subnet          = ["192.168.0.0/22"]
  remote_subnet         = ["192.161.0.0/20"]
  dpd_action            = "none"
  nat_traversal         = true
  ike_config_psk        = "tftest@!3"
  ike_config_version    = "ikev1"
  ike_config_mode       = "main"
  ike_config_enc_alg    = "aes"
  ike_config_auth_alg   = "md5"
  ike_config_dh_group   = "group2"
  ike_config_lifetime   = 100
  ike_config_local_id   = "tf_test"
  ike_config_remote_id  = "tf_test"
  ipsec_config_enc_alg  = "aes"
  ipsec_config_auth_alg = "sha256"
  ipsec_config_dh_group = "group2"
  ipsec_config_lifetime = 100
  project_name          = "default"
}
```
## Argument Reference
The following arguments are supported:
* `customer_gateway_id` - (Required, ForceNew) The ID of the customer gateway.
* `ike_config_psk` - (Required) The psk of the ike config of the VPN connection.
* `local_subnet` - (Required) The local subnet of the VPN connection.
* `remote_subnet` - (Required) The remote subnet of the VPN connection.
* `vpn_gateway_id` - (Required, ForceNew) The ID of the vpn gateway.
* `attach_type` - (Optional, ForceNew) The attach type of the VPN connection, the value can be VpnGateway or TransitRouter.
* `description` - (Optional) The description of the VPN connection.
* `dpd_action` - (Optional) The dpd action of the VPN connection.
* `ike_config_auth_alg` - (Optional) The auth alg of the ike config of the VPN connection.
* `ike_config_dh_group` - (Optional) The dk group of the ike config of the VPN connection.
* `ike_config_enc_alg` - (Optional) The enc alg of the ike config of the VPN connection.
* `ike_config_lifetime` - (Optional) The lifetime of the ike config of the VPN connection.
* `ike_config_local_id` - (Optional) The local_id of the ike config of the VPN connection.
* `ike_config_mode` - (Optional) The mode of the ike config of the VPN connection.
* `ike_config_remote_id` - (Optional) The remote id of the ike config of the VPN connection.
* `ike_config_version` - (Optional) The version of the ike config of the VPN connection.
* `ipsec_config_auth_alg` - (Optional) The auth alg of the ipsec config of the VPN connection.
* `ipsec_config_dh_group` - (Optional) The dh group of the ipsec config of the VPN connection.
* `ipsec_config_enc_alg` - (Optional) The enc alg of the ipsec config of the VPN connection.
* `ipsec_config_lifetime` - (Optional) The ipsec config of the ike config of the VPN connection.
* `nat_traversal` - (Optional) The nat traversal of the VPN connection.
* `negotiate_instantly` - (Optional) Whether to initiate negotiation mode immediately.
* `project_name` - (Optional) The project name of the VPN connection.
* `vpn_connection_name` - (Optional) The name of the VPN connection.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `account_id` - The account ID of the VPN connection.
* `attach_status` - The IPsec attach status.
* `connect_status` - The connect status of the VPN connection.
* `creation_time` - The create time of VPN connection.
* `status` - The status of the VPN connection.
* `update_time` - The update time of VPN connection.
* `vpn_connection_id` - The ID of the VPN connection.


## Import
VpnConnection can be imported using the id, e.g.
```
$ terraform import volcengine_vpn_connection.default vgc-3tex2x1cwd4c6c0v****
```

