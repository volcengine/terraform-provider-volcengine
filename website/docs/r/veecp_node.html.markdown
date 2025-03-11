---
subcategory: "VEECP"
layout: "volcengine"
page_title: "Volcengine: volcengine_veecp_node"
sidebar_current: "docs-volcengine-resource-veecp_node"
description: |-
  Provides a resource to manage veecp node
---
# volcengine_veecp_node
Provides a resource to manage veecp node
## Example Usage
```hcl
resource "volcengine_veecp_node" "foo" {
  cluster_id  = ""
  instance_id = ""
}
```
## Argument Reference
The following arguments are supported:
* `cluster_id` - (Required, ForceNew) The cluster id.
* `instance_id` - (Required, ForceNew) The instance id.
* `additional_container_storage_enabled` - (Optional, ForceNew) The flag of additional container storage enable, the value is `true` or `false`. This field is valid only when adding an existing instance to the default node pool.
* `client_token` - (Optional, ForceNew) The client token.
* `container_storage_path` - (Optional, ForceNew) The container storage path. This field is valid only when adding an existing instance to the default node pool.
* `image_id` - (Optional, ForceNew) The ImageId of NodeConfig. This field is valid only when adding an existing instance to the default node pool.
* `initialize_script` - (Optional, ForceNew) The initializeScript of Node. This field is valid only when adding an existing instance to the default node pool.
* `keep_instance_name` - (Optional) The flag of keep instance name, the value is `true` or `false`.
* `kubernetes_config` - (Optional, ForceNew) The KubernetesConfig of Node. This field is valid only when adding an existing instance to the default node pool.
* `node_pool_id` - (Optional, ForceNew) The node pool id. This field is used to specify the custom node pool to which you want to add nodes. If not filled in, it means added to the default node pool.

The `kubernetes_config` object supports the following:

* `cordon` - (Optional, ForceNew) The Cordon of KubernetesConfig.
* `labels` - (Optional, ForceNew) The Labels of KubernetesConfig.
* `taints` - (Optional, ForceNew) The Taints of KubernetesConfig.

The `labels` object supports the following:

* `key` - (Optional, ForceNew) The Key of Labels.
* `value` - (Optional, ForceNew) The Value of Labels.

The `taints` object supports the following:

* `effect` - (Optional, ForceNew) The Effect of Taints, the value can be `NoSchedule` or `NoExecute` or `PreferNoSchedule`.
* `key` - (Optional, ForceNew) The Key of Taints.
* `value` - (Optional, ForceNew) The Value of Taints.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
VeecpNode can be imported using the id, e.g.
```
$ terraform import volcengine_veecp_node.default resource_id
```

