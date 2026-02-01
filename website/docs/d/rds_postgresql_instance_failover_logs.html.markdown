---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_instance_failover_logs"
sidebar_current: "docs-volcengine-datasource-rds_postgresql_instance_failover_logs"
description: |-
  Use this data source to query detailed information of rds postgresql instance failover logs
---
# volcengine_rds_postgresql_instance_failover_logs
Use this data source to query detailed information of rds postgresql instance failover logs
## Example Usage
```hcl
data "volcengine_rds_postgresql_instance_failover_logs" "example" {
  instance_id      = "postgres-72******9f58"
  query_start_time = "2025-12-10T16:00:00Z"
  query_end_time   = "2025-12-12T17:00:00Z"
  limit            = 1000
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The ID of the PostgreSQL instance.
* `limit` - (Optional) The number of records per page. Max: 1000, Min: 1.
* `output_file` - (Optional) File name where to save data source results.
* `query_end_time` - (Optional) The end time of the query. Format: yyyy-MM-ddTHH:mmZ (UTC time).
* `query_start_time` - (Optional) The start time of the query. Format: yyyy-MM-ddTHH:mmZ (UTC time).

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `failover_logs` - The collection of failover logs.
    * `failover_time` - The time when the failover occurred. Format: yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).
    * `failover_type` - The type of the failover, such as User or System.
    * `new_master_node_id` - The node ID of the new master after failover.
    * `old_master_node_id` - The node ID of the old master before failover.
* `total_count` - The total count of query.


