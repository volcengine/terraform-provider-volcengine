---
subcategory: "VPN(BETA)"
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
resource "volcengine_vpn_gateway_route" "foo" {
  vpn_gateway_id         = "vgw-2c012ea9fm5mo2dx0efxg46qi"
  destination_cidr_block = "192.168.0.0/20"
  next_hop_id            = "vgc-2d5ww3ww2lwcg58ozfe61ppc3"
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

