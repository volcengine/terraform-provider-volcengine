---
subcategory: "REDIS"
layout: "volcengine"
page_title: "Volcengine: volcengine_redis_parameter_group"
sidebar_current: "docs-volcengine-resource-redis_parameter_group"
description: |-
  Provides a resource to manage redis parameter group
---
# volcengine_redis_parameter_group
Provides a resource to manage redis parameter group
## Example Usage
```hcl
resource "volcengine_redis_parameter_group" "foo" {
  name           = "tf-test"
  engine_version = "5.0"
  description    = "tf-test-description"
  param_values {
    name  = "active-defrag-cycle-max"
    value = "30"
  }
  param_values {
    name  = "active-defrag-cycle-min"
    value = "15"
  }
}
```
## Argument Reference
The following arguments are supported:
* `engine_version` - (Required) The Redis database version adapted to the parameter template. The value range is as follows; 7.0: Redis 7.0. 6.0: Redis 6.0. 5.0: Redis 5.0.
* `name` - (Required) Parameter template name. The name needs to meet the following requirements simultaneously: It cannot start with a number or a hyphen (-). Only Chinese characters, letters, numbers, underscores (_) and hyphens (-) can be included. The length should be 2 to 64 characters.
* `param_values` - (Required) The list of parameter information that needs to be included in the new parameter template. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `description` - (Optional) The remarks information of the parameter template should not exceed 200 characters in length.

The `param_values` object supports the following:

* `name` - (Required) The parameter names that need to be included in the parameter template.
* `value` - (Required) The parameter values set for the corresponding parameters.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
ParameterGroup can be imported using the id, e.g.
```
$ terraform import volcengine_parameter_group.default resource_id
```

