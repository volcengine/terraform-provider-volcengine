---
subcategory: "ECS"
layout: "volcengine"
page_title: "Volcengine: volcengine_zones"
sidebar_current: "docs-volcengine-datasource-zones"
description: |-
  Use this data source to query detailed information of zones
---
**❗Notice:**
The current provider is no longer being maintained. We recommend that you use the [volcenginecc](https://registry.terraform.io/providers/volcengine/volcenginecc/latest/docs) instead.
# volcengine_zones
Use this data source to query detailed information of zones
## Example Usage
```hcl
data "volcengine_zones" "default" {
  ids = ["cn-beijing-a"]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of zone ids.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of zone query.
* `zones` - The collection of zone query.
    * `id` - The id of the zone.
    * `zone_id` - The id of the zone.


