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
resource "volcengine_transit_router_vpc_attachment" "foo" {
  transit_router_id = "tr-2d6fr7f39unsw58ozfe1ow21x"
  vpc_id            = "vpc-2bysvq1xx543k2dx0eeulpeiv"
  attach_points {
    subnet_id = "subnet-3refsrxdswsn45zsk2hmdg4zx"
    zone_id   = "cn-beijing-a"
  }
  attach_points {
    subnet_id = "subnet-2d68bh74345q858ozfekrm8fj"
    zone_id   = "cn-beijing-a"
  }
  transit_router_attachment_name = "tfname1"
  description                    = "desc"
}
```
## Argument Reference
The following arguments are supported:
* `attach_points` - (Required) The attach points of transit router vpc attachment.
* `transit_router_id` - (Required, ForceNew) The id of the transit router.
* `vpc_id` - (Required, ForceNew) The ID of vpc.
* `description` - (Optional) The description of the transit router vpc attachment.
* `transit_router_attachment_name` - (Optional) The name of the transit router vpc attachment.

The `attach_points` object supports the following:

* `subnet_id` - (Required) The id of subnet.
* `zone_id` - (Required) The id of zone.

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

