---
subcategory: "TRANSIT_ROUTER"
layout: "volcengine"
page_title: "Volcengine: volcengine_transit_router_bandwidth_package"
sidebar_current: "docs-volcengine-resource-transit_router_bandwidth_package"
description: |-
  Provides a resource to manage transit router bandwidth package
---
# volcengine_transit_router_bandwidth_package
Provides a resource to manage transit router bandwidth package
## Example Usage
```hcl
resource "volcengine_transit_router_bandwidth_package" "foo" {
  transit_router_bandwidth_package_name = "acc-tf-test"
  description                           = "acc-test"
  bandwidth                             = 2
  period                                = 1
  renew_type                            = "Manual"
}
```
## Argument Reference
The following arguments are supported:
* `bandwidth` - (Optional) The bandwidth peak of the transit router bandwidth package. Unit: Mbps. Valid values: 2-10000. Default is 2 Mbps.
* `description` - (Optional) The description of the transit router bandwidth package.
* `period` - (Optional) The period of the transit router bandwidth package, the valid value range in 1~9 or 12 or 36. Default value is 12. The period unit defaults to `Month`.The modification of this field only takes effect when the value of the `renew_type` is `Manual`.
* `remain_renew_times` - (Optional) The remaining renewal times of of the transit router bandwidth package. Valid values: -1 or 1~100. Default value is -1, means unlimited renewal.This field is only effective when the value of the `renew_type` is `Auto`.
* `renew_period` - (Optional) The auto renewal period of the transit router bandwidth package. Valid values: 1,2,3,6,12. Default value is 1. Unit: Month.This field is only effective when the value of the `renew_type` is `Auto`. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `renew_type` - (Optional) The renewal type of the transit router bandwidth package. Valid values: `Manual`, `Auto`, `NoRenew`. Default is `Manual`.This field is only effective when modifying the bandwidth package.
* `transit_router_bandwidth_package_name` - (Optional) The name of the transit router bandwidth package.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `allocations` - The detailed information on cross regional connections associated with bandwidth packets.
    * `allocate_time` - The delete time of the transit router bandwidth package.
    * `delete_time` - The peer region id of the transit router.
    * `local_region_id` - The local region id of the transit router.
    * `transit_router_peer_attachment_id` - The ID of the peer attachment.
* `business_status` - The business status of the transit router bandwidth package.
* `creation_time` - The create time of the transit router bandwidth package.
* `delete_time` - The delete time of the transit router bandwidth package.
* `expired_time` - The expired time of the transit router bandwidth package.
* `remaining_bandwidth` - The remaining bandwidth of the transit router bandwidth package. Unit: Mbps.
* `status` - The status of the transit router bandwidth package.
* `update_time` - The update time of the transit router bandwidth package.


## Import
TransitRouterBandwidthPackage can be imported using the Id, e.g.
```
$ terraform import volcengine_transit_router_bandwidth_package.default tbp-cd-2felfww0i6pkw59gp68bq****
```

