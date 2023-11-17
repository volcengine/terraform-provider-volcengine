---
subcategory: "ALB"
layout: "volcengine"
page_title: "Volcengine: volcengine_alb_certificate"
sidebar_current: "docs-volcengine-resource-alb_certificate"
description: |-
  Provides a resource to manage alb certificate
---
# volcengine_alb_certificate
Provides a resource to manage alb certificate
## Example Usage
```hcl
resource "volcengine_alb_certificate" "foo" {
  description = "test123"
  public_key  = "public key"
  private_key = "private key"
}
```
## Argument Reference
The following arguments are supported:
* `private_key` - (Required, ForceNew) The private key of the Certificate.
* `public_key` - (Required, ForceNew) The public key of the Certificate.
* `certificate_name` - (Optional) The name of the Certificate.
* `description` - (Optional) The description of the Certificate.
* `project_name` - (Optional) The project name of the Certificate.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `certificate_type` - The type of the Certificate.
* `create_time` - The create time of the Certificate.
* `domain_name` - The domain name of the Certificate.
* `expired_at` - The expire time of the Certificate.
* `listeners` - The ID list of the Listener.
* `status` - The status of the Certificate.


## Import
Certificate can be imported using the id, e.g.
```
$ terraform import volcengine_alb_certificate.default cert-2fe5k****c16o5oxruvtk3qf5
```

