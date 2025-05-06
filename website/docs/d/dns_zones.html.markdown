---
subcategory: "DNS"
layout: "volcengine"
page_title: "Volcengine: volcengine_dns_zones"
sidebar_current: "docs-volcengine-datasource-dns_zones"
description: |-
  Use this data source to query detailed information of dns zones
---
# volcengine_dns_zones
Use this data source to query detailed information of dns zones
## Example Usage
```hcl
data "volcengine_dns_zones" "foo" {
  tags {
    key    = "xx"
    values = ["xx"]
  }
}
```
## Argument Reference
The following arguments are supported:
* `key` - (Optional) The keyword included in domains.
* `name_regex` - (Optional) A Name Regex of Resource.
* `order_key` - (Optional) The key for sorting the results.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The ProjectName of the domain.
* `search_mode` - (Optional) The matching mode for the Key parameter.
* `search_order` - (Optional) The sorting order of the results.
* `stage` - (Optional) The status of the domain.
* `tags` - (Optional) Tags.
* `trade_code` - (Optional) The edition of the domain.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `values` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of query.
* `zones` - The collection of query.
    * `allocate_dns_server_list` - The list of DNS servers allocated to the domain by BytePlus DNS.
    * `auto_renew` - Whether automatic domain renewal is enabled.
    * `cache_stage` - The most recent update time of the domain.
    * `created_at` - The creation time of the domain.
    * `dns_security` - The version of DNS DDoS protection service.
    * `expired_time` - The expiration time of the domain.
    * `id` - The id of the zone.
    * `instance_id` - The ID of the instance.
    * `instance_no` - The ID of the instance. For free edition, the value of this field is null.
    * `is_ns_correct` - Indicates whether the configuration of NS servers is correct. If the configuration is correct, the status of the domain in BytePlus DNS is Active.
    * `is_sub_domain` - Whether the domain is a subdomain.
    * `last_operator` - The ID of the account that last updated this domain.
    * `project_name` - The ProjectName of the domain.
    * `real_dns_server_list` - The list of DNS servers actually used by the domain.
    * `record_count` - The total number of DNS records contained in the domain.
    * `remark` - The remarks for the domain.
    * `stage` - The status of the domain.
    * `sub_domain_host` - The domain prefix of the subdomain. If the domain is not a subdomain, this parameter is null.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `trade_code` - The edition of the domain.
    * `updated_at` - The most recent update time of the domain.
    * `zid` - The ID of the domain.
    * `zone_name` - The domain name.


