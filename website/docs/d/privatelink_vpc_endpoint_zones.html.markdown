---
subcategory: "PRIVATELINK"
layout: "volcengine"
page_title: "Volcengine: volcengine_privatelink_vpc_endpoint_zones"
sidebar_current: "docs-volcengine-datasource-privatelink_vpc_endpoint_zones"
description: |-
  Use this data source to query detailed information of privatelink vpc endpoint zones
---
# volcengine_privatelink_vpc_endpoint_zones
Use this data source to query detailed information of privatelink vpc endpoint zones
## Example Usage
```hcl
data "volcengine_privatelink_vpc_endpoint_zones" "default" {
  endpoint_id = "ep-2byz5npiuu1hc2dx0efkv****"
}
```
## Argument Reference
The following arguments are supported:
* `endpoint_id` - (Optional) The endpoint id of query.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - Returns the total amount of the data list.
* `vpc_endpoint_zones` - The collection of query.
    * `id` - The Id of vpc endpoint zone.
    * `network_interface_id` - The network interface id of vpc endpoint.
    * `network_interface_ip` - The network interface ip of vpc endpoint.
    * `service_status` - The status of vpc endpoint service.
    * `subnet_id` - The subnet id of vpc endpoint.
    * `zone_domain` - The domain of vpc endpoint zone.
    * `zone_id` - The Id of vpc endpoint zone.
    * `zone_status` - The status of vpc endpoint zone.


