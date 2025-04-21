---
subcategory: "DNS"
layout: "volcengine"
page_title: "Volcengine: volcengine_dns_zone"
sidebar_current: "docs-volcengine-resource-dns_zone"
description: |-
  Provides a resource to manage dns zone
---
# volcengine_dns_zone
Provides a resource to manage dns zone
## Example Usage
```hcl
resource "volcengine_dns_zone" "foo" {
  zone_name = "xxxx.com"
  tags {
    key   = "xx"
    value = "xx"
  }
  project_name = "xxx"
  remark       = "xxx"
}
```
## Argument Reference
The following arguments are supported:
* `zone_name` - (Required, ForceNew) The domain to be created. The domain must be a second-level domain and cannot be a wildcard domain.
* `project_name` - (Optional) The project to which the domain name belongs. The default value is default.
* `remark` - (Optional) The remark for the domain.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `allocate_dns_server_list` - The list of DNS servers allocated to the domain by BytePlus DNS.
* `auto_renew` - Whether automatic domain renewal is enabled.
* `dns_security` - The version of DNS DDoS protection service.
* `expired_time` - The expiration time of the domain.
* `instance_no` - The ID of the instance. For free edition, the value of this field is null.
* `is_ns_correct` - Indicates whether the configuration of NS servers is correct. If the configuration is correct, the status of the domain in BytePlus DNS is Active.
* `is_sub_domain` - Whether the domain is a subdomain.
* `real_dns_server_list` - The list of DNS servers actually used by the domain.
* `record_count` - The total number of DNS records under the domain.
* `stage` - The status of the domain.
* `sub_domain_host` - The domain prefix of the subdomain. If the domain is not a subdomain, this parameter is null.
* `trade_code` - The edition of the domain.
* `updated_at` - The update time of the domain.
* `zid` - The ID of the domain.


## Import
Zone can be imported using the id, e.g.
```
$ terraform import volcengine_zone.default resource_id
```

