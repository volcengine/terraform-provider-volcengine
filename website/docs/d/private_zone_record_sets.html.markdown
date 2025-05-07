---
subcategory: "PRIVATE_ZONE"
layout: "volcengine"
page_title: "Volcengine: volcengine_private_zone_record_sets"
sidebar_current: "docs-volcengine-datasource-private_zone_record_sets"
description: |-
  Use this data source to query detailed information of private zone record sets
---
# volcengine_private_zone_record_sets
Use this data source to query detailed information of private zone record sets
## Example Usage
```hcl
data "volcengine_private_zone_record_sets" "foo" {
  zid = 2450000
}
```
## Argument Reference
The following arguments are supported:
* `zid` - (Required) The zid of Private Zone.
* `host` - (Optional) The host of Private Zone Record Set.
* `output_file` - (Optional) File name where to save data source results.
* `record_set_id` - (Optional) The id of Private Zone Record Set.
* `search_mode` - (Optional) The search mode of query `host`. Valid values: `LIKE`, `EXACT`. Default is `LIKE`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `record_sets` - The collection of query.
    * `fqdn` - The Complete domain name of the private zone record.
    * `host` - The host of the private zone record.
    * `line` - The subnet id of the private zone record. This field is only effected when the `intelligent_mode` of the private zone is true.
    * `record_set_id` - The id of the private zone record set.
    * `type` - The type of the private zone record.
    * `weight_enabled` - Whether to enable the load balance of the private zone record set.
* `total_count` - The total count of query.


