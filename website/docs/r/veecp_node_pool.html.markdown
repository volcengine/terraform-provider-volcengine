---
subcategory: "VEECP"
layout: "volcengine"
page_title: "Volcengine: volcengine_veecp_node_pool"
sidebar_current: "docs-volcengine-resource-veecp_node_pool"
description: |-
  Provides a resource to manage veecp node pool
---
# volcengine_veecp_node_pool
Provides a resource to manage veecp node pool
## Example Usage
```hcl
data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-project1"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-subnet-test-2"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_security_group" "foo" {
  vpc_id              = volcengine_vpc.foo.id
  security_group_name = "acc-test-security-group2"
}

resource "volcengine_veecp_cluster" "foo" {
  name                      = "acc-test-1"
  description               = "created by terraform"
  delete_protection_enabled = false
  profile                   = "Edge"
  cluster_config {
    subnet_ids                       = [volcengine_subnet.foo.id]
    api_server_public_access_enabled = true
    api_server_public_access_config {
      public_access_network_config {
        billing_type = "PostPaidByBandwidth"
        bandwidth    = 1
      }
    }
    resource_public_access_default_enabled = true
  }
  pods_config {
    pod_network_mode = "Flannel"
    flannel_config {
      pod_cidrs         = ["172.22.224.0/20"]
      max_pods_per_node = 64
    }
  }
  services_config {
    service_cidrsv4 = ["172.30.0.0/18"]
  }
}

resource "volcengine_veecp_node_pool" "foo" {
  cluster_id   = volcengine_veecp_cluster.foo.id
  name         = "acc-test-node-pool-9505"
  client_token = "FGAHIxa23412FGAIOHioj"
  auto_scaling {
    enabled          = true
    min_replicas     = 0
    max_replicas     = 5
    desired_replicas = 0
    priority         = 5
    subnet_policy    = "ZoneBalance"
  }
  node_config {
    instance_type_ids = ["ecs.c1ie.xlarge"]
    subnet_ids        = [volcengine_subnet.foo.id]
    image_id          = ""
    system_volume {
      type = "ESSD_PL0"
      size = 80
    }
    data_volumes {
      type        = "ESSD_PL0"
      size        = 80
      mount_point = "/tf1"
    }
    data_volumes {
      type        = "ESSD_PL0"
      size        = 60
      mount_point = "/tf2"
    }
    initialize_script = "ZWNobyBoZWxsbyB0ZXJyYWZvcm0h"
    security {
      login {
        password = "UHdkMTIzNDU2"
      }
      security_strategies = ["Hids"]
      security_group_ids  = [volcengine_security_group.foo.id]
    }
    additional_container_storage_enabled = false
    instance_charge_type                 = "PostPaid"
    name_prefix                          = "acc-test"
    ecs_tags {
      key   = "ecs_k1"
      value = "ecs_v1"
    }
  }
  kubernetes_config {
    labels {
      key   = "label1"
      value = "value1"
    }
    taints {
      key    = "taint-key/node-type"
      value  = "taint-value"
      effect = "NoSchedule"
    }
    cordon = true
    #auto_sync_disabled = false
  }
}
```
## Argument Reference
The following arguments are supported:
* `kubernetes_config` - (Required) The KubernetesConfig of NodeConfig.
* `node_config` - (Required) The Config of NodePool.
* `auto_scaling` - (Optional) The node pool elastic scaling configuration information.
* `client_token` - (Optional) The ClientToken of NodePool.
* `cluster_id` - (Optional, ForceNew) The ClusterId of NodePool.
* `instance_ids` - (Optional) The list of existing ECS instance ids. Add existing instances with same type of security group under the same cluster VPC to the custom node pool.
Note that removing instance ids from the list will only remove the nodes from cluster and not release the ECS instances. But deleting node pool will release the ECS instances in it.
It is not recommended to use this field, it is recommended to use `volcengine_veecp_node` resource to add an existing instance to a custom node pool.
* `keep_instance_name` - (Optional) Whether to keep instance name when adding an existing instance to a custom node pool, the value is `true` or `false`.
This field is valid only when adding new instances to the custom node pool.
* `name` - (Optional) The Name of NodePool.

The `auto_scaling` object supports the following:

* `desired_replicas` - (Optional) The DesiredReplicas of AutoScaling, default 0, range in min_replicas to max_replicas.
* `enabled` - (Optional) Whether to enable the auto scaling function of the node pool. When a node needs to be manually added to the node pool, the value of this field must be `false`.
* `max_replicas` - (Optional) The MaxReplicas of AutoScaling, default 10, range in 1~2000. This field is valid when the value of `enabled` is `true`.
* `min_replicas` - (Optional) The MinReplicas of AutoScaling, default 0. This field is valid when the value of `enabled` is `true`.
* `priority` - (Optional) The Priority of AutoScaling, default 10, rang in 0~100. This field is valid when the value of `enabled` is `true` and the value of `subnet_policy` is `Priority`.
* `subnet_policy` - (Optional) Multi-subnet scheduling strategy for nodes. The value can be `ZoneBalance` or `Priority`.

