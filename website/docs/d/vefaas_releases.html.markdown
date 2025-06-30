---
subcategory: "VEFAAS"
layout: "volcengine"
page_title: "Volcengine: volcengine_vefaas_releases"
sidebar_current: "docs-volcengine-datasource-vefaas_releases"
description: |-
  Use this data source to query detailed information of vefaas releases
---
# volcengine_vefaas_releases
Use this data source to query detailed information of vefaas releases
## Example Usage
```hcl
data "volcengine_vefaas_releases" "foo" {
  function_id = "g79asxxx"
}
```
## Argument Reference
The following arguments are supported:
* `function_id` - (Required) The ID of Function.
* `filters` - (Optional) Query the filtering conditions.
* `name_regex` - (Optional) A Name Regex of Resource.
* `order_by` - (Optional) Query the sorting parameters.
* `output_file` - (Optional) File name where to save data source results.

The `filters` object supports the following:

* `name` - (Optional) Filter key enumeration.
* `values` - (Optional) The filtering value of the query.

The `order_by` object supports the following:

* `ascend` - (Optional) Whether the sorting result is sorted in ascending order.
* `key` - (Optional) Key names used for sorting.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `items` - The list of function publication records.
    * `creation_time` - The create time of the published information.
    * `description` - The description of the published information.
    * `finish_time` - Finish time.
    * `function_id` - The ID of Function.
    * `id` - The ID of function release.
    * `last_update_time` - The last update time of the published information.
    * `source_revision_number` - The historical version numbers released.
    * `status` - The status of function release.
    * `target_revision_number` - The target version number released.
* `total_count` - The total count of query.


