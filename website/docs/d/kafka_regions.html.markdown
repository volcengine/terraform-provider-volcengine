---
subcategory: "KAFKA"
layout: "volcengine"
page_title: "Volcengine: volcengine_kafka_regions"
sidebar_current: "docs-volcengine-datasource-kafka_regions"
description: |-
  Use this data source to query detailed information of kafka regions
---
**❗Notice:**
The current provider is no longer being maintained. We recommend that you use the [volcenginecc](https://registry.terraform.io/providers/volcengine/volcenginecc/latest/docs) instead.
# volcengine_kafka_regions
Use this data source to query detailed information of kafka regions
## Example Usage
```hcl
data "volcengine_kafka_regions" "default" {
}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `regions` - The collection of region query.
    * `description` - The description of region.
    * `region_id` - The id of the region.
    * `region_name` - The name of region.
    * `status` - The status of region.
* `total_count` - The total count of region query.


