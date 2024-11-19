---
subcategory: "VPN"
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
}

data "volcengine_vpn_connections" "foo" {
  ids = [volcengine_vpn_connection.foo.id]
}
```
## Argument Reference
The following arguments are supported:
* `attach_status` - (Optional) The attach status of VPN connection.
* `attach_type` - (Optional) The attach type of VPN connection. Valid values: `VpnGateway`, `TransitRouter`.
* `customer_gateway_id` - (Optional) An ID of customer gateway.
* `ids` - (Optional) A list of VPN connection ids.
* `name_regex` - (Optional) A Name Regex of VPN connection.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name of VPN connection.
* `spec` - (Optional) The spec of IPSec connection. Valid values: `default`, `large`.
* `status` - (Optional) The status of IPSec connection. Valid values: `Creating`, `Deleting`, `Pending`, `Available`.
* `transit_router_id` - (Optional) An ID of transit router.
* `vpn_connection_names` - (Optional) A list of VPN connection names.
* `vpn_gateway_id` - (Optional) An ID of VPN gateway.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of VPN connection query.
* `vpn_connections` - The collection of VPN connection query.
    * `account_id` - The account ID of the VPN connection.
    * `attach_status` - The IPsec attach status.
    * `attach_type` - The IPsec attach type.
    * `business_status` - The business status of IPsec connection, valid when the attach type is 'TransitRouter'.
    * `connect_status` - The connect status of the VPN connection.
    * `creation_time` - The create time of VPN connection.
    * `customer_gateway_id` - The ID of the customer gateway.
    * `deleted_time` - The delete time of resource, valid when the attach type is 'TransitRouter'.
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
    * `ip_address` - The ip address of transit router, valid when the attach type is 'TransitRouter'.
    * `ipsec_config_auth_alg` - The auth alg of the ipsec config of the VPN connection.
    * `ipsec_config_dh_group` - The dh group of the ipsec config of the VPN connection.
    * `ipsec_config_enc_alg` - The enc alg of the ipsec config of the VPN connection.
    * `ipsec_config_lifetime` - The lifetime of the ike config of the VPN connection.
    * `local_subnet` - The local subnet of the VPN connection.
    * `log_enabled` - Whether to enable the connection log.
    * `nat_traversal` - The nat traversal of the VPN connection.
    * `negotiate_instantly` - Whether to initiate negotiation mode immediately.
    * `overdue_time` - The overdue time of resource, valid when the attach type is 'TransitRouter'.
    * `remote_subnet` - The remote subnet of the VPN connection.
    * `status` - The status of the VPN connection.
    * `transit_router_id` - The id of transit router, valid when the attach type is 'TransitRouter'.
    * `update_time` - The update time of VPN connection.
    * `vpn_connection_id` - The ID of the VPN connection.
    * `vpn_connection_name` - The name of the VPN connection.
    * `vpn_gateway_id` - The ID of the vpn gateway.
    * `zone_id` - The zone id of transit router, valid when the attach type is 'TransitRouter'.


