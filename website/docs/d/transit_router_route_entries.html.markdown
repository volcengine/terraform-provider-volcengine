---
subcategory: "TRANSIT_ROUTER"
layout: "volcengine"
page_title: "Volcengine: volcengine_transit_router_route_entries"
sidebar_current: "docs-volcengine-datasource-transit_router_route_entries"
description: |-
  Use this data source to query detailed information of transit router route entries
---
# volcengine_transit_router_route_entries
Use this data source to query detailed information of transit router route entries
## Example Usage
```hcl
data "volcengine_transit_router_route_entries" "default" {
  transit_router_route_table_id = "tr-rtb-12b7qd3fmzf2817q7y2jkbd55"
  //destination_cidr_block = ""
  //status = ""
  //transit_router_route_entry_name = ""
  ids = ["tr-rte-12b7qd5eo3h1c17q7y1sq5ixv"]
}
```
## Argument Reference
The following arguments are supported:
* `transit_router_route_table_id` - (Required) The id of the route table.
* `destination_cidr_block` - (Optional) The target network segment of the route entry.
* `ids` - (Optional) The ids of the transit router route entry.
* `output_file` - (Optional) File name where to save data source results.
* `status` - (Optional) The status of the route entry.
* `transit_router_route_entry_name` - (Optional) The name of the route entry.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `entries` - The list of route entries.
    * `as_path` - The as path of the route entry.
    * `creation_time` - The creation time of the route entry.
    * `description` - Description of the transit router route entry.
    * `destination_cidr_block` - The target network segment of the route entry.
    * `status` - The status of the route entry.
    * `transit_router_route_entry_id` - The id of the route entry.
    * `transit_router_route_entry_name` - The name of the route entry.
    * `transit_router_route_entry_next_hop_id` - The next hot id of the routing entry.
    * `transit_router_route_entry_next_hop_type` - The next hop type of the routing entry. The value can be Attachment or BlackHole.
    * `transit_router_route_entry_type` - The type of the route entry.
    * `update_time` - The update time of the route entry.
* `total_count` - The total count of data query.


