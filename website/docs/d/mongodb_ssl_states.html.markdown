---
subcategory: "MONGODB"
layout: "volcengine"
page_title: "Volcengine: volcengine_mongodb_ssl_states"
sidebar_current: "docs-volcengine-datasource-mongodb_ssl_states"
description: |-
  Use this data source to query detailed information of mongodb ssl states
---
# volcengine_mongodb_ssl_states
Use this data source to query detailed information of mongodb ssl states
## Example Usage
```hcl
data "volcengine_mongodb_ssl_states" "foo" {
  instance_id = "mongo-shard-xxx"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Optional) The mongodb instance ID to query.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `ssl_state` - The collection of mongodb ssl state query.
    * `instance_id` - The mongodb instance id.
    * `is_valid` - Whetehr SSL is valid.
    * `ssl_enable` - Whether SSL is enabled.
    * `ssl_expired_time` - The expire time of SSL.