The `data_volumes` object supports the following:

* `mount_point` - (Optional) The target mount directory of the disk. Must start with `/`.
* `size` - (Optional) The Size of DataVolumes, the value range in 20~32768. Default value is `20`.
* `type` - (Optional) The Type of DataVolumes, the value can be `PTSSD` or `ESSD_PL0` or `ESSD_FlexPL`. Default value is `ESSD_PL0`.

The `ecs_tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

The `kubernetes_config` object supports the following:

* `cordon` - (Required) The Cordon of KubernetesConfig.
* `auto_sync_disabled` - (Optional) Whether to disable the function of automatically synchronizing labels and taints to existing nodes. Default is false.
* `labels` - (Optional) The Labels of KubernetesConfig.
* `name_prefix` - (Optional) The NamePrefix of node metadata.
* `taints` - (Optional) The Taints of KubernetesConfig.

The `labels` object supports the following:

* `key` - (Optional) The Key of Labels.
* `value` - (Optional) The Value of Labels.

The `login` object supports the following:

* `password` - (Optional) The Password of Security, this field must be encoded with base64.
* `ssh_key_pair_name` - (Optional) The SshKeyPairName of Security.

The `node_config` object supports the following:

* `instance_type_ids` - (Required) The InstanceTypeIds of NodeConfig. The value can get from volcengine_veecp_support_resource_types datasource.
* `security` - (Required) The Security of NodeConfig.
* `subnet_ids` - (Required) The SubnetIds of NodeConfig.
* `additional_container_storage_enabled` - (Optional) The AdditionalContainerStorageEnabled of NodeConfig.
* `auto_renew_period` - (Optional) The AutoRenewPeriod of PrePaid instance of NodeConfig. Valid values: 1, 2, 3, 6, 12. Unit: month. when InstanceChargeType is PrePaid and AutoRenew enable, default value is 1.
* `auto_renew` - (Optional) Is AutoRenew of PrePaid instance of NodeConfig. Valid values: true, false. when InstanceChargeType is PrePaid, default value is true.
* `data_volumes` - (Optional) The DataVolumes of NodeConfig.
* `ecs_tags` - (Optional) Tags for Ecs.
* `hpc_cluster_ids` - (Optional) The IDs of HpcCluster, only one ID is supported currently.
* `image_id` - (Optional) The ImageId of NodeConfig.
* `initialize_script` - (Optional) The initializeScript of NodeConfig.
* `instance_charge_type` - (Optional, ForceNew) The InstanceChargeType of PrePaid instance of NodeConfig. Valid values: PostPaid, PrePaid. Default value: PostPaid.
* `name_prefix` - (Optional) The NamePrefix of NodeConfig.
* `period` - (Optional) The Period of PrePaid instance of NodeConfig. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36. Unit: month. when InstanceChargeType is PrePaid, default value is 12.
* `system_volume` - (Optional) The SystemVolume of NodeConfig.

The `security` object supports the following:

* `login` - (Optional) The Login of Security.
* `security_group_ids` - (Optional) The SecurityGroupIds of Security.
* `security_strategies` - (Optional) The SecurityStrategies of Security, the value can be empty or `Hids`.

The `system_volume` object supports the following:

* `size` - (Optional) The Size of SystemVolume, the value range in 20~2048.
* `type` - (Optional) The Type of SystemVolume, the value can be `PTSSD` or `ESSD_PL0` or `ESSD_FlexPL`.

The `taints` object supports the following:

* `effect` - (Optional) The Effect of Taints, the value can be `NoSchedule` or `NoExecute` or `PreferNoSchedule`.
* `key` - (Optional) The Key of Taints.
* `value` - (Optional) The Value of Taints.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `node_statistics` - The NodeStatistics of NodeConfig.
    * `creating_count` - The CreatingCount of Node.
    * `deleting_count` - The DeletingCount of Node.
    * `failed_count` - The FailedCount of Node.
    * `running_count` - The RunningCount of Node.
    * `starting_count` - (**Deprecated**) This field has been deprecated and is not recommended for use. The StartingCount of Node.
    * `stopped_count` - (**Deprecated**) This field has been deprecated and is not recommended for use. The StoppedCount of Node.
    * `stopping_count` - (**Deprecated**) This field has been deprecated and is not recommended for use. The StoppingCount of Node.
    * `total_count` - The TotalCount of Node.
    * `updating_count` - The UpdatingCount of Node.


## Import
VeecpNodePool can be imported using the id, e.g.
```
$ terraform import volcengine_veecp_node_pool.default resource_id
```

