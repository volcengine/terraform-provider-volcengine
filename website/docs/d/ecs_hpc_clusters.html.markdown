---
subcategory: "ECS"
layout: "volcengine"
page_title: "Volcengine: volcengine_ecs_hpc_clusters"
sidebar_current: "docs-volcengine-datasource-ecs_hpc_clusters"
description: |-
  Use this data source to query detailed information of ecs hpc clusters
---
# volcengine_ecs_hpc_clusters
Use this data source to query detailed information of ecs hpc clusters
## Example Usage
```hcl
data "volcengine_ecs_hpc_clusters" "foo" {
  zone_id = "cn-beijing-a"
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of hpc cluster ids.
* `name_regex` - (Optional) A Name Regex of Resource.
* `name` - (Optional) The name of the hpc cluster.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name of the hpc cluster.
* `tags` - (Optional) Tags.
* `zone_id` - (Optional) The zone id of the hpc cluster.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `hpc_clusters` - The collection of query.
    * `created_at` - The created time of the hpc cluster.
    * `description` - The description of the hpc cluster.
    * `hpc_cluster_id` - The id of the hpc cluster.
    * `id` - The id of the hpc cluster.
    * `name` - The name of the hpc cluster.
    * `project_name` - The project name of the hpc cluster.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `updated_at` - The updated time of the hpc cluster.
    * `vpc_id` - The vpc id of the hpc cluster.
    * `zone_id` - The zone id of the hpc cluster.
* `total_count` - The total count of query.


