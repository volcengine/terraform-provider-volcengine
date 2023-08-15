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
```
## Argument Reference
The following arguments are supported:
* `customer_gateway_id` - (Required, ForceNew) The ID of the customer gateway.
* `ike_config_psk` - (Required) The psk of the ike config of the VPN connection. The length does not exceed 100 characters, and only uppercase and lowercase letters, special symbols and numbers are allowed.
* `local_subnet` - (Required) The local subnet of the VPN connection. Up to 5 network segments are supported.
* `remote_subnet` - (Required) The remote subnet of the VPN connection. Up to 5 network segments are supported.
* `attach_type` - (Optional, ForceNew) The attach type of the VPN connection, the value can be `VpnGateway` or `TransitRouter`.
* `description` - (Optional) The description of the VPN connection.
* `dpd_action` - (Optional) The dpd action of the VPN connection.
* `ike_config_auth_alg` - (Optional) The auth alg of the ike config of the VPN connection. Valid value are `sha1`, `md5`, `sha256`, `sha384`, `sha512`, `sm3`. The default value is `sha1`.
* `ike_config_dh_group` - (Optional) The dk group of the ike config of the VPN connection. Valid value are `group1`, `group2`, `group5`, `group14`. The default value is `group2`.
* `ike_config_enc_alg` - (Optional) The enc alg of the ike config of the VPN connection. Valid value are `aes`, `aes192`, `aes256`, `des`, `3des`, `sm4`. The default value is `aes`.
* `ike_config_lifetime` - (Optional) The lifetime of the ike config of the VPN connection. Value: 900~86400.
* `ike_config_local_id` - (Optional) The local_id of the ike config of the VPN connection.
* `ike_config_mode` - (Optional) The mode of the ike config of the VPN connection. Valid values are `main`, `aggressive`, and default value is `main`.
* `ike_config_remote_id` - (Optional) The remote id of the ike config of the VPN connection.
* `ike_config_version` - (Optional) The version of the ike config of the VPN connection. The value can be `ikev1` or `ikev2`. The default value is `ikev1`.
* `ipsec_config_auth_alg` - (Optional) The auth alg of the ipsec config of the VPN connection. Valid value are `sha1`, `md5`, `sha256`, `sha384`, `sha512`, `sm3`. The default value is `sha1`.
* `ipsec_config_dh_group` - (Optional) The dh group of the ipsec config of the VPN connection. Valid value are `group1`, `group2`, `group5`, `group14` and `disable`. The default value is `group2`.
* `ipsec_config_enc_alg` - (Optional) The enc alg of the ipsec config of the VPN connection. Valid value are `aes`, `aes192`, `aes256`, `des`, `3des`, `sm4`. The default value is `aes`.
* `ipsec_config_lifetime` - (Optional) The ipsec config of the ike config of the VPN connection. Value: 900~86400.
* `log_enabled` - (Optional) Whether to enable connection logging. After enabling Connection Day, you can view and download IPsec connection logs, and use the log information to troubleshoot IPsec connection problems yourself.
* `nat_traversal` - (Optional) The nat traversal of the VPN connection.
* `negotiate_instantly` - (Optional) Whether to initiate negotiation mode immediately.
* `project_name` - (Optional) The project name of the VPN connection.
* `vpn_connection_name` - (Optional) The name of the VPN connection.
* `vpn_gateway_id` - (Optional, ForceNew) The ID of the vpn gateway. If the `AttachType` is not passed or the passed value is `VpnGateway`, this parameter must be filled. If the value of `AttachType` is `TransitRouter`, this parameter does not need to be filled.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `account_id` - The account ID of the VPN connection.
* `attach_status` - The IPsec attach status.
* `business_status` - The business status of IPsec connection, valid when the attach type is 'TransitRouter'.
* `connect_status` - The connect status of the VPN connection.
* `creation_time` - The create time of VPN connection.
* `deleted_time` - The delete time of resource, valid when the attach type is 'TransitRouter'.
* `ip_address` - The ip address of transit router, valid when the attach type is 'TransitRouter'.
* `overdue_time` - The overdue time of resource, valid when the attach type is 'TransitRouter'.
* `status` - The status of the VPN connection.
* `transit_router_id` - The id of transit router, valid when the attach type is 'TransitRouter'.
* `update_time` - The update time of VPN connection.
* `vpn_connection_id` - The ID of the VPN connection.
* `zone_id` - The zone id of transit router, valid when the attach type is 'TransitRouter'.


## Import
VpnConnection can be imported using the id, e.g.
```
$ terraform import volcengine_vpn_connection.default vgc-3tex2x1cwd4c6c0v****
```

