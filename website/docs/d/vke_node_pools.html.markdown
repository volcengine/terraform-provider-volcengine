---
subcategory: "VKE"
layout: "volcengine"
page_title: "Volcengine: volcengine_vke_node_pools"
sidebar_current: "docs-volcengine-datasource-vke_node_pools"
description: |-
  Use this data source to query detailed information of vke node pools
---
# volcengine_vke_node_pools
Use this data source to query detailed information of vke node pools
## Example Usage
```hcl
data "volcengine_vke_node_pools" "vke_test" {
  cluster_ids = ["ccabe57fqtofgrbln3dog"]
  name        = "demo"
}
```
## Argument Reference
The following arguments are supported:
* `auto_scaling_enabled` - (Optional) The Switch of AutoScaling.
* `cluster_id` - (Optional) The ClusterId of NodePool.
* `cluster_ids` - (Optional) The ClusterIds of NodePool IDs.
* `create_client_token` - (Optional) The create client token of NodePool.
* `ids` - (Optional) A list of NodePool IDs.
* `name_regex` - (Optional) A Name Regex of NodePool.
* `name` - (Optional) The Name of NodePool.
* `output_file` - (Optional) File name where to save data source results.
* `statuses` - (Optional) The Status of NodePool.
* `update_client_token` - (Optional) The update client token of NodePool.

The `statuses` object supports the following:

* `conditions_type` - (Optional) The Type of Status.
* `phase` - (Optional) The Phase of Status.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `node_pools` - The collection of NodePools query.
  * `additional_container_storage_enabled` - The additionalContainerStorageEnabled of NodeConfig.
  * `cluster_id` - The ClusterId of NodePool.
  * `condition_types` - The Condition of Status.
  * `cordon` - The Cordon of KubernetesConfig.
  * `create_client_token` - The CreateClientToken of NodePool.
  * `create_time` - The CreateTime time of NodePool.
  * `data_volumes` - The DataVolume of NodeConfig.
    * `size` - The size of DataVolume.
    * `type` - The type of DataVolume.
  * `description` - The description of NodePool.
  * `desired_replicas` - The DesiredReplicas of AutoScaling.
  * `enabled` - The switch of AutoScaling.
  * `id` - The ID of NodePool.
  * `initialize_script` - The InitializeScript of NodeConfig.
  * `instance_type_ids` - The InstanceTypeIds of NodeConfig.
  * `label_content` - The Labels of KubernetesConfig.
    * `key` - The Key of KubernetesConfig.
    * `value` - The Value of KubernetesConfig.
  * `max_replicas` - The MaxReplicas of AutoScaling.
  * `min_replicas` - The MinReplicas of AutoScaling.
  * `name` - The Name of NodePool.
  * `node_statistics` - The NodeStatistics of NodeConfig.
    * `creating_count` - The creatingCount of Node.
    * `deleting_count` - The deletingCount of Node.
    * `failed_count` - The failedCount of Node.
    * `running_count` - The runningCount of Node.
    * `total_count` - The totalCount of Node.
    * `updating_count` - The updatingCount of Node.
  * `phase` - The Phase of Status.
  * `priority` - The Priority of AutoScaling.
  * `subnet_ids` - The SubnetId of NodeConfig.
  * `system_volume` - The SystemVolume of NodeConfig.
    * `size` - The size of SystemVolume.
    * `type` - The type of SystemVolume.
  * `taint_content` - The taintContent of NodeConfig.
    * `effect` - The effect of Taint.
    * `key` - The key of Taint.
    * `value` - The value of Taint.
  * `update_client_token` - The UpdateClientToken of NodePool.
  * `update_time` - The UpdateTime time of NodePool.
* `total_count` - The total count of NodePools query.


