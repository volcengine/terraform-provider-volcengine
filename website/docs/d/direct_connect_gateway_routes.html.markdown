---
subcategory: "DIRECT_CONNECT"
layout: "volcengine"
page_title: "Volcengine: volcengine_direct_connect_gateway_routes"
sidebar_current: "docs-volcengine-datasource-direct_connect_gateway_routes"
description: |-
  Use this data source to query detailed information of direct connect gateway routes
---
# volcengine_direct_connect_gateway_routes
Use this data source to query detailed information of direct connect gateway routes
## Example Usage
```hcl
data "volcengine_direct_connect_gateway_routes" "foo" {
  ids = ["dcr-638ry33wmzggn3gd6gv****", "dcr-20d6tkadi2k8w65sqhgbj****"]
}
```
## Argument Reference
The following arguments are supported:
* `destination_cidr_block` - (Optional) The cidr block.
* `direct_connect_gateway_id` - (Optional) The id of direct connect gateway.
* `ids` - (Optional) A list of IDs.
* `next_hop_id` - (Optional) The id of next hop.
* `next_hop_type` - (Optional) The type of next hop.
* `output_file` - (Optional) File name where to save data source results.
* `route_type` - (Optional) The type of route. The value can be BGP or CEN or Static.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `direct_connect_gateway_routes` - The collection of query.
    * `account_id` - The id of account.
    * `creation_time` - The create time.
    * `destination_cidr_block` - The cidr block.
    * `direct_connect_gateway_id` - The id of direct connect gateway.
    * `direct_connect_gateway_route_id` - The id of direct connect gateway route.
    * `next_hop_id` - The id of next hop.
    * `next_hop_type` - The type of next hop.
    * `route_type` - The type of route.
    * `status` - The status info.
* `total_count` - The total count of query.


