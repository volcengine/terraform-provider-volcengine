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
  instance_type          = "ShardedCluster"
  super_account_password = "@acc-test-123"
  node_spec              = "mongo.shard.1c2g"
  mongos_node_spec       = "mongo.mongos.1c2g"
  instance_name          = "acc-test-mongo-shard"
  charge_type            = "PostPaid"
  project_name           = "default"
  mongos_node_number     = 2
  shard_number           = 2
  storage_space_gb       = 20
  subnet_id              = volcengine_subnet.foo.id
  zone_id                = data.volcengine_zones.foo.zones[0].id
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_eip_address" "foo" {
  billing_type = "PostPaidByBandwidth"
  bandwidth    = 1
  isp          = "ChinaUnicom"
  name         = "acc-eip-${count.index}"
  description  = "acc-test"
  project_name = "default"
  count        = 2
}

resource "volcengine_mongodb_endpoint" "foo_public" {
  instance_id     = volcengine_mongodb_instance.foo.id
  network_type    = "Public"
  object_id       = volcengine_mongodb_instance.foo.mongos_id
  mongos_node_ids = [volcengine_mongodb_instance.foo.mongos[0].mongos_node_id, volcengine_mongodb_instance.foo.mongos[1].mongos_node_id]
  eip_ids         = volcengine_eip_address.foo[*].id
}

resource "volcengine_mongodb_endpoint" "foo_private" {
  instance_id  = volcengine_mongodb_instance.foo.id
  network_type = "Private"
  object_id    = volcengine_mongodb_instance.foo.config_servers_id
}

data "volcengine_mongodb_endpoints" "foo" {
  instance_id = volcengine_mongodb_instance.foo.id
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


