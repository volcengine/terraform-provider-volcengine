---
subcategory: "VPN"
layout: "volcengine"
page_title: "Volcengine: volcengine_ssl_vpn_server"
sidebar_current: "docs-volcengine-resource-ssl_vpn_server"
description: |-
  Provides a resource to manage ssl vpn server
---
# volcengine_ssl_vpn_server
Provides a resource to manage ssl vpn server
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
```
## Argument Reference
The following arguments are supported:
* `client_ip_pool` - (Required, ForceNew) SSL client network segment.
* `local_subnets` - (Required, ForceNew) The local network segment of the SSL server. The local network segment is the address segment that the client accesses through the SSL VPN connection.
* `vpn_gateway_id` - (Required, ForceNew) The vpn gateway id.
* `auth` - (Optional, ForceNew) The authentication algorithm of the SSL server.
Values:
`SHA1` (default)
`MD5`
`None` (do not use encryption).
* `cipher` - (Optional, ForceNew) The encryption algorithm of the SSL server.
Values:
`AES-128-CBC` (default)
`AES-192-CBC`
`AES-256-CBC`
`None` (do not use encryption).
* `compress` - (Optional, ForceNew) Whether to compress the transmitted data. The default value is false.
* `description` - (Optional) The description of the ssl server.
* `protocol` - (Optional, ForceNew) The protocol used by the SSL server. Valid values are `TCP`, `UDP`. Default Value: `UDP`.
* `ssl_vpn_server_name` - (Optional) The name of the SSL server.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `ssl_vpn_server_id` - The id of the ssl vpn server.


## Import
SSL VPN server can be imported using the id, e.g.
```
$ terraform import volcengine_ssl_vpn_gateway.default vss-zm55pqtvk17oq32zd****
```

