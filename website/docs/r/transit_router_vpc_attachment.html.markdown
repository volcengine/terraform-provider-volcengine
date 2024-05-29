---
subcategory: "TRANSIT_ROUTER"
layout: "volcengine"
page_title: "Volcengine: volcengine_transit_router_vpc_attachment"
sidebar_current: "docs-volcengine-resource-transit_router_vpc_attachment"
description: |-
  Provides a resource to manage transit router vpc attachment
---
# volcengine_transit_router_vpc_attachment
Provides a resource to manage transit router vpc attachment
## Example Usage
```hcl
resource "volcengine_transit_router" "foo" {
  transit_router_name = "test-tf-acc"
  description         = "test-tf-acc"
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
  tags {
    key   = "k1"
    value = "v1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `attach_points` - (Required) The attach points of transit router vpc attachment.
* `transit_router_id` - (Required, ForceNew) The id of the transit router.
* `vpc_id` - (Required, ForceNew) The ID of vpc.
* `description` - (Optional) The description of the transit router vpc attachment.
* `tags` - (Optional) Tags.
* `transit_router_attachment_name` - (Optional) The name of the transit router vpc attachment.

The `attach_points` object supports the following:

* `subnet_id` - (Required) The id of subnet.
* `zone_id` - (Required) The id of zone.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `creation_time` - The create time.
* `status` - The status of the transit router.
* `transit_router_attachment_id` - The id of the transit router attachment.
* `update_time` - The update time.


## Import
TransitRouterVpcAttachment can be imported using the transitRouterId:attachmentId, e.g.
```
$ terraform import volcengine_transit_router_vpc_attachment.default tr-2d6fr7mzya2gw58ozfes5g2oh:tr-attach-7qthudw0ll6jmc****
```

