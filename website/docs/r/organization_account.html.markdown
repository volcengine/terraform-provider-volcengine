---
subcategory: "ORGANIZATION"
layout: "volcengine"
page_title: "Volcengine: volcengine_organization_account"
sidebar_current: "docs-volcengine-resource-organization_account"
description: |-
  Provides a resource to manage organization account
---
# volcengine_organization_account
Provides a resource to manage organization account
## Example Usage
```hcl
resource "volcengine_organization_unit" "foo" {
  name        = "acc-test-org-unit"
  parent_id   = "730671013833632****"
  description = "acc-test"
}

resource "volcengine_organization_account" "foo" {
  account_name = "acc-test-account"
  show_name    = "acc-test-account"
  description  = "acc-test"
  org_unit_id  = volcengine_organization_unit.foo.id

  tags {
    key   = "k1"
    value = "v1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `account_name` - (Required) The name of the account.
* `show_name` - (Required) The show name of the account.
* `description` - (Optional) The description of the account.
* `org_unit_id` - (Optional) The id of the organization unit. Default is root organization.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `iam_role` - The name of the iam role.
* `org_id` - The id of the organization.
* `org_unit_name` - The name of the organization unit.
* `org_verification_id` - The id of the organization verification.
* `owner` - The owner id of the account.


## Import
OrganizationAccount can be imported using the id, e.g.
```
$ terraform import volcengine_organization_account.default resource_id
```

