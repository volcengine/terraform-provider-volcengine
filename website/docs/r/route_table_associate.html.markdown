---
subcategory: "VPC"
layout: "vestack"
page_title: "Vestack: vestack_route_table_associate"
sidebar_current: "docs-vestack-resource-route_table_associate"
description: |-
  Provides a resource to manage route table associate
---
# vestack_route_table_associate
Provides a resource to manage route table associate
## Example Usage
```hcl
resource "vestack_route_table_associate" "foo" {
  route_table_id = "vtb-274e19skkuhog7fap8u4i8ird"
  subnet_id      = "subnet-2744ht7fhjthc7fap8tm10eqg"
}
```
## Argument Reference
The following arguments are supported:
* `route_table_id` - (Required, ForceNew) The id of the route table.
* `subnet_id` - (Required, ForceNew) The id of the subnet.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Route table associate address can be imported using the route_table_id:subnet_id, e.g.
```
$ terraform import vestack_route_table_associate.default vtb-2fdzao4h726f45******:subnet-2fdzaou4liw3k5oxruv******
```

