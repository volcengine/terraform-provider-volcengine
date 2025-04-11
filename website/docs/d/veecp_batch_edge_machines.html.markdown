---
subcategory: "VEECP"
layout: "volcengine"
page_title: "Volcengine: volcengine_veecp_batch_edge_machines"
sidebar_current: "docs-volcengine-datasource-veecp_batch_edge_machines"
description: |-
  Use this data source to query detailed information of veecp batch edge machines
---
# volcengine_veecp_batch_edge_machines
Use this data source to query detailed information of veecp batch edge machines
## Example Usage
```hcl
resource "volcengine_veecp_batch_edge_machine" "foo" {
  cluster_id   = "ccvd7mte6t101fno98u60"
  name         = "tf-test"
  node_pool_id = "pcvd90uacnsr73g6bjic0"
  ttl_hours    = 1
}

data "volcengine_veecp_batch_edge_machines" "foo" {
  cluster_ids = [volcengine_veecp_batch_edge_machine.foo.cluster_id]
  #    create_client_token = ""
  ids = [volcengine_veecp_batch_edge_machine.foo.id]
  #    ips = []
  #    name = ""
  #    need_bootstrap_script = ""
  #    zone_ids = []
}
```
## Argument Reference
The following arguments are supported:
* `cluster_ids` - (Optional) The ClusterIds of NodePool IDs.
* `create_client_token` - (Optional) The ClientToken when successfully created.
* `ids` - (Optional) A list of IDs.
* `ips` - (Optional) The IPs.
* `name` - (Optional) The Name of NodePool.
* `need_bootstrap_script` - (Optional) Whether it is necessary to query the node management script.
* `output_file` - (Optional) File name where to save data source results.
* `statuses` - (Optional) The Status of NodePool.
* `zone_ids` - (Optional) The Zone Ids.

The `statuses` object supports the following:

* `edge_node_status_condition_type` - (Optional) Indicates the status condition of the node pool in the active state. The value can be `Progressing` or `Ok` or `VersionPartlyUpgraded` or `StockOut` or `LimitedByQuota` or `Balance` or `Degraded` or `ClusterVersionUpgrading` or `Cluster` or `ResourceCleanupFailed` or `Unknown` or `ClusterNotRunning` or `SetByProvider`.
* `phase` - (Optional) The Phase of Status. The value can be `Creating` or `Running` or `Updating` or `Deleting` or `Failed` or `Scaling`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `machines` - The collection of query.
    * `bootstrap_script` - The bootstrap script.
    * `cluster_id` - The ClusterId of NodePool.
    * `condition_types` - The Condition of Status.
    * `create_client_token` - The ClientToken when successfully created.
    * `create_time` - The CreateTime of NodePool.
    * `edge_node_type` - Edge node type.
    * `id` - The Id of NodePool.
    * `name` - The Name of NodePool.
    * `phase` - The Phase of Status.
    * `profile` - Edge: Edge node pool. If the return value is empty, it is the central node pool.
    * `ttl_time` - The TTL time.
    * `update_time` - The UpdateTime time of NodePool.
* `total_count` - The total count of query.


