---
subcategory: "VKE"
layout: "volcengine"
page_title: "Volcengine: volcengine_vke_default_node_pool"
sidebar_current: "docs-volcengine-resource-vke_default_node_pool"
description: |-
  Provides a resource to manage vke default node pool
---
# volcengine_vke_default_node_pool
Provides a resource to manage vke default node pool
## Example Usage
```hcl
resource "volcengine_vke_default_node_pool" "default" {
  cluster_id = "ccc2umdnqtoflv91lqtq0"
  node_config {
    security {
      login {
        password = "amw4WTdVcTRJVVFsUXpVTw=="
      }
      security_group_ids  = ["sg-2d6t6djr2wge858ozfczv41xq", "sg-3re6v4lz76yv45zsk2hjvvwcj"]
      security_strategies = ["Hids"]
    }
    initialize_script = "ISMvYmluL2Jhc2gKZWNobyAx"
    ecs_tags {
      key   = "ecs_k1"
      value = "ecs_v1"
    }
  }
  kubernetes_config {
    labels {
      key   = "aa"
      value = "bb"
    }
    labels {
      key   = "cccc"
      value = "dddd"
    }
    taints {
      key    = "cccc"
      value  = "dddd"
      effect = "NoSchedule"
    }
    taints {
      key    = "aa11"
      value  = "111"
      effect = "NoSchedule"
    }
    cordon = true
  }
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
  tags {
    key   = "k1"
    value = "v1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `cluster_id` - (Required, ForceNew) The ClusterId of NodePool.
* `kubernetes_config` - (Required) The KubernetesConfig of NodeConfig.
* `node_config` - (Required) The Config of NodePool.
* `instances` - (Optional) The ECS InstanceIds add to NodePool.
* `tags` - (Optional) Tags.

The `ecs_tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

The `instances` object supports the following:

* `instance_id` - (Required) The instance id.
* `additional_container_storage_enabled` - (Optional) The flag of additional container storage enable, the value is `true` or `false`..Default is `false`.
* `container_storage_path` - (Optional) The container storage path.When additional_container_storage_enabled is `false` will ignore.
* `image_id` - (Optional) The Image Id to the ECS Instance.
* `keep_instance_name` - (Optional) The flag of keep instance name, the value is `true` or `false`.Default is `false`.

The `kubernetes_config` object supports the following:

* `cordon` - (Required) The Cordon of KubernetesConfig.
* `labels` - (Optional) The Labels of KubernetesConfig.
* `taints` - (Optional) The Taints of KubernetesConfig.

The `labels` object supports the following:

* `key` - (Optional) The Key of Labels.
* `value` - (Optional) The Value of Labels.

The `login` object supports the following:

* `password` - (Optional) The Password of Security.
* `ssh_key_pair_name` - (Optional) The SshKeyPairName of Security.

The `node_config` object supports the following:

* `security` - (Required) The Security of NodeConfig.
* `ecs_tags` - (Optional) Tags for Ecs.
* `initialize_script` - (Optional) The initializeScript of NodeConfig.
* `name_prefix` - (Optional) The NamePrefix of NodeConfig.

The `security` object supports the following:

* `login` - (Required) The Login of Security.
* `security_group_ids` - (Optional) The SecurityGroupIds of Security.
* `security_strategies` - (Optional) The SecurityStrategies of Security.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

The `taints` object supports the following:

* `effect` - (Optional) The Effect of Taints.
* `key` - (Optional) The Key of Taints.
* `value` - (Optional) The Value of Taints.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `is_import` - Is import of the DefaultNodePool. It only works when imported, set to true.


## Import
VKE default node can be imported using the node id, e.g.
```
$ terraform import volcengine_vke_default_node.default nc5t5epmrsf****
```

