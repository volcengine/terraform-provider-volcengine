---
subcategory: "TOS(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_tos_bucket_transfer_acceleration"
sidebar_current: "docs-volcengine-resource-tos_bucket_transfer_acceleration"
description: |-
  Provides a resource to manage tos bucket transfer acceleration
---
# volcengine_tos_bucket_transfer_acceleration
Provides a resource to manage tos bucket transfer acceleration
## Example Usage
```hcl
# Test bucket transfer acceleration

resource "volcengine_tos_bucket_transfer_acceleration" "default" {
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
TosBucketTransferAcceleration can be imported using the bucketName, e.g.
```
$ terraform import volcengine_tos_bucket_transfer_acceleration.default bucket_name
```

