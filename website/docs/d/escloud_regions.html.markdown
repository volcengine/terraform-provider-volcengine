---
subcategory: "ESCLOUD"
layout: "volcengine"
page_title: "Volcengine: volcengine_escloud_regions"
sidebar_current: "docs-volcengine-datasource-escloud_regions"
description: |-
  Use this data source to query detailed information of escloud regions
---
# volcengine_escloud_regions
(Deprecated! Recommend use volcengine_escloud_instance_v2 replace) Use this data source to query detailed information of escloud regions
## Example Usage
```hcl
data "volcengine_escloud_regions" "default" {
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
* `total_count` - The total count of region query.


