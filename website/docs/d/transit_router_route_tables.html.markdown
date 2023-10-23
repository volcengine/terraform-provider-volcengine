---
subcategory: "TRANSIT_ROUTER"
layout: "volcengine"
page_title: "Volcengine: volcengine_transit_router_route_tables"
sidebar_current: "docs-volcengine-datasource-transit_router_route_tables"
description: |-
  Use this data source to query detailed information of transit router route tables
---
# volcengine_transit_router_route_tables
Use this data source to query detailed information of transit router route tables
## Example Usage
```hcl
data "volcengine_transit_router_route_tables" "default" {
  transit_router_id = "tr-2ff4v69tkxji859gp684cm14e"
  ids               = ["tr-rtb-12b7qd3fmzf2817q7y2jkbd55"]
  //transit_router_route_table_type = ""
}
```
## Argument Reference
The following arguments are supported:
* `transit_router_id` - (Required) The id of the transit router.
* `ids` - (Optional) The ids of the transit router route table.
* `output_file` - (Optional) File name where to save data source results.
* `transit_router_route_table_type` - (Optional) The type of the route table. The value can be System or Custom.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `route_tables` - The list of route tables query.
    * `creation_time` - The creation time of the route table.
    * `description` - The description.
    * `status` - The status of the route table.
    * `transit_router_route_table_id` - The id of the route table.
    * `transit_router_route_table_name` - The name of the route table.
    * `transit_router_route_table_type` - The type of route table.
    * `update_time` - The update time of the route table.
* `total_count` - The total count of data query.


