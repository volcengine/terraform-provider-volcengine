---
subcategory: "MONGODB"
layout: "volcengine"
page_title: "Volcengine: volcengine_mongodb_endpoints"
sidebar_current: "docs-volcengine-datasource-mongodb_endpoints"
description: |-
  Use this data source to query detailed information of mongodb endpoints
---
# volcengine_mongodb_endpoints
Use this data source to query detailed information of mongodb endpoints
## Example Usage
```hcl
data "volcengine_mongodb_endpoints" "foo" {
  instance_id = "mongo-shard-xxx"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Optional) The instance ID to query.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `endpoints` - The collection of mongodb endpoints query.
    * `db_addresses` - The list of mongodb addresses.
        * `address_domain` - The domain of mongodb connection.
        * `address_ip` - The IP of mongodb connection.
        * `address_port` - The port of mongodb connection.
        * `address_type` - The connection type of mongodb.
        * `eip_id` - The EIP ID bound to the instance's public network address.
        * `node_id` - The node ID.
    * `endpoint_id` - The ID of endpoint.
    * `endpoint_str` - The endpoint information.
    * `endpoint_type` - The node type corresponding to the endpoint.
    * `network_type` - The network type of endpoint.
    * `object_id` - The object ID corresponding to the endpoint.
    * `subnet_id` - The subnet ID.
    * `vpc_id` - The VPC ID.
* `total_count` - The total count of mongodb endpoint query.


