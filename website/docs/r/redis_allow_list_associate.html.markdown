---
subcategory: "REDIS"
layout: "volcengine"
page_title: "Volcengine: volcengine_redis_allow_list_associate"
sidebar_current: "docs-volcengine-resource-redis_allow_list_associate"
description: |-
  Provides a resource to manage redis allow list associate
---
# volcengine_redis_allow_list_associate
Provides a resource to manage redis allow list associate
## Example Usage
```hcl
resource "volcengine_redis_allow_list_associate" "default" {
  instance_id   = "redis-cnlfbzifs7bpvundz"
  allow_list_id = "acl-cnlfc5zsxscu1gg2ajh"
}

resource "volcengine_redis_allow_list_associate" "default1" {
  instance_id   = "redis-cnlfbzifs7bpvundz"
  allow_list_id = "acl-cnlff2mb31zo85p5am7"
}
```
## Argument Reference
The following arguments are supported:
* `allow_list_id` - (Required, ForceNew) Id of allow list to associate.
* `instance_id` - (Required, ForceNew) Id of instance to associate.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Redis AllowList Association can be imported using the instanceId:allowListId, e.g.
```
$ terraform import volcengine_redis_allow_list_associate.default redis-asdljioeixxxx:acl-cn03wk541s55c376xxxx
```

