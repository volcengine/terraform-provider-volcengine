---
subcategory: "VEECP"
layout: "volcengine"
page_title: "Volcengine: volcengine_veecp_edge_nodes"
sidebar_current: "docs-volcengine-datasource-veecp_edge_nodes"
description: |-
  Use this data source to query detailed information of veecp edge nodes
---
# volcengine_veecp_edge_nodes
Use this data source to query detailed information of veecp edge nodes
## Example Usage
```hcl
resource "volcengine_veecp_edge_node" "foo" {
  cluster_id   = "ccvmf49t1ndqeechmj8p0"
  name         = "test-node"
  node_pool_id = "pcvpkdn7ic26jjcjsa20g"
  auto_complete_config {
    enable     = true
    direct_add = true
    direct_add_instances {
      cloud_server_identity = "cloudserver-wvvflw9qdns2qrk"
      instance_identity     = "veen91912104432151420041"
    }
  }
}

data "volcengine_veecp_edge_nodes" "foo" {
  ids = [volcengine_veecp_edge_node.foo.id]
}
```
## Argument Reference
The following arguments are supported:
* `cluster_ids` - (Optional) A list of Cluster IDs.
* `create_client_token` - (Optional) The Create Client Token.
* `ids` - (Optional) A list of Node IDs.
* `ips` - (Optional) The node ips.
* `name_regex` - (Optional) A Name Regex of Node.
* `name` - (Optional) The Name of Node.
* `need_bootstrap_script` - (Optional) Whether to query the node management script is needed.
* `node_pool_ids` - (Optional) The Node Pool IDs.
* `output_file` - (Optional) File name where to save data source results.
* `statuses` - (Optional) The Status of filter.
* `zone_ids` - (Optional) The Zone IDs.

The `statuses` object supports the following:

* `edge_node_status_condition_type` - (Optional) The Type of Node Condition, the value is `Progressing` or `Ok` or `Unschedulable` or `InitilizeFailed` or `Unknown` or `NotReady` or `Security` or `Balance` or `ResourceCleanupFailed`.
* `phase` - (Optional) The Phase of Node, the value is `Creating` or `Running` or `Updating` or `Deleting` or `Failed` or `Starting` or `Stopping` or `Stopped`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `nodes` - The collection of query.
    * `bootstrap_script` - The bootstrap script of node.
    * `cluster_id` - The cluster id of node.
    * `condition_types` - The Condition of Node.
    * `create_client_token` - The create client token of node.
    * `create_time` - The create time of Node.
    * `edge_node_type` - The edge node type of node.
    * `id` - The ID of Node.
    * `instance_id` - The instance id of node.
    * `name` - The name of Node.
    * `node_pool_id` - The node pool id.
    * `phase` - The Phase of Node.
    * `profile` - The profile of node. Distinguish between edge and central nodes.
    * `provider_id` - The provider id of node.
    * `update_time` - The update time of Node.
* `total_count` - The total count of Node query.


