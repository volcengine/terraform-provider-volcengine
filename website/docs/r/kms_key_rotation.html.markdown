---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_key_rotation"
sidebar_current: "docs-volcengine-resource-kms_key_rotation"
description: |-
  Provides a resource to manage kms key rotation
---
# volcengine_kms_key_rotation
Provides a resource to manage kms key rotation
## Example Usage
```hcl
resource "volcengine_kms_key_rotation" "foo" {
  key_id = "m_cn-guilin-boe_63c08fe9-42e8-4c10-a09e-8e8e6xxxxxx"
}
```
## Argument Reference
The following arguments are supported:
* `key_id` - (Optional) The id of the CMK.
* `key_name` - (Optional) The name of the CMK.
* `keyring_name` - (Optional) The name of the keyring.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `rotation_state` - The state of the key rotation.


## Import
KmsKeyRotation can be imported using the id, e.g.
```
$ terraform import volcengine_kms_key_rotation.default resource_id
or
$ terraform import volcengine_kms_key_rotation.default key_name:keyring_name
```

