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
resource "volcengine_transit_router" "foo" {
  transit_router_name = "test-tf-acc"
  description         = "test-tf-acc"
}

resource "volcengine_transit_router_route_table" "foo" {
  description                     = "tf-test-acc-description"
  transit_router_route_table_name = "tf-table-test-acc"
  transit_router_id               = volcengine_transit_router.foo.id
  tags {
    key   = "k1"
    value = "v1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `transit_router_id` - (Required, ForceNew) Id of the transit router.
* `description` - (Optional) Description of the transit router route table.
* `tags` - (Optional) Tags.
* `transit_router_route_table_name` - (Optional) The name of the route table.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

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

