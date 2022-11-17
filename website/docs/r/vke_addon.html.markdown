---
subcategory: "VKE"
layout: "volcengine"
page_title: "Volcengine: volcengine_vke_addon"
sidebar_current: "docs-volcengine-resource-vke_addon"
description: |-
  Provides a resource to manage vke addon
---
# volcengine_vke_addon
Provides a resource to manage vke addon
## Example Usage
```hcl
resource "volcengine_vke_addon" "foo" {
  cluster_id       = "cccctv1vqtofp49d96ujg"
  name             = "csi-nas"
  version          = "v0.1.3"
  deploy_node_type = "Node"
  deploy_mode      = "Unmanaged"
  config           = "{\"xxx\":\"true\"}"
}
```
## Argument Reference
The following arguments are supported:
* `cluster_id` - (Required, ForceNew) The cluster id of the addon.
* `name` - (Required, ForceNew) The name of the addon.
* `config` - (Optional) The config info of addon. Please notice that `ingress-nginx` component prohibits updating config, can only works on the web console.
* `deploy_mode` - (Optional, ForceNew) The deploy mode.
* `deploy_node_type` - (Optional, ForceNew) The deploy node type.
* `version` - (Optional, ForceNew) The version info of the cluster.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
VkeAddon can be imported using the clusterId:Name, e.g.
```
$ terraform import volcengine_vke_addon.default cc9l74mvqtofjnoj5****:nginx-ingress
```

Notice
Some kind of VKEAddon can not be removed from volcengine, and it will make a forbidden error when try to destroy.
If you want to remove it from terraform state, please use command
```
$ terraform state rm volcengine_vke_addon.${name}
```

