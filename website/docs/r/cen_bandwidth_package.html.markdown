---
subcategory: "CEN"
layout: "volcengine"
page_title: "Volcengine: volcengine_cen_bandwidth_package"
sidebar_current: "docs-volcengine-resource-cen_bandwidth_package"
description: |-
  Provides a resource to manage cen bandwidth package
---
# volcengine_cen_bandwidth_package
Provides a resource to manage cen bandwidth package
## Notice
When Destroy this resource,If the resource charge type is PrePaid,Please unsubscribe the resource 
in  [Volcengine Console](https://console.volcengine.com/finance/unsubscribe/),when complete console operation,yon can
use 'terraform state rm ${resourceId}' to remove.
## Example Usage
```hcl
resource "volcengine_cen_bandwidth_package" "foo" {
  local_geographic_region_set_id = "China"
  peer_geographic_region_set_id  = "China"
  bandwidth                      = 2
  cen_bandwidth_package_name     = "acc-test-cen-bp"
  description                    = "acc-test"
  billing_type                   = "PrePaid"
  period_unit                    = "Month"
  period                         = 1
  project_name                   = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `bandwidth` - (Optional) The bandwidth of the cen bandwidth package. Value: 2~10000.
* `billing_type` - (Optional, ForceNew) The billing type of the cen bandwidth package. Only support `PrePaid` and default value is `PrePaid`.
* `cen_bandwidth_package_name` - (Optional) The name of the cen bandwidth package.
* `description` - (Optional) The description of the cen bandwidth package.
* `local_geographic_region_set_id` - (Optional, ForceNew) The local geographic region set id of the cen bandwidth package. Valid value: `China`, `Asia`.
* `peer_geographic_region_set_id` - (Optional, ForceNew) The peer geographic region set id of the cen bandwidth package. Valid value: `China`, `Asia`.
* `period_unit` - (Optional) The period unit of the cen bandwidth package. Value: `Month`, `Year`. Default value is `Month`.
* `period` - (Optional) The period of the cen bandwidth package. Default value is 1.
* `project_name` - (Optional) The ProjectName of the cen bandwidth package.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

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

