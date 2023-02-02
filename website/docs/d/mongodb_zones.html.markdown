---
subcategory: "MONGODB"
layout: "volcengine"
page_title: "Volcengine: volcengine_mongodb_zones"
sidebar_current: "docs-volcengine-datasource-mongodb_zones"
description: |-
  Use this data source to query detailed information of mongodb zones
---
# volcengine_mongodb_zones
Use this data source to query detailed information of mongodb zones
## Example Usage
```hcl
data "volcengine_mongodb_zones" "default" {
  region_id = "XXX"
}
```
## Argument Reference
The following arguments are supported:
* `region_id` - (Required) The Id of Region.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of zone query.
* `zones` - The collection of zone query.
    * `id` - The id of the zone.
    * `zone_id` - The id of the zone.
    * `zone_name` - The name of the zone.


