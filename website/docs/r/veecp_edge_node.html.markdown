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
  cluster_id   = "ccvmf49t1ndqeechmj8p0"
  name         = "test-node"
  node_pool_id = "pcvpkdn7ic26jjcjsa20g"
  auto_complete_config {
    enable     = true
    direct_add = true
    direct_add_instances {
      cloud_server_identity = "cloudserver-wvvflw9qdns2qrk"
      instance_identity     = "veen91912104432151420041"
    }
  }
}
```
## Argument Reference
The following arguments are supported:
* `auto_complete_config` - (Required, ForceNew) Machine information to be managed.
* `cluster_id` - (Required, ForceNew) The cluster id.
* `name` - (Required, ForceNew) The name of node.
* `node_pool_id` - (Required, ForceNew) The node pool id.
* `client_token` - (Optional, ForceNew) The client token.

The `auto_complete_config` object supports the following:

* `enable` - (Required, ForceNew) Enable/Disable automatic management.
* `address` - (Optional, ForceNew) The address of the machine to be managed.
* `direct_add_instances` - (Optional, ForceNew) Edge computing instance ID on Volcano Engine.
* `direct_add` - (Optional, ForceNew) Directly managed through the edge computing instance ID. When it is true, there is no need to provide Address. Only DirectAddInstances needs to be provided.
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

