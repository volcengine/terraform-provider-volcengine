---
subcategory: "TRANSIT_ROUTER"
layout: "volcengine"
page_title: "Volcengine: volcengine_transit_router_route_table"
sidebar_current: "docs-volcengine-resource-transit_router_route_table"
description: |-
  Provides a resource to manage transit router route table
---
# volcengine_transit_router_route_table
Provides a resource to manage transit router route table
## Example Usage
```hcl
resource "volcengine_transit_router_route_table" "foo" {
  transit_router_id               = "tr-2ff4v69tkxji859gp684cm14e"
  description                     = "tf test23"
  transit_router_route_table_name = "tf-table-23"
}
```
## Argument Reference
The following arguments are supported:
* `transit_router_id` - (Required, ForceNew) Id of the transit router.
* `description` - (Optional) Description of the transit router route table.
* `transit_router_route_table_name` - (Optional) The name of the route table.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `creation_time` - The creation time of the route table.
* `status` - The status of the route table.
* `transit_router_route_table_id` - The id of the route table.
* `transit_router_route_table_type` - The type of route table.
* `update_time` - The update time of the route table.


## Import
transit router route table can be imported using the router id and route table id, e.g.
```
$ terraform import volcengine_transit_router_route_table.default tr-2ff4v69tkxji859gp684cm14e:tr-rtb-hy13n2l4c6c0v****
```

