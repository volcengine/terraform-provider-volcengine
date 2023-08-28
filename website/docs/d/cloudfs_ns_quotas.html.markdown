---
subcategory: "CLOUDFS"
layout: "volcengine"
page_title: "Volcengine: volcengine_cloudfs_ns_quotas"
sidebar_current: "docs-volcengine-datasource-cloudfs_ns_quotas"
description: |-
  Use this data source to query detailed information of cloudfs ns quotas
---
# volcengine_cloudfs_ns_quotas
Use this data source to query detailed information of cloudfs ns quotas
## Example Usage
```hcl
data "volcengine_cloudfs_ns_quotas" "default" {
  fs_names = ["tffile", "tftest2"]
}
```
## Argument Reference
The following arguments are supported:
* `fs_names` - (Required) A list of fs name.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `quotas` - The collection of query.
    * `account_id` - The ID of account.
    * `fs_name` - The name of fs.
    * `ns_count_per_fs` - This file stores the number of namespaces under the instance.
    * `ns_count` - The count of cloud fs namespace.
    * `ns_quota_per_fs` - This file stores the total namespace quota under the instance.
    * `ns_quota` - The quota of cloud fs namespace.
    * `quota_enough` - Whether is enough of cloud fs namespace.
* `total_count` - The total count of cloud fs quota query.


