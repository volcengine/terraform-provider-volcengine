---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_instance_recoverable_times"
sidebar_current: "docs-volcengine-datasource-rds_postgresql_instance_recoverable_times"
description: |-
  Use this data source to query detailed information of rds postgresql instance recoverable times
---
# volcengine_rds_postgresql_instance_recoverable_times
Use this data source to query detailed information of rds postgresql instance recoverable times
## Example Usage
```hcl
data "volcengine_rds_postgresql_instance_recoverable_times" "example" {
  instance_id = "postgres-72715e0d9f58"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The id of the Postgresql instance.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `recoverable_time_info` - The earliest and latest recoverable times of the instance (UTC time). If it is empty, it indicates that the instance is currently unrecoverable.
    * `earliest_recoverable_time` - The earliest recoverable time of the instance (UTC time).
    * `latest_recoverable_time` - The latest recoverable time of the instance (UTC time).
* `total_count` - The total count of query.


