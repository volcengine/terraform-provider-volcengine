---
subcategory: "VKE"
layout: "volcengine"
page_title: "Volcengine: volcengine_vke_support_resource_types"
sidebar_current: "docs-volcengine-datasource-vke_support_resource_types"
description: |-
  Use this data source to query detailed information of vke support resource types
---
# volcengine_vke_support_resource_types
Use this data source to query detailed information of vke support resource types
## Example Usage
```hcl
data "volcengine_vke_support_resource_types" "default" {
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


