---
subcategory: "NAS"
layout: "volcengine"
page_title: "Volcengine: volcengine_nas_regions"
sidebar_current: "docs-volcengine-datasource-nas_regions"
description: |-
  Use this data source to query detailed information of nas regions
---
# volcengine_nas_regions
Use this data source to query detailed information of nas regions
## Example Usage
```hcl
data "volcengine_nas_regions" "default" {

}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `regions` - The collection of region query.
    * `region_id` - The id of the region.
    * `region_name` - The name of region.
    * `status` - The status info of region.
* `total_count` - The total count of region query.


