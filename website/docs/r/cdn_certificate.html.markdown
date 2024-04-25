---
subcategory: "CDN"
layout: "volcengine"
page_title: "Volcengine: volcengine_cdn_certificate"
sidebar_current: "docs-volcengine-resource-cdn_certificate"
description: |-
  Provides a resource to manage cdn certificate
---
# volcengine_cdn_certificate
Provides a resource to manage cdn certificate
## Example Usage
```hcl
resource "volcengine_cdn_certificate" "foo" {
  certificate = ""
  private_key = ""
  desc        = "tftest"
  source      = "cdn_cert_hosting"
}
```
## Argument Reference
The following arguments are supported:
* `certificate` - (Required, ForceNew) Content of the specified certificate public key file. Line breaks in the content should be replaced with `\r\n`. The file extension for the certificate public key is `.crt` or `.pem`. The public key must include the complete certificate chain. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `desc` - (Required, ForceNew) Note on the certificate.
* `private_key` - (Required, ForceNew) The content of the specified certificate private key file. Replace line breaks in the content with `\r\n`. The file extension for the certificate private key is `.key` or `.pem`. The private key must be unencrypted. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `source` - (Required, ForceNew) Specify the location for storing the certificate. The parameter can take the following values: `volc_cert_center`: indicates that the certificate will be stored in the certificate center.`cdn_cert_hosting`: indicates that the certificate will be hosted on the content delivery network.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
CdnCertificate can be imported using the id, e.g.
```
$ terraform import volcengine_cdn_certificate.default resource_id
```
You can delete the certificate hosted on the content delivery network.
You can configure the HTTPS module to associate the certificate and domain name through the domain_config field of volcengine_cdn_domain.
If the certificate to be deleted is already associated with a domain name, the deletion will fail.
To remove the association between the domain name and the certificate, you can disable the HTTPS function for the domain name in the Content Delivery Network console.

