---
subcategory: "ORGANIZATION"
layout: "volcengine"
page_title: "Volcengine: volcengine_organization"
sidebar_current: "docs-volcengine-resource-organization"
description: |-
  Provides a resource to manage organization
---
# volcengine_organization
Provides a resource to manage organization
## Example Usage
```hcl
resource "volcengine_organization" "foo" {

}
```
## Argument Reference
The following arguments are supported:


## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `account_id` - The account id of the organization owner.
* `account_name` - The account name of the organization owner.
* `created_time` - The created time of the organization.
* `delete_uk` - The delete uk of the organization.
* `deleted_time` - The deleted time of the organization.
* `description` - The description of the organization.
* `main_name` - The main name of the organization owner.
* `name` - The name of the organization.
* `owner` - The owner id of the organization.
* `status` - The status of the organization.
* `type` - The type of the organization.
* `updated_time` - The updated time of the organization.


## Import
Organization can be imported using the id, e.g.
```
$ terraform import volcengine_organization.default resource_id
```

