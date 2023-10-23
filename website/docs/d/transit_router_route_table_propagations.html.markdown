---
subcategory: "TRANSIT_ROUTER"
layout: "volcengine"
page_title: "Volcengine: volcengine_transit_router_route_table_propagations"
sidebar_current: "docs-volcengine-datasource-transit_router_route_table_propagations"
description: |-
  Use this data source to query detailed information of transit router route table propagations
---
# volcengine_transit_router_route_table_propagations
Use this data source to query detailed information of transit router route table propagations
## Example Usage
```hcl
data "volcengine_transit_router_route_table_propagations" "default" {
  transit_router_attachment_id  = "tr-attach-im73ng3n5kao8gbssz2ddpuq"
  transit_router_route_table_id = "tr-rtb-12b7qd3fmzf2817q7y2jkbd55"
}
```
## Argument Reference
The following arguments are supported:
* `transit_router_route_table_id` - (Required) The ID of the routing table associated with the transit router instance.
* `output_file` - (Optional) File name where to save data source results.
* `transit_router_attachment_id` - (Optional) The ID of the network instance connection.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `propagations` - The list of route table propagations.
    * `creation_time` - The creation time of the route table propagation.
    * `status` - The status of the route table.
    * `transit_router_attachment_id` - The ID of the network instance connection.
    * `transit_router_route_table_id` - The ID of the routing table associated with the transit router instance.
* `total_count` - The total count of data query.


