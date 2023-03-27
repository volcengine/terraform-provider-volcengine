---
subcategory: "MONGODB"
layout: "volcengine"
page_title: "Volcengine: volcengine_mongodb_instance_parameters"
sidebar_current: "docs-volcengine-datasource-mongodb_instance_parameters"
description: |-
  Use this data source to query detailed information of mongodb instance parameters
---
# volcengine_mongodb_instance_parameters
Use this data source to query detailed information of mongodb instance parameters
## Example Usage
```hcl
data "volcengine_mongodb_instance_parameters" "foo" {
  instance_id     = "mongo-replica-f16e9298b121" // 必填
  parameter_role  = "Node"                       // 选填
  parameter_names = "connPoolMaxConnsPerHost"    // 选填
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The instance ID to query.
* `output_file` - (Optional) File name where to save data source results.
* `parameter_names` - (Optional) The parameter names, support fuzzy query, case insensitive.
* `parameter_role` - (Optional) The node type of instance parameter, valid value contains `Node`, `Shard`, `ConfigServer`, `Mongos`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `parameters` - The collection of parameter query.
    * `db_engine_version` - The database engine version.
    * `db_engine` - The database engine.
    * `instance_id` - The instance ID.
    * `instance_parameters` - The list of parameters.
        * `checking_code` - The checking code of parameter.
        * `force_modify` - Whether the parameter supports modifying.
        * `force_restart` - Does the new parameter value need to restart the instance to take effect after modification.
        * `parameter_default_value` - The default value of parameter.
        * `parameter_description` - The description of parameter.
        * `parameter_names` - The name of parameter.
        * `parameter_role` - The node type to which the parameter belongs.
        * `parameter_type` - The type of parameter value.
        * `parameter_value` - The value of parameter.
    * `total` - The total parameters queried.
* `total_count` - The total count of mongodb instance parameter query.


