---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_engine_version_parameters"
sidebar_current: "docs-volcengine-datasource-rds_postgresql_engine_version_parameters"
description: |-
  Use this data source to query detailed information of rds postgresql engine version parameters
---
# volcengine_rds_postgresql_engine_version_parameters
Use this data source to query detailed information of rds postgresql engine version parameters
## Example Usage
```hcl
data "volcengine_rds_postgresql_engine_version_parameters" "pg12" {
  db_engine         = "PostgreSQL"
  db_engine_version = "PostgreSQL_12"
}
```
## Argument Reference
The following arguments are supported:
* `db_engine_version` - (Required) The database engine version of the RDS PostgreSQL instance. Valid value: PostgreSQL_11, PostgreSQL_12, PostgreSQL_13, PostgreSQL_14, PostgreSQL_15, PostgreSQL_16, PostgreSQL_17.
* `db_engine` - (Required) The type of the parameter template. The value can only be PostgreSQL.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `db_engine_version_parameters` - The collection of query.
    * `db_engine_version` - The database engine version of the RDS PostgreSQL instance.
    * `parameter_count` - The number of parameters that users can set under the specified database engine version.
    * `parameters` - The collection of parameters that users can set under the specified database engine version.
        * `checking_code` - The value range of the parameter.
        * `default_value` - Parameter default value. Refers to the default value provided in the default template corresponding to this instance.
        * `force_restart` - Indicates whether a restart is required after the parameter is modified.
        * `name` - The name of the parameter.
        * `type` - The type of the parameter.
* `total_count` - The total count of query.


