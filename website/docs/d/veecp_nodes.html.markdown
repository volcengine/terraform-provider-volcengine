---
subcategory: "VEECP"
layout: "volcengine"
page_title: "Volcengine: volcengine_veecp_nodes"
sidebar_current: "docs-volcengine-datasource-veecp_nodes"
description: |-
  Use this data source to query detailed information of veecp nodes
---
# volcengine_veecp_nodes
Use this data source to query detailed information of veecp nodes
## Example Usage
```hcl
data "volcengine_veecp_nodes" "foo" {
  cluster_ids         = []
  create_client_token = ""
  ids                 = []
  name                = ""
  node_pool_ids       = []
  zone_ids            = []
}
```
## Argument Reference
The following arguments are supported:
* `cluster_ids` - (Optional) A list of Cluster IDs.
* `create_client_token` - (Optional) The Create Client Token.
* `ids` - (Optional) A list of Node IDs.
* `name_regex` - (Optional) A Name Regex of Node.
* `name` - (Optional) The Name of Node.
* `node_pool_ids` - (Optional) The Node Pool IDs.
* `output_file` - (Optional) File name where to save data source results.
* `statuses` - (Optional) The Status of filter.
* `zone_ids` - (Optional) The Zone IDs.

The `statuses` object supports the following:

* `conditions_type` - (Optional) The Type of Node Condition, the value is `Progressing` or `Ok` or `Unschedulable` or `InitilizeFailed` or `Unknown` or `NotReady` or `Security` or `Balance` or `ResourceCleanupFailed`.
* `phase` - (Optional) The Phase of Node, the value is `Creating` or `Running` or `Updating` or `Deleting` or `Failed` or `Starting` or `Stopping` or `Stopped`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `nodes` - The collection of Node query.
    * `additional_container_storage_enabled` - Is Additional Container storage enables.
    * `cluster_id` - The cluster id of node.
    * `condition_types` - The Condition of Node.
    * `container_storage_path` - The Storage Path.
    * `cordon` - The Cordon of KubernetesConfig.
    * `create_client_token` - The create client token of node.
    * `create_time` - The create time of Node.
    * `id` - The ID of Node.
    * `image_id` - The ImageId of NodeConfig.
    * `initialize_script` - The InitializeScript of NodeConfig.
    * `instance_id` - The instance id of node.
    * `is_virtual` - Is virtual node.
    * `labels` - The Label of KubernetesConfig.
        * `key` - The Key of KubernetesConfig.
        * `value` - The Value of KubernetesConfig.
    * `name` - The name of Node.
    * `node_pool_id` - The node pool id.
    * `phase` - The Phase of Node.
    * `roles` - The roles of node.
    * `taints` - The Taint of KubernetesConfig.
        * `effect` - The Effect of Taint.
        * `key` - The Key of Taint.
        * `value` - The Value of Taint.
    * `update_time` - The update time of Node.
    * `zone_id` - The zone id.
* `total_count` - The total count of Node query.


