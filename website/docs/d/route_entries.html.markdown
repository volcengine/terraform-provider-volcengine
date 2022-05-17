---
subcategory: "VPC"
layout: "vestack"
page_title: "Vestack: vestack_route_entries"
sidebar_current: "docs-vestack-datasource-route_entries"
description: |-
  Use this data source to query detailed information of route entries
---
# vestack_route_entries
Use this data source to query detailed information of route entries
## Example Usage
```hcl
data "vestack_route_entries" "default" {
  ids            = []
  route_table_id = "vtb-274e19skkuhog7fap8u4i8ird"
}
```
## Argument Reference
The following arguments are supported:
* `route_table_id` - (Required) An id of route table.
* `destination_cidr_block` - (Optional) A destination CIDR block of route entry.
* `ids` - (Optional) A list of route entry ids.
* `next_hop_id` - (Optional) An id of next hop.
* `next_hop_type` - (Optional) A type of next hop.
* `output_file` - (Optional) File name where to save data source results.
* `route_entry_name` - (Optional) A name of route entry.
* `route_entry_type` - (Optional) A type of route entry.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `route_entries` - The collection of route tables.
  * `description` - The description of the route entry.
  * `destination_cidr_block` - The destination CIDR block of the route entry.
  * `id` - The id of the route entry.
  * `next_hop_id` - The id of the next hop.
  * `next_hop_name` - The name of the next hop.
  * `next_hop_type` - The type of the next hop.
  * `route_entry_id` - The id of the route entry.
  * `route_entry_name` - The name of the route entry.
  * `route_table_id` - The id of the route table to which the route entry belongs.
  * `status` - The status of the route entry.
  * `type` - The type of the route entry.
  * `vpc_id` - The id of the virtual private cloud (VPC) to which the route entry belongs.
* `total_count` - The total count of route entry query.


