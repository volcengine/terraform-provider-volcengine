---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_ip_list"
sidebar_current: "docs-volcengine-resource-rds_ip_list"
description: |-
  Provides a resource to manage rds ip list
---
# volcengine_rds_ip_list
Provides a resource to manage rds ip list
## Example Usage
```hcl
resource "volcengine_rds_ip_list" "foo" {
  instance_id = "mysql-0fdd3bab2e7c"
  group_name  = "foo"
  ip_list     = ["1.1.1.1", "2.2.2.2"]
}
```
## Argument Reference
The following arguments are supported:
* `group_name` - (Required, ForceNew) The name of the RDS ip list.
* `instance_id` - (Required, ForceNew) The ID of the RDS instance.
* `ip_list` - (Required) The list of IP address.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
RDSIPList can be imported using the id, e.g.
```
$ terraform import volcengine_rds_ip_list.default mysql-42b38c769c4b:group_name
```

