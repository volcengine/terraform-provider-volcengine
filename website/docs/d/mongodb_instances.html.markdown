---
subcategory: "MONGODB"
layout: "volcengine"
page_title: "Volcengine: volcengine_mongodb_instances"
sidebar_current: "docs-volcengine-datasource-mongodb_instances"
description: |-
  Use this data source to query detailed information of mongodb instances
---
# volcengine_mongodb_instances
Use this data source to query detailed information of mongodb instances
## Example Usage
```hcl
data "volcengine_mongodb_instances" "foo" {
  instance_id = "mongo-replica-xxx"
}
```
## Argument Reference
The following arguments are supported:
* `create_end_time` - (Optional) The end time of creation to query.
* `create_start_time` - (Optional) The start time of creation to query.
* `db_engine_version` - (Optional) The version of db engine to query, valid value contains `MongoDB_4_0`.
* `db_engine` - (Optional) The db engine to query, valid value contains `MongoDB`.
* `instance_id` - (Optional) The instance ID to query.
* `instance_name` - (Optional) The instance name to query.
* `instance_status` - (Optional) The instance status to query.
* `instance_type` - (Optional) The type of instance to query, the valid value contains `ReplicaSet` or `ShardedCluster`.
* `name_regex` - (Optional) A Name Regex of DB instance.
* `output_file` - (Optional) File name where to save data source results.
* `tags` - (Optional) Tags.
* `update_end_time` - (Optional) The end time of update to query.
* `update_start_time` - (Optional) The start time of update to query.
* `vpc_id` - (Optional) The vpc id of instance to query.
* `zone_id` - (Optional) The zone ID to query.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `instances` - The collection of mongodb instances query.
    * `auto_renew` - Whether to enable automatic renewal.
    * `charge_status` - The charge status.
    * `charge_type` - The charge type of instance.
    * `closed_time` - The planned close time.
    * `config_servers_id` - The ID of config servers.
    * `config_servers` - The list of config servers.
        * `config_server_node_id` - The config server node ID.
        * `node_role` - The config server node role.
        * `node_status` - The config server node status.
        * `total_memory_gb` - The total memory in GB.
        * `total_vcpu` - The total vCPU.
        * `used_memory_gb` - The used memory in GB.
        * `used_vcpu` - The used vCPU.
        * `zone_id` - The zone ID of node.
    * `create_time` - The creation time of instance.
    * `db_engine_version_str` - The version string of database engine.
    * `db_engine_version` - The version of database engine.
    * `db_engine` - The db engine.
    * `expired_time` - The expired time of instance.
    * `instance_id` - The instance ID.
    * `instance_name` - The instance name.
    * `instance_status` - The instance status.
    * `instance_type` - The instance type.
    * `mongos_id` - The ID of mongos.
    * `mongos` - The list of mongos.
        * `mongos_node_id` - The mongos node ID.
        * `node_spec` - The node spec.
        * `node_status` - The node status.
        * `total_memory_gb` - The total memory in GB.
        * `total_vcpu` - The total vCPU.
        * `used_memory_gb` - The used memory in GB.
        * `used_vcpu` - The used vCPU.
        * `zone_id` - The zone ID of node.
    * `nodes` - The node information.
        * `node_delay_time` - The master-slave delay time.
        * `node_id` - The node ID.
        * `node_role` - The node role.
        * `node_spec` - The node spec.
        * `node_status` - The node status.
        * `total_memory_gb` - The total memory in GB.
        * `total_storage_gb` - The total storage in GB.
        * `total_vcpu` - The total vCPU.
        * `used_memory_gb` - The used memory in GB.
        * `used_storage_gb` - The used storage in GB.
        * `used_vcpu` - The used vCPU.
        * `zone_id` - The zone ID of node.
    * `project_name` - The project name to which the instance belongs.
    * `reclaim_time` - The planned reclaim time of instance.
    * `shards` - The list of shards.
        * `nodes` - The node information.
            * `node_delay_time` - The master-slave delay time.
            * `node_id` - The node ID.
            * `node_role` - The nod role.
            * `node_spec` - The node spec.
            * `node_status` - The node status.
            * `total_memory_gb` - The total memory in GB.
            * `total_storage_gb` - The total storage in GB.
            * `total_vcpu` - The total vCPU.
            * `used_memory_gb` - The used memory in GB.
            * `used_storage_gb` - The used storage in GB.
            * `used_vcpu` - The used vCPU.
            * `zone_id` - The zone ID of node.
        * `shard_id` - The shard ID.
    * `ssl_enable` - Whether ssl enabled.
    * `ssl_expire_time` - The ssl expire time.
    * `ssl_is_valid` - Whether ssl is valid.
    * `storage_type` - The storage type of instance.
    * `subnet_id` - The subnet id of instance.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `update_time` - The update time of instance.
    * `vpc_id` - The vpc ID.
    * `zone_id` - The zone ID of instance.
* `total_count` - The total count of mongodb instances query.


