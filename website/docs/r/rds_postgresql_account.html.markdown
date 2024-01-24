---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_account"
sidebar_current: "docs-volcengine-resource-rds_postgresql_account"
description: |-
  Provides a resource to manage rds postgresql account
---
# volcengine_rds_postgresql_account
Provides a resource to manage rds postgresql account
## Example Usage
```hcl
resource "volcengine_rds_postgresql_account" "foo" {
  account_name     = "acc-test-account"
  account_password = "93c@*****!ab12"
  account_type     = "Super"
  instance_id      = "postgres-954*****7233"
}

resource "volcengine_rds_postgresql_account" "foo1" {
  account_name       = "acc-test-account1"
  account_password   = "9wc@****b12"
  account_type       = "Normal"
  instance_id        = "postgres-95*****7233"
  account_privileges = "Inherit,Login,CreateRole,CreateDB"
}
```
## Argument Reference
The following arguments are supported:
* `account_name` - (Required, ForceNew) Database account name.
* `account_password` - (Required) The password of the database account. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `account_type` - (Required, ForceNew) Database account type, value:
Super: A high-privilege account. Only one database account can be created for an instance.
Normal: An account with ordinary privileges.
* `instance_id` - (Required, ForceNew) The ID of the RDS instance.
* `account_privileges` - (Optional) The privilege information of account. When the account type is a super account, there is no need to pass in this parameter, and all privileges are supported by default. When the account type is a normal account, this parameter can be passed in, the default values are Login and Inherit.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `account_status` - The status of the database account.


## Import
RDS postgresql account can be imported using the instance_id:account_name, e.g.
```
$ terraform import volcengine_rds_postgresql_account.default postgres-ca7b7019****:accountName
```

