---
subcategory: "PRIVATE_ZONE"
layout: "volcengine"
page_title: "Volcengine: volcengine_private_zone_records"
sidebar_current: "docs-volcengine-datasource-private_zone_records"
description: |-
  Use this data source to query detailed information of private zone records
---
# volcengine_private_zone_records
Use this data source to query detailed information of private zone records
## Example Usage
```hcl
data "volcengine_private_zone_records" "foo" {
  zid       = 2450000
  record_id = "907925684878276****"
}
```
## Argument Reference
The following arguments are supported:
* `host` - (Optional) The host of Private Zone Record.
* `last_operator` - (Optional) The last operator account id of Private Zone Record.
* `line` - (Optional) The subnet id of Private Zone Record. This field is only effected when the `intelligent_mode` of the private zone is true.
* `name` - (Optional) The domain name of Private Zone Record.
* `output_file` - (Optional) File name where to save data source results.
* `record_id` - (Optional) The id of Private Zone Record.
* `search_mode` - (Optional) The search mode of query `host`. Valid values: `LIKE`, `EXACT`. Default is `LIKE`.
* `type` - (Optional) The type of Private Zone Record.
* `value` - (Optional) The value of Private Zone Record.
* `zid` - (Optional) The zid of Private Zone.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `records` - The collection of query.
    * `created_at` - The created time of the private zone record.
    * `enable` - Whether the private zone record is enabling.
    * `host` - The host of the private zone record.
    * `last_operator` - The last operator account id of the private zone record.
    * `line` - The subnet id of the private zone record. This field is only effected when the `intelligent_mode` of the private zone is true.
    * `record_id` - The id of the private zone record.
    * `remark` - The remark of the private zone record.
    * `ttl` - The ttl of the private zone record. Unit: second.
    * `type` - The type of the private zone record.
    * `updated_at` - The updated time of the private zone record.
    * `value` - The value of the private zone record.
    * `weight` - The weight of the private zone record.
    * `zid` - The zid of the private zone record.
* `total_count` - The total count of query.


