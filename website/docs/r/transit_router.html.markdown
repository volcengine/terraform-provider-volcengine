---
subcategory: "TRANSIT_ROUTER"
layout: "volcengine"
page_title: "Volcengine: volcengine_transit_router"
sidebar_current: "docs-volcengine-resource-transit_router"
description: |-
  Provides a resource to manage transit router
---
# volcengine_transit_router
Provides a resource to manage transit router
## Example Usage
```hcl
resource "volcengine_transit_router" "foo" {
  transit_router_name = "acc-test-tr"
  description         = "acc-test"
  asn                 = 4294967294
  project_name        = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `asn` - (Optional, ForceNew) The asn of the transit router. Valid value range in 64512-65534 and 4200000000-4294967294. Default is 64512.
* `description` - (Optional) The description of the transit router.
* `project_name` - (Optional) The ProjectName of the transit router.
* `tags` - (Optional) Tags.
* `transit_router_name` - (Optional) The name of the transit router.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `account_id` - The ID of account.
* `business_status` - The business status of the transit router.
* `creation_time` - The create time.
* `overdue_time` - The overdue time.
* `status` - The status of the transit router.
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
* `update_time` - The update time.


## Import
TransitRouter can be imported using the id, e.g.
```
$ terraform import volcengine_transit_router.default tr-2d6fr7mzya2gw58ozfes5g2oh
```

