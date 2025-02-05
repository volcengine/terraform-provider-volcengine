---
subcategory: "MONGODB"
layout: "volcengine"
page_title: "Volcengine: volcengine_mongodb_specs"
sidebar_current: "docs-volcengine-datasource-mongodb_specs"
description: |-
  Use this data source to query detailed information of mongodb specs
---
# volcengine_mongodb_specs
Use this data source to query detailed information of mongodb specs
## Example Usage
```hcl
data "volcengine_mongodb_specs" "foo" {
  #region_id="cn-xxx" //选填
}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.
* `region_id` - (Optional) The region ID to query.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `specs` - A list of supported node specification information for MongoDB instances.
    * `config_server_node_specs` - The collection of config server node specs.
        * `cpu_num` - The cpu cores.
        * `max_conn` - The max connections.
        * `max_storage` - The max storage.
        * `mem_in_gb` - The memory in GB.
        * `min_storage` - The min storage.
        * `spec_name` - The shard node spec name.
    * `mongos_node_specs` - The collection of mongos node specs.
        * `cpu_num` - The max cpu cores.
        * `max_conn` - The max connections.
        * `mem_in_gb` - The memory in GB.
        * `spec_name` - The mongos node spec name.
    * `node_specs` - The collection of node specs.
        * `cpu_num` - The cpu cores.
        * `max_conn` - The max connections.
        * `max_storage` - The max storage.
        * `mem_in_db` - The memory in GB.
        * `min_storage` - The min storage.
        * `spec_name` - The node spec name.
    * `shard_node_specs` - The collection of shard node specs.
        * `cpu_num` - The cpu cores.
        * `max_conn` - The max connections.
        * `max_storage` - The max storage.
        * `mem_in_gb` - The memory in GB.
        * `min_storage` - The min storage.
        * `spec_name` - The shard node spec name.
* `total_count` - The total count of region query.


