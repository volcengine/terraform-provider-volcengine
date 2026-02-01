---
subcategory: "TOS(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_tos_bucket_request_payment"
sidebar_current: "docs-volcengine-resource-tos_bucket_request_payment"
description: |-
  Provides a resource to manage tos bucket request payment
---
# volcengine_tos_bucket_request_payment
Provides a resource to manage tos bucket request payment
## Example Usage
```hcl
resource "volcengine_tos_bucket_request_payment" "foo" {
  bucket_name = "tflyb7"
}
```
## Argument Reference
The following arguments are supported:
* `bucket_name` - (Required, ForceNew) The name of the TOS bucket.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
TosBucketRequestPayment can be imported using the bucketName, e.g.
```
$ terraform import volcengine_tos_bucket_request_payment.default bucket_name
```

