---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_database_endpoints"
sidebar_current: "docs-volcengine-datasource-rds_postgresql_database_endpoints"
description: |-
  Use this data source to query detailed information of rds postgresql database endpoints
---
# volcengine_rds_postgresql_database_endpoints
Use this data source to query detailed information of rds postgresql database endpoints
## Example Usage
```hcl
data "volcengine_rds_postgresql_database_endpoints" "example" {
  instance_id = "postgres-72715e0d9f58"
  name_regex  = "默认.*"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Optional) The ID of the RDS PostgreSQL instance.
* `name_regex` - (Optional) The name of the endpoint to filter.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `endpoints` - The collection of query.
    * `cross_region_domain` - Cross-region domain for private address.
    * `dns_visibility` - Whether to enable public network resolution.
    * `domain` - Connect domain name.
    * `endpoint_id` - The ID of the RDS PostgreSQL database endpoint.
    * `endpoint_name` - The name of the RDS PostgreSQL database endpoint.
    * `endpoint_type` - The type of the RDS PostgreSQL database endpoint. Valid values: `Custom`(custom endpoint), `Cluster`(default endpoint).
    * `port` - The endpoint port.
    * `read_only_node_distribution_type` - The distribution type of the read-only nodes.
    * `read_only_node_max_delay_time` - ReadOnly node max delay seconds.
    * `read_write_mode` - ReadWrite or ReadOnly. Default value is ReadOnly.
    * `read_write_proxy_connection` - The number of proxy connections set for the terminal.
    * `write_node_halt_writing` - Whether the endpoint sends write requests to the write node.
* `total_count` - The total count of query.


