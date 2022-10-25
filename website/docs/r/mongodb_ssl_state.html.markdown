---
subcategory: "MONGODB"
layout: "volcengine"
page_title: "Volcengine: volcengine_mongodb_ssl_state"
sidebar_current: "docs-volcengine-resource-mongodb_ssl_state"
description: |-
  Provides a resource to manage mongodb ssl state
---
# volcengine_mongodb_ssl_state
Provides a resource to manage mongodb ssl state
## Example Usage
```hcl
resource "volcengine_mongodb_ssl_state" "foo" {
  instance_id = "mongo-shard-xxx"
  ssl_action  = "Update"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The ID of mongodb instance.
* `ssl_action` - (Optional) The action of ssl,valid value contains `Update`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
mongosdb ssl state can be imported using the ssl:instanceId, e.g.
set `ssl_action` to `Update` will update ssl always when terraform apply.
```
$ terraform import volcengine_mongosdb_ssl_state.default ssl:mongo-shard-d050db19xxx
```

