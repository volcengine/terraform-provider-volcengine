---
subcategory: "REDIS"
layout: "volcengine"
page_title: "Volcengine: volcengine_redis_zones"
sidebar_current: "docs-volcengine-datasource-redis_zones"
description: |-
  Use this data source to query detailed information of redis zones
---
**❗Notice:**
The current provider is no longer being maintained. We recommend that you use the [volcenginecc](https://registry.terraform.io/providers/volcengine/volcenginecc/latest/docs) instead.
# volcengine_redis_zones
Use this data source to query detailed information of redis zones
## Example Usage
```hcl
data "volcengine_redis_zones" "default" {
  region_id = "cn-north-3"
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
    * `zone_status` - The status of the zone.


