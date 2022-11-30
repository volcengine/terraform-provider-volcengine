---
subcategory: "CEN(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_cens"
sidebar_current: "docs-volcengine-datasource-cens"
description: |-
  Use this data source to query detailed information of cens
---
# volcengine_cens
Use this data source to query detailed information of cens
## Example Usage
```hcl
data "volcengine_cens" "foo" {
  ids = ["cen-2bzrl3srxsv0g2dx0efyoojn3"]
}
```
## Argument Reference
The following arguments are supported:
* `cen_names` - (Optional) A list of cen names.
* `ids` - (Optional) A list of cen IDs.
* `name_regex` - (Optional) A Name Regex of cen.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `cens` - The collection of cen query.
    * `account_id` - The account ID of the cen.
    * `cen_bandwidth_package_ids` - A list of bandwidth package IDs of the cen.
    * `cen_id` - The ID of the cen.
    * `cen_name` - The name of the cen.
    * `creation_time` - The create time of the cen.
    * `description` - The description of the cen.
    * `id` - The ID of the cen.
    * `status` - The status of the cen.
    * `update_time` - The update time of the cen.
* `total_count` - The total count of cen query.


