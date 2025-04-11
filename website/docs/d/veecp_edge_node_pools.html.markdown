---
subcategory: "VEECP"
layout: "volcengine"
page_title: "Volcengine: volcengine_veecp_edge_node_pools"
sidebar_current: "docs-volcengine-datasource-veecp_edge_node_pools"
description: |-
  Use this data source to query detailed information of veecp edge node pools
---
# volcengine_veecp_edge_node_pools
Use this data source to query detailed information of veecp edge node pools
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

resource "volcengine_veecp_edge_node_pool" "foo" {
  cluster_id = volcengine_veecp_cluster.foo.id
  name       = "acc-test-tf"
}

data "volcengine_veecp_edge_node_pools" "foo" {
  cluster_ids = [volcengine_veecp_cluster.foo.id]
  ids         = [volcengine_veecp_edge_node_pool.foo.id]
}
```
## Argument Reference
The following arguments are supported:
* `add_by_auto` - (Optional) Managed by auto.
* `add_by_list` - (Optional) Managed by list.
* `add_by_script` - (Optional) Managed by script.
* `auto_scaling_enabled` - (Optional) Is enabled of AutoScaling.
* `cluster_ids` - (Optional) The ClusterIds of NodePool IDs.
* `create_client_token` - (Optional) The ClientToken when successfully created.
* `ids` - (Optional) A list of IDs.
* `name_regex` - (Optional) A Name Regex of Resource.
* `node_pool_types` - (Optional) The NodePoolTypes of NodePool.
* `output_file` - (Optional) File name where to save data source results.
* `statuses` - (Optional) The Status of NodePool.
* `update_client_token` - (Optional) The ClientToken when last update was successful.

The `statuses` object supports the following:

* `conditions_type` - (Optional) Indicates the status condition of the node pool in the active state. The value can be `Progressing` or `Ok` or `VersionPartlyUpgraded` or `StockOut` or `LimitedByQuota` or `Balance` or `Degraded` or `ClusterVersionUpgrading` or `Cluster` or `ResourceCleanupFailed` or `Unknown` or `ClusterNotRunning` or `SetByProvider`.
* `phase` - (Optional) The Phase of Status. The value can be `Creating` or `Running` or `Updating` or `Deleting` or `Failed` or `Scaling`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `node_pools` - The collection of query.
    * `billing_configs` - The billing configuration.
        * `auto_renew` - Whether to automatically renew the node pool.
        * `pre_paid_period_number` - Prepaid period number.
        * `pre_paid_period` - The pre-paid period of the node pool, in months. The value range is 1-9. This parameter takes effect only when the billing_type is PrePaid.
    * `cluster_id` - The ClusterId of NodePool.
    * `condition_types` - The Condition of Status.
    * `create_client_token` - The ClientToken when successfully created.
    * `create_time` - The CreateTime of NodePool.
    * `elastic_config` - Elastic scaling configuration of node pool.
        * `auto_scale_config` - The auto scaling configuration.
            * `desired_replicas` - The DesiredReplicas of AutoScaling.
            * `enabled` - Whether to enable auto scaling.
            * `max_replicas` - The maximum number of nodes.
            * `min_replicas` - The minimum number of nodes.
            * `priority` - The Priority of AutoScaling.
        * `cloud_server_identity` - Cloud server identity.
        * `instance_area` - The information of instance area.
            * `area_name` - Region name. You can obtain the regions and operators supported by instance specifications through the ListAvailableResourceInfo interface.
            * `cluster_name` - Cluster name.
            * `default_isp` - Default operator. When using three-line nodes, this parameter can be configured. After configuration, this operator will be used as the default export.
            * `external_network_mode` - Public network configuration of three-line nodes. If it is a single-line node, this parameter will be ignored. Value range: single_interface_multi_ip: Single network card with multiple IPs. single_interface_cmcc_ip: Single network card with China Mobile IP. Relevant permissions need to be opened by submitting a work order. single_interface_cucc_ip: Single network card with China Unicom IP. Relevant permissions need to be opened by submitting a work order. single_interface_ctcc_ip: Single network card with China Telecom IP. Relevant permissions need to be opened by submitting a work order. multi_interface_multi_ip: Multiple network cards with multiple IPs. Relevant permissions need to be opened by submitting a work order. no_interface: No public network network card. Relevant permissions need to be opened by submitting a work order. If this parameter is not configured: When there is a public network network card, single_interface_multi_ip is used by default. When there is no public network network card, no_interface is used by default.
            * `isp` - Operator. You can obtain the regions and operators supported by the instance specification through the ListAvailableResourceInfo interface.
            * `subnet_identity` - Subnet ID.
            * `vpc_identity` - VPC ID.
    * `id` - The Id of NodePool.
    * `label_content` - The LabelContent of KubernetesConfig.
        * `key` - The Key of KubernetesConfig.
        * `value` - The Value of KubernetesConfig.
    * `name` - The Name of NodePool.
    * `node_add_methods` - The method of adding nodes to the node pool.
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
    * `phase` - The Phase of Status.
    * `profile` - Edge: Edge node pool. If the return value is empty, it is the central node pool.
    * `taint_content` - The TaintContent of NodeConfig.
        * `effect` - The Effect of Taint.
        * `key` - The Key of Taint.
        * `value` - The Value of Taint.
    * `type` - Node pool type, machine-set: central node pool. edge-machine-set: edge node pool. edge-machine-pool: edge elastic node pool.
    * `update_client_token` - The ClientToken when last update was successful.
    * `update_time` - The UpdateTime time of NodePool.
    * `vpc_id` - The static node pool specifies the node pool to associate with the VPC.
* `total_count` - The total count of query.


