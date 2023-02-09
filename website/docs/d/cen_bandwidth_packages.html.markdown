---
subcategory: "CEN(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_cen_bandwidth_packages"
sidebar_current: "docs-volcengine-datasource-cen_bandwidth_packages"
description: |-
  Use this data source to query detailed information of cen bandwidth packages
---
# volcengine_cen_bandwidth_packages
Use this data source to query detailed information of cen bandwidth packages
## Example Usage
```hcl
data "volcengine_cen_bandwidth_packages" "foo" {
  ids    = ["cbp-2bzeew3s8p79c2dx0eeohej4x"]
  cen_id = "cen-2bzrl3srxsv0g2dx0efyoojn3"
}
```
## Argument Reference
The following arguments are supported:
* `cen_bandwidth_package_names` - (Optional) A list of cen bandwidth package names.
* `cen_id` - (Optional) A cen id.
* `ids` - (Optional) A list of cen bandwidth package IDs.
* `local_geographic_region_set_id` - (Optional) A local geographic region set id.
* `name_regex` - (Optional) A Name Regex of cen bandwidth package.
* `output_file` - (Optional) File name where to save data source results.
* `peer_geographic_region_set_id` - (Optional) A peer geographic region set id.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `bandwidth_packages` - The collection of cen bandwidth package query.
    * `account_id` - The account ID of the cen bandwidth package.
    * `bandwidth` - The bandwidth of the cen bandwidth package.
    * `billing_type` - The billing type of the cen bandwidth package.
    * `business_status` - The business status of the cen bandwidth package.
    * `cen_bandwidth_package_id` - The ID of the cen bandwidth package.
    * `cen_bandwidth_package_name` - The name of the cen bandwidth package.
    * `cen_ids` - The cen IDs of the bandwidth package.
    * `creation_time` - The create time of the cen bandwidth package.
    * `deleted_time` - The deleted time of the cen bandwidth package.
    * `description` - The description of the cen bandwidth package.
    * `expired_time` - The expired time of the cen bandwidth package.
    * `id` - The ID of the cen bandwidth package.
    * `local_geographic_region_set_id` - The local geographic region set id of the cen bandwidth package.
    * `peer_geographic_region_set_id` - The peer geographic region set id of the cen bandwidth package.
    * `remaining_bandwidth` - The remain bandwidth of the cen bandwidth package.
    * `status` - The status of the cen bandwidth package.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `update_time` - The update time of the cen bandwidth package.
* `total_count` - The total count of cen bandwidth package query.


