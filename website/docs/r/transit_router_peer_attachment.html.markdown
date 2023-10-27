---
subcategory: "TRANSIT_ROUTER"
layout: "volcengine"
page_title: "Volcengine: volcengine_transit_router_peer_attachment"
sidebar_current: "docs-volcengine-resource-transit_router_peer_attachment"
description: |-
  Provides a resource to manage transit router peer attachment
---
# volcengine_transit_router_peer_attachment
Provides a resource to manage transit router peer attachment
## Example Usage
```hcl
resource "volcengine_transit_router_peer_attachment" "foo" {
  transit_router_id                   = "tr-12bbdsa6ode6817q7y1f5****"
  transit_router_attachment_name      = "tf-test-tra"
  description                         = "tf-test"
  peer_transit_router_id              = "tr-3jgsfiktn0feo3pncmfb5****"
  peer_transit_router_region_id       = "cn-beijing"
  transit_router_bandwidth_package_id = "tbp-cd-2felfww0i6pkw59gp68bq****"
  bandwidth                           = 2
}
```
## Argument Reference
The following arguments are supported:
* `peer_transit_router_id` - (Required, ForceNew) The id of the peer transit router.
* `peer_transit_router_region_id` - (Required, ForceNew) The region id of the peer transit router.
* `transit_router_id` - (Required, ForceNew) The id of the local transit router.
* `bandwidth` - (Optional) The bandwidth of the transit router peer attachment. Unit: Mbps.
* `description` - (Optional) The description of the transit router peer attachment.
* `transit_router_attachment_name` - (Optional) The name of the transit router peer attachment.
* `transit_router_bandwidth_package_id` - (Optional) The bandwidth package id of the transit router peer attachment. When specifying this field, the field `bandwidth` must also be specified.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `creation_time` - The creation time of the transit router peer attachment.
* `status` - The status of the transit router peer attachment.
* `transit_router_route_table_id` - The route table id of the transit router peer attachment.
* `update_time` - The update time of the transit router peer attachment.


## Import
TransitRouterPeerAttachment can be imported using the id, e.g.
```
$ terraform import volcengine_transit_router_peer_attachment.default tr-attach-12be67d0yh2io17q7y1au****
```

