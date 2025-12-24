---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_regions"
sidebar_current: "docs-volcengine-datasource-rds_postgresql_regions"
description: |-
  Use this data source to query detailed information of rds postgresql regions
---
# volcengine_rds_postgresql_regions
Use this data source to query detailed information of rds postgresql regions
## Example Usage
```hcl
data "volcengine_rds_postgresql_regions" "example" {
}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `regions` - The collection of region query.
    * `region_id` - The ID of the region.
    * `region_name` - The name of the region.
* `total_count` - The total count of query.


