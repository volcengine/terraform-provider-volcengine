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
  cluster_id = "ccc2umdnqtoflv91lqtq0"
  name       = "tf-test"
  node_config {
    instance_type_ids = ["ecs.r1.large"]
    subnet_ids        = ["subnet-3reyr9ld3obnk5zsk2iqb1kk3"]
    security {
      login {
        #      ssh_key_pair_name = "ssh-6fbl66fxqm"
        password = "UHdkMTIzNDU2"
      }
      security_group_ids = ["sg-2bz8cga08u48w2dx0eeym1fzy", "sg-2d6t6djr2wge858ozfczv41xq"]
    }
    data_volumes {
      type = "ESSD_PL0"
      size = "60"
    }
    instance_charge_type = "PrePaid"
    period               = 1
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
* `node_config` - (Required) The Config of NodePool.
* `auto_scaling` - (Optional) The node pool elastic scaling configuration information.
* `client_token` - (Optional) The ClientToken of NodePool.
* `cluster_id` - (Optional, ForceNew) The ClusterId of NodePool.
* `kubernetes_config` - (Optional) The KubernetesConfig of NodeConfig.
* `name` - (Optional) The Name of NodePool.

The `auto_scaling` object supports the following:

* `desired_replicas` - (Optional) The DesiredReplicas of AutoScaling, default 0.
* `enabled` - (Optional) Is Enabled of AutoScaling.
* `max_replicas` - (Optional) The MaxReplicas of AutoScaling, default 10, range in 1~1000.
* `min_replicas` - (Optional) The MinReplicas of AutoScaling, default 0.
* `priority` - (Optional) The Priority of AutoScaling, default 10, rang in 0~100.

The `data_volumes` object supports the following:

* `size` - (Optional, ForceNew) The Size of DataVolumes, the value range in 20~32768.
* `type` - (Optional, ForceNew) The Type of DataVolumes, the value can be `PTSSD` or `ESSD_PL0`.

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

* `instance_type_ids` - (Required, ForceNew) The InstanceTypeIds of NodeConfig.
* `security` - (Required) The Security of NodeConfig.
* `subnet_ids` - (Required, ForceNew) The SubnetIds of NodeConfig.
* `additional_container_storage_enabled` - (Optional) The AdditionalContainerStorageEnabled of NodeConfig.
* `auto_renew_period` - (Optional, ForceNew) The AutoRenewPeriod of PrePaid instance of NodeConfig. Valid values: 1, 2, 3, 6, 12. Unit: month. when InstanceChargeType is PrePaid and AutoRenew enable, default value is 1.
* `auto_renew` - (Optional, ForceNew) Is AutoRenew of PrePaid instance of NodeConfig. Valid values: true, false. when InstanceChargeType is PrePaid, default value is true.
* `data_volumes` - (Optional, ForceNew) The DataVolumes of NodeConfig.
* `image_id` - (Optional, ForceNew) The ImageId of NodeConfig.
* `initialize_script` - (Optional) The initializeScript of NodeConfig.
* `instance_charge_type` - (Optional, ForceNew) The InstanceChargeType of PrePaid instance of NodeConfig. Valid values: PostPaid, PrePaid. Default value: PostPaid.
* `period` - (Optional, ForceNew) The Period of PrePaid instance of NodeConfig. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36. Unit: month. when InstanceChargeType is PrePaid, default value is 12.
* `system_volume` - (Optional, ForceNew) The SystemVolume of NodeConfig.

The `security` object supports the following:

* `login` - (Optional) The Login of Security.
* `security_group_ids` - (Optional) The SecurityGroupIds of Security.
* `security_strategies` - (Optional) The SecurityStrategies of Security, the value can be empty or `Hids`.

The `system_volume` object supports the following:

* `size` - (Optional, ForceNew) The Size of SystemVolume, the value range in 20~2048.
* `type` - (Optional, ForceNew) The Type of SystemVolume, the value can be `PTSSD` or `ESSD_PL0`.

The `taints` object supports the following:

* `effect` - (Optional) The Effect of Taints, the value can be `NoSchedule` or `NoExecute` or `PreferNoSchedule`.
* `key` - (Optional) The Key of Taints.
* `value` - (Optional) The Value of Taints.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
NodePool can be imported using the id, e.g.
```
$ terraform import volcengine_node_pools.default pcabe57vqtofgrbln3dp0
```

