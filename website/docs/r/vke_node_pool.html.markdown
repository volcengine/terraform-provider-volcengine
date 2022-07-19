---
subcategory: "VKE"
layout: "volcengine"
page_title: "Volcengine: volcengine_vke_node_pool"
sidebar_current: "docs-volcengine-resource-vke_node_pool"
description: |-
  Provides a resource to manage vke node pool
---
# volcengine_vke_node_pool
Provides a resource to manage vke node pool
## Example Usage
```hcl
resource "volcengine_vke_node_pool" "vke_test" {
  cluster_id = "ccah01nnqtofnluts98j0"
  name       = "demo20"
  node_config {
    instance_type_ids = ["ecs.r1.large"]
    subnet_ids        = ["subnet-3recgzi7hfim85zsk2i8l9ve7"]
    security {
      login {
        #      ssh_key_pair_name = "ssh-6fbl66fxqm"
        password = "UHdkMTIzNDU2"
      }
    }
    data_volumes {
      type = "ESSD_PL0"
      size = "60"
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
  }
}
```
## Argument Reference
The following arguments are supported:
* `auto_scaling_enabled` - (Optional) Is enabled of AutoScaling.
* `auto_scaling` - (Optional) The node pool elastic scaling configuration information.
* `client_token` - (Optional) The ClientToken of NodePool.
* `cluster_id` - (Optional, ForceNew) The ClusterId of NodePool.
* `cluster_ids` - (Optional) The ClusterIds of NodePool.
* `ids` - (Optional) The IDs of NodePool.
* `kubernetes_config` - (Optional) The KubernetesConfig of NodeConfig.
* `name` - (Optional) The Name of NodePool.
* `node_config` - (Optional) The Config of NodePool.
* `statuses` - (Optional) The Status of NodePool.

The `auto_scaling` object supports the following:

* `desired_replicas` - (Optional) The DesiredReplicas of AutoScaling.
* `enabled` - (Optional) Is Enabled of AutoScaling.
* `max_replicas` - (Optional) The MaxReplicas of AutoScaling.
* `min_replicas` - (Optional) The MinReplicas of AutoScaling.
* `priority` - (Optional) The Priority of AutoScaling.

The `data_volumes` object supports the following:

* `size` - (Optional, ForceNew) The Size of DataVolumes.
* `type` - (Optional, ForceNew) The Type of DataVolumes.

The `kubernetes_config` object supports the following:

* `cordon` - (Optional) The Cordon of KubernetesConfig.
* `labels` - (Optional) The Labels of KubernetesConfig.
* `taints` - (Optional) The Taints of KubernetesConfig.

The `labels` object supports the following:

* `key` - (Optional) The Key of Labels.
* `value` - (Optional) The Value of Labels.

The `login` object supports the following:

* `password` - (Optional) The Password of Security.
* `ssh_key_pair_name` - (Optional) The SshKeyPairName of Security.

The `node_config` object supports the following:

* `additional_container_storage_enabled` - (Optional) The AdditionalContainerStorageEnabled of NodeConfig.
* `data_volumes` - (Optional, ForceNew) The DataVolumes of NodeConfig.
* `initialize_script` - (Optional) The initializeScript of NodeConfig.
* `instance_type_ids` - (Optional, ForceNew) The InstanceTypeIds of NodeConfig.
* `security` - (Optional) The Security of NodeConfig.
* `subnet_ids` - (Optional, ForceNew) The SubnetIds of NodeConfig.
* `system_volume` - (Optional, ForceNew) The SystemVolume of NodeConfig.

The `security` object supports the following:

* `login` - (Optional) The Login of Security.
* `security_group_ids` - (Optional) The SecurityGroupIds of Security.
* `security_strategies` - (Optional) The SecurityStrategies of Security.

The `statuses` object supports the following:

* `conditions_type` - (Optional) Indicates the status condition of the node pool in the active state.
* `phase` - (Optional) The Phase of Status.

The `system_volume` object supports the following:

* `size` - (Optional, ForceNew) The Size of SystemVolume.
* `type` - (Optional, ForceNew) The Type of SystemVolume.

The `taints` object supports the following:

* `effect` - (Optional) The Effect of Taints.
* `key` - (Optional) The Key of Taints.
* `value` - (Optional) The Value of Taints.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `create_client_token` - The CreateClientToken of NodePool.
* `update_client_token` - The UpdateClientToken of NodePool.


## Import
NodePool can be imported using the id, e.g.
```
$ terraform import volcengine_node_pools.default pcabe57vqtofgrbln3dp0
```

