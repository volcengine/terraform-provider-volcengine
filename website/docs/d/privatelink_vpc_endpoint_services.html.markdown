---
subcategory: "PRIVATELINK"
layout: "volcengine"
page_title: "Volcengine: volcengine_privatelink_vpc_endpoint_services"
sidebar_current: "docs-volcengine-datasource-privatelink_vpc_endpoint_services"
description: |-
  Use this data source to query detailed information of privatelink vpc endpoint services
---
# volcengine_privatelink_vpc_endpoint_services
Use this data source to query detailed information of privatelink vpc endpoint services
## Example Usage
```hcl
data "volcengine_privatelink_vpc_endpoint_services" "default" {
  ids = ["epsvc-3rel73uf2ewao5zsk2j2l58ro", "epsvc-2d72mxjgq02yo58ozfe5tndeh"]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) The IDs of vpc endpoint service.
* `name_regex` - (Optional) A Name Regex of vpc endpoint service.
* `output_file` - (Optional) File name where to save data source results.
* `service_name` - (Optional) The name of vpc endpoint service.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `services` - The collection of query.
    * `auto_accept_enabled` - Whether auto accept node connect.
    * `creation_time` - The create time of service.
    * `description` - The description of service.
    * `id` - The Id of service.
    * `resources` - The resources info.
        * `resource_id` - The id of resource.
        * `resource_type` - The type of resource.
        * `zone_id` - The zone id of resource.
    * `service_domain` - The domain of service.
    * `service_id` - The Id of service.
    * `service_name` - The name of service.
    * `service_resource_type` - The resource type of service.
    * `service_type` - The type of service.
    * `status` - The status of service.
    * `update_time` - The update time of service.
    * `zone_ids` - The IDs of zones.
* `total_count` - Returns the total amount of the data list.


