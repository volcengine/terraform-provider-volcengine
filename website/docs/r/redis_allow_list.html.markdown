---
subcategory: "REDIS"
layout: "volcengine"
page_title: "Volcengine: volcengine_redis_allow_list"
sidebar_current: "docs-volcengine-resource-redis_allow_list"
description: |-
  Provides a resource to manage redis allow list
---
# volcengine_redis_allow_list
Provides a resource to manage redis allow list
## Example Usage
```hcl
resource "volcengine_redis_allow_list" "foo" {
  allow_list_name = "rx_test_tf_allowlist_create"
  allow_list      = ["0.0.0.0/0", "192.168.0.0/24", "192.168.1.1", "192.168.2.22"]
  allow_list_desc = "renxin terraform测试白xxxxxxx"
}
```
## Argument Reference
The following arguments are supported:
* `allow_list_name` - (Required) Name of allow list.
* `allow_list` - (Required) Ip list of allow list.
* `allow_list_desc` - (Optional) Description of allow list.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `allow_list_id` - Id of allow list.
* `allow_list_ip_num` - The IP number of allow list.
* `allow_list_type` - Type of allow list.
* `associated_instance_num` - The number of instance that associated to allow list.
* `associated_instances` - Instances associated by this allow list.
    * `instance_id` - Id of instance.
    * `instance_name` - Name of instance.
    * `vpc` - Id of virtual private cloud.


## Import
Redis AllowList can be imported using the id, e.g.
```
$ terraform import volcengine_redis_allow_list.default acl-cn03wk541s55c376xxxx
```

