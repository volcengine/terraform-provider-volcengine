---
subcategory: "CLB"
layout: "volcengine"
page_title: "Volcengine: volcengine_server_group"
sidebar_current: "docs-volcengine-resource-server_group"
description: |-
  Provides a resource to manage server group
---
# volcengine_server_group
Provides a resource to manage server group
## Example Usage
```hcl
resource "volcengine_server_group" "foo" {
  load_balancer_id  = "clb-273z7d4r8tvk07fap8tsniyfe"
  server_group_name = "demo-demo11"
  description       = "hello demo11"
}
```
## Argument Reference
The following arguments are supported:
* `load_balancer_id` - (Required, ForceNew) The ID of the Clb.
* `description` - (Optional) The description of ServerGroup.
* `server_group_id` - (Optional) The ID of the ServerGroup.
* `server_group_name` - (Optional) The name of the ServerGroup.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
ServerGroupServer can be imported using the id, e.g.
```
$ terraform import volcengine_server_group_server.default rs-3ciynux6i1x4w****rszh49sj
```

