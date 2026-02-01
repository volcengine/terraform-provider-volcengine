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
resource "volcengine_transit_router" "foo" {
  transit_router_name = "test-tf-acc"
  description         = "test-tf-acc"
}

resource "volcengine_transit_router_route_table" "foo" {
  description                     = "tf-test-acc-description"
  transit_router_route_table_name = "tf-table-test-acc"
  transit_router_id               = volcengine_transit_router.foo.id
}


data "volcengine_transit_router_route_tables" "default" {
  transit_router_id = volcengine_transit_router.foo.id
  ids               = [volcengine_transit_router_route_table.foo.transit_router_route_table_id]
  //transit_router_route_table_type = ""
}
```
## Argument Reference
The following arguments are supported:
* `transit_router_id` - (Required) The id of the transit router.
* `ids` - (Optional) The ids of the transit router route table.
* `output_file` - (Optional) File name where to save data source results.
* `tags` - (Optional) Tags.
* `transit_router_route_table_type` - (Optional) The type of the route table. The value can be System or Custom.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `route_tables` - The list of route tables query.
    * `creation_time` - The creation time of the route table.
    * `description` - The description.
    * `status` - The status of the route table.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `transit_router_route_table_id` - The id of the route table.
    * `transit_router_route_table_name` - The name of the route table.
    * `transit_router_route_table_type` - The type of route table.
    * `update_time` - The update time of the route table.
* `total_count` - The total count of data query.


