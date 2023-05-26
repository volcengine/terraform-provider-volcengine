---
subcategory: "REDIS"
layout: "volcengine"
page_title: "Volcengine: volcengine_redis_endpoint"
sidebar_current: "docs-volcengine-resource-redis_endpoint"
description: |-
  Provides a resource to manage redis endpoint
---
# volcengine_redis_endpoint
Provides a resource to manage redis endpoint
## Example Usage
```hcl
resource "volcengine_redis_endpoint" "foo" {
  instance_id = "redis-cn03bb67g3tr2****"
  eip_id      = "eip-274ho3mtx543k7fap8tyi****"
}
```
## Argument Reference
The following arguments are supported:
* `eip_id` - (Required, ForceNew) Id of eip.
* `instance_id` - (Required, ForceNew) Id of instance.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Redis Endpoint can be imported using the instanceId:eipId, e.g.
```
$ terraform import volcengine_redis_endpoint.default redis-asdljioeixxxx:eip-2fef2qcfbfw8w5oxruw3w****
```

