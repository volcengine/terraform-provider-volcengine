---
subcategory: "MONGODB"
layout: "volcengine"
page_title: "Volcengine: volcengine_mongodb_instance_parameters"
sidebar_current: "docs-volcengine-datasource-mongodb_instance_parameters"
description: |-
  Use this data source to query detailed information of mongodb instance parameters
---
# volcengine_mongodb_instance_parameters
Use this data source to query detailed information of mongodb instance parameters
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
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_mongodb_instance" "foo" {
  db_engine_version      = "MongoDB_4_0"
  instance_type          = "ReplicaSet"
  super_account_password = "@acc-test-123"
  node_spec              = "mongo.2c4g"
  mongos_node_spec       = "mongo.mongos.2c4g"
  instance_name          = "acc-test-mongo-replica"
  charge_type            = "PostPaid"
  project_name           = "default"
  mongos_node_number     = 32
  shard_number           = 3
  storage_space_gb       = 20
  subnet_id              = volcengine_subnet.foo.id
  zone_id                = data.volcengine_zones.foo.zones[0].id
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_mongodb_instance_parameter" "foo" {
  instance_id     = volcengine_mongodb_instance.foo.id
  parameter_name  = "cursorTimeoutMillis"
  parameter_role  = "Node"
  parameter_value = "600111"
}

data "volcengine_mongodb_instance_parameters" "foo" {
  instance_id     = volcengine_mongodb_instance.foo.id
  parameter_names = "cursorTimeoutMillis"
  parameter_role  = "Node"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The instance ID to query.
* `output_file` - (Optional) File name where to save data source results.
* `parameter_names` - (Optional) The parameter names, support fuzzy query, case insensitive.
* `parameter_role` - (Optional) The node type of instance parameter, valid value contains `Node`, `Shard`, `ConfigServer`, `Mongos`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `instance_parameters` - The collection of parameter query.
    * `checking_code` - The checking code of parameter.
    * `force_modify` - Whether the parameter supports modifying.
    * `force_restart` - Does the new parameter value need to restart the instance to take effect after modification.
    * `parameter_default_value` - The default value of parameter.
    * `parameter_description` - The description of parameter.
    * `parameter_name` - The name of parameter.
    * `parameter_role` - The node type to which the parameter belongs.
    * `parameter_type` - The type of parameter value.
    * `parameter_value` - The value of parameter.
* `parameters` - (**Deprecated**) This field has been deprecated and it is recommended to use instance_parameters. The collection of parameter query.
    * `db_engine_version` - The database engine version.
    * `db_engine` - The database engine.
    * `instance_id` - The instance ID.
    * `instance_parameters` - The list of parameters.
        * `checking_code` - The checking code of parameter.
        * `force_modify` - Whether the parameter supports modifying.
        * `force_restart` - Does the new parameter value need to restart the instance to take effect after modification.
        * `parameter_default_value` - The default value of parameter.
        * `parameter_description` - The description of parameter.
        * `parameter_name` - The name of parameter.
        * `parameter_role` - The node type to which the parameter belongs.
        * `parameter_type` - The type of parameter value.
        * `parameter_value` - The value of parameter.
    * `total` - The total parameters queried.
* `total_count` - The total count of mongodb instance parameter query.


