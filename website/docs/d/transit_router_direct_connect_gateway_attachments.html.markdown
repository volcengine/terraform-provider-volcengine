---
subcategory: "TRANSIT_ROUTER"
layout: "volcengine"
page_title: "Volcengine: volcengine_transit_router_direct_connect_gateway_attachments"
sidebar_current: "docs-volcengine-datasource-transit_router_direct_connect_gateway_attachments"
description: |-
  Use this data source to query detailed information of transit router direct connect gateway attachments
---
# volcengine_transit_router_direct_connect_gateway_attachments
Use this data source to query detailed information of transit router direct connect gateway attachments
## Example Usage
```hcl
data "volcengine_transit_router_direct_connect_gateway_attachments" "foo" {
  transit_router_id = "tr-2bzy39x27qtxc2dx0eg5qaj05"
}
```
## Argument Reference
The following arguments are supported:
* `transit_router_id` - (Required) The id of the transit router.
* `direct_connect_gateway_id` - (Optional) ID of the direct connection gateway.
* `output_file` - (Optional) File name where to save data source results.
* `transit_router_attachment_ids` - (Optional) ID of the network instance connection.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `attachments` - The collection of query.
    * `account_id` - The account id.
    * `creation_time` - The create time.
    * `description` - The description info.
    * `direct_connect_gateway_id` - The direct connect gateway id.
    * `status` - The status of the network instance connection.
    * `transit_router_attachment_id` - The id of the transit router attachment.
    * `transit_router_attachment_name` - The name of the transit router attachment.
    * `transit_router_id` - The id of the transit router.
    * `update_time` - The update time.
* `total_count` - The total count of query.


