---
subcategory: "BIOOS"
layout: "volcengine"
page_title: "Volcengine: volcengine_bioos_clusters"
sidebar_current: "docs-volcengine-datasource-bioos_clusters"
description: |-
  Use this data source to query detailed information of bioos clusters
---
# volcengine_bioos_clusters
Use this data source to query detailed information of bioos clusters
## Example Usage
```hcl
data "volcengine_bioos_clusters" "default" {

}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of cluster ids.
* `output_file` - (Optional) File name where to save data source results.
* `public` - (Optional) whether it is a public cluster.
* `status` - (Optional) The status of the clusters.
* `type` - (Optional) The type of the clusters.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `items` - The list of cluster.
    * `bound` - Whether there is a bound workspace.
    * `description` - The description of the cluster.
    * `external_config_filesystem` - Workflow computing engine file system (currently supports tos, local).
    * `external_config_jupyterhub_endpoint` - The endpoint of jupyterhub.
    * `external_config_jupyterhub_jwt_secret` - The jupyterhub jwt secret.
    * `external_config_resource_scheduler` - External Resource Scheduler.
    * `external_config_wes_endpoint` - The WES endpoint.
    * `id` - The id of the bioos cluster.
    * `name` - The name of the cluster.
    * `public` - whether it is a public cluster.
    * `start_time` - The start time of the cluster.
    * `stopped_time` - The end time of the cluster.
    * `vke_config_id` - The id of the vke cluster id.
    * `vke_config_storage_class` - The name of the StorageClass that the vke cluster has installed.
* `total_count` - The total count of Vpc query.


