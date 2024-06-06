---
subcategory: "FINANCIAL_RELATION"
layout: "volcengine"
page_title: "Volcengine: volcengine_financial_relations"
sidebar_current: "docs-volcengine-datasource-financial_relations"
description: |-
  Use this data source to query detailed information of financial relations
---
# volcengine_financial_relations
Use this data source to query detailed information of financial relations
## Example Usage
```hcl
data "volcengine_financial_relations" "foo" {
  account_ids = ["210026****"]
  relation    = ["1"]
  status      = ["200"]
}
```
## Argument Reference
The following arguments are supported:
* `account_ids` - (Optional) A list of sub account IDs.
* `output_file` - (Optional) File name where to save data source results.
* `relation` - (Optional) A list of relation. Valid values: `1`, `4`.
* `status` - (Optional) A list of status. Valid values: `100`, `200`, `250`, `300`, `400`, `500`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `financial_relations` - The collection of query.
    * `account_alias` - The display name of the sub account.
    * `auth_info` - The authorization info of the financial relation.
        * `auth_id` - The auth id of the financial relation.
        * `auth_list` - The auth list of the financial relation.
        * `auth_status` - The auth status of the financial relation.
    * `filiation_desc` - The filiation description of the financial relation.
    * `filiation` - The filiation of the financial relation.
    * `major_account_id` - The id of the major account.
    * `major_account_name` - The name of the major account.
    * `relation_desc` - The relation description of the financial.
    * `relation_id` - The id of the financial relation.
    * `relation` - The relation of the financial.
    * `status_desc` - The status description of the financial relation.
    * `status` - The status of the financial relation.
    * `sub_account_id` - The id of the sub account.
    * `sub_account_name` - The name of the sub account.
    * `update_time` - The update time of the financial relation.
* `total_count` - The total count of query.


