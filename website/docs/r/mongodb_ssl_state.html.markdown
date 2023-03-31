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
  instance_id = "mongo-replica-f16e9298b121" // 必填
  ssl_action  = "Update"                     // 选填 仅支持Update 
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required, ForceNew) The ID of mongodb instance.
* `ssl_action` - (Optional) The action of ssl, valid value contains `Update`. Set `ssl_action` to `Update` will update ssl always when terraform apply.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `is_valid` - Whetehr SSL is valid.
* `ssl_enable` - Whether SSL is enabled.
* `ssl_expired_time` - The expire time of SSL.


## Import
mongodb ssl state can be imported using the ssl:instanceId, e.g.
```
$ terraform import volcengine_mongodb_ssl_state.default ssl:mongo-shard-d050db19xxx
```
Set `ssl_action` to `Update` will update ssl always when terraform apply.

