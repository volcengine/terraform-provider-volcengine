---
subcategory: "TRANSIT_ROUTER"
layout: "volcengine"
page_title: "Volcengine: volcengine_transit_router_vpn_attachments"
sidebar_current: "docs-volcengine-datasource-transit_router_vpn_attachments"
description: |-
  Use this data source to query detailed information of transit router vpn attachments
---
# volcengine_transit_router_vpn_attachments
Use this data source to query detailed information of transit router vpn attachments
## Example Usage
```hcl
data "volcengine_transit_router_vpn_attachments" "default" {
  ids               = ["tr-attach-3rf2xi7ae6y9s5zsk2hm6pibt"]
  transit_router_id = "tr-2d6fr7f39unsw58ozfe1ow21x"
}
```
## Argument Reference
The following arguments are supported:
* `transit_router_id` - (Required) The id of the transit router.
* `ids` - (Optional) The ID list of the VPN attachment.
* `output_file` - (Optional) File name where to save data source results.
* `vpn_connection_id` - (Optional) The ID of the IPSec connection.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `attachments` - The collection of query.
    * `creation_time` - The create time.
    * `description` - The description info.
    * `status` - The status of the transit router.
    * `transit_router_attachment_id` - The id of the transit router attachment.
    * `transit_router_attachment_name` - The name of the transit router attachment.
    * `transit_router_id` - The id of the transit router.
    * `update_time` - The update time.
    * `vpn_connection_id` - The ID of the IPSec connection.
    * `zone_id` - The ID of the availability zone.
* `total_count` - The total count of query.


