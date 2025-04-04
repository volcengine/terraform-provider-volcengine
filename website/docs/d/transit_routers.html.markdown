---
subcategory: "TRANSIT_ROUTER"
layout: "volcengine"
page_title: "Volcengine: volcengine_transit_routers"
sidebar_current: "docs-volcengine-datasource-transit_routers"
description: |-
  Use this data source to query detailed information of transit routers
---
# volcengine_transit_routers
Use this data source to query detailed information of transit routers
## Example Usage
```hcl
resource "volcengine_transit_router" "foo" {
  transit_router_name = "test-tf-acc"
  description         = "test-tf-acc"
}

data "volcengine_transit_routers" "default" {
  ids                 = [volcengine_transit_router.foo.id]
  transit_router_name = "test"
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of Transit Router ids.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The ProjectName of the transit router.
* `tags` - (Optional) Tags.
* `transit_router_name` - (Optional) The name info.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of query.
* `transit_routers` - The collection of query.
    * `account_id` - The ID of account.
    * `asn` - The asn of the transit router.
    * `business_status` - The business status of the transit router.
    * `creation_time` - The create time.
    * `description` - The description info.
    * `grant_status` - The grant status of the transit router.
    * `id` - The ID of the transit router.
    * `overdue_time` - The overdue time.
    * `project_name` - The ProjectName of the transit router.
    * `status` - The status of the transit router.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `transit_router_attachments` - The attachments of transit router.
        * `creation_time` - The create time.
        * `resource_id` - The id of resource.
        * `resource_type` - The type of resource.
        * `status` - The status of the transit router.
        * `transit_router_attachment_id` - The id of transit router attachment.
        * `transit_router_attachment_name` - The name of transit router attachment.
        * `transit_router_route_table_id` - The id of transit router route table.
        * `update_time` - The update time.
    * `transit_router_id` - The ID of the transit router.
    * `transit_router_name` - The name of the transit router.
    * `update_time` - The update time.


