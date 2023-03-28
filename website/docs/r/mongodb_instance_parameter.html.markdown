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
/*
    该资源无法创建，需先import资源
    $ terraform import volcengine_mongodb_instance_parameter.default param:mongo-replica-f16e9298b121:connPoolMaxConnsPerHost
    请注意instance_id和parameter_name需与上述import的ID对应
*/
resource "volcengine_mongodb_instance_parameter" "default" {
  instance_id     = "mongo-replica-f16e9298b121" // 必填 import之后不允许修改
  parameter_name  = "connPoolMaxConnsPerHost"    // 必填 import之后不允许修改
  parameter_role  = "Node"                       // 必填
  parameter_value = "600"                        // 必填

}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The instance ID. This field cannot be modified after the resource is imported.
* `parameter_name` - (Required) The name of parameter. This field cannot be modified after the resource is imported.
* `parameter_role` - (Required) The node type to which the parameter belongs. The value range is as follows: Node, Shard, ConfigServer, Mongos.
* `parameter_value` - (Required) The value of parameter.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
mongodb parameter can be imported using the param:instanceId:parameterName, e.g.
```
$ terraform import volcengine_mongodb_instance_parameter.default param:mongo-replica-e405f8e2****:connPoolMaxConnsPerHost
```
Note: This resource must be imported before it can be used.
Please note that instance_id and parameter_name must correspond to the ID of the above import.

