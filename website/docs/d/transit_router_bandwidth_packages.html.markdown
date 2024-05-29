---
subcategory: "TRANSIT_ROUTER"
layout: "volcengine"
page_title: "Volcengine: volcengine_transit_router_bandwidth_packages"
sidebar_current: "docs-volcengine-datasource-transit_router_bandwidth_packages"
description: |-
  Use this data source to query detailed information of transit router bandwidth packages
---
# volcengine_transit_router_bandwidth_packages
Use this data source to query detailed information of transit router bandwidth packages
## Example Usage
```hcl
resource "volcengine_transit_router_bandwidth_package" "foo" {
  transit_router_bandwidth_package_name = "acc-tf-test"
  description                           = "acc-test"
  bandwidth                             = 2
  period                                = 1
  renew_type                            = "Manual"
}

data "volcengine_transit_router_bandwidth_packages" "foo" {
  ids = [volcengine_transit_router_bandwidth_package.foo.id]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) The ID list of the TransitRouter bandwidth package.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The ProjectName of the TransitRouter bandwidth package.
* `tags` - (Optional) Tags.
* `transit_router_bandwidth_package_name` - (Optional) The name of the TransitRouter bandwidth package.
* `transit_router_peer_attachment_id` - (Optional) The ID of the peer attachment.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `bandwidth_packages` - The collection of query.
    * `account_id` - The account id.
    * `allocations` - The detailed information on cross regional connections associated with bandwidth packets.
        * `allocate_time` - The delete time of the transit router bandwidth package.
        * `delete_time` - The peer region id of the transit router.
        * `local_region_id` - The local region id of the transit router.
        * `transit_router_peer_attachment_id` - The ID of the peer attachment.
    * `bandwidth` - The bandwidth peak of the transit router bandwidth package. Unit: Mbps.
    * `billing_type` - The billing type of the transit router bandwidth package.
    * `business_status` - The business status of the transit router bandwidth package.
    * `creation_time` - The create time of the transit router bandwidth package.
    * `delete_time` - The delete time of the transit router bandwidth package.
    * `description` - The description of the transit router bandwidth package.
    * `expired_time` - The expired time of the transit router bandwidth package.
    * `id` - The id of the transit router bandwidth package.
    * `local_geographic_region_set_id` - The local geographic region set ID.
    * `peer_geographic_region_set_id` - The peer geographic region set ID.
    * `project_name` - The ProjectName of the transit router bandwidth package.
    * `remaining_bandwidth` - The remaining bandwidth of the transit router bandwidth package. Unit: Mbps.
    * `status` - The status of the transit router bandwidth package.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `transit_router_bandwidth_package_id` - The id of the transit router attachment.
    * `transit_router_bandwidth_package_name` - The name of the transit router bandwidth package.
    * `update_time` - The update time of the transit router bandwidth package.
* `total_count` - The total count of query.


