---
subcategory: "VEECP"
layout: "volcengine"
page_title: "Volcengine: volcengine_veecp_node_pools"
sidebar_current: "docs-volcengine-datasource-veecp_node_pools"
description: |-
  Use this data source to query detailed information of veecp node pools
---
# volcengine_veecp_node_pools
Use this data source to query detailed information of veecp node pools
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

data "volcengine_veecp_node_pools" "foo" {
  ids = [volcengine_veecp_node_pool.foo.id]
}
```
## Argument Reference
The following arguments are supported:
* `auto_scaling_enabled` - (Optional) Is enabled of AutoScaling.
* `cluster_id` - (Optional) The ClusterId of NodePool.
* `cluster_ids` - (Optional) The ClusterIds of NodePool IDs.
* `create_client_token` - (Optional) The ClientToken when successfully created.
* `ids` - (Optional) The IDs of NodePool.
* `name_regex` - (Optional) A Name Regex of Resource.
* `name` - (Optional) The Name of NodePool.
* `output_file` - (Optional) File name where to save data source results.
* `statuses` - (Optional) The Status of NodePool.
* `update_client_token` - (Optional) The ClientToken when last update was successful.

The `statuses` object supports the following:

* `conditions_type` - (Optional) Indicates the status condition of the node pool in the active state. The value can be `Progressing` or `Ok` or `VersionPartlyUpgraded` or `StockOut` or `LimitedByQuota` or `Balance` or `Degraded` or `ClusterVersionUpgrading` or `Cluster` or `ResourceCleanupFailed` or `Unknown` or `ClusterNotRunning` or `SetByProvider`.
* `phase` - (Optional) The Phase of Status. The value can be `Creating` or `Running` or `Updating` or `Deleting` or `Failed` or `Scaling`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `node_pools` - The collection of NodePools query.
    * `additional_container_storage_enabled` - Is AdditionalContainerStorageEnabled of NodeConfig.
    * `auto_renew_period` - The AutoRenewPeriod of the PrePaid instance of NodeConfig.
    * `auto_renew` - Is auto renew of the PrePaid instance of NodeConfig.
    * `cluster_id` - The ClusterId of NodePool.
    * `condition_types` - The Condition of Status.
    * `cordon` - The Cordon of KubernetesConfig.
    * `create_client_token` - The ClientToken when successfully created.
    * `create_time` - The CreateTime of NodePool.
    * `data_volumes` - The DataVolume of NodeConfig.
        * `mount_point` - The target mount directory of the disk.
        * `size` - The Size of DataVolume.
        * `type` - The Type of DataVolume.
    * `desired_replicas` - The DesiredReplicas of AutoScaling.
    * `ecs_tags` - Tags for Ecs.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `enabled` - Is Enabled of AutoScaling.
    * `hpc_cluster_ids` - The IDs of HpcCluster.
    * `id` - The Id of NodePool.
    * `image_id` - The ImageId of NodeConfig.
    * `initialize_script` - The InitializeScript of NodeConfig.
    * `instance_charge_type` - The InstanceChargeType of NodeConfig.
    * `instance_type_ids` - The InstanceTypeIds of NodeConfig.
    * `kube_config_name_prefix` - The NamePrefix of node metadata.
    * `label_content` - The LabelContent of KubernetesConfig.
        * `key` - The Key of KubernetesConfig.
        * `value` - The Value of KubernetesConfig.
    * `login_key_pair_name` - The login SshKeyPairName of NodeConfig.
    * `login_type` - The login type of NodeConfig.
    * `max_replicas` - The MaxReplicas of AutoScaling.
    * `min_replicas` - The MinReplicas of AutoScaling.
    * `name_prefix` - The NamePrefix of NodeConfig.
    * `name` - The Name of NodePool.
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
    * `period` - The period of the PrePaid instance of NodeConfig.
    * `phase` - The Phase of Status.
    * `priority` - The Priority of AutoScaling.
    * `security_group_ids` - The SecurityGroupIds of NodeConfig.
    * `security_strategies` - The SecurityStrategies of NodeConfig.
    * `security_strategy_enabled` - The SecurityStrategyEnabled of NodeConfig.
    * `subnet_ids` - The SubnetId of NodeConfig.
    * `subnet_policy` - Multi-subnet scheduling strategy for nodes. The value can be `ZoneBalance` or `Priority`.
    * `system_volume` - The SystemVolume of NodeConfig.
        * `size` - The Size of SystemVolume.
        * `type` - The Type of SystemVolume.
    * `tags` - Tags of the NodePool.
        * `key` - The Key of Tags.
        * `type` - The Type of Tags.
        * `value` - The Value of Tags.
    * `taint_content` - The TaintContent of NodeConfig.
        * `effect` - The Effect of Taint.
        * `key` - The Key of Taint.
        * `value` - The Value of Taint.
    * `update_client_token` - The ClientToken when last update was successful.
    * `update_time` - The UpdateTime time of NodePool.
* `total_count` - The total count of query.


