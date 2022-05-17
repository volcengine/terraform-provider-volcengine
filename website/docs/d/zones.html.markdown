---
subcategory: "ECS"
layout: "vestack"
page_title: "Vestack: vestack_zones"
sidebar_current: "docs-vestack-datasource-zones"
description: |-
  Use this data source to query detailed information of zones
---
# vestack_zones
Use this data source to query detailed information of zones
## Example Usage
```hcl
data "vestack_zones" "default" {
  ids = ["cn-lingqiu-a"]
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


