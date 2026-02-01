---
subcategory: "VEDB_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_vedb_mysql_databases"
sidebar_current: "docs-volcengine-datasource-vedb_mysql_databases"
description: |-
  Use this data source to query detailed information of vedb mysql databases
---
# volcengine_vedb_mysql_databases
Use this data source to query detailed information of vedb mysql databases
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

data "volcengine_vedb_mysql_databases" "foo" {
  db_name     = volcengine_vedb_mysql_database.foo.db_name
  instance_id = volcengine_vedb_mysql_instance.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The instance id.
* `db_name` - (Optional) Database name.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `databases` - The collection of query.
    * `character_set_name` - Database character set: utf8mb4 (default), utf8, latin1, ascii.
    * `db_name` - The name of the database. Naming rules:
 Unique name. Start with a lowercase letter and end with a letter or number. The length is within 2 to 64 characters.
 Consist of lowercase letters, numbers, underscores (_), or hyphens (-).
 The name cannot contain certain reserved words.
* `total_count` - The total count of query.


