---
subcategory: "RDS_MSSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_mssql_regions"
sidebar_current: "docs-volcengine-datasource-rds_mssql_regions"
description: |-
  Use this data source to query detailed information of rds mssql regions
---
# volcengine_rds_mssql_regions
Use this data source to query detailed information of rds mssql regions
## Example Usage
```hcl
data "volcengine_rds_mssql_regions" "foo" {
}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `regions` - The collection of region query.
    * `region_id` - The id of the region.
    * `region_name` - The name of region.
* `total_count` - The total count of region query.


