---
subcategory: "VKE"
layout: "volcengine"
page_title: "Volcengine: volcengine_vke_nodes"
sidebar_current: "docs-volcengine-datasource-vke_nodes"
description: |-
  Use this data source to query detailed information of vke nodes
---
# volcengine_vke_nodes
Use this data source to query detailed information of vke nodes
## Example Usage
```hcl
data "volcengine_vke_nodes" "default" {
  ids         = ["ncaa3e5mrsferqkomi190"]
  cluster_ids = ["c123", "c456"]
  statuses {
    phase           = "Creating"
    conditions_type = "Progressing"
  }
  statuses {
    phase           = "Creating123"
    conditions_type = "Progressing123"
  }
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

The `statuses` object supports the following:

* `conditions_type` - (Optional) The Type of Node Condition, the value is `Progressing` or `Ok` or `Unschedulable` or `InitilizeFailed` or `Unknown` or `NotReady` or `Security` or `Balance` or `ResourceCleanupFailed`.
* `phase` - (Optional) The Phase of Node, the value is `Creating` or `Running` or `Updating` or `Deleting` or `Failed`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `nodes` - The collection of Node query.
    * `additional_container_storage_enabled` - Is Additional Container storage enables.
    * `cluster_id` - The cluster id of node.
    * `condition_types` - The Condition of Node.
    * `container_storage_path` - The Storage Path.
    * `create_client_token` - The create client token of node.
    * `create_time` - The create time of Node.
    * `id` - The ID of Node.
    * `instance_id` - The instance id of node.
    * `is_virtual` - Is virtual node.
    * `name` - The name of Node.
    * `node_pool_id` - The node pool id.
    * `phase` - The Phase of Node.
    * `roles` - The roles of node.
    * `update_time` - The update time of Node.
* `total_count` - The total count of Node query.


