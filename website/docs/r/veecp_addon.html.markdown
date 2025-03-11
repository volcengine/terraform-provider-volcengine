---
subcategory: "VEECP"
layout: "volcengine"
page_title: "Volcengine: volcengine_veecp_addon"
sidebar_current: "docs-volcengine-resource-veecp_addon"
description: |-
  Provides a resource to manage veecp addon
---
# volcengine_veecp_addon
Provides a resource to manage veecp addon
## Example Usage
```hcl
resource "volcengine_veecp_addon" "foo" {
  cluster_id = ""
  name       = ""
}
```
## Argument Reference
The following arguments are supported:
* `cluster_id` - (Required, ForceNew) The cluster id of the addon.
* `name` - (Required, ForceNew) The name of the addon.
* `config` - (Optional) The config info of addon. Please notice that `ingress-nginx` component prohibits updating config, can only works on the web console.
* `deploy_mode` - (Optional, ForceNew) The deploy mode.
* `deploy_node_type` - (Optional, ForceNew) The deploy node type.
* `version` - (Optional) The version info of the cluster.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
VeecpAddon can be imported using the id, e.g.
```
$ terraform import volcengine_veecp_addon.default resource_id
```

Notice
Some kind of VeecpAddon can not be removed from volcengine, and it will make a forbidden error when try to destroy.
If you want to remove it from terraform state, please use command
```
$ terraform state rm volcengine_veecp_addon.${name}
```

