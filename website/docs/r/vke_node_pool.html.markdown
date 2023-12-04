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
data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_security_group" "foo" {
  security_group_name = "acc-test-security-group"
  vpc_id              = volcengine_vpc.foo.id
}

data "volcengine_images" "foo" {
  name_regex = "veLinux 1.0 CentOS兼容版 64位"
}

resource "volcengine_vke_cluster" "foo" {
  name                      = "acc-test-cluster"
  description               = "created by terraform"
  delete_protection_enabled = false
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
    pod_network_mode = "VpcCniShared"
    vpc_cni_config {
      subnet_ids = [volcengine_subnet.foo.id]
    }
  }
  services_config {
    service_cidrsv4 = ["172.30.0.0/18"]
  }
  tags {
    key   = "tf-k1"
    value = "tf-v1"
  }
}

resource "volcengine_vke_node_pool" "foo" {
  cluster_id = volcengine_vke_cluster.foo.id
  name       = "acc-test-node-pool"
  auto_scaling {
    enabled          = true
    min_replicas     = 0
    max_replicas     = 5
    desired_replicas = 0
    priority         = 5
    subnet_policy    = "ZoneBalance"
  }
  node_config {
    instance_type_ids = ["ecs.g1ie.xlarge"]
    subnet_ids        = [volcengine_subnet.foo.id]
    image_id          = [for image in data.volcengine_images.foo.images : image.image_id if image.image_name == "veLinux 1.0 CentOS兼容版 64位"][0]
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
  }
  tags {
    key   = "node-pool-k1"
    value = "node-pool-v1"
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
* `name` - (Optional) The Name of NodePool.
* `tags` - (Optional) Tags.

The `auto_scaling` object supports the following:

* `desired_replicas` - (Optional) The DesiredReplicas of AutoScaling, default 0, range in min_replicas to max_replicas.
* `enabled` - (Optional) Is Enabled of AutoScaling.
* `max_replicas` - (Optional) The MaxReplicas of AutoScaling, default 10, range in 1~2000.
* `min_replicas` - (Optional) The MinReplicas of AutoScaling, default 0.
* `priority` - (Optional) The Priority of AutoScaling, default 10, rang in 0~100.
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
* `labels` - (Optional) The Labels of KubernetesConfig.
* `taints` - (Optional) The Taints of KubernetesConfig.

The `labels` object supports the following:

* `key` - (Optional) The Key of Labels.
* `value` - (Optional) The Value of Labels.

The `login` object supports the following:

* `password` - (Optional) The Password of Security, this field must be encoded with base64.
* `ssh_key_pair_name` - (Optional) The SshKeyPairName of Security.

The `node_config` object supports the following:

* `instance_type_ids` - (Required) The InstanceTypeIds of NodeConfig. The value can get from volcengine_vke_support_resource_types datasource.
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

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

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
    * `starting_count` - The StartingCount of Node.
    * `stopped_count` - The StoppedCount of Node.
    * `stopping_count` - The StoppingCount of Node.
    * `total_count` - The TotalCount of Node.
    * `updating_count` - The UpdatingCount of Node.


## Import
NodePool can be imported using the id, e.g.
```
$ terraform import volcengine_vke_node_pool.default pcabe57vqtofgrbln3dp0
```

