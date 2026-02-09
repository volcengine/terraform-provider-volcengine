---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_access_key_last_useds"
sidebar_current: "docs-volcengine-datasource-iam_access_key_last_useds"
description: |-
  Use this data source to query detailed information of iam access key last useds
---
# volcengine_iam_access_key_last_useds
Use this data source to query detailed information of iam access key last useds
## Example Usage
```hcl
data "volcengine_iam_access_key_last_useds" "default" {
  access_key_id = "AKLxxxxxxxxxxxxxxxxxxxxxxxxx"
}
```
## Argument Reference
The following arguments are supported:
* `access_key_id` - (Required) The access key id.
* `output_file` - (Optional) File name where to save data source results.
* `user_name` - (Optional) The user name.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `access_key_last_useds` - The collection of access key last used.
    * `region` - The region of the last used.
    * `request_time` - The request time of the last used.
    * `service` - The service of the last used.
* `total_count` - The total count of query.


