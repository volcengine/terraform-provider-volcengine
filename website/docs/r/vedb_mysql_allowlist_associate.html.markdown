---
subcategory: "VEDB_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_vedb_mysql_allowlist_associate"
sidebar_current: "docs-volcengine-resource-vedb_mysql_allowlist_associate"
description: |-
  Provides a resource to manage vedb mysql allowlist associate
---
# volcengine_vedb_mysql_allowlist_associate
Provides a resource to manage vedb mysql allowlist associate
## Example Usage
```hcl
resource "volcengine_vedb_mysql_allowlist" "foo" {
  allow_list_name = "acc-test-allowlist"
  allow_list_desc = "acc-test"
  allow_list_type = "IPv4"
  allow_list      = ["192.168.0.0/24", "192.168.1.0/24", "192.168.2.0/24"]
}

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

resource "volcengine_vedb_mysql_allowlist_associate" "foo" {
  allow_list_id = volcengine_vedb_mysql_allowlist.foo.id
  instance_id   = volcengine_vedb_mysql_instance.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `allow_list_id` - (Required, ForceNew) The id of the allow list.
* `instance_id` - (Required, ForceNew) The id of the mysql instance.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
VedbMysqlAllowlistAssociate can be imported using the instance id and allow list id, e.g.
```
$ terraform import volcengine_vedb_mysql_allowlist_associate.default vedbm-iqnh3a7z****:acl-d1fd76693bd54e658912e7337d5b****
```

