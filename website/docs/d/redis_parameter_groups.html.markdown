---
subcategory: "REDIS"
layout: "volcengine"
page_title: "Volcengine: volcengine_redis_parameter_groups"
sidebar_current: "docs-volcengine-datasource-redis_parameter_groups"
description: |-
  Use this data source to query detailed information of redis parameter groups
---
# volcengine_redis_parameter_groups
Use this data source to query detailed information of redis parameter groups
## Example Usage
```hcl

```
## Argument Reference
The following arguments are supported:
* `engine_version` - (Optional) The Redis database version applicable to the parameter template.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `source` - (Optional) The source of creating the parameter template.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `parameter_groups` - The details of the parameter template.
    * `create_time` - The creation time of the parameter template, in the format of yyyy-MM-ddTHH:mm:ssZ (UTC).
    * `default` - Whether it is the default parameter template.
    * `description` - The description the parameter template.
    * `engine_version` - The database version applicable to the parameter template.
    * `name` - The name of the parameter template.
    * `parameter_group_id` - The ID of the parameter template.
    * `parameter_num` - The number of parameters contained in the parameter template.
    * `parameters` - The list of parameter information contained in the parameter template.
        * `current_value` - The current running value of the parameter.
        * `description` - The description the parameter.
        * `need_reboot` - Whether to restart the instance to take effect after modifying this parameter.
        * `options` - The optional list of selector type parameters.
            * `description` - The description the Optional parameters.
            * `value` - Optional selector type parameters.
        * `param_name` - The name of parameter.
        * `range` - The value range of numerical type parameters.
        * `type` - The type of the parameter.
        * `unit` - The unit of the numerical type parameter.
    * `source` - The source of creating the parameter template.
    * `update_time` - The last update time of the parameter template, in the format of yyyy-MM-ddTHH:mm:ssZ (UTC).
* `total_count` - The total count of query.


