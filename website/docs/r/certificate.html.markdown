---
subcategory: "CLB"
layout: "volcengine"
page_title: "Volcengine: volcengine_certificate"
sidebar_current: "docs-volcengine-resource-certificate"
description: |-
  Provides a resource to manage certificate
---
# volcengine_certificate
Provides a resource to manage certificate
## Example Usage
```hcl
variable "certificate" {
  type = object({
    public_key  = string
    private_key = string
  })
}

resource "volcengine_certificate" "foo" {
  certificate_name = "demo-certificate"
  description      = "This is a clb certificate"
  public_key       = var.certificate.public_key
  private_key      = var.certificate.private_key
}
```
## Argument Reference
The following arguments are supported:
* `private_key` - (Required, ForceNew) The private key of the Certificate.
* `public_key` - (Required, ForceNew) The public key of the Certificate.
* `certificate_name` - (Optional, ForceNew) The name of the Certificate.
* `description` - (Optional, ForceNew) The description of the Certificate.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Certificate can be imported using the id, e.g.
```
$ terraform import volcengine_certificate.default cert-2fe5k****c16o5oxruvtk3qf5
```

