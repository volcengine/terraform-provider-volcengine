---
subcategory: "VEDB_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_vedb_mysql_endpoint_public_address"
sidebar_current: "docs-volcengine-resource-vedb_mysql_endpoint_public_address"
description: |-
  Provides a resource to manage vedb mysql endpoint public address
---
# volcengine_vedb_mysql_endpoint_public_address
Provides a resource to manage vedb mysql endpoint public address
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
  project_name        = "default"
  tags {
    key   = "tftest"
    value = "tftest"
  }
  tags {
    key   = "tftest2"
    value = "tftest2"
  }
}
data "volcengine_vedb_mysql_instances" "foo" {
  instance_id = volcengine_vedb_mysql_instance.foo.id
}

resource "volcengine_vedb_mysql_endpoint" "foo" {
  endpoint_type               = "Custom"
  instance_id                 = volcengine_vedb_mysql_instance.foo.id
  node_ids                    = [data.volcengine_vedb_mysql_instances.foo.instances[0].nodes[0].node_id, data.volcengine_vedb_mysql_instances.foo.instances[0].nodes[1].node_id]
  read_write_mode             = "ReadWrite"
  endpoint_name               = "tf-test"
  description                 = "tf test"
  master_accept_read_requests = true
  distributed_transaction     = true
  consist_level               = "Session"
  consist_timeout             = 100000
  consist_timeout_action      = "ReadMaster"
}

resource "volcengine_eip_address" "foo" {
  billing_type = "PostPaidByBandwidth"
  bandwidth    = 1
  isp          = "ChinaUnicom"
  name         = "acc-eip"
  description  = "acc-test"
  project_name = "default"
}

resource "volcengine_vedb_mysql_endpoint_public_address" "foo" {
  eip_id      = volcengine_eip_address.foo.id
  endpoint_id = volcengine_vedb_mysql_endpoint.foo.endpoint_id
  instance_id = volcengine_vedb_mysql_instance.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `eip_id` - (Required, ForceNew) EIP ID that needs to be bound to the instance.
* `endpoint_id` - (Required, ForceNew) The endpoint id.
* `instance_id` - (Required, ForceNew) The instance id.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
VedbMysqlEndpointPublicAddress can be imported using the instance id, endpoint id and the eip id, e.g.
```
$ terraform import volcengine_vedb_mysql_endpoint_public_address.default vedbm-iqnh3a7z****:vedbm-2pf2xk5v****-Custom-50yv:eip-xxxx
```

