---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_access_keys"
sidebar_current: "docs-volcengine-datasource-iam_access_keys"
description: |-
  Use this data source to query detailed information of iam access keys
---
# volcengine_iam_access_keys
Use this data source to query detailed information of iam access keys
## Example Usage
```hcl
data "volcengine_iam_access_keys" "default" {
}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.
* `user_name` - (Optional) The user name.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `access_key_metadata` - The collection of access keys.
    * `access_key_id` - The user access key id.
    * `create_date` - The user access key create date.
    * `status` - The user access key status.
    * `update_date` - The user access key update date.
    * `user_name` - The user name.
* `total_count` - The total count of user query.


