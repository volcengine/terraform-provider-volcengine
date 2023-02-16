---
subcategory: "MONGODB"
layout: "volcengine"
page_title: "Volcengine: volcengine_mongodb_instance_parameter"
sidebar_current: "docs-volcengine-resource-mongodb_instance_parameter"
description: |-
  Provides a resource to manage mongodb instance parameter
---
# volcengine_mongodb_instance_parameter
Provides a resource to manage mongodb instance parameter
## Example Usage
```hcl
resource "volcengine_mongodb_instance_parameter" "foo" {
  instance_id     = "mongo-replica-b2xxx"
  parameter_name  = "connPoolMaxConnsPerHost"
  parameter_role  = "Node"
  parameter_value = "800"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The instance ID.
* `parameter_role` - (Required) The node type to which the parameter belongs.
* `parameter_value` - (Required) The value of parameter.
* `parameter_name` - (Optional) The name of parameter.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
mongosdb parameter can be imported using the param:instanceId:parameterName, e.g.
```
$ terraform import volcengine_mongodb_instance_parameter.default param:mongo-replica-e405f8e2****:connPoolMaxConnsPerHost
```

