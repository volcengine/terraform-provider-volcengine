---
subcategory: "CLOUDFS"
layout: "volcengine"
page_title: "Volcengine: volcengine_cloudfs_quotas"
sidebar_current: "docs-volcengine-datasource-cloudfs_quotas"
description: |-
  Use this data source to query detailed information of cloudfs quotas
---
# volcengine_cloudfs_quotas
Use this data source to query detailed information of cloudfs quotas
## Example Usage
```hcl
data "volcengine_cloudfs_quotas" "default" {
}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `quotas` - The collection of query.
    * `account_id` - The ID of account.
    * `fs_count` - The count of cloud fs.
    * `fs_quota` - The quota of cloud fs.
    * `quota_enough` - Whether is enough of cloud fs.
* `total_count` - The total count of cloud fs quota query.


