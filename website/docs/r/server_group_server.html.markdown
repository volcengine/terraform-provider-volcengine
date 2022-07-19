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
  server_group_id = "rsp-273zn4ewlhkw07fap8tig9ujz"
  instance_id     = "i-72q1zvko6i5lnawvg940"
  type            = "ecs"
  weight          = 100
  ip              = "192.168.100.99"
  port            = 80
  description     = "This is a server"
}
```
## Argument Reference
The following arguments are supported:
* `description` - (Optional) The description of the instance.
* `instance_id` - (Optional, ForceNew) The ID of ecs instance or the network card bound to ecs instance.
* `ip` - (Optional, ForceNew) The private ip of the instance.
* `port` - (Optional) The port receiving request.
* `server_group_id` - (Optional, ForceNew) The ID of the ServerGroup.
* `type` - (Optional, ForceNew) The type of instance. Optional choice contains `ecs`, `eni`.
* `weight` - (Optional) The weight of the instance.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `server_id` - The server id of instance in ServerGroup.


## Import
ServerGroupServer can be imported using the id, e.g.
```
$ terraform import volcengine_server_group_server.default rs-3ciynux6i1x4w****rszh49sj
```

