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
  instance_id = "mongo-replica-f16e9298b121" // 必填
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The mongodb instance ID to query.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `ssl_state` - The collection of mongodb ssl state query.
    * `instance_id` - The mongodb instance id.
    * `is_valid` - Whetehr SSL is valid.
    * `ssl_enable` - Whether SSL is enabled.
    * `ssl_expired_time` - The expire time of SSL.
* `total_count` - The total count of mongodb ssl state query.


