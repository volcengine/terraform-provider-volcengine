---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_caller_identities"
sidebar_current: "docs-volcengine-datasource-iam_caller_identities"
description: |-
  Use this data source to query detailed information of iam caller identities
---
# volcengine_iam_caller_identities
Use this data source to query detailed information of iam caller identities
## Example Usage
```hcl
data "volcengine_iam_caller_identities" "default" {
}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `caller_identities` - The collection of caller identities.
    * `account_id` - The account id.
    * `identity_id` - The identity id.
    * `identity_type` - The identity type.
    * `trn` - The trn.
* `total_count` - The total count of query.


