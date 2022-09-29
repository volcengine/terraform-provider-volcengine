---
subcategory: "CLB"
layout: "volcengine"
page_title: "Volcengine: volcengine_server_group_server"
sidebar_current: "docs-volcengine-resource-server_group_server"
description: |-
  Provides a resource to manage server group server
---
# volcengine_server_group_server
Provides a resource to manage server group server
## Example Usage
```hcl
resource "volcengine_server_group_server" "foo" {
  server_group_id = "rsp-274xltv2sjoxs7fap8tlv3q3s"
  instance_id     = "i-ybp1scasbe72q1vq35wv"
  type            = "ecs"
  weight          = 100
  port            = 80
  description     = "This is a server"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required, ForceNew) The ID of ecs instance or the network card bound to ecs instance.
* `port` - (Required) The port receiving request.
* `server_group_id` - (Required, ForceNew) The ID of the ServerGroup.
* `type` - (Required, ForceNew) The type of instance. Optional choice contains `ecs`, `eni`.
* `description` - (Optional) The description of the instance.
* `ip` - (Optional, ForceNew) The private ip of the instance.
* `weight` - (Optional) The weight of the instance, range in 0~100.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `server_id` - The server id of instance in ServerGroup.


## Import
ServerGroupServer can be imported using the id, e.g.
```
$ terraform import volcengine_server_group_server.default rsp-274xltv2*****8tlv3q3s:rs-3ciynux6i1x4w****rszh49sj
```

