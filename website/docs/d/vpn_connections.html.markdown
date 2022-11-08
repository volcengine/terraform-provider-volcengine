---
subcategory: "VPN(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_vpn_connections"
sidebar_current: "docs-volcengine-datasource-vpn_connections"
description: |-
  Use this data source to query detailed information of vpn connections
---
# volcengine_vpn_connections
Use this data source to query detailed information of vpn connections
## Example Usage
```hcl
data "volcengine_vpn_connections" "default" {
  ids = ["vgc-2d5wwids8cdts58ozfe63k2uq"]
}
```
## Argument Reference
The following arguments are supported:
* `customer_gateway_id` - (Optional) An ID of customer gateway.
* `ids` - (Optional) A list of VPN connection ids.
* `name_regex` - (Optional) A Name Regex of VPN connection.
* `output_file` - (Optional) File name where to save data source results.
* `vpn_connection_names` - (Optional) A list of VPN connection names.
* `vpn_gateway_id` - (Optional) An ID of VPN gateway.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of VPN connection query.
* `vpn_connections` - The collection of VPN connection query.
    * `account_id` - The account ID of the VPN connection.
    * `connect_status` - The connect status of the VPN connection.
    * `creation_time` - The create time of VPN connection.
    * `customer_gateway_id` - The ID of the customer gateway.
    * `description` - The description of the VPN connection.
    * `dpd_action` - The dpd action of the VPN connection.
    * `id` - The ID of the VPN connection.
    * `ike_config_auth_alg` - The auth alg of the ike config of the VPN connection.
    * `ike_config_dh_group` - The dk group of the ike config of the VPN connection.
    * `ike_config_enc_alg` - The enc alg of the ike config of the VPN connection.
    * `ike_config_lifetime` - The lifetime of the ike config of the VPN connection.
    * `ike_config_local_id` - The local_id of the ike config of the VPN connection.
    * `ike_config_mode` - The mode of the ike config of the VPN connection.
    * `ike_config_psk` - The psk of the ike config of the VPN connection.
    * `ike_config_remote_id` - The remote id of the ike config of the VPN connection.
    * `ike_config_version` - The version of the ike config of the VPN connection.
    * `ipsec_config_auth_alg` - The auth alg of the ipsec config of the VPN connection.
    * `ipsec_config_dh_group` - The dh group of the ipsec config of the VPN connection.
    * `ipsec_config_enc_alg` - The enc alg of the ipsec config of the VPN connection.
    * `ipsec_config_lifetime` - The lifetime of the ike config of the VPN connection.
    * `local_subnet` - The local subnet of the VPN connection.
    * `nat_traversal` - The nat traversal of the VPN connection.
    * `remote_subnet` - The remote subnet of the VPN connection.
    * `status` - The status of the VPN connection.
    * `update_time` - The update time of VPN connection.
    * `vpn_connection_id` - The ID of the VPN connection.
    * `vpn_connection_name` - The name of the VPN connection.
    * `vpn_gateway_id` - The ID of the vpn gateway.


