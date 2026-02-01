---
subcategory: "MONGODB"
layout: "volcengine"
page_title: "Volcengine: volcengine_mongodb_ssl_states"
sidebar_current: "docs-volcengine-datasource-mongodb_ssl_states"
description: |-
  Use this data source to query detailed information of mongodb ssl states
---
# volcengine_mongodb_ssl_states
Use this data source to query detailed information of mongodb ssl states
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
  mongos_node_number     = 2
  shard_number           = 3
  storage_space_gb       = 20
  subnet_id              = volcengine_subnet.foo.id
  zone_id                = data.volcengine_zones.foo.zones[0].id
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_mongodb_ssl_state" "foo" {
  instance_id = volcengine_mongodb_instance.foo.id
}

data "volcengine_mongodb_ssl_states" "foo" {
  instance_id = volcengine_mongodb_instance.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The mongodb instance ID to query.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `ssl_state` - The collection of mongodb ssl state query.
    * `instance_id` - The mongodb instance id.
    * `is_valid` - Whetehr SSL is valid.
    * `ssl_enable` - Whether SSL is enabled.
    * `ssl_expired_time` - The expire time of SSL.
* `total_count` - The total count of mongodb ssl state query.


