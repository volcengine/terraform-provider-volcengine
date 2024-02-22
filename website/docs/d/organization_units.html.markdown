---
subcategory: "ORGANIZATION"
layout: "volcengine"
page_title: "Volcengine: volcengine_organization_units"
sidebar_current: "docs-volcengine-datasource-organization_units"
description: |-
  Use this data source to query detailed information of organization units
---
# volcengine_organization_units
Use this data source to query detailed information of organization units
## Example Usage
```hcl
data "volcengine_organization_units" "foo" {
}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of query.
* `units` - The collection of query.
    * `created_time` - The created time of the organization unit.
    * `delete_uk` - Delete marker.
    * `deleted_time` - The deleted time of the organization unit.
    * `depth` - The depth of the organization unit.
    * `description` - The description of the organization unit.
    * `id` - The id of the organization unit.
    * `name` - The name of the organization unit.
    * `org_id` - The id of the organization.
    * `org_type` - The organization type.
    * `owner` - The owner of the organization unit.
    * `parent_id` - Parent Unit ID.
    * `updated_time` - The updated time of the organization unit.


