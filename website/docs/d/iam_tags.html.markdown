---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_tags"
sidebar_current: "docs-volcengine-datasource-iam_tags"
description: |-
  Use this data source to query detailed information of iam tags
---
# volcengine_iam_tags
Use this data source to query detailed information of iam tags
## Example Usage
```hcl
data "volcengine_iam_tags" "default" {
  resource_type = "Role"
}
```
## Argument Reference
The following arguments are supported:
* `resource_type` - (Required) The type of the resource. Valid values: User, Role.
* `output_file` - (Optional) File name where to save data source results.
* `resource_names` - (Optional) The names of the resource.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `next_token` - The next token of query.
* `resource_tags` - The collection of query.
    * `resource_name` - The name of the resource.
    * `resource_type` - The type of the resource.
    * `tag_key` - The key of the tag.
    * `tag_value` - The value of the tag.
* `total_count` - The total count of query.


