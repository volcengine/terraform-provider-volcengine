---
subcategory: "DNS"
layout: "volcengine"
page_title: "Volcengine: volcengine_dns_record_sets"
sidebar_current: "docs-volcengine-datasource-dns_record_sets"
description: |-
  Use this data source to query detailed information of dns record sets
---
# volcengine_dns_record_sets
Use this data source to query detailed information of dns record sets
## Example Usage
```hcl
data "volcengine_dns_zones" "foo" {
  key         = "xxx"
  search_mode = "xx"
}

data "volcengine_dns_record_sets" "foo" {
  zid = data.volcengine_dns_zones.foo.zones[0].zid
}
```
## Argument Reference
The following arguments are supported:
* `zid` - (Required) The domain ID.
* `host` - (Optional) The domain prefix of the record set.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `record_set_id` - (Optional) The record set ID.
* `search_mode` - (Optional) The matching mode for Host.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `record_sets` - The collection of query.
    * `host` - The host record contained in the DNS record set.
    * `id` - The ID of the DNS record set.
    * `line` - The line code corresponding to the DNS record set.
    * `pqdn` - The domain prefix contained in the DNS record set, in PQDN (Partially Qualified Domain Name) format.
    * `type` - The type of DNS records in the DNS record set.
    * `weight_enabled` - Indicates whether load balancing is enabled for the DNS record set.
* `total_count` - The total count of query.


