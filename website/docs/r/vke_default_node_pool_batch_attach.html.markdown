---
subcategory: "VKE"
layout: "volcengine"
page_title: "Volcengine: volcengine_vke_default_node_pool_batch_attach"
sidebar_current: "docs-volcengine-resource-vke_default_node_pool_batch_attach"
description: |-
  Provides a resource to manage vke default node pool batch attach
---
# volcengine_vke_default_node_pool_batch_attach
Provides a resource to manage vke default node pool batch attach
## Example Usage
```hcl
resource "volcengine_vke_default_node_pool_batch_attach" "default" {
  cluster_id           = "ccc2umdnqtoflv91lqtq0"
  default_node_pool_id = "11111"
  instances {
    instance_id                          = "i-ybvza90ohwexzk8emaa3"
    keep_instance_name                   = false
    additional_container_storage_enabled = false
  }
  instances {
    instance_id                          = "i-ybvza90ohxexzkm4zihf"
    keep_instance_name                   = false
    additional_container_storage_enabled = true
    container_storage_path               = "/"
  }
}
```
## Argument Reference
The following arguments are supported:
* `cluster_id` - (Required, ForceNew) The ClusterId of NodePool.
* `default_node_pool_id` - (Required, ForceNew) The default NodePool ID.
* `instances` - (Optional) The ECS InstanceIds add to NodePool.

The `instances` object supports the following:

* `instance_id` - (Required) The instance id.
* `additional_container_storage_enabled` - (Optional) The flag of additional container storage enable, the value is `true` or `false`..Default is `false`.
* `container_storage_path` - (Optional) The container storage path.When additional_container_storage_enabled is `false` will ignore.
* `image_id` - (Optional) The Image Id to the ECS Instance.
* `keep_instance_name` - (Optional) The flag of keep instance name, the value is `true` or `false`.Default is `false`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `is_import` - Is import of the DefaultNodePool. It only works when imported, set to true.
* `kubernetes_config` - The KubernetesConfig of NodeConfig.
    * `cordon` - The Cordon of KubernetesConfig.
    * `labels` - The Labels of KubernetesConfig.
        * `key` - The Key of Labels.
        * `value` - The Value of Labels.
    * `taints` - The Taints of KubernetesConfig.
        * `effect` - The Effect of Taints.
        * `key` - The Key of Taints.
        * `value` - The Value of Taints.
* `node_config` - The Config of NodePool.
    * `initialize_script` - The initializeScript of NodeConfig.
    * `security` - The Security of NodeConfig.
        * `login` - The Login of Security.
            * `password` - The Password of Security.
            * `ssh_key_pair_name` - The SshKeyPairName of Security.
        * `security_group_ids` - The SecurityGroupIds of Security.
        * `security_strategies` - The SecurityStrategies of Security.


