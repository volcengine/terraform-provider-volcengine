---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_cancel_secret_deletion"
sidebar_current: "docs-volcengine-resource-kms_cancel_secret_deletion"
description: |-
  Provides a resource to manage kms cancel secret deletion
---
# volcengine_kms_cancel_secret_deletion
Provides a resource to manage kms cancel secret deletion
## Example Usage
```hcl
resource "volcengine_kms_cancel_secret_deletion" "default" {
  secret_name = "tf-test2"
}
```
## Argument Reference
The following arguments are supported:
* `secret_name` - (Required, ForceNew) The name of the secret.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `secret_state` - The state of the secret.


## Import
KmsCancelSecretDeletion can be imported using the id, e.g.
```
$ terraform import volcengine_kms_cancel_secret_deletion.default secret_name
```

