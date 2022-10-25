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
resource "volcengine_mongodb_account" "foo" {
  instance_id      = "mongo-replicxxx"
  account_name     = "root"
  account_password = "1xxx"
}
```
## Argument Reference
The following arguments are supported:
* `account_name` - (Required, ForceNew) The account name.
* `account_password` - (Required) The account password.
* `instance_id` - (Required, ForceNew) The instance ID.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `account_privileges` - The privilege info of mongo instance.
    * `db_name` - The Name of DB.
    * `role_name` - The Name of role.
* `account_type` - The type of account.


## Import
mongodb account can be imported using the instanceId:accountName, e.g.
```
$ terraform import volcengine_mongodb_instance_account.default mongo-replica-e405f8e2****:root
```

