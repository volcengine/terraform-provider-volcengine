---
subcategory: "ORGANIZATION"
layout: "volcengine"
page_title: "Volcengine: volcengine_organization_accounts"
sidebar_current: "docs-volcengine-datasource-organization_accounts"
description: |-
  Use this data source to query detailed information of organization accounts
---
# volcengine_organization_accounts
Use this data source to query detailed information of organization accounts
## Example Usage
```hcl
data "volcengine_organization_accounts" "foo" {
  search = "210061****"
  #  org_unit_id = "730662904425309****"
  #  verification_id = "730671013833631****"
}
```
## Argument Reference
The following arguments are supported:
* `name_regex` - (Optional) A Name Regex of Resource.
* `org_unit_id` - (Optional) The id of the organization unit.
* `output_file` - (Optional) File name where to save data source results.
* `search` - (Optional) The id or the show name of the account. This field supports fuzzy query.
* `verification_id` - (Optional) The id of the verification.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `accounts` - The collection of query.
    * `account_id` - The id of the account.
    * `account_name` - The name of the account.
    * `allow_console` - Whether to allow the account enable console. `0` means allowed, `1` means not allowed.
    * `allow_exit` - Whether to allow exit the organization. `0` means allowed, `1` means not allowed.
    * `created_time` - The created time of the account.
    * `delete_uk` - The delete uk of the account.
    * `deleted_time` - The deleted time of the account.
    * `description` - The description of the account.
    * `iam_role` - The name of the iam role.
    * `id` - The id of the account.
    * `is_owner` - Whether the account is owner. `0` means not owner, `1` means owner.
    * `join_type` - The join type of the account. `0` means create, `1` means invitation.
    * `org_id` - The id of the organization.
    * `org_type` - The type of the organization. `1` means business organization.
    * `org_unit_id` - The id of the organization unit.
    * `org_unit_name` - The name of the organization unit.
    * `org_verification_id` - The id of the organization verification.
    * `owner` - The owner id of the account.
    * `show_name` - The show name of the account.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `updated_time` - The updated time of the account.
* `total_count` - The total count of query.


