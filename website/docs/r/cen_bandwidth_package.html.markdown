---
subcategory: "CEN(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_cen_bandwidth_package"
sidebar_current: "docs-volcengine-resource-cen_bandwidth_package"
description: |-
  Provides a resource to manage cen bandwidth package
---
# volcengine_cen_bandwidth_package
Provides a resource to manage cen bandwidth package
## Example Usage
```hcl
resource "volcengine_cen_bandwidth_package" "foo" {
  local_geographic_region_set_id = "China"
  peer_geographic_region_set_id  = "China"
  bandwidth                      = 32
  cen_bandwidth_package_name     = "tf-test"
  description                    = "tf-test1"
  billing_type                   = "PrePaid"
  period_unit                    = "Year"
  period                         = 1
}
```
## Argument Reference
The following arguments are supported:
* `bandwidth` - (Optional) The bandwidth of the cen bandwidth package.
* `billing_type` - (Optional, ForceNew) The billing type of the cen bandwidth package. Terraform will only remove the PrePaid cen bandwidth package from the state file, not actually remove.
* `cen_bandwidth_package_name` - (Optional) The name of the cen bandwidth package.
* `description` - (Optional) The description of the cen bandwidth package.
* `local_geographic_region_set_id` - (Optional, ForceNew) The local geographic region set id of the cen bandwidth package.
* `peer_geographic_region_set_id` - (Optional, ForceNew) The peer geographic region set id of the cen bandwidth package.
* `period_unit` - (Optional) The period unit of the cen bandwidth package.
* `period` - (Optional) The period of the cen bandwidth package.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `account_id` - The account ID of the cen bandwidth package.
* `business_status` - The business status of the cen bandwidth package.
* `cen_bandwidth_package_id` - The ID of the cen bandwidth package.
* `cen_ids` - The cen IDs of the bandwidth package.
* `creation_time` - The create time of the cen bandwidth package.
* `deleted_time` - The deleted time of the cen bandwidth package.
* `expired_time` - The expired time of the cen bandwidth package.
* `remaining_bandwidth` - The remain bandwidth of the cen bandwidth package.
* `status` - The status of the cen bandwidth package.
* `update_time` - The update time of the cen bandwidth package.


## Import
CenBandwidthPackage can be imported using the id, e.g.
```
$ terraform import volcengine_cen_bandwidth_package.default cbp-4c2zaavbvh5f42****
```

