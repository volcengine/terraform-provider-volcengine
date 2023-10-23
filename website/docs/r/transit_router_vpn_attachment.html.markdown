---
subcategory: "TRANSIT_ROUTER"
layout: "volcengine"
page_title: "Volcengine: volcengine_transit_router_vpn_attachment"
sidebar_current: "docs-volcengine-resource-transit_router_vpn_attachment"
description: |-
  Provides a resource to manage transit router vpn attachment
---
# volcengine_transit_router_vpn_attachment
Provides a resource to manage transit router vpn attachment
## Example Usage
```hcl
resource "volcengine_transit_router_vpn_attachment" "foo" {
  transit_router_id              = "tr-2d6frp10q687458ozfep4****"
  vpn_connection_id              = "vgc-3reidwjf1t1c05zsk2hik****"
  zone_id                        = "cn-beijing-a"
  transit_router_attachment_name = "tf-test"
  description                    = "desc"
}
```
## Argument Reference
The following arguments are supported:
* `transit_router_id` - (Required, ForceNew) The id of the transit router.
* `vpn_connection_id` - (Required, ForceNew) The ID of the IPSec connection.
* `zone_id` - (Required, ForceNew) The ID of the availability zone.
* `description` - (Optional) The description of the transit router vpn attachment.
* `transit_router_attachment_name` - (Optional) The name of the transit router vpn attachment.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `creation_time` - The create time.
* `status` - The status of the transit router.
* `transit_router_attachment_id` - The id of the transit router vpn attachment.
* `update_time` - The update time.


## Import
TransitRouterVpnAttachment can be imported using the transitRouterId:attachmentId, e.g.
```
$ terraform import volcengine_transit_router_vpn_attachment.default tr-2d6fr7mzya2gw58ozfes5g2oh:tr-attach-7qthudw0ll6jmc****
```

