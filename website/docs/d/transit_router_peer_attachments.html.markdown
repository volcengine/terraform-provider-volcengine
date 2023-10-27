---
subcategory: "TRANSIT_ROUTER"
layout: "volcengine"
page_title: "Volcengine: volcengine_transit_router_peer_attachments"
sidebar_current: "docs-volcengine-datasource-transit_router_peer_attachments"
description: |-
  Use this data source to query detailed information of transit router peer attachments
---
# volcengine_transit_router_peer_attachments
Use this data source to query detailed information of transit router peer attachments
## Example Usage
```hcl
data "volcengine_transit_router_peer_attachments" "foo" {
  ids = ["tr-attach-12be67d0yh2io17q7y1au****"]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of IDs.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `peer_transit_router_id` - (Optional) The id of peer transit router.
* `peer_transit_router_region_id` - (Optional) The region id of peer transit router.
* `transit_router_attachment_name` - (Optional) The name of transit router peer attachment.
* `transit_router_id` - (Optional) The id of local transit router.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of query.
* `transit_router_attachments` - The collection of query.
    * `bandwidth` - The bandwidth of the transit router peer attachment.
    * `creation_time` - The creation time of the transit router peer attachment.
    * `description` - The description of the transit router peer attachment.
    * `id` - The id of the transit router peer attachment.
    * `peer_transit_router_id` - The id of the peer transit router.
    * `peer_transit_router_region_id` - The region id of the peer transit router.
    * `status` - The status of the transit router peer attachment.
    * `transit_router_attachment_id` - The id of the transit router peer attachment.
    * `transit_router_attachment_name` - The name of the transit router peer attachment.
    * `transit_router_bandwidth_package_id` - The bandwidth package id of the transit router peer attachment.
    * `transit_router_id` - The id of the local transit router.
    * `transit_router_route_table_id` - The route table id of the transit router peer attachment.
    * `update_time` - The update time of the transit router peer attachment.


