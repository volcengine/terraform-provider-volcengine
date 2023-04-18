---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_mysql_allowlist_associate"
sidebar_current: "docs-volcengine-resource-rds_mysql_allowlist_associate"
description: |-
  Provides a resource to manage rds mysql allowlist associate
---
# volcengine_rds_mysql_allowlist_associate
Provides a resource to manage rds mysql allowlist associate
## Example Usage
```hcl
resource "volcengine_rds_mysql_allowlist_associate" "foo" {
  instance_id   = "mysql-1b2c7b2d7583"
  allow_list_id = "acl-15451212dcfa473baeda24be4baa02fe"
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
RDS AllowList Associate can be imported using the instance id and allow list id, e.g.
```
$ terraform import volcengine_rds_mysql_allowlist_associate.default rds-mysql-h441603c68aaa:acl-d1fd76693bd54e658912e7337d5b****
```

