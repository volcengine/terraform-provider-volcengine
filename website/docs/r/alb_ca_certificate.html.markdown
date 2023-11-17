---
subcategory: "ALB"
layout: "volcengine"
page_title: "Volcengine: volcengine_alb_ca_certificate"
sidebar_current: "docs-volcengine-resource-alb_ca_certificate"
description: |-
  Provides a resource to manage alb ca certificate
---
# volcengine_alb_ca_certificate
Provides a resource to manage alb ca certificate
## Example Usage
```hcl
resource "volcengine_alb_ca_certificate" "foo" {
  ca_certificate_name = "acc-test-1"
  ca_certificate      = "-----BEGIN CERTIFICATE-----\n-----END CERTIFICATE-----"
  description         = "acc-test-1"
  project_name        = "default"
}
```
## Argument Reference
The following arguments are supported:
* `ca_certificate` - (Required, ForceNew) The content of the CA certificate.
* `ca_certificate_name` - (Optional) The name of the CA certificate.
* `description` - (Optional) The description of the CA certificate.
* `project_name` - (Optional) The project name of the CA certificate.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `certificate_type` - The type of the CA Certificate.
* `create_time` - The create time of the CA Certificate.
* `domain_name` - The domain name of the CA Certificate.
* `expired_at` - The expire time of the CA Certificate.
* `listeners` - The ID list of the Listener.
* `status` - The status of the CA Certificate.


## Import
AlbCaCertificate can be imported using the id, e.g.
```
$ terraform import volcengine_alb_ca_certificate.default cert-*****
```

