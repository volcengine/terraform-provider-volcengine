---
subcategory: "VMP"
layout: "volcengine"
page_title: "Volcengine: volcengine_vmp_integration_tasks"
sidebar_current: "docs-volcengine-datasource-vmp_integration_tasks"
description: |-
  Use this data source to query detailed information of vmp integration tasks
---
# volcengine_vmp_integration_tasks
Use this data source to query detailed information of vmp integration tasks
## Example Usage
```hcl
data "volcengine_vmp_integration_tasks" "foo" {
  ids = ["xxxxxx"]
}
```
## Argument Reference
The following arguments are supported:
* `environment` - (Optional) The deployment environment. Valid values: `Vke` or `Managed`.
* `ids` - (Optional) A list of integration task IDs.
* `name` - (Optional) The name of the integration task.
* `statuses` - (Optional) The status of the integration task. Valid values: `Creating`, `Updating`, `Active`, `Error`, `Deleting`.
* `vke_cluster_ids` - (Optional) The ID of the VKE cluster.
* `workspace_id` - (Optional) The workspace ID.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `integration_tasks` - The list of integration tasks.
    * `environment` - The deployment environment.
    * `id` - The ID of the integration task.
    * `name` - The name of the integration task.
    * `status` - The status of the integration task.
    * `type` - The type of the integration task.
    * `vke_cluster_ids` - The ID of the VKE cluster.
    * `vke_cluster_info` - The information of the VKE cluster.
        * `name` - The name of the VKE cluster.
        * `status` - The status of the VKE cluster.
    * `workspace_id` - The workspace ID.


