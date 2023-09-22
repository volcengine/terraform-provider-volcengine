---
subcategory: "NAS"
layout: "volcengine"
page_title: "Volcengine: volcengine_nas_zones"
sidebar_current: "docs-volcengine-datasource-nas_zones"
description: |-
  Use this data source to query detailed information of nas zones
---
# volcengine_nas_zones
Use this data source to query detailed information of nas zones
## Example Usage
```hcl
data "volcengine_nas_zones" "default" {

}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of zone query.
* `zones` - The collection of zone query.
    * `id` - The id of the zone.
    * `sales` - The collection of sales info.
        * `file_system_type` - The type of file system.
        * `protocol_type` - The type of protocol.
        * `status` - The status info.
        * `storage_type` - The type of storage.
    * `status` - The status info.
    * `zone_id` - The id of the zone.
    * `zone_name` - The name of the zone.


