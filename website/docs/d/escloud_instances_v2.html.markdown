---
subcategory: "ESCLOUD"
layout: "volcengine"
page_title: "Volcengine: volcengine_escloud_instances_v2"
sidebar_current: "docs-volcengine-datasource-escloud_instances_v2"
description: |-
  Use this data source to query detailed information of escloud instances v2
---
# volcengine_escloud_instances_v2
Use this data source to query detailed information of escloud instances v2
## Example Usage
```hcl
data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  description = "tfdesc"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_escloud_instance_v2" "foo" {
  instance_name       = "acc-test-escloud-instance"
  version             = "V7_10"
  zone_ids            = [data.volcengine_zones.foo.zones[0].id, data.volcengine_zones.foo.zones[1].id, data.volcengine_zones.foo.zones[2].id]
  subnet_id           = volcengine_subnet.foo.id
  enable_https        = false
  admin_password      = "Password@@123"
  charge_type         = "PostPaid"
  auto_renew          = false
  period              = 1
  configuration_code  = "es.standard"
  enable_pure_master  = true
  deletion_protection = false
  project_name        = "default"

  node_specs_assigns {
    type               = "Master"
    number             = 3
    resource_spec_name = "es.x2.medium"
    storage_spec_name  = "es.volume.essd.pl0"
    storage_size       = 20
  }
  node_specs_assigns {
    type               = "Hot"
    number             = 6
    resource_spec_name = "es.x2.medium"
    storage_spec_name  = "es.volume.essd.flexpl-standard"
    storage_size       = 500
    extra_performance {
      throughput = 65
    }
  }
  node_specs_assigns {
    type               = "Kibana"
    number             = 1
    resource_spec_name = "kibana.x2.small"
    storage_spec_name  = ""
    storage_size       = 0
  }

  network_specs {
    type      = "Elasticsearch"
    bandwidth = 1
    is_open   = true
    spec_name = "es.eip.bgp_fixed_bandwidth"
  }

  network_specs {
    type      = "Kibana"
    bandwidth = 1
    is_open   = true
    spec_name = "es.eip.bgp_fixed_bandwidth"
  }

  tags {
    key   = "k1"
    value = "v1"
  }
}

data "volcengine_escloud_instances_v2" "foo" {
  ids = [volcengine_escloud_instance_v2.foo.id]
}
```
## Argument Reference
The following arguments are supported:
* `charge_types` - (Optional) The charge types of instance.
* `ids` - (Optional) A list of instance IDs.
* `instance_names` - (Optional) The names of instance.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name of instance.
* `statuses` - (Optional) The status of instance.
* `tags` - (Optional) The tags of instance.
* `versions` - (Optional) The versions of instance.
* `zone_ids` - (Optional) The available zone IDs of instance.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `values` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `instances` - The collection of query.
    * `cerebro_enabled` - Whether to enable cerebro.
    * `cerebro_private_domain` - The cerebro private domain of instance.
    * `cerebro_public_domain` - The cerebro public domain of instance.
    * `charge_enabled` - The charge status of instance.
    * `cluster_id` - The cluster id of instance.
    * `create_time` - The create time of instance.
    * `deletion_protection` - Whether enable deletion protection for ESCloud instance.
    * `enable_es_private_domain_public` - whether enable es private domain public.
    * `enable_es_private_network` - whether enable es private network.
    * `enable_es_public_network` - whether enable es public network.
    * `enable_kibana_private_domain_public` - whether enable kibana private domain public.
    * `enable_kibana_private_network` - whether enable kibana private network.
    * `enable_kibana_public_network` - whether enable kibana public network.
    * `es_eip_id` - The eip id associated with the instance.
    * `es_eip` - The eip address of instance.
    * `es_inner_endpoint` - The es inner endpoint of instance.
    * `es_private_domain` - The es private domain of instance.
    * `es_private_endpoint` - The es private endpoint of instance.
    * `es_private_ip_whitelist` - The whitelist of es private ip.
    * `es_public_domain` - The es public domain of instance.
    * `es_public_endpoint` - The es public endpoint of instance.
    * `es_public_ip_whitelist` - The whitelist of es public ip.
    * `expire_date` - The expire time of instance.
    * `id` - The id of instance.
    * `instance_configuration` - The configuration of instance.
        * `admin_user_name` - The user name of instance.
        * `charge_type` - The charge type of instance.
        * `cold_node_number` - The node number of cold.
        * `cold_node_resource_spec` - The node resource spec of cold.
            * `cpu` - The cpu info of resource spec.
            * `description` - The description of resource spec.
            * `display_name` - The show name of resource spec.
            * `memory` - The memory info of resource spec.
            * `name` - The name of resource spec.
        * `cold_node_storage_spec` - The node storage spec of cold.
            * `description` - The description of storage spec.
            * `display_name` - The show name of storage spec.
            * `max_size` - The max size of storage spec.
            * `min_size` - The min size of storage spec.
            * `name` - The name of storage spec.
            * `size` - The size of storage spec.
        * `coordinator_node_number` - The node number of coordinator.
        * `coordinator_node_resource_spec` - The node resource spec of coordinator.
            * `cpu` - The cpu info of resource spec.
            * `description` - The description of resource spec.
            * `display_name` - The show name of resource spec.
            * `memory` - The memory info of resource spec.
            * `name` - The name of resource spec.
        * `coordinator_node_storage_spec` - The node storage spec of coordinator.
            * `description` - The description of storage spec.
            * `display_name` - The show name of storage spec.
            * `max_size` - The max size of storage spec.
            * `min_size` - The min size of storage spec.
            * `name` - The name of storage spec.
            * `size` - The size of storage spec.
        * `enable_https` - whether enable https.
        * `enable_pure_master` - Whether enable pure master.
        * `hot_node_number` - The node number of hot.
        * `hot_node_resource_spec` - The node resource spec of hot.
            * `cpu` - The cpu info of resource spec.
            * `description` - The description of resource spec.
            * `display_name` - The show name of resource spec.
            * `memory` - The memory info of resource spec.
            * `name` - The name of resource spec.
        * `hot_node_storage_spec` - The node storage spec of hot.
            * `description` - The description of storage spec.
            * `display_name` - The show name of storage spec.
            * `max_size` - The max size of storage spec.
            * `min_size` - The min size of storage spec.
            * `name` - The name of storage spec.
            * `size` - The size of storage spec.
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
        * `warm_node_number` - The node number of warm.
        * `warm_node_resource_spec` - The node resource spec of warm.
            * `cpu` - The cpu info of resource spec.
            * `description` - The description of resource spec.
            * `display_name` - The show name of resource spec.
            * `memory` - The memory info of resource spec.
            * `name` - The name of resource spec.
        * `warm_node_storage_spec` - The node storage spec of warm.
            * `description` - The description of storage spec.
            * `display_name` - The show name of storage spec.
            * `max_size` - The max size of storage spec.
            * `min_size` - The min size of storage spec.
            * `name` - The name of storage spec.
            * `size` - The size of storage spec.
        * `zone_id` - The zoneId of instance.
        * `zone_number` - The zone number of instance.
    * `instance_id` - The id of instance.
    * `kibana_eip_id` - The eip id associated with kibana.
    * `kibana_eip` - The eip address of kibana.
    * `kibana_private_domain` - The kibana private domain of instance.
    * `kibana_private_ip_whitelist` - The whitelist of kibana private ip.
    * `kibana_public_domain` - The kibana public domain of instance.
    * `kibana_public_ip_whitelist` - The whitelist of kibana public ip.
    * `main_zone_id` - The main zone id of instance.
    * `maintenance_day` - The maintenance day of instance.
    * `maintenance_time` - The maintenance time of instance.
    * `nodes` - The nodes info of instance.
        * `is_cold` - Is cold node.
        * `is_coordinator` - Is coordinator node.
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
        * `restart_number` - The restart times of node.
        * `start_time` - The start time of node.
        * `status` - The status of node.
        * `storage_spec` - The node storage spec of master.
            * `description` - The description of storage spec.
            * `display_name` - The show name of storage spec.
            * `max_size` - The max size of storage spec.
            * `min_size` - The min size of storage spec.
    * `plugins` - The plugin info of instance.
        * `description` - The description of plugin.
        * `plugin_name` - The name of plugin.
        * `status` - The status of plugin.
        * `version` - The version of plugin.
    * `status` - The status of instance.
    * `support_code_node` - Whether support code node.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `total_nodes` - The total nodes of instance.
    * `user_id` - The user id of instance.
* `total_count` - The total count of query.


