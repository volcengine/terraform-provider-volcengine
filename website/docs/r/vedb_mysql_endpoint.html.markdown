---
subcategory: "VEDB_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_vedb_mysql_endpoint"
sidebar_current: "docs-volcengine-resource-vedb_mysql_endpoint"
description: |-
  Provides a resource to manage vedb mysql endpoint
---
# volcengine_vedb_mysql_endpoint
Provides a resource to manage vedb mysql endpoint
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
  zone_id     = data.volcengine_zones.foo.zones[2].id
  vpc_id      = volcengine_vpc.foo.id
}


resource "volcengine_vedb_mysql_instance" "foo" {
  charge_type         = "PostPaid"
  storage_charge_type = "PostPaid"
  db_engine_version   = "MySQL_8_0"
  db_minor_version    = "3.0"
  node_number         = 2
  node_spec           = "vedb.mysql.x4.large"
  subnet_id           = volcengine_subnet.foo.id
  instance_name       = "tf-test"
  project_name        = "testA"
  tags {
    key   = "tftest"
    value = "tftest"
  }
  tags {
    key   = "tftest2"
    value = "tftest2"
  }
}
data "volcengine_vedb_mysql_instances" "foo" {
  instance_id = volcengine_vedb_mysql_instance.foo.id
}

resource "volcengine_vedb_mysql_endpoint" "foo" {
  endpoint_type               = "Custom"
  instance_id                 = volcengine_vedb_mysql_instance.foo.id
  node_ids                    = [data.volcengine_vedb_mysql_instances.foo.instances[0].nodes[0].node_id, data.volcengine_vedb_mysql_instances.foo.instances[0].nodes[1].node_id]
  read_write_mode             = "ReadWrite"
  endpoint_name               = "tf-test"
  description                 = "tf test"
  master_accept_read_requests = true
  distributed_transaction     = true
  consist_level               = "Session"
  consist_timeout             = 100000
  consist_timeout_action      = "ReadMaster"
}
```
## Argument Reference
The following arguments are supported:
* `endpoint_type` - (Required, ForceNew) Connect endpoint type. The value is fixed as Custom, indicating a custom endpoint.
* `instance_id` - (Required, ForceNew) The id of the instance.
* `node_ids` - (Required) Connect the node IDs associated with the endpoint.The filling rules are as follows:
When the value of ReadWriteMode is ReadWrite, at least two nodes must be passed in, and the master node must be passed in.
When the value of ReadWriteMode is ReadOnly, one or more read-only nodes can be passed in.
* `consist_level` - (Optional) Consistency level. For detailed introduction of consistency level, please refer to consistency level. Value range:
Eventual: eventual consistency.
Session: session consistency.
Global: global consistency.
Description
When the value of ReadWriteMode is ReadWrite, the selectable consistency levels are Eventual, Session (default), and Global.
When the value of ReadWriteMode is ReadOnly, the consistency level is Eventual by default and cannot be changed.
* `consist_timeout_action` - (Optional) Timeout policy after data synchronization timeout of read-only nodes supports the following two policies:
ReturnError: Return SQL error (wait replication complete timeout, please retry).
ReadMaster: Send a request to the master node (default).
Description
 This parameter takes effect only when the value of ConsistLevel is Global or Session.
* `consist_timeout` - (Optional) When there is a large delay, the timeout period for read-only nodes to synchronize the latest data, in us. The value range is from 1us to 100000000us, and the default value is 10000us.
Explanation
 This parameter takes effect only when the value of ConsistLevel is Global or Session.
* `description` - (Optional) Description information for connecting endpoint. The length cannot exceed 200 characters.
* `distributed_transaction` - (Optional) Set whether to enable transaction splitting. For detailed introduction to transaction splitting, please refer to transaction splitting. Value range:
true: Enabled (default).
false: Disabled.
Description
Only when the value of ReadWriteMode is ReadWrite, is enabling transaction splitting supported.
* `endpoint_name` - (Optional) Connect the endpoint name. The setting rules are as follows:
 It cannot start with a number or a hyphen (-).
 It can only contain Chinese characters, letters, numbers, underscores (_), and hyphens (-).
 The length is 1 to 64 characters.
* `master_accept_read_requests` - (Optional) The master node accepts read requests. Value range:
true: (default) After enabling the master node to accept read functions, non-transactional read requests will be sent to the master node or read-only nodes in a load-balanced mode according to the number of active requests.
false: After disabling the master node from accepting read requests, at this time, the master node only accepts transactional read requests, and non-transactional read requests will not be sent to the master node.
Description
Only when the value of ReadWriteMode is ReadWrite, enabling the master node to accept reads is supported.
* `read_write_mode` - (Optional) Endpoint read-write mode. Values:
 ReadWrite: Read and write endpoint.
 ReadOnly: Read-only endpoint (default).

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `endpoint_id` - The id of the endpoint.


## Import
VedbMysqlEndpoint can be imported using the instance id:endpoint id, e.g.
```
$ terraform import volcengine_vedb_mysql_endpoint.default vedbm-iqnh3a7z****:vedbm-2pf2xk5v****-Custom-50yv
```
Note: The master node endpoint only supports modifying the EndpointName and Description parameters. If values are passed in for other parameters, these values will be ignored without generating an error.
The default endpoint does not support modifying the ReadWriteMode, AutoAddNewNodes, and Nodes parameters. If values are passed in for these parameters, these values will be ignored without generating an error.

