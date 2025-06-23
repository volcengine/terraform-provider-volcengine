---
subcategory: "APIG"
layout: "volcengine"
page_title: "Volcengine: volcengine_apig_upstream_versions"
sidebar_current: "docs-volcengine-datasource-apig_upstream_versions"
description: |-
  Use this data source to query detailed information of apig upstream versions
---
# volcengine_apig_upstream_versions
Use this data source to query detailed information of apig upstream versions
## Example Usage
```hcl
data "volcengine_apig_upstream_versions" "foo" {
  upstream_id = "ud18p5krj5ce3htvrd0v0"
}
```
## Argument Reference
The following arguments are supported:
* `upstream_id` - (Required) The id of the apig upstream.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of query.
* `versions` - The collection of query.
    * `labels` - The labels of apig upstream version.
        * `key` - The key of apig upstream version label.
        * `value` - The value of apig upstream version label.
    * `name` - The name of apig upstream version.
    * `update_time` - The update time of apig upstream version.


