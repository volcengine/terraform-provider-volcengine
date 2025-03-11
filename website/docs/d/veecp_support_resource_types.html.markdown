---
subcategory: "VEECP"
layout: "volcengine"
page_title: "Volcengine: volcengine_veecp_support_resource_types"
sidebar_current: "docs-volcengine-datasource-veecp_support_resource_types"
description: |-
  Use this data source to query detailed information of veecp support resource types
---
# volcengine_veecp_support_resource_types
Use this data source to query detailed information of veecp support resource types
## Example Usage
```hcl
data "volcengine_veecp_support_resource_types" "foo" {
  resource_types = []
  zone_ids       = []
}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.
* `resource_types` - (Optional) A list of resource types. Support Ecs or Zone.
* `zone_ids` - (Optional) A list of zone ids. If no parameter value, all available regions is returned.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `resources` - The collection of query.
    * `resource_scope` - The scope of resource.
    * `resource_specifications` - The resource specifications info.
    * `resource_type` - The type of resource.
    * `zone_id` - The id of zone.
* `total_count` - The total count of query.


