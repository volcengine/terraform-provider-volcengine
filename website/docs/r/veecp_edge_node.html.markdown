---
subcategory: "VEECP"
layout: "volcengine"
page_title: "Volcengine: volcengine_veecp_edge_node"
sidebar_current: "docs-volcengine-resource-veecp_edge_node"
description: |-
  Provides a resource to manage veecp edge node
---
# volcengine_veecp_edge_node
Provides a resource to manage veecp edge node
## Example Usage
```hcl
resource "volcengine_veecp_edge_node" "foo" {
  cluster_id   = ""
  name         = ""
  node_pool_id = ""
}
```
## Argument Reference
The following arguments are supported:
* `cluster_id` - (Required, ForceNew) The cluster id.
* `name` - (Required, ForceNew) The name of node.
* `node_pool_id` - (Required, ForceNew) The node pool id.
* `auto_complete_config` - (Optional, ForceNew) Machine information to be managed.
* `client_token` - (Optional, ForceNew) The client token.

The `auto_complete_config` object supports the following:

* `address` - (Optional, ForceNew) The address of the machine to be managed.
* `direct_add_instances` - (Optional, ForceNew) Edge computing instance ID on Volcano Engine.
* `direct_add` - (Optional, ForceNew) Directly managed through the edge computing instance ID. When it is true, there is no need to provide Address. Only DirectAddInstances needs to be provided.
* `enable` - (Optional, ForceNew) Enable/Disable automatic management.
* `machine_auth` - (Optional, ForceNew) Login credentials.

The `direct_add_instances` object supports the following:

* `cloud_server_identity` - (Required, ForceNew) Edge service ID.
* `instance_identity` - (Required, ForceNew) Edge computing instance ID.

The `machine_auth` object supports the following:

* `auth_type` - (Required, ForceNew) Authentication method. Currently only Password is open.
* `ssh_port` - (Required, ForceNew) SSH port, default 22.
* `user` - (Required, ForceNew) Login username.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
VeecpNode can be imported using the id, e.g.
```
$ terraform import volcengine_veecp_node.default resource_id
```

