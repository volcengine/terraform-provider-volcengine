---
subcategory: "ECS"
layout: "volcengine"
page_title: "Volcengine: volcengine_regions"
sidebar_current: "docs-volcengine-datasource-regions"
description: |-
  Use this data source to query detailed information of regions
---
# volcengine_regions
Use this data source to query detailed information of regions
## Example Usage
```hcl
data "volcengine_regions" "default" {
  ids = ["cn-beijing"]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of region ids.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `regions` - The collection of region query.
    * `id` - The id of the region.
    * `region_id` - The id of the region.
* `total_count` - The total count of region query.


