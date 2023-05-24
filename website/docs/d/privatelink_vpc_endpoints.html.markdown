---
subcategory: "PRIVATELINK"
layout: "volcengine"
page_title: "Volcengine: volcengine_privatelink_vpc_endpoints"
sidebar_current: "docs-volcengine-datasource-privatelink_vpc_endpoints"
description: |-
  Use this data source to query detailed information of privatelink vpc endpoints
---
# volcengine_privatelink_vpc_endpoints
Use this data source to query detailed information of privatelink vpc endpoints
## Example Usage
```hcl
data "volcengine_privatelink_vpc_endpoints" "default" {
  ids = ["ep-3rel74u229dz45zsk2i6l****"]
}
```
## Argument Reference
The following arguments are supported:
* `endpoint_name` - (Optional) The name of vpc endpoint.
* `ids` - (Optional) The IDs of vpc endpoint.
* `name_regex` - (Optional) A Name Regex of vpc endpoint.
* `output_file` - (Optional) File name where to save data source results.
* `service_name` - (Optional) The name of vpc endpoint service.
* `status` - (Optional) The status of vpc endpoint. Valid values: `Creating`, `Pending`, `Available`, `Deleting`, `Inactive`.
* `vpc_id` - (Optional) The vpc id of vpc endpoint.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - Returns the total amount of the data list.
* `vpc_endpoints` - The collection of query.
    * `business_status` - Whether the vpc endpoint is locked.
    * `connection_status` - The connection  status of vpc endpoint.
    * `creation_time` - The create time of vpc endpoint.
    * `deleted_time` - The delete time of vpc endpoint.
    * `description` - The description of vpc endpoint.
    * `endpoint_domain` - The domain of vpc endpoint.
    * `endpoint_id` - The Id of vpc endpoint.
    * `endpoint_name` - The name of vpc endpoint.
    * `endpoint_type` - The type of vpc endpoint.
    * `id` - The Id of vpc endpoint.
    * `service_id` - The Id of vpc endpoint service.
    * `service_name` - The name of vpc endpoint service.
    * `status` - The status of vpc endpoint.
    * `update_time` - The update time of vpc endpoint.
    * `vpc_id` - The vpc id of vpc endpoint.


