---
subcategory: "TRANSIT_ROUTER"
layout: "volcengine"
page_title: "Volcengine: volcengine_transit_router_vpc_attachments"
sidebar_current: "docs-volcengine-datasource-transit_router_vpc_attachments"
description: |-
  Use this data source to query detailed information of transit router vpc attachments
---
# volcengine_transit_router_vpc_attachments
Use this data source to query detailed information of transit router vpc attachments
## Example Usage
```hcl
data "volcengine_transit_router_vpc_attachments" "default" {
  transit_router_id             = "tr-2d6fr7f39unsw58ozfe1ow21x"
  transit_router_attachment_ids = ["tr-attach-3rf2xi7ae6y9s5zsk2hm6pibt"]
}
```
## Argument Reference
The following arguments are supported:
* `transit_router_id` - (Required) The id of transit router.
* `output_file` - (Optional) File name where to save data source results.
* `transit_router_attachment_ids` - (Optional) A list of Transit Router Attachment ids.
* `vpc_id` - (Optional) The id of vpc.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `attachments` - The collection of query.
    * `attach_points` - The collection of attach points.
        * `network_interface_id` - The ID of network interface.
        * `subnet_id` - The ID of subnet.
        * `zone_id` - The ID of zone.
    * `creation_time` - The create time.
    * `description` - The description info.
    * `status` - The status of the transit router.
    * `transit_router_attachment_id` - The id of the transit router attachment.
    * `transit_router_attachment_name` - The name of the transit router attachment.
    * `transit_router_id` - The id of the transit router.
    * `update_time` - The update time.
    * `vpc_id` - The ID of vpc.
* `total_count` - The total count of query.


