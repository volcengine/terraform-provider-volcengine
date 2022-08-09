---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_policies"
sidebar_current: "docs-volcengine-datasource-iam_policies"
description: |-
  Use this data source to query detailed information of iam policies
---
# volcengine_iam_policies
Use this data source to query detailed information of iam policies
## Example Usage
```hcl
data "volcengine_iam_policies" "default" {
  query = "AdministratorAccess"
}
```
## Argument Reference
The following arguments are supported:
* `name_regex` - (Optional) A Name Regex of Policy.
* `output_file` - (Optional) File name where to save data source results.
* `query` - (Optional) Query policies, support policy name or description.
* `scope` - (Optional) The scope of the Policy.
* `status` - (Optional) The status of policy.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `policies` - The collection of Policy query.
    * `create_date` - The create time of the Policy.
    * `description` - The description of the Policy.
    * `id` - The ID of the Policy.
    * `policy_document` - The document of the Policy.
    * `policy_name` - The name of the Policy.
    * `policy_trn` - The resource name of the Policy.
    * `policy_type` - The type of the Policy.
    * `update_date` - The update time of the Policy.
* `total_count` - The total count of Policy query.


