---
subcategory: "CDN"
layout: "volcengine"
page_title: "Volcengine: volcengine_cdn_certificates"
sidebar_current: "docs-volcengine-datasource-cdn_certificates"
description: |-
  Use this data source to query detailed information of cdn certificates
---
# volcengine_cdn_certificates
Use this data source to query detailed information of cdn certificates
## Example Usage
```hcl
resource "volcengine_cdn_certificate" "foo" {
  certificate = ""
  private_key = ""
  desc        = "tftest"
  source      = "cdn_cert_hosting"
}

data "volcengine_cdn_certificates" "foo" {
  source = volcengine_cdn_certificate.foo.source
}
```
## Argument Reference
The following arguments are supported:
* `source` - (Required) Specify the location for storing the certificate. The parameter can take the following values: `volc_cert_center`: indicates that the certificate will be stored in the certificate center.`cdn_cert_hosting`: indicates that the certificate will be hosted on the content delivery network.
* `name` - (Optional) Specify a domain to obtain certificates that include that domain in the SAN field. The domain can be a wildcard domain. For example, specifying *.example.com will obtain certificates that include img.example.com or www.example.com in the SAN field.
* `output_file` - (Optional) File name where to save data source results.
* `status` - (Optional) Specify one or more states to retrieve certificates in those states. By default, all certificates in all states are returned. You can specify the following states. Multiple states are separated by commas. running: Retrieves certificates with a validity period greater than 30 days. expired: Retrieves certificates that have already expired. expiring_soon: Retrieves certificates with a validity period less than or equal to 30 days but have not yet expired.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `cert_info` - The collection of query.
    * `cert_id` - ID indicating the certificate.
    * `cert_name` - The domain name to which the certificate is issued.
    * `configured_domain` - The domain name associated with the certificate. If the certificate is not yet associated with any domain name, the parameter value is null.
    * `desc` - The remark of the cert.
    * `dns_name` - The domain names included in the SAN field of the certificate.
    * `effective_time` - The issuance time of the certificate is indicated. The unit is Unix timestamp.
    * `expire_time` - The expiration time of the certificate is indicated. The unit is Unix timestamp.
    * `source` - Specify the location for storing the certificate. The parameter can take the following values: `volc_cert_center`: indicates that the certificate will be stored in the certificate center.`cdn_cert_hosting`: indicates that the certificate will be hosted on the content delivery network.
    * `status` - Specify one or more states to retrieve certificates in those states. By default, all certificates in all states are returned. You can specify the following states. Multiple states are separated by commas. running: Retrieves certificates with a validity period greater than 30 days. expired: Retrieves certificates that have already expired. expiring_soon: Retrieves certificates with a validity period less than or equal to 30 days but have not yet expired.
* `total_count` - The total count of query.


