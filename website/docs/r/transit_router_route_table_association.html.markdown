---
subcategory: "TRANSIT_ROUTER"
layout: "volcengine"
page_title: "Volcengine: volcengine_transit_router_route_table_association"
sidebar_current: "docs-volcengine-resource-transit_router_route_table_association"
description: |-
  Provides a resource to manage transit router route table association
---
# volcengine_transit_router_route_table_association
Provides a resource to manage transit router route table association
## Example Usage
```hcl
resource "volcengine_transit_router_route_table_association" "foo" {
  transit_router_attachment_id  = "tr-attach-im73ng3n5kao8gbssz2ddpuq"
  transit_router_route_table_id = "tr-rtb-12b7qd3fmzf2817q7y2jkbd55"
}
```
## Argument Reference
The following arguments are supported:
* `transit_router_attachment_id` - (Required, ForceNew) The ID of the network instance connection.
* `transit_router_route_table_id` - (Required, ForceNew) The ID of the routing table associated with the transit router instance.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
TransitRouterRouteTableAssociation can be imported using the TransitRouterAttachmentId:TransitRouterRouteTableId, e.g.
```
$ terraform import volcengine_transit_router_route_table_association.default tr-attach-13n2l4c****:tr-rt-1i5i8khf9m58gae5kcx6****
```

