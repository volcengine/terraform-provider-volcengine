---
subcategory: "TRANSIT_ROUTER"
layout: "volcengine"
page_title: "Volcengine: volcengine_transit_router_route_entry"
sidebar_current: "docs-volcengine-resource-transit_router_route_entry"
description: |-
  Provides a resource to manage transit router route entry
---
# volcengine_transit_router_route_entry
Provides a resource to manage transit router route entry
## Example Usage
```hcl
resource "volcengine_transit_router_route_entry" "foo" {
  transit_router_route_table_id            = "tr-rtb-12b7qd3fmzf2817q7y2jkbd55"
  destination_cidr_block                   = "192.168.0.0/24"
  transit_router_route_entry_next_hop_type = "BlackHole"
  //transit_router_route_entry_next_hop_id = ""
  description                     = "tf test 23"
  transit_router_route_entry_name = "tf-entry-23"
}
```
## Argument Reference
The following arguments are supported:
* `destination_cidr_block` - (Required, ForceNew) The target network segment of the route entry.
* `transit_router_route_entry_next_hop_type` - (Required, ForceNew) The next hop type of the routing entry. The value can be Attachment or BlackHole.
* `transit_router_route_table_id` - (Required, ForceNew) The id of the route table.
* `description` - (Optional) Description of the transit router route entry.
* `transit_router_route_entry_name` - (Optional) The name of the route entry.
* `transit_router_route_entry_next_hop_id` - (Optional, ForceNew) The next hot id of the routing entry. When the parameter TransitRouterRouteEntryNextHopType is Attachment, this parameter must be filled.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `creation_time` - The creation time of the route entry.
* `status` - The status of the route entry.
* `transit_router_route_entry_id` - The id of the route entry.
* `transit_router_route_entry_type` - The type of the route entry.
* `update_time` - The update time of the route entry.


## Import
transit router route entry can be imported using the table and entry id, e.g.
```
$ terraform import volcengine_transit_router_route_entry.default tr-rtb-12b7qd3fmzf2817q7y2jkbd55:tr-rte-1i5i8khf9m58gae5kcx6***
```

