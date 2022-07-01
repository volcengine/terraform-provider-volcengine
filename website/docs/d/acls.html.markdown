---
subcategory: "CLB"
layout: "volcengine"
page_title: "Volcengine: volcengine_acls"
sidebar_current: "docs-volcengine-datasource-acls"
description: |-
  Use this data source to query detailed information of acls
---
# volcengine_acls
Use this data source to query detailed information of acls
## Example Usage
```hcl
data "volcengine_acls" "default" {
  ids = ["acl-3ti8n0rurx4bwbh9jzdy"]
}
```
## Argument Reference
The following arguments are supported:
* `acl_name` - (Optional) The name of acl.
* `ids` - (Optional) A list of Acl IDs.
* `name_regex` - (Optional) A Name Regex of Acl.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `acls` - The collection of Acl query.
  * `acl_entry_count` - The count of acl entry.
  * `acl_id` - The ID of Acl.
  * `acl_name` - The Name of Acl.
  * `create_time` - Creation time of Acl.
  * `description` - The description of Acl.
  * `id` - The ID of Acl.
  * `listeners` - The listeners of Acl.
  * `update_time` - Update time of Acl.
* `total_count` - The total count of Acl query.


