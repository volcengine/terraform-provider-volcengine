---
subcategory: "VPN"
layout: "volcengine"
page_title: "Volcengine: volcengine_ssl_vpn_client_cert"
sidebar_current: "docs-volcengine-resource-ssl_vpn_client_cert"
description: |-
  Provides a resource to manage ssl vpn client cert
---
# volcengine_ssl_vpn_client_cert
Provides a resource to manage ssl vpn client cert
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

resource "volcengine_ssl_vpn_client_cert" "foo" {
  ssl_vpn_server_id        = volcengine_ssl_vpn_server.foo.id
  ssl_vpn_client_cert_name = "acc-test-client-cert"
  description              = "acc-test"
}
```
## Argument Reference
The following arguments are supported:
* `ssl_vpn_server_id` - (Required, ForceNew) The id of the ssl vpn server.
* `description` - (Optional) The description of the ssl vpn client cert.
* `ssl_vpn_client_cert_name` - (Optional) The name of the ssl vpn client cert.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `ca_certificate` - The CA certificate.
* `certificate_status` - The status of the ssl vpn client cert.
* `client_certificate` - The client certificate.
* `client_key` - The key of the ssl vpn client.
* `creation_time` - The creation time of the ssl vpn client cert.
* `expired_time` - The expired time of the ssl vpn client cert.
* `open_vpn_client_config` - The config of the open vpn client.
* `status` - The status of the ssl vpn client.
* `update_time` - The update time of the ssl vpn client cert.


## Import
SSL VPN Client Cert can be imported using the id, e.g.
```
$ terraform import volcengine_ssl_vpn_client_cert.default vsc-2d6b7gjrzc2yo58ozfcx2****
```

