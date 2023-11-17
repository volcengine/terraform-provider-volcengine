---
subcategory: "ALB"
layout: "volcengine"
page_title: "Volcengine: volcengine_alb_server_group_server"
sidebar_current: "docs-volcengine-resource-alb_server_group_server"
description: |-
  Provides a resource to manage alb server group server
---
# volcengine_alb_server_group_server
Provides a resource to manage alb server group server
## Example Usage
```hcl
resource "volcengine_alb_server_group_server" "foo" {
  server_group_id = "rsp-1g7317vrcx3pc2zbhq4c3i6a2"
  instance_id     = "i-ycony2kef4ygp2f8cgmk"
  type            = "ecs"
  weight          = 30
  ip              = "172.16.0.3"
  port            = 5679
  description     = "test add server group server ecs1"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required, ForceNew) The ID of ecs instance or the network card bound to ecs instance.
* `ip` - (Required, ForceNew) The private ip of the instance.
* `port` - (Required) The port receiving request.
* `server_group_id` - (Required, ForceNew) The ID of the ServerGroup.
* `type` - (Required, ForceNew) The type of instance. Optional choice contains `ecs`, `eni`.
* `description` - (Optional) The description of the instance.
* `weight` - (Optional) The weight of the instance, range in 0~100.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `server_id` - The server id of instance in ServerGroup.


## Import
AlbServerGroupServer can be imported using the server_group_id:server_id, e.g.
```
$ terraform import volcengine_alb_server_group_server.default rsp-274xltv2*****8tlv3q3s:rs-3ciynux6i1x4w****rszh49sj
```

