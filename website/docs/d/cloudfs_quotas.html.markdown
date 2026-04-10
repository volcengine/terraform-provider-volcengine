---
subcategory: "CLOUDFS"
layout: "volcengine"
page_title: "Volcengine: volcengine_cloudfs_quotas"
sidebar_current: "docs-volcengine-datasource-cloudfs_quotas"
description: |-
  Use this data source to query detailed information of cloudfs quotas
---
**❗Notice:**
The current provider is no longer being maintained. We recommend that you use the [volcenginecc](https://registry.terraform.io/providers/volcengine/volcenginecc/latest/docs) instead.
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


