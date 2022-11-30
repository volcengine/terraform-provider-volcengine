---
subcategory: "ESCLOUD"
layout: "volcengine"
page_title: "Volcengine: volcengine_escloud_instances"
sidebar_current: "docs-volcengine-datasource-escloud_instances"
description: |-
  Use this data source to query detailed information of escloud instances
---
# volcengine_escloud_instances
Use this data source to query detailed information of escloud instances
## Example Usage
```hcl
data "volcengine_escloud_instances" "default" {
  ids      = ["d3gftqjvnah74eie"]
  statuses = ["Running"]
}
```
## Argument Reference
The following arguments are supported:
* `charge_types` - (Optional) The charge types of instance.
* `ids` - (Optional) A list of instance IDs.
* `names` - (Optional) The names of instance.
* `output_file` - (Optional) File name where to save data source results.
* `statuses` - (Optional) The list status of instance.
* `versions` - (Optional) The versions of instance.
* `zone_ids` - (Optional) The available zone IDs of instance.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `instances` - The collection of instance query.
    * `charge_enabled` - The charge status of instance.
    * `create_time` - The create time of instance.
    * `enable_es_private_network` - whether enable es private network.
    * `enable_es_public_network` - whether enable es public network.
    * `enable_kibana_private_network` - whether enable kibana private network.
    * `enable_kibana_public_network` - whether enable kibana public network.
    * `es_inner_endpoint` - The es inner endpoint of instance.
    * `es_private_domain` - The es private domain of instance.
    * `es_private_endpoint` - The es private endpoint of instance.
    * `es_public_domain` - The es public domain of instance.
    * `es_public_endpoint` - The es public endpoint of instance.
    * `expire_date` - The expire time of instance.
    * `id` - The Id of instance.
    * `instance_configuration` - The configuration of instance.
        * `admin_user_name` - The user name of instance.
        * `charge_type` - The charge type of instance.
        * `enable_https` - whether enable https.
        * `enable_pure_master` - Whether enable pure master.
        * `hot_node_number` - The node number of host.
        * `hot_node_resource_spec` - The node resource spec of host.
            * `cpu` - The cpu info of resource spec.
            * `description` - The description of resource spec.
            * `display_name` - The show name of resource spec.
            * `memory` - The memory info of resource spec.
            * `name` - The name of resource spec.
        * `hot_node_storage_spec` - The node storage spec of host.
            * `description` - The description of storage spec.
            * `display_name` - The show name of storage spec.
            * `max_size` - The max size of storage spec.
            * `min_size` - The min size of storage spec.
            * `name` - The name of storage spec.
            * `size` - The size of storage spec.
            * `type` - The type of storage spec.
        * `instance_name` - The name of instance.
        * `kibana_node_number` - The node number of kibana.
        * `kibana_node_resource_spec` - The node resource spec of kibana.
            * `cpu` - The cpu info of resource spec.
            * `description` - The description of resource spec.
            * `display_name` - The show name of resource spec.
            * `memory` - The memory info of resource spec.
            * `name` - The name of resource spec.
        * `kibana_node_storage_spec` - The node storage spec of kibana.
            * `description` - The description of storage spec.
            * `display_name` - The show name of storage spec.
            * `max_size` - The max size of storage spec.
            * `min_size` - The min size of storage spec.
            * `name` - The name of storage spec.
            * `size` - The size of storage spec.
            * `type` - The type of storage spec.
        * `master_node_number` - The node number of master.
        * `master_node_resource_spec` - The node resource spec of master.
            * `cpu` - The cpu info of resource spec.
            * `description` - The description of resource spec.
            * `display_name` - The show name of resource spec.
            * `memory` - The memory info of resource spec.
            * `name` - The name of resource spec.
        * `master_node_storage_spec` - The node storage spec of master.
            * `description` - The description of storage spec.
            * `display_name` - The show name of storage spec.
            * `max_size` - The max size of storage spec.
            * `min_size` - The min size of storage spec.
            * `name` - The name of storage spec.
            * `size` - The size of storage spec.
            * `type` - The type of storage spec.
        * `period` - The period of project.
        * `project_name` - The name of project.
        * `region_id` - The region info of instance.
        * `subnet` - The subnet info.
            * `subnet_id` - The id of subnet.
            * `subnet_name` - The name of subnet.
        * `version` - The version of instance.
        * `vpc` - The vpc info.
            * `vpc_id` - The id of vpc.
            * `vpc_name` - The name of vpc.
        * `zone_id` - The zoneId of instance.
        * `zone_number` - The zone number of instance.
    * `instance_id` - The Id of instance.
    * `kibana_private_domain` - The kibana private domain of instance.
    * `kibana_public_domain` - The kibana public domain of instance.
    * `maintenance_day` - The maintenance day of instance.
    * `maintenance_time` - The maintenance time of instance.
    * `namespace` - The namespace of instance.
    * `nodes` - The nodes info of instance.
        * `is_cold` - Is cold node.
        * `is_hot` - Is hot node.
        * `is_kibana` - Is kibana node.
        * `is_master` - Is master node.
        * `is_warm` - Is warm node.
        * `node_display_name` - The show name of node.
        * `node_name` - The name of node.
        * `resource_spec` - The node resource spec of master.
            * `cpu` - The cpu info of resource spec.
            * `description` - The description of resource spec.
            * `display_name` - The show name of resource spec.
            * `memory` - The memory info of resource spec.
            * `name` - The name of resource spec.
        * `restart_number` - The restart times of node.
        * `start_time` - The start time of node.
        * `status` - The status of node.
        * `storage_spec` - The node storage spec of master.
            * `description` - The description of storage spec.
            * `display_name` - The show name of storage spec.
            * `max_size` - The max size of storage spec.
            * `min_size` - The min size of storage spec.
            * `name` - The name of storage spec.
            * `size` - The size of storage spec.
            * `type` - The type of storage spec.
    * `plugins` - The plugin info of instance.
        * `description` - The description of plugin.
        * `plugin_name` - The name of plugin.
        * `version` - The version of plugin.
    * `status` - The status of instance.
    * `total_nodes` - The total nodes of instance.
    * `user_id` - The user id of instance.
* `total_count` - The total count of instance query.


