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
resource "volcengine_transit_router_bandwidth_package" "foo" {
  transit_router_bandwidth_package_name = "acc-tf-test"
  description                           = "acc-test"
  bandwidth                             = 2
  period                                = 1
  renew_type                            = "Manual"
  renew_period                          = 1
  remain_renew_times                    = -1
}

resource "volcengine_transit_router" "foo" {
  transit_router_name = "acc-test-tf"
  description         = "acc-test-tf"
}

resource "volcengine_transit_router_peer_attachment" "foo" {
  transit_router_id                   = volcengine_transit_router.foo.id
  transit_router_attachment_name      = "acc-test-tf"
  description                         = "tf-test"
  peer_transit_router_id              = "tr-xxx"
  peer_transit_router_region_id       = "cn-xx"
  transit_router_bandwidth_package_id = volcengine_transit_router_bandwidth_package.foo.id
  bandwidth                           = 2
  tags {
    key   = "k1"
    value = "v1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `peer_transit_router_id` - (Required, ForceNew) The id of the peer transit router.
* `peer_transit_router_region_id` - (Required, ForceNew) The region id of the peer transit router.
* `transit_router_id` - (Required, ForceNew) The id of the local transit router.
* `bandwidth` - (Optional) The bandwidth of the transit router peer attachment. Unit: Mbps.
* `description` - (Optional) The description of the transit router peer attachment.
* `tags` - (Optional) Tags.
* `transit_router_attachment_name` - (Optional) The name of the transit router peer attachment.
* `transit_router_bandwidth_package_id` - (Optional) The bandwidth package id of the transit router peer attachment. When specifying this field, the field `bandwidth` must also be specified.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

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

