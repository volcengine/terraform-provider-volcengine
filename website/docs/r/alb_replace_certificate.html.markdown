---
subcategory: "ALB"
layout: "volcengine"
page_title: "Volcengine: volcengine_alb_replace_certificate"
sidebar_current: "docs-volcengine-resource-alb_replace_certificate"
description: |-
  Provides a resource to manage alb replace certificate
---
# volcengine_alb_replace_certificate
Provides a resource to manage alb replace certificate
## Example Usage
```hcl
# replace server certificate
resource "volcengine_alb_replace_certificate" "foo1" {
  certificate_type   = "server"
  old_certificate_id = "cert-bdde0znk524g8dv40or*****"
  update_mode        = "new"
  certificate_name   = "replaced-server-cert"
  description        = "Replaced server certificate"
  project_name       = "default"
  public_key         = file("/path/server_certificate.pem")
  private_key        = file("/path/private_key_rsa.pem")
}
resource "volcengine_alb_replace_certificate" "foo2" {
  certificate_type   = "server"
  old_certificate_id = "cert-1pf4a8k8tokcg845wfar*****"
  update_mode        = "stock"
  certificate_source = "alb"
  certificate_id     = "cert-bdde0znk524g8dv40or*****"
  certificate_name   = "replaced-server-cert-stock"
  description        = "Replaced server certificate (stock)"
  project_name       = "default"
}

# replace ca certificate
resource "volcengine_alb_replace_certificate" "foo3" {
  certificate_type   = "ca"
  old_certificate_id = "cert-xoekc6lpu9s054ov5eo*****"
  update_mode        = "new"
  certificate_name   = "acc-test-replace"
  ca_certificate     = file("/path/server_certificate.pem")
  description        = "acc-test-replace"
  project_name       = "default"
}
```
## Argument Reference
The following arguments are supported:
* `certificate_type` - (Required, ForceNew) The type of the certificate. Valid values: 'server' for server certificates, 'ca' for CA certificates.
* `old_certificate_id` - (Required, ForceNew) The ID of the old certificate to be replaced.
* `update_mode` - (Required, ForceNew) The mode of certificate replacement. Valid values: 'new' for uploading new certificate, 'stock' for using existing certificate.
* `ca_certificate` - (Optional, ForceNew) The content of the CA certificate. Required when certificate_type is 'ca' and update_mode is 'new'.
* `cert_center_certificate_id` - (Optional, ForceNew) The ID of the new certificate. Required when certificate_source is 'cert_center' and update_mode is 'stock'.
* `certificate_id` - (Optional, ForceNew) The ID of the new certificate or CA certificate. Required when certificate_source is 'alb' and update_mode is 'stock'.
* `certificate_name` - (Optional, ForceNew) The name of the certificate.
* `certificate_source` - (Optional, ForceNew) The source of the server certificate. Valid values: `alb`, `cert_center`. Required when update_mode is 'stock'.
* `description` - (Optional, ForceNew) The description of the certificate.
* `private_key` - (Optional, ForceNew) The private key of the server certificate. Required when certificate_type is 'server' and update_mode is 'new'.
* `project_name` - (Optional, ForceNew) The project name of the certificate.
* `public_key` - (Optional, ForceNew) The public key of the server certificate. Required when certificate_type is 'server' and update_mode is 'new'.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
The AlbReplaceCertificate is not support import.

