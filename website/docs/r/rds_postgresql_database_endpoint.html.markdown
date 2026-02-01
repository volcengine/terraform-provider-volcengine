---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_database_endpoint"
sidebar_current: "docs-volcengine-resource-rds_postgresql_database_endpoint"
description: |-
  Provides a resource to manage rds postgresql database endpoint
---
# volcengine_rds_postgresql_database_endpoint
Provides a resource to manage rds postgresql database endpoint
## Example Usage
```hcl
resource "volcengine_rds_postgresql_database_endpoint" "cluster" {
  endpoint_id                      = "postgres-72715e0d9f58-cluster"
  endpoint_name                    = "默认终端"
  endpoint_type                    = "Cluster"
  instance_id                      = "postgres-72715e0d9f58"
  read_only_node_distribution_type = "Custom"
  read_only_node_max_delay_time    = 40
  read_write_mode                  = "ReadWrite"
  read_write_proxy_connection      = 20
  write_node_halt_writing          = false
  read_write_splitting             = true
  read_only_node_weight {
    node_id   = null
    node_type = "Primary"
    weight    = 200
  }
  dns_visibility = true
  port           = 5432
}

resource "volcengine_rds_postgresql_database_endpoint" "example" {
  instance_id     = "postgres-72715e0d9f58"
  endpoint_name   = "tf-test"
  endpoint_type   = "Custom"
  nodes           = "Primary"
  read_write_mode = "ReadWrite"
}
```
## Argument Reference
The following arguments are supported:
* `endpoint_name` - (Required) The name of the connection endpoint. If not provided, the connection endpoint will be automatically named Custom Endpoint.
* `instance_id` - (Required, ForceNew) The ID of the RDS PostgreSQL instance.
* `dns_visibility` - (Optional) Whether to enable public network resolution. false: Default value, Volcano Engine private network resolution. true: Volcano Engine private network and public network resolution. Do not set this field when creating a endpoint.
* `domain_prefix` - (Optional) Private address domain prefix to modify. Do not set this field when creating a endpoint.
* `endpoint_id` - (Optional) The ID of the connection endpoint. The ID of the default endpoint is in the form of instance_id-cluster.
* `endpoint_type` - (Optional) Type of the connection endpoint. Valid values: `Custom`(custom endpoint), `Cluster`(default endpoint). When create a new endpoint, the value must be `Custom`. The default cluster endpoint does not support creation; you can use import to bring it under Terraform management.
* `global_read_only` - (Optional) Whether to enable the global read-only mode for the instance. There is no default value. If no value is passed, the request will be ignored. Do not set this field when creating a endpoint.
* `nodes` - (Optional) List of nodes configured for the connection endpoint. Required when EndpointType is Custom. The primary node does not need to pass the node ID; it is sufficient to pass the string "Primary".
* `port` - (Optional) Private address port to modify. The value range is 1000~65534. Do not set this field when creating a endpoint.
* `read_only_node_distribution_type` - (Optional) Read-only weight distribution mode, Default or Custom. Default: Standard weight allocation. Custom: Custom weight allocation.
* `read_only_node_max_delay_time` - (Optional) The maximum delay threshold for read-only nodes. When the delay time of a read-only node exceeds this value, read traffic will not be sent to that node. The value range is 0~3600. Default value is 30 seconds.
* `read_only_node_weight` - (Optional) Custom read weight allocation. This parameter needs to be set when the value of read_only_node_distribution_type is Custom.
* `read_write_mode` - (Optional) ReadWrite or ReadOnly.
* `read_write_proxy_connection` - (Optional) The number of proxy connections set for the terminal after enabling read-write separation. The minimum value of the proxy connection count is 20.
* `read_write_splitting` - (Optional) Whether to enable read-write separation. Only default endpoint supports this feature.
* `write_node_halt_writing` - (Optional) Whether to prohibit the terminal from sending write requests to the write node. To avoid having no available connection endpoints to carry write operations, this configuration can only be enabled when the instance has other read-write endpoints.

The `read_only_node_weight` object supports the following:

* `node_id` - (Optional) A read-only node requires passing in the NodeId. A primary node does not need to pass in the NodeId.
* `node_type` - (Optional) Node type. Primary or ReadOnly.
* `weight` - (Optional) Custom read weight allocation. Increases by 100, with a maximum value of 40000. Weights cannot all be set to 0.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
RdsPostgresqlDatabaseEndpoint can be imported using the id, e.g.
```
$ terraform import volcengine_rds_postgresql_database_endpoint.default resource_id
```

