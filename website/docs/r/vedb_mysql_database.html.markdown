---
subcategory: "VEDB_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_vedb_mysql_database"
sidebar_current: "docs-volcengine-resource-vedb_mysql_database"
description: |-
  Provides a resource to manage vedb mysql database
---
# volcengine_vedb_mysql_database
Provides a resource to manage vedb mysql database
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

resource "volcengine_vedb_mysql_database" "foo" {
  db_name     = "tf-table"
  instance_id = volcengine_vedb_mysql_instance.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `db_name` - (Required, ForceNew) The name of the database. Naming rules:
 Unique name. Start with a lowercase letter and end with a letter or number. The length is within 2 to 64 characters.
 Consist of lowercase letters, numbers, underscores (_), or hyphens (-).
 The name cannot contain certain reserved words.
* `instance_id` - (Required, ForceNew) The id of the instance.
* `character_set_name` - (Optional, ForceNew) Database character set: utf8mb4 (default), utf8, latin1, ascii.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
VedbMysqlDatabase can be imported using the instance id and database name, e.g.
```
$ terraform import volcengine_vedb_mysql_database.default vedbm-r3xq0zdl****:testdb

```

