---
subcategory: "MONGODB"
layout: "volcengine"
page_title: "Volcengine: volcengine_mongodb_instance_parameter_logs"
sidebar_current: "docs-volcengine-datasource-mongodb_instance_parameter_logs"
description: |-
  Use this data source to query detailed information of mongodb instance parameter logs
---
# volcengine_mongodb_instance_parameter_logs
Use this data source to query detailed information of mongodb instance parameter logs
## Example Usage
```hcl
data "volcengine_mongodb_instance_parameter_logs" "foo" {
  instance_id = "mongo-replica-xxx"
  start_time  = "2022-11-14 00:00Z"
  end_time    = "2022-11-14 18:15Z"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The instance ID to query.
* `end_time` - (Optional) The end time to query.
* `output_file` - (Optional) File name where to save data source results.
* `start_time` - (Optional) The start time to query.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `parameter_change_logs` - The collection of parameter change log query.
    * `modify_time` - The modifying time of parameter.
    * `new_parameter_value` - The new parameter value.
    * `old_parameter_value` - The old parameter value.
    * `parameter_name` - The parameter name.
    * `parameter_role` - The node type to which the parameter belongs.
    * `parameter_status` - The status of parameter change.
* `total_count` - The total count of mongodb instance parameter log query.


