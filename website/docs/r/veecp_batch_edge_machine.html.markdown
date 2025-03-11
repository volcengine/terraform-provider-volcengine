---
subcategory: "VEECP"
layout: "volcengine"
page_title: "Volcengine: volcengine_veecp_batch_edge_machine"
sidebar_current: "docs-volcengine-resource-veecp_batch_edge_machine"
description: |-
  Provides a resource to manage veecp batch edge machine
---
# volcengine_veecp_batch_edge_machine
Provides a resource to manage veecp batch edge machine
## Example Usage
```hcl
resource "volcengine_veecp_batch_edge_machine" "foo" {
  cluster_id   = ""
  name         = ""
  node_pool_id = ""
  ttl_hours    = 1
}
```
## Argument Reference
The following arguments are supported:
* `cluster_id` - (Required, ForceNew) The cluster id.
* `name` - (Required, ForceNew) The name of the node.
* `node_pool_id` - (Required, ForceNew) The node pool id.
* `ttl_hours` - (Required, ForceNew) Effective hours of the managed script are counted from the creation time.
* `client_token` - (Optional) The client token.
* `expiration_date` - (Optional) Expiration date of the managed script, UTC time point, in seconds. If the expiration time is set, TTLHours will be ignored.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
VeecpBatchEdgeMachine can be imported using the id, e.g.
```
$ terraform import volcengine_veecp_batch_edge_machine.default resource_id
```

