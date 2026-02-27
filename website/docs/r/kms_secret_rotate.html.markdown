---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_secret_rotate"
sidebar_current: "docs-volcengine-resource-kms_secret_rotate"
description: |-
  Provides a resource to manage kms secret rotate
---
# volcengine_kms_secret_rotate
Provides a resource to manage kms secret rotate
## Example Usage
```hcl
resource "volcengine_kms_secret_rotate" "default" {
  secret_name  = "ecs-secret-test"
  version_name = "v1"
}
```
## Argument Reference
The following arguments are supported:
* `secret_name` - (Required) The name of the secret to manually rotate.
* `version_name` - (Optional) The version alias after rotation. Manual rotation can be triggered by modifying version_name.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
KmsSecretRotate can be imported using the secret_name, e.g.
```
$ terraform import volcengine_kms_secret_rotate.default ecs-secret-test
```

