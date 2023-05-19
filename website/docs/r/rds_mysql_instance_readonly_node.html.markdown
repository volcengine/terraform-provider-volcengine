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
resource "volcengine_rds_mysql_instance_readonly_node" "foo" {
  instance_id = "mysql-b3fca7f571d6"
  node_spec   = "rds.mysql.1c2g"
  zone_id     = "cn-guilin-b"
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

