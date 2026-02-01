---
subcategory: "DIRECT_CONNECT"
layout: "volcengine"
page_title: "Volcengine: volcengine_direct_connect_gateways"
sidebar_current: "docs-volcengine-datasource-direct_connect_gateways"
description: |-
  Use this data source to query detailed information of direct connect gateways
---
# volcengine_direct_connect_gateways
Use this data source to query detailed information of direct connect gateways
## Example Usage
```hcl
data "volcengine_direct_connect_gateways" "foo" {
  direct_connect_gateway_name = "tf-test"
}
```
## Argument Reference
The following arguments are supported:
* `cen_id` - (Optional) The CEN ID which direct connect gateway belongs.
* `direct_connect_gateway_name` - (Optional) The direst connect gateway name.
* `ids` - (Optional) A list of IDs.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `tag_filters` - (Optional) The filter tag of direct connect.

The `tag_filters` object supports the following:

* `key` - (Optional) The tag key of cloud resource instance.
* `value` - (Optional) The tag value of cloud resource instance.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `direct_connect_gateways` - The collection of query.
    * `account_id` - The account ID that direct connect gateway belongs.
    * `associate_cens` - The CEN information associated with the direct connect gateway.
        * `cen_id` - The cen ID.
        * `cen_owner_id` - The CEN owner's ID.
        * `cen_status` - The CEN status.
    * `business_status` - The business status of direct connect gateway.
    * `creation_time` - The creation time of direct connect gateway.
    * `deleted_time` - The expected resource force collection time. Only when the resource is frozen due to arrears, this parameter will have a return value, otherwise it will return a null value.
    * `description` - The description of direct connect gateway.
    * `direct_connect_gateway_id` - The direct connect gateway ID.
    * `direct_connect_gateway_name` - The direct connect gateway name.
    * `lock_reason` - The reason of the direct connect gateway locked.
    * `overdue_time` - The resource freeze time. Only when the resource is frozen due to arrears, this parameter will have a return value, otherwise it will return a null value.
    * `status` - The status of direct connect gateway.
    * `tags` - The tags that direct connect gateway added.
        * `key` - The tag key.
        * `value` - The tag value.
    * `update_time` - The update time of direct connect gateway.
* `total_count` - The total count of query.


