---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_zones"
sidebar_current: "docs-volcengine-datasource-rds_postgresql_zones"
description: |-
  Use this data source to query detailed information of rds postgresql zones
---
# volcengine_rds_postgresql_zones
Use this data source to query detailed information of rds postgresql zones
## Example Usage
```hcl
data "volcengine_rds_postgresql_zones" "example" {
}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.
* `region_id` - (Optional) The region id of the resource.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of query.
* `zones` - The collection of zone query.
    * `description` - The description of the zone.
    * `zone_id` - The id of the zone.
    * `zone_name` - The name of the zone.


