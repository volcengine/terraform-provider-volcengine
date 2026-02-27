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
  key_id          = "c44870c3-f33b-421a-****-a2bba37c993e"
  rotate_interval = 90
}
```
## Argument Reference
The following arguments are supported:
* `key_id` - (Optional, ForceNew) The id of the key. When key_id is not specified, both keyring_name and key_name must be specified.
* `key_name` - (Optional, ForceNew) The name of the key.
* `keyring_name` - (Optional, ForceNew) The name of the keyring.
* `rotate_interval` - (Optional) Key rotation period, unit: days; value range: [90, 2560].

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

