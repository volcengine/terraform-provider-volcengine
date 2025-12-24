---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_instance_parameter_logs"
sidebar_current: "docs-volcengine-datasource-rds_postgresql_instance_parameter_logs"
description: |-
  Use this data source to query detailed information of rds postgresql instance parameter logs
---
# volcengine_rds_postgresql_instance_parameter_logs
Use this data source to query detailed information of rds postgresql instance parameter logs
## Example Usage
```hcl
data "volcengine_rds_postgresql_instance_parameter_logs" "example" {
  instance_id = "postgres-72715e0d9f58"
  start_time  = "2025-12-01T00:00:00.000Z"
  end_time    = "2025-12-15T23:59:59.999Z"
}
```
## Argument Reference
The following arguments are supported:
* `end_time` - (Required) The end time of the query. Format: yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).
* `instance_id` - (Required) The ID of the PostgreSQL instance.
* `start_time` - (Required) The start time of the query. Format: yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `parameter_change_logs` - The collection of parameter change logs.
    * `modify_time` - The time when the parameter was last modified. Format: yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).
    * `name` - The name of the parameter.
    * `new_value` - The new value of the parameter.
    * `old_value` - The old value of the parameter.
    * `status` - The status of the parameter. Applied: Already in effect. Invalid: Not in effect. Syncing: Being applied, not yet in effect.
* `total_count` - The total count of query.


