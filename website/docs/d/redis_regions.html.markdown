---
subcategory: "REDIS"
layout: "volcengine"
page_title: "Volcengine: volcengine_redis_regions"
sidebar_current: "docs-volcengine-datasource-redis_regions"
description: |-
  Use this data source to query detailed information of redis regions
---
# volcengine_redis_regions
Use this data source to query detailed information of redis regions
## Example Usage
```hcl
data "volcengine_redis_regions" "default" {
  region_id = "cn-north-3"
}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.
* `region_id` - (Optional) Target region info.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `regions` - The collection of region query.
    * `region_id` - The id of the region.
    * `region_name` - The name of region.
* `total_count` - The total count of region query.


