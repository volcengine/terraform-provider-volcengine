---
subcategory: "ORGANIZATION"
layout: "volcengine"
page_title: "Volcengine: volcengine_organizations"
sidebar_current: "docs-volcengine-datasource-organizations"
description: |-
  Use this data source to query detailed information of organizations
---
# volcengine_organizations
Use this data source to query detailed information of organizations
## Example Usage
```hcl
data "volcengine_organizations" "foo" {

}
```
## Argument Reference
The following arguments are supported:
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `organizations` - The collection of query.
    * `account_id` - The account id of the organization owner.
    * `account_name` - The account name of the organization owner.
    * `created_time` - The created time of the organization.
    * `delete_uk` - The delete uk of the organization.
    * `deleted_time` - The deleted time of the organization.
    * `description` - The description of the organization.
    * `id` - The id of the organization.
    * `main_name` - The main name of the organization owner.
    * `name` - The name of the organization.
    * `owner` - The owner id of the organization.
    * `status` - The status of the organization.
    * `type` - The type of the organization.
    * `updated_time` - The updated time of the organization.
* `total_count` - The total count of query.


