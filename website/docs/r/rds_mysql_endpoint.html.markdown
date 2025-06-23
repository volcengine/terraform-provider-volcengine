---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_mysql_endpoint"
sidebar_current: "docs-volcengine-resource-rds_mysql_endpoint"
description: |-
  Provides a resource to manage rds mysql endpoint
---
# volcengine_rds_mysql_endpoint
Provides a resource to manage rds mysql endpoint
## Example Usage
```hcl
resource "volcengine_rds_mysql_endpoint" "foo" {
  instance_id                      = "mysql-b51d37110dd1"
  endpoint_name                    = "tf-test-1"
  read_write_mode                  = "ReadWrite"
  description                      = "tf-test-1"
  nodes                            = ["Primary"]
  auto_add_new_nodes               = true
  read_write_spliting              = true
  read_only_node_max_delay_time    = 30
  read_only_node_distribution_type = "RoundRobinAuto"

  read_only_node_weight {
    node_type = "Primary"
    weight    = 100
  }
  dns_visibility = false
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required, ForceNew) The id of the mysql instance.
* `nodes` - (Required) List of node IDs configured for the endpoint. Required when EndpointType is Custom. To add a master node to the terminal, there is no need to fill in the master node ID, just fill in `Primary`.
* `auto_add_new_nodes` - (Optional) When the terminal type is a read-write terminal or a read-only terminal, support is provided for setting whether new nodes are automatically added. The values are:
true: Automatically add.
false: Do not automatically add (default).
* `description` - (Optional) The description of the endpoint.
* `dns_visibility` - (Optional) Values:
false: Volcano Engine private network resolution (default).
true: Volcano Engine private and public network resolution.
* `domain` - (Optional) Connection address, Please note that the connection address can only modify the prefix. In one call, it is not possible to modify both the connection address prefix and the port at the same time.
* `endpoint_id` - (Optional, ForceNew) The id of the endpoint. Import an exist endpoint, usually for import a default endpoint generated with instance creating.
* `endpoint_name` - (Optional) The name of the endpoint.
* `port` - (Optional) The port. Cannot modify public network port. In one call, it is not possible to modify both the connection address prefix and the port at the same time.
* `read_only_node_distribution_type` - (Optional) Read weight distribution mode. This parameter needs to be passed in when the read-write separation setting is true. When used as a request parameter in the CreateDBEndpoint and ModifyDBEndpoint interfaces, the value range is as follows: LoadSchedule: Load scheduling. RoundRobinCustom: Polling scheduling with custom weights. RoundRobinAuto: Polling scheduling with automatically allocated weights.
* `read_only_node_max_delay_time` - (Optional) The maximum delay threshold for read-only nodes, when the delay time of a read-only node exceeds this value, the read traffic will not be sent to that node, unit: seconds. Value range: 0~3600. Default value: 30.
* `read_only_node_weight` - (Optional) Customize read weight distribution, that is, pass in the read request weight of the master node and read-only nodes. It increases by 100 and the maximum value is 10000. When the ReadOnlyNodeDistributionType value is Custom, this parameter needs to be passed in.
* `read_write_mode` - (Optional) Reading and writing mode: ReadWrite, ReadOnly(Default).
* `read_write_spliting` - (Optional) Whether to enable read-write splitting. Values: true: Yes. Default value. false: No.

The `read_only_node_weight` object supports the following:

* `weight` - (Required) The read weight of the node increases by 100, with a maximum value of 10000.
* `node_id` - (Optional) Read-only nodes require NodeId to be passed, while primary nodes do not require it.
* `node_type` - (Optional) The primary node needs to pass in the NodeType as Primary, while the read-only node does not need to pass it in.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
RdsMysqlEndpoint can be imported using the instance id and endpoint id, e.g.
```
$ terraform import volcengine_rds_mysql_endpoint.default mysql-3c25f219***:mysql-3c25f219****-custom-eeb5
```

