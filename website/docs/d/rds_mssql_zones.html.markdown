---
subcategory: "RDS_MSSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_mssql_zones"
sidebar_current: "docs-volcengine-datasource-rds_mssql_zones"
description: |-
  Use this data source to query detailed information of rds mssql zones
---
# volcengine_rds_mssql_zones
Use this data source to query detailed information of rds mssql zones
## Example Usage
```hcl
data "volcengine_rds_mssql_zones" "foo" {

}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.
* `region_id` - (Optional) The Id of Region.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of zone query.
* `zones` - The collection of zone query.
    * `description` - The description of the zone.
    * `id` - The id of the zone.
    * `zone_id` - The id of the zone.
    * `zone_name` - The name of the zone.


