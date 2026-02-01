---
subcategory: "DIRECT_CONNECT"
layout: "volcengine"
page_title: "Volcengine: volcengine_direct_connect_gateway_route"
sidebar_current: "docs-volcengine-resource-direct_connect_gateway_route"
description: |-
  Provides a resource to manage direct connect gateway route
---
# volcengine_direct_connect_gateway_route
Provides a resource to manage direct connect gateway route
## Example Usage
```hcl
resource "volcengine_direct_connect_gateway_route" "foo" {
  direct_connect_gateway_id = "dcg-172frxs5utjb44d1w33op****"
  destination_cidr_block    = "192.168.40.0/24"
  next_hop_id               = "dcv-1729lrbfx7fuo4d1w34pk****"
}
```
## Argument Reference
The following arguments are supported:
* `destination_cidr_block` - (Required, ForceNew) The cidr block.
* `direct_connect_gateway_id` - (Required, ForceNew) The id of direct connect gateway.
* `next_hop_id` - (Required, ForceNew) The id of next hop.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `account_id` - The id of account.
* `creation_time` - The create time.
* `direct_connect_gateway_route_id` - The id of direct connect gateway route.
* `next_hop_type` - The type of next hop.
* `route_type` - The type of route.
* `status` - The status info.


## Import
DirectConnectGatewayRoute can be imported using the id, e.g.
```
$ terraform import volcengine_direct_connect_gateway_route.default resource_id
```

