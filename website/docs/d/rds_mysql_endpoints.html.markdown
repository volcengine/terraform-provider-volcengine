---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_mysql_endpoints"
sidebar_current: "docs-volcengine-datasource-rds_mysql_endpoints"
description: |-
  Use this data source to query detailed information of rds mysql endpoints
---
# volcengine_rds_mysql_endpoints
Use this data source to query detailed information of rds mysql endpoints
## Example Usage
```hcl
data "volcengine_rds_mysql_endpoints" "foo" {
  instance_id = "mysql-38c3d4f05f6e"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The id of the mysql instance.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `endpoints` - The collection of query.
    * `addresses` - Address list.
        * `dns_visibility` - DNS Visibility.
        * `domain` - Connect domain name.
        * `eip_id` - The ID of the EIP, only valid for Public addresses.
        * `ip_address` - The IP Address.
        * `network_type` - Network address type, temporarily Private, Public, PublicService.
        * `port` - The Port.
        * `subnet_id` - Subnet ID, valid only for private addresses.
    * `auto_add_new_nodes` - When the terminal type is read-write terminal or read-only terminal, it supports setting whether new nodes are automatically added.
    * `description` - The description of the mysql endpoint.
    * `enable_read_only` - Whether global read-only is enabled, value: Enable: Enable. Disable: Disabled.
    * `enable_read_write_splitting` - Whether read-write separation is enabled, value: Enable: Enable. Disable: Disabled.
    * `endpoint_id` - The id of the mysql endpoint.
    * `endpoint_name` - The name of the mysql endpoint.
    * `endpoint_type` - The endpoint type of the mysql endpoint.
    * `id` - The id of the mysql endpoint.
    * `read_only_node_weight` - The list of nodes configured by the connection terminal and the corresponding read-only weights.
        * `node_id` - The ID of the node.
        * `node_type` - The type of the node.
        * `weight` - The weight of the node.
    * `read_write_mode` - The read write mode.
* `total_count` - The total count of query.


