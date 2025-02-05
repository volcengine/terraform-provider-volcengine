---
subcategory: "MONGODB"
layout: "volcengine"
page_title: "Volcengine: volcengine_mongodb_account"
sidebar_current: "docs-volcengine-resource-mongodb_account"
description: |-
  Provides a resource to manage mongodb account
---
# volcengine_mongodb_account
Provides a resource to manage mongodb account
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
  zone_ids               = [data.volcengine_zones.foo.zones[0].id]
  db_engine_version      = "MongoDB_4_0"
  instance_type          = "ReplicaSet"
  node_spec              = "mongo.2c4g"
  storage_space_gb       = 20
  subnet_id              = volcengine_subnet.foo.id
  instance_name          = "acc-test-mongodb-replica"
  charge_type            = "PostPaid"
  super_account_password = "93f0cb0614Aab12"
  project_name           = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_mongodb_account" "foo" {
  instance_id      = volcengine_mongodb_instance.foo.id
  account_name     = "acc-test-mongodb-account"
  auth_db          = "admin"
  account_password = "93f0cb0614Aab12"
  account_desc     = "acc-test"
  account_privileges {
    db_name    = "admin"
    role_names = ["userAdmin", "clusterMonitor"]
  }
  account_privileges {
    db_name    = "config"
    role_names = ["read"]
  }
  account_privileges {
    db_name    = "local"
    role_names = ["read"]
  }
}
```
## Argument Reference
The following arguments are supported:
* `account_name` - (Required, ForceNew) The name of the mongodb account.
* `account_password` - (Required) The password of the mongodb account. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `instance_id` - (Required, ForceNew) The id of the mongodb instance.
* `account_desc` - (Optional) The description of the mongodb account.
* `account_privileges` - (Optional) The privilege information of account.
* `auth_db` - (Optional, ForceNew) The database of the mongodb account.

The `account_privileges` object supports the following:

* `db_name` - (Required) The name of database.
* `role_names` - (Required) The role names of the account.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `account_type` - The type of the account.


## Import
MongodbAccount can be imported using the instance_id:account_name, e.g.
```
$ terraform import volcengine_mongodb_account.default resource_id
```

