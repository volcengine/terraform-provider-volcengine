---
subcategory: "VPN(BETA)"
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
data "volcengine_vpn_gateway_routes" "default" {
  ids = ["vgr-2byssu52dktts2dx0ee90r5hp]"]
}
```
## Argument Reference
The following arguments are supported:
* `destination_cidr_block` - (Optional) A destination cidr block.
* `ids` - (Optional) A list of VPN gateway route ids.
* `next_hop_id` - (Optional) An ID of next hop.
* `output_file` - (Optional) File name where to save data source results.
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


