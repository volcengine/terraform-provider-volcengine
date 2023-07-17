---
subcategory: "CEN(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_cen_service_route_entries"
sidebar_current: "docs-volcengine-datasource-cen_service_route_entries"
description: |-
  Use this data source to query detailed information of cen service route entries
---
# volcengine_cen_service_route_entries
Use this data source to query detailed information of cen service route entries
## Example Usage
```hcl
data "volcengine_cen_service_route_entries" "default" {
  cen_id = "cen-12ar8uclj68sg17q7y20v9gil"
}
```
## Argument Reference
The following arguments are supported:
* `cen_id` - (Optional) A cen ID.
* `destination_cidr_block` - (Optional) A destination cidr block.
* `output_file` - (Optional) File name where to save data source results.
* `service_region_id` - (Optional) A service region id.
* `service_vpc_id` - (Optional) A service VPC id.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `service_route_entries` - The collection of cen service route entry query.
    * `cen_id` - The cen ID of the cen service route entry.
    * `creation_time` - The create time of the cen service route entry.
    * `description` - The description of the cen service route entry.
    * `destination_cidr_block` - The destination cidr block of the cen service route entry.
    * `publish_mode` - Publishing scope of cloud service access routes. Valid values are `LocalDCGW`(default), `Custom`.
    * `publish_to_instances` - The publish instances. A maximum of 100 can be uploaded in one request.
        * `instance_id` - Cloud service access routes need to publish the network instance ID.
        * `instance_region_id` - The region where the cloud service access route needs to be published.
        * `instance_type` - The network instance type that needs to be published for cloud service access routes. The values are as follows: `VPC`, `DCGW`.
    * `service_region_id` - The service region id of the cen service route entry.
    * `service_vpc_id` - The service VPC id of the cen service route entry.
    * `status` - The status of the cen service route entry.
* `total_count` - The total count of cen service route entry.


