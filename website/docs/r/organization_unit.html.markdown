---
subcategory: "ORGANIZATION"
layout: "volcengine"
page_title: "Volcengine: volcengine_organization_unit"
sidebar_current: "docs-volcengine-resource-organization_unit"
description: |-
  Provides a resource to manage organization unit
---
# volcengine_organization_unit
Provides a resource to manage organization unit
## Example Usage
```hcl
resource "volcengine_organization" "foo" {

}

data "volcengine_organization_units" "foo" {
  depends_on = [volcengine_organization.foo]
}

resource "volcengine_organization_unit" "foo" {
  name        = "tf-test-unit"
  parent_id   = [for unit in data.volcengine_organization_units.foo.units : unit.id if unit.parent_id == "0"][0]
  description = "tf-test"
}
```
## Argument Reference
The following arguments are supported:
* `name` - (Required) Name of the organization unit.
* `parent_id` - (Required, ForceNew) Parent Organization Unit ID.
* `description` - (Optional) Description of the organization unit.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `depth` - The depth of the organization unit.
* `org_id` - The id of the organization.
* `org_type` - The organization type.
* `owner` - The owner of the organization unit.


## Import
OrganizationUnit can be imported using the id, e.g.
```
$ terraform import volcengine_organization_unit.default ID
```

