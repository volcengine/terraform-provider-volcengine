---
subcategory: "TOS(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_tos_buckets"
sidebar_current: "docs-volcengine-datasource-tos_buckets"
description: |-
  Use this data source to query detailed information of tos buckets
---
# volcengine_tos_buckets
Use this data source to query detailed information of tos buckets
## Example Usage
```hcl
data "volcengine_tos_buckets" "default" {
  name_regex = "test"
}
```
## Argument Reference
The following arguments are supported:
* `bucket_name` - (Optional) The name the TOS bucket.
* `name_regex` - (Optional) A Name Regex of TOS bucket.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `buckets` - The collection of TOS bucket query.
    * `creation_date` - The create date of the TOS bucket.
    * `extranet_endpoint` - The extranet endpoint of the TOS bucket.
    * `intranet_endpoint` - The intranet endpoint the TOS bucket.
    * `is_truncated` - (**Deprecated**) The Field is Deprecated. The truncated the TOS bucket.
    * `location` - The location of the TOS bucket.
    * `marker` - (**Deprecated**) The Field is Deprecated. The marker the TOS bucket.
    * `max_keys` - (**Deprecated**) The Field is Deprecated. The max keys the TOS bucket.
    * `name` - The name the TOS bucket.
    * `prefix` - (**Deprecated**) The Field is Deprecated. The prefix the TOS bucket.
* `total_count` - The total count of TOS bucket query.


