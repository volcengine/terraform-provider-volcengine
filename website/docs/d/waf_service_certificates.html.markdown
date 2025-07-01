---
subcategory: "WAF"
layout: "volcengine"
page_title: "Volcengine: volcengine_waf_service_certificates"
sidebar_current: "docs-volcengine-datasource-waf_service_certificates"
description: |-
  Use this data source to query detailed information of waf service certificates
---
# volcengine_waf_service_certificates
Use this data source to query detailed information of waf service certificates
## Example Usage
```hcl
data "volcengine_waf_service_certificates" "foo" {
}
```
## Argument Reference
The following arguments are supported:
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `data` - The Information of the certificate.
    * `applicable_domains` - Associate the domain name of this certificate.
    * `description` - The description of the certificate.
    * `expire_time` - The expiration time of the certificate.
    * `id` - The ID of the certificate.
    * `insert_time` - The time when the certificate was added.
    * `name` - The name of the certificate.
* `total_count` - The total count of query.


