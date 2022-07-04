---
subcategory: "VKE"
layout: "volcengine"
page_title: "Volcengine: volcengine_vke_node"
sidebar_current: "docs-volcengine-resource-vke_node"
description: |-
  Provides a resource to manage vke node
---
# volcengine_vke_node
Provides a resource to manage vke node
## Example Usage
```hcl
resource "volcengine_vke_node" "foo" {
  cluster_id                           = "ccahbr0nqtofhiuuuajn0"
  instance_id                          = "i-ybrfa2vu2t7grbv8qa0j"
  keep_instance_name                   = true
  additional_container_storage_enabled = false
  container_storage_path               = ""
}
```
## Argument Reference
The following arguments are supported:
* `cluster_id` - (Required, ForceNew) The cluster id.
* `instance_id` - (Required, ForceNew) The instance id.
* `additional_container_storage_enabled` - (Optional, ForceNew) The flag of additional container storage enable.
* `client_token` - (Optional, ForceNew) The client token.
* `container_storage_path` - (Optional, ForceNew) The container storage path.
* `keep_instance_name` - (Optional) The flag of keep instance name.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `node_pool_id` - The node pool id.


## Import
VKE node can be imported using the node id, e.g.
```
$ terraform import volcengine_vke_node.default nc5t5epmrsf****
```

