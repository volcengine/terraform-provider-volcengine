---
subcategory: "VEENEDGE"
layout: "volcengine"
page_title: "Volcengine: volcengine_veenedge_available_resources"
sidebar_current: "docs-volcengine-datasource-veenedge_available_resources"
description: |-
  Use this data source to query detailed information of veenedge available resources
---
# volcengine_veenedge_available_resources
Use this data source to query detailed information of veenedge available resources
## Example Usage
```hcl
data "volcengine_veenedge_available_resources" "default" {
  instance_type   = "ve******rge"
  bandwith_limit  = 20
  cloud_disk_type = "CloudSSD"
}
```
## Argument Reference
The following arguments are supported:
* `bandwith_limit` - (Required) The limit of bandwidth.
* `cloud_disk_type` - (Required) The type of storage. The value can be `CloudHDD` or `CloudSSD`.
* `instance_type` - (Required) The type of instance.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `regions` - The collection of resource query.
    * `area` - The config of area.
        * `en_name` - The english name of region.
        * `id` - The id of region.
        * `name` - The name of region.
    * `city` - The config of city.
        * `en_name` - The english name of region.
        * `id` - The id of region.
        * `name` - The name of region.
    * `cluster` - The config of cluster.
        * `en_name` - The english name of region.
        * `id` - The id of region.
        * `name` - The name of region.
    * `country` - The config of country.
        * `en_name` - The english name of region.
        * `id` - The id of region.
        * `name` - The name of region.
    * `isp` - The config of isp.
        * `en_name` - The english name of region.
        * `id` - The id of region.
        * `name` - The name of region.
* `total_count` - The total count of resource query.


