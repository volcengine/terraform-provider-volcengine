---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_roles"
sidebar_current: "docs-volcengine-datasource-iam_roles"
description: |-
  Use this data source to query detailed information of iam roles
---
# volcengine_iam_roles
Use this data source to query detailed information of iam roles
## Example Usage
```hcl
data "volcengine_iam_roles" "default" {
}
```
## Argument Reference
The following arguments are supported:
* `name_regex` - (Optional) A Name Regex of Role.
* `output_file` - (Optional) File name where to save data source results.
* `query` - (Optional) The query field of Role.
* `role_name` - (Optional) The name of the Role, comma separated.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `roles` - The collection of Role query.
  * `create_date` - The create time of the Role.
  * `description` - The description of the Role.
  * `id` - The ID of the Role.
  * `role_name` - The name of the Role.
  * `trn` - The resource name of the Role.
  * `trust_policy_document` - The trust policy document of the Role.
* `total_count` - The total count of Role query.


