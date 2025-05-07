---
subcategory: "PRIVATE_ZONE"
layout: "volcengine"
page_title: "Volcengine: volcengine_private_zones"
sidebar_current: "docs-volcengine-datasource-private_zones"
description: |-
  Use this data source to query detailed information of private zones
---
# volcengine_private_zones
Use this data source to query detailed information of private zones
## Example Usage
```hcl
data "volcengine_private_zones" "foo" {
  zid            = 770000
  zone_name      = "volces.com"
  search_mode    = "EXACT"
  recursion_mode = true
  line_mode      = 3
}
```
## Argument Reference
The following arguments are supported:
* `line_mode` - (Optional) The line mode of Private Zone, specified whether the intelligent mode and the load balance function is enabled.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `recursion_mode` - (Optional) Whether the recursion mode of Private Zone is enabled.
* `region` - (Optional) The region of Private Zone.
* `search_mode` - (Optional) The search mode of query. Valid values: `LIKE`, `EXACT`. Default is `LIKE`.
* `vpc_id` - (Optional) The vpc id associated to Private Zone.
* `zid` - (Optional) The zid of Private Zone.
* `zone_name` - (Optional) The name of Private Zone.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `private_zones` - The collection of query.
    * `bind_vpcs` - The Bind vpc info of the private zone.
        * `account_id` - The account id of the bind vpc.
        * `id` - The id of the bind vpc.
        * `region_name` - The region name of the bind vpc.
        * `region` - The region of the bind vpc.
    * `created_at` - The created time of the private zone.
    * `id` - The id of the private zone.
    * `last_operator` - The account id of the last operator who created the private zone.
    * `line_mode` - The line mode of the private zone, specified whether the intelligent mode and the load balance function is enabled.
    * `record_count` - The record count of the private zone.
    * `recursion_mode` - Whether the recursion mode of the private zone is enabled.
    * `region` - The region of the private zone.
    * `remark` - The remark of the private zone.
    * `updated_at` - The updated time of the private zone.
    * `zid` - The id of the private zone.
    * `zone_name` - The id of the private zone.
* `total_count` - The total count of query.


