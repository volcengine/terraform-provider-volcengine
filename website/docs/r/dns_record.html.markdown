---
subcategory: "DNS"
layout: "volcengine"
page_title: "Volcengine: volcengine_dns_record"
sidebar_current: "docs-volcengine-resource-dns_record"
description: |-
  Provides a resource to manage dns record
---
# volcengine_dns_record
Provides a resource to manage dns record
## Example Usage
```hcl
resource "volcengine_dns_record" "foo" {
  zid   = 58846
  host  = "a.com"
  type  = "A"
  value = "1.1.1.2"
}
```
## Argument Reference
The following arguments are supported:
* `host` - (Required, ForceNew) The host record, which is the domain prefix of the subdomain.
* `type` - (Required) The record type.
* `value` - (Required) The value of the DNS record.
* `zid` - (Required, ForceNew) The ID of the domain to which you want to add a DNS record.
* `line` - (Optional, ForceNew) The value of the DNS record.
* `remark` - (Optional) The remark for the DNS record.
* `ttl` - (Optional) The Time-To-Live (TTL) of the DNS record, in seconds.
* `weight` - (Optional) The weight of the DNS record.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `created_at` - The creation time of the domain.
* `enable` - Whether the DNS record is enabled.
* `operators` - The account ID that called this API.
* `pqdn` - The account ID that called this API.
* `record_id` - The ID of the DNS record.
* `record_set_id` - The ID of the record set where the DNS record is located.
* `tags` - The tag information of the DNS record.
* `updated_at` - The update time of the domain.


## Import
DnsRecord can be imported using the id, e.g.
```
$ terraform import volcengine_dns_record.default ZID:recordId
```

