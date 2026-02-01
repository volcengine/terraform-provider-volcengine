---
subcategory: "ESCLOUD"
layout: "volcengine"
page_title: "Volcengine: volcengine_escloud_zones_v2"
sidebar_current: "docs-volcengine-datasource-escloud_zones_v2"
description: |-
  Use this data source to query detailed information of escloud zones v2
---
# volcengine_escloud_zones_v2
Use this data source to query detailed information of escloud zones v2
## Example Usage
```hcl
data "volcengine_escloud_zones_v2" "foo" {

}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of query.
* `zones` - The collection of query.
    * `region_id` - The region ID of zone.
    * `zone_id` - The ID of zone.
    * `zone_name` - The name of zone.
    * `zone_status` - The status of zone.


