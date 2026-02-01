---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_instance_parameters"
sidebar_current: "docs-volcengine-datasource-rds_postgresql_instance_parameters"
description: |-
  Use this data source to query detailed information of rds postgresql instance parameters
---
# volcengine_rds_postgresql_instance_parameters
Use this data source to query detailed information of rds postgresql instance parameters
## Example Usage
```hcl
data "volcengine_rds_postgresql_instance_parameters" "example" {
  instance_id    = "postgres-72715e0d9f58"
  parameter_name = "wal_level"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The ID of the PostgreSQL instance.
* `output_file` - (Optional) File name where to save data source results.
* `parameter_name` - (Optional) The name of the parameter, supports fuzzy query. If no value is passed or a null value is passed, all parameters under the specified instance will be queried.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `instance_parameters` - The collection of query.
    * `db_engine_version` - The version of the PostgreSQL engine.
    * `instance_id` - The ID of the PostgreSQL instance.
    * `parameter_count` - The total count of parameters.
    * `parameters` - The current parameter configuration of the instance (kernel parameters).
        * `checking_code` - The value range of the parameter.
        * `default_value` - Parameter default value. Refers to the default value provided in the default template corresponding to this instance.
        * `description_zh` - The description of the parameter in Chinese.
        * `description` - The description of the parameter in English.
        * `force_restart` - Indicates whether a restart is required after the parameter is modified.
        * `name` - The name of the parameter.
        * `type` - The type of the parameter.
        * `value` - The current value of the parameter.
* `total_count` - The total count of query.


