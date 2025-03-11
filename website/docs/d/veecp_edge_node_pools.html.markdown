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
data "volcengine_veecp_edge_node_pools" "foo" {
  auto_scaling_enabled = true
  cluster_ids          = []
  create_client_token  = ""
  ids                  = []
  node_pool_types      = []
  update_client_token  = ""
}
```
## Argument Reference
The following arguments are supported:
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
    * `cordon` - The Cordon of KubernetesConfig.
    * `create_client_token` - The ClientToken when successfully created.
    * `create_time` - The CreateTime of NodePool.
    * `desired_replicas` - The DesiredReplicas of AutoScaling.
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
    * `enabled` - Is Enabled of AutoScaling.
    * `id` - The Id of NodePool.
    * `kube_config_name_prefix` - The NamePrefix of node metadata.
    * `label_content` - The LabelContent of KubernetesConfig.
        * `key` - The Key of KubernetesConfig.
        * `value` - The Value of KubernetesConfig.
    * `max_replicas` - The MaxReplicas of AutoScaling.
    * `min_replicas` - The MinReplicas of AutoScaling.
    * `name` - The Name of NodePool.
    * `node_add_methods` - The method of adding nodes to the node pool.
    * `node_pool_type` - The NodePoolType of NodePool.
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
    * `priority` - The Priority of AutoScaling.
    * `profile` - Edge: Edge node pool. If the return value is empty, it is the central node pool.
    * `subnet_policy` - Multi-subnet scheduling strategy for nodes. The value can be `ZoneBalance` or `Priority`.
    * `tags` - Tags of the NodePool.
        * `key` - The Key of Tags.
        * `type` - The Type of Tags.
        * `value` - The Value of Tags.
    * `taint_content` - The TaintContent of NodeConfig.
        * `effect` - The Effect of Taint.
        * `key` - The Key of Taint.
        * `value` - The Value of Taint.
    * `type` - Node pool type, machine-set: central node pool. edge-machine-set: edge node pool. edge-machine-pool: edge elastic node pool.
    * `update_client_token` - The ClientToken when last update was successful.
    * `update_time` - The UpdateTime time of NodePool.
    * `vpc_id` - The static node pool specifies the node pool to associate with the VPC.
* `total_count` - The total count of query.


