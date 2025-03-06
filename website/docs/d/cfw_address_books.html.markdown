---
subcategory: "CLOUD_FIREWALL"
layout: "volcengine"
page_title: "Volcengine: volcengine_cfw_address_books"
sidebar_current: "docs-volcengine-datasource-cfw_address_books"
description: |-
  Use this data source to query detailed information of cfw address books
---
# volcengine_cfw_address_books
Use this data source to query detailed information of cfw address books
## Example Usage
```hcl
data "volcengine_cfw_address_books" "foo" {
  group_type = "ip"
  group_name = "acc-test"
}
```
## Argument Reference
The following arguments are supported:
* `address` - (Optional) The group type of address book. This field support fuzzy query.
* `description` - (Optional) The group type of address book. This field support fuzzy query.
* `group_name` - (Optional) The group name of address book. This field support fuzzy query.
* `group_type` - (Optional) The group type of address book. Valid values: `ip`, `port`, `domain`.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `address_books` - The collection of query.
    * `address_list` - The address list of the address book.
    * `description` - The description of the address book.
    * `group_name` - The name of the address book.
    * `group_type` - The type of the address book.
    * `group_uuid` - The uuid of the address book.
    * `id` - The uuid of the address book.
    * `ref_cnt` - The reference count of the address book.
* `total_count` - The total count of query.


