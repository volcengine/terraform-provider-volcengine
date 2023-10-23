---
subcategory: "VPN(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_ssl_vpn_servers"
sidebar_current: "docs-volcengine-datasource-ssl_vpn_servers"
description: |-
  Use this data source to query detailed information of ssl vpn servers
---
# volcengine_ssl_vpn_servers
Use this data source to query detailed information of ssl vpn servers
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
  vpc_id              = volcengine_vpc.foo.id
  subnet_id           = volcengine_subnet.foo.id
  bandwidth           = 5
  vpn_gateway_name    = "acc-test1"
  description         = "acc-test1"
  period              = 7
  project_name        = "default"
  ssl_enabled         = true
  ssl_max_connections = 5
}

resource "volcengine_ssl_vpn_server" "foo" {
  vpn_gateway_id      = volcengine_vpn_gateway.foo.id
  local_subnets       = [volcengine_subnet.foo.cidr_block]
  client_ip_pool      = "172.16.2.0/24"
  ssl_vpn_server_name = "acc-test-ssl"
  description         = "acc-test"
  protocol            = "UDP"
  cipher              = "AES-128-CBC"
  auth                = "SHA1"
  compress            = true
}

data "volcengine_ssl_vpn_servers" "foo" {
  ids = [volcengine_ssl_vpn_server.foo.id]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) The ids list.
* `output_file` - (Optional) File name where to save data source results.
* `ssl_vpn_server_name` - (Optional) The name of the ssl vpn server.
* `vpn_gateway_id` - (Optional) The id of the vpn gateway.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `ssl_vpn_servers` - List of SSL VPN servers.
    * `auth` - The authentication algorithm of the SSL server.
Values:
`SHA1` (default)
`MD5`
`None` (do not use encryption).
    * `cipher` - The encryption algorithm of the SSL server.
Values:
`AES-128-CBC` (default)
`AES-192-CBC`
`AES-256-CBC`
`None` (do not use encryption).
    * `client_ip_pool` - SSL client network segment.
    * `compress` - Whether to compress the transmitted data. The default value is false.
    * `creation_time` - The creation time.
    * `description` - The description of the ssl server.
    * `id` - The SSL VPN server id.
    * `local_subnets` - The local network segment of the SSL server. The local network segment is the address segment that the client accesses through the SSL VPN connection.
    * `protocol` - The protocol used by the SSL server. Valid values are `TCP`, `UDP`. Default Value: `UDP`.
    * `ssl_vpn_server_id` - The id of the ssl vpn server.
    * `ssl_vpn_server_name` - The name of the SSL server.
    * `status` - The status of the ssl vpn server.
    * `update_time` - The update time.
    * `vpn_gateway_id` - The vpn gateway id.
* `total_count` - The total count of SSL VPN server query.


