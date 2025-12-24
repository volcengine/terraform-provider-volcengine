---
subcategory: "TOS(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_tos_bucket_customdomain"
sidebar_current: "docs-volcengine-resource-tos_bucket_customdomain"
description: |-
  Provides a resource to manage tos bucket customdomain
---
# volcengine_tos_bucket_customdomain
Provides a resource to manage tos bucket customdomain
## Example Usage
```hcl
# Create a custom domain for TOS bucket
resource "volcengine_tos_bucket_customdomain" "default" {
  bucket_name = "tflyb7"
  custom_domain_rule {
    domain   = "www.163.com"
    protocol = "tos"
  }
}

resource "volcengine_tos_bucket_customdomain" "default1" {
  bucket_name = "tflyb7"
  custom_domain_rule {
    domain   = "www.2345.com"
    protocol = "tos"
  }
}
```
## Argument Reference
The following arguments are supported:
* `bucket_name` - (Required, ForceNew) The name of the TOS bucket.
* `custom_domain_rule` - (Required) The custom domain role for the bucket.

The `custom_domain_rule` object supports the following:

* `domain` - (Required, ForceNew) The custom domain name for the bucket.
* `cert_id` - (Optional) The certificate id.
* `protocol` - (Optional) Custom domain access protocol.tos|s3.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
TosBucketCustomDomain can be imported using the bucketName:domain, e.g.
```
$ terraform import volcengine_tos_bucket_customdomain.default bucket_name:custom_domain
```

