---
subcategory: "CEN(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_cen_route_entries"
sidebar_current: "docs-volcengine-datasource-cen_route_entries"
description: |-
  Use this data source to query detailed information of cen route entries
---
# volcengine_cen_route_entries
Use this data source to query detailed information of cen route entries
## Example Usage
```hcl
data "volcengine_cen_route_entries" "foo" {
  cen_id = "cen-12ar8uclj68sg17q7y20v9gil"
}
```
## Argument Reference
The following arguments are supported:
* `cen_id` - (Required) A cen ID.
* `destination_cidr_block` - (Optional) A destination cidr block.
* `instance_id` - (Optional) An instance ID.
* `instance_region_id` - (Optional) An instance region ID.
* `instance_type` - (Optional) An instance type.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `cen_route_entries` - The collection of cen route entry query.
    * `as_path` - The AS path of the cen route entry.
    * `cen_id` - The cen ID of the cen route entry.
    * `destination_cidr_block` - The destination cidr block of the cen route entry.
    * `instance_id` - The instance id of the next hop of the cen route entry.
    * `instance_region_id` - The instance region id of the next hop of the cen route entry.
    * `instance_type` - The instance type of the next hop of the cen route entry.
    * `publish_status` - The publish status of the cen route entry.
    * `status` - The status of the cen route entry.
* `total_count` - The total count of cen route entry.


