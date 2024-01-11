---
subcategory: "TRANSIT_ROUTER"
layout: "volcengine"
page_title: "Volcengine: volcengine_transit_router_direct_connect_gateway_attachment"
sidebar_current: "docs-volcengine-resource-transit_router_direct_connect_gateway_attachment"
description: |-
  Provides a resource to manage transit router direct connect gateway attachment
---
# volcengine_transit_router_direct_connect_gateway_attachment
Provides a resource to manage transit router direct connect gateway attachment
## Example Usage
```hcl
resource "volcengine_transit_router" "foo" {
  transit_router_name = "acc-test-tf-acc"
  description         = "acc-test-tf-acc"
}

resource "volcengine_direct_connect_gateway" "foo" {
  direct_connect_gateway_name = "acc-test-gateway-acc"
  description                 = "acc-test-acc"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_transit_router_direct_connect_gateway_attachment" "foo" {
  description                    = "acc-test-tf"
  transit_router_attachment_name = "acc-test-tf"
  transit_router_id              = volcengine_transit_router.foo.id
  direct_connect_gateway_id      = volcengine_direct_connect_gateway.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `direct_connect_gateway_id` - (Required, ForceNew) The id of the direct connect gateway.
* `transit_router_id` - (Required, ForceNew) The id of the transit router.
* `description` - (Optional) The description.
* `transit_router_attachment_name` - (Optional) The name of the transit router direct connect gateway attachment.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `transit_router_attachment_id` - The id of the transit router direct connect gateway attachment.


## Import
TransitRouterDirectConnectGatewayAttachment can be imported using the transitRouterId:attachmentId, e.g.
```
$ terraform import volcengine_transit_router_direct_connect_gateway_attachment.default tr-2d6fr7mzya2gw58ozfes5g2oh:tr-attach-7qthudw0ll6jmc****
```

