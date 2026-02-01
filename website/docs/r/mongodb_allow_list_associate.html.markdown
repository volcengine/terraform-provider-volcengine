---
subcategory: "MONGODB"
layout: "volcengine"
page_title: "Volcengine: volcengine_mongodb_allow_list_associate"
sidebar_current: "docs-volcengine-resource-mongodb_allow_list_associate"
description: |-
  Provides a resource to manage mongodb allow list associate
---
# volcengine_mongodb_allow_list_associate
Provides a resource to manage mongodb allow list associate
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

resource "volcengine_mongodb_allow_list" "foo" {
  allow_list_name = "acc-test"
  allow_list_desc = "acc-test"
  allow_list_type = "IPv4"
  allow_list      = "10.1.1.3,10.2.3.0/24,10.1.1.1"
}

resource "volcengine_mongodb_allow_list_associate" "foo" {
  allow_list_id = volcengine_mongodb_allow_list.foo.id
  instance_id   = volcengine_mongodb_instance.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `allow_list_id` - (Required, ForceNew) Id of allow list to associate.
* `instance_id` - (Required, ForceNew) Id of instance to associate.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
mongodb allow list associate can be imported using the instanceId:allowListId, e.g.
```
$ terraform import volcengine_mongodb_allow_list_associate.default mongo-replica-e405f8e2****:acl-d1fd76693bd54e658912e7337d5b****
```

