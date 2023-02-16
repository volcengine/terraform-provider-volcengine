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
  region_id = "cn-xxx"
}
```
## Argument Reference
The following arguments are supported:
* `region_id` - (Required) The region ID to query.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `specs` - The collection of mongos spec query.
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
        * `spec_name` - The node spec name.
    * `shard_node_specs` - The collection of shard node specs.
        * `cpu_num` - The cpu cores.
        * `max_conn` - The max connections.
        * `max_storage` - The max storage.
        * `mem_in_gb` - The memory in GB.
        * `spec_name` - The shard node spec name.
* `total_count` - The total count of region query.


