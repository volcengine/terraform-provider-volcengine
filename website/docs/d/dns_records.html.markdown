---
subcategory: "DNS"
layout: "volcengine"
page_title: "Volcengine: volcengine_dns_records"
sidebar_current: "docs-volcengine-datasource-dns_records"
description: |-
  Use this data source to query detailed information of dns records
---
# volcengine_dns_records
Use this data source to query detailed information of dns records
## Example Usage
```hcl
data "volcengine_dns_records" "foo" {
  zid = 58857
}
```
## Argument Reference
The following arguments are supported:
* `zid` - (Required) The ID of the domain.
* `host` - (Optional) Domain prefix of the DNS record.
* `line` - (Optional) Line of the DNS record.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `search_mode` - (Optional) The matching mode for the Host parameter.
* `search_order` - (Optional) The Method to sort the returned list of DNS records.
* `type` - (Optional) Type of the DNS record.
* `value` - (Optional) Value of the DNS record.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `records` - The collection of query.
    * `created_at` - The creation time of the domain.
    * `enable` - Indicates whether the DNS record is enabled.
    * `host` - The host record included in the DNS record.
    * `line` - The line code corresponding to the DNS record.
    * `operators` - The account ID that called this API.
    * `pqdn` - The hostname included in the DNS record, in PQDN (Partially Qualified Domain Name) format.
    * `record_id` - The ID of the DNS record.
    * `record_set_id` - The ID of the record set to which the DNS record belongs.
    * `remark` - The remark of the DNS record.
    * `tags` - The tag information of the DNS record.
    * `ttl` - The Time to Live (TTL) of the DNS record. The unit is seconds.
    * `type` - The type of the DNS record.
    * `updated_at` - The most recent update time of the domain.
    * `value` - The record value contained in the DNS record.
    * `weight` - The weight of the DNS record.
* `total_count` - The total count of query.


