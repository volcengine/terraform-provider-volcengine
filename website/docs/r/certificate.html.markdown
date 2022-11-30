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
resource "volcengine_certificate" "foo" {
  certificate_name = "demo-certificate"
  description      = "This is a clb certificate"
  public_key       = "public-key"
  private_key      = "private-key"
}
```
## Argument Reference
The following arguments are supported:
* `private_key` - (Required, ForceNew) The private key of the Certificate. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `public_key` - (Required, ForceNew) The public key of the Certificate. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
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

