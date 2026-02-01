---
subcategory: "TRANSIT_ROUTER"
layout: "volcengine"
page_title: "Volcengine: volcengine_transit_router_route_table_propagation"
sidebar_current: "docs-volcengine-resource-transit_router_route_table_propagation"
description: |-
  Provides a resource to manage transit router route table propagation
---
# volcengine_transit_router_route_table_propagation
Provides a resource to manage transit router route table propagation
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

data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc-acc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  vpc_id      = volcengine_vpc.foo.id
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  subnet_name = "acc-test-subnet"
}

resource "volcengine_subnet" "foo2" {
  vpc_id      = volcengine_vpc.foo.id
  cidr_block  = "172.16.255.0/24"
  zone_id     = data.volcengine_zones.foo.zones[1].id
  subnet_name = "acc-test-subnet2"
}


resource "volcengine_transit_router_vpc_attachment" "foo" {
  transit_router_id = volcengine_transit_router.foo.id
  vpc_id            = volcengine_vpc.foo.id
  attach_points {
    subnet_id = volcengine_subnet.foo.id
    zone_id   = "cn-beijing-a"
  }
  attach_points {
    subnet_id = volcengine_subnet.foo2.id
    zone_id   = "cn-beijing-b"
  }
  transit_router_attachment_name = "tf-test-acc-name1"
  description                    = "tf-test-acc-description"
}

resource "volcengine_transit_router_route_table_propagation" "foo" {
  transit_router_attachment_id  = volcengine_transit_router_vpc_attachment.foo.transit_router_attachment_id
  transit_router_route_table_id = volcengine_transit_router_route_table.foo.transit_router_route_table_id
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
TransitRouterRouteTablePropagation can be imported using the propagation:TransitRouterAttachmentId:TransitRouterRouteTableId, e.g.
```
$ terraform import volcengine_transit_router_route_table_propagation.default propagation:tr-attach-13n2l4c****:tr-rt-1i5i8khf9m58gae5kcx6****
```

