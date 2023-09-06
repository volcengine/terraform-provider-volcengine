---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_mysql_instance_readonly_node"
sidebar_current: "docs-volcengine-resource-rds_mysql_instance_readonly_node"
description: |-
  Provides a resource to manage rds mysql instance readonly node
---
# volcengine_rds_mysql_instance_readonly_node
Provides a resource to manage rds mysql instance readonly node
## Example Usage
```hcl
data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-project1"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-subnet-test-2"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_rds_mysql_instance" "foo" {
  db_engine_version      = "MySQL_5_7"
  node_spec              = "rds.mysql.1c2g"
  primary_zone_id        = data.volcengine_zones.foo.zones[0].id
  secondary_zone_id      = data.volcengine_zones.foo.zones[0].id
  storage_space          = 80
  subnet_id              = volcengine_subnet.foo.id
  instance_name          = "acc-test"
  lower_case_table_names = "1"

  charge_info {
    charge_type = "PostPaid"
  }

  parameters {
    parameter_name  = "auto_increment_increment"
    parameter_value = "2"
  }
  parameters {
    parameter_name  = "auto_increment_offset"
    parameter_value = "4"
  }
}
resource "volcengine_rds_mysql_instance_readonly_node" "foo" {
  instance_id = volcengine_rds_mysql_instance.foo.id
  node_spec   = "rds.mysql.2c4g"
  zone_id     = data.volcengine_zones.foo.zones[0].id
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required, ForceNew) The RDS mysql instance id of the readonly node.
* `node_spec` - (Required) The specification of readonly node.
* `zone_id` - (Required, ForceNew) The available zone of readonly node.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `node_id` - The id of the readonly node.


## Import
Rds Mysql Instance Readonly Node can be imported using the instance_id:node_id, e.g.
```
$ terraform import volcengine_rds_mysql_instance_readonly_node.default mysql-72da4258c2c7:mysql-72da4258c2c7-r7f93
```

