---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_route_entry"
sidebar_current: "docs-volcengine-resource-route_entry"
description: |-
  Provides a resource to manage route entry
---
# volcengine_route_entry
Provides a resource to manage route entry
## Example Usage
```hcl
resource "volcengine_route_entry" "foo" {
  route_table_id         = "vtb-2744hslq5b7r47fap8tjomgnj"
  destination_cidr_block = "0.0.0.0/2"
  next_hop_type          = "NatGW"
  next_hop_id            = "ngw-274gwbqe340zk7fap8spkzo7x"
  route_entry_name       = "tf-test-up"
  description            = "tf-test-up"
}
```
## Argument Reference
The following arguments are supported:
* `destination_cidr_block` - (Required, ForceNew) The destination CIDR block of the route entry.
* `next_hop_id` - (Required, ForceNew) The id of the next hop.
* `next_hop_type` - (Required, ForceNew) The type of the next hop, Optional choice contains `Instance`, `NetworkInterface`, `NatGW`, `VpnGW`.
* `route_table_id` - (Required, ForceNew) The id of the route table.
* `description` - (Optional) The description of the route entry.
* `route_entry_name` - (Optional) The name of the route entry.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `route_entry_id` - The id of the route entry.
* `status` - The description of the route entry.


## Import
Route entry can be imported using the route_table_id:route_entry_id, e.g.
```
$ terraform import volcengine_route_entry.default vtb-274e19skkuhog7fap8u4i8ird:rte-274e1g9ei4k5c7fap8sp974fq
```

