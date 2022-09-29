---
subcategory: "VKE"
layout: "volcengine"
page_title: "Volcengine: volcengine_vke_addons"
sidebar_current: "docs-volcengine-datasource-vke_addons"
description: |-
  Use this data source to query detailed information of vke addons
---
# volcengine_vke_addons
Use this data source to query detailed information of vke addons
## Example Usage
```hcl
data "volcengine_vke_addons" "default" {
  cluster_ids = ["cccctv1vqtofp49d96ujg"]
}
```
## Argument Reference
The following arguments are supported:
* `cluster_ids` - (Optional) The IDs of Cluster.
* `create_client_token` - (Optional) ClientToken when the addon is created successfully. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.
* `deploy_modes` - (Optional) The deploy model, the value is `Managed` or `Unmanaged`.
* `deploy_node_types` - (Optional) The deploy node types, the value is `Node` or `VirtualNode`. Only effected when deploy_mode is `Unmanaged`.
* `name_regex` - (Optional) A Name Regex of addon.
* `names` - (Optional) The Names of addons.
* `output_file` - (Optional) File name where to save data source results.
* `statuses` - (Optional) Array of addon states to filter.
* `update_client_token` - (Optional) The ClientToken when the last addon update succeeded. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.

The `statuses` object supports the following:

* `conditions_type` - (Optional) The state condition in the current main state of the addon, that is, the reason for entering the main state, there can be multiple reasons, the value contains `Progressing`, `Ok`, `Degraded`,`Unknown`, `ClusterNotRunning`, `CrashLoopBackOff`, `SchedulingFailed`, `NameConflict`, `ResourceCleanupFailed`, `ClusterVersionUpgrading`.
* `phase` - (Optional) The status of addon. the value contains `Creating`, `Running`, `Updating`, `Deleting`, `Failed`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `addons` - The collection of addon query.
    * `cluster_id` - The ID of the Cluster.
    * `config` - The config of addon.
    * `create_time` - Addon creation time. UTC+0 time in standard RFC3339 format.
    * `deploy_mode` - The deploy mode.
    * `deploy_node_type` - The deploy node type.
    * `name` - The name of the cluster.
    * `status` - The status of the addon.
        * `conditions` - The state condition in the current primary state of the cluster, that is, the reason for entering the primary state.
            * `type` - The state condition in the current main state of the addon, that is, the reason for entering the main state, there can be multiple reasons, the value contains `Progressing`, `Ok`, `Degraded`,`Unknown`, `ClusterNotRunning`, `CrashLoopBackOff`, `SchedulingFailed`, `NameConflict`, `ResourceCleanupFailed`, `ClusterVersionUpgrading`.
        * `phase` - The status of addon. the value contains `Creating`, `Running`, `Updating`, `Deleting`, `Failed`.
    * `update_time` - The last time a request was accepted by the addon and executed or completed. UTC+0 time in standard RFC3339 format.
    * `version` - The name of the cluster.
* `total_count` - The total count of addon query.


