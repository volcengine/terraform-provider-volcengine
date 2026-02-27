---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_cancel_key_deletion"
sidebar_current: "docs-volcengine-resource-kms_cancel_key_deletion"
description: |-
  Provides a resource to manage kms cancel key deletion
---
# volcengine_kms_cancel_key_deletion
Provides a resource to manage kms cancel key deletion
## Example Usage
```hcl
resource "volcengine_kms_cancel_key_deletion" "foo" {
  key_id = "50f588aa-32e1-4cd1-****-63afcbc7d523"
}
```
## Argument Reference
The following arguments are supported:
* `key_id` - (Optional, ForceNew) The id of the key. When key_id is not specified, both keyring_name and key_name must be specified.
* `key_name` - (Optional, ForceNew) The name of the key.
* `keyring_name` - (Optional, ForceNew) The name of the keyring.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `key_state` - The state of the key.


## Import
KmsCancelKeyDeletion can be imported using the id, e.g.
```
$ terraform import volcengine_kms_cancel_key_deletion.default resource_id
or
$ terraform import volcengine_kms_cancel_key_deletion.default key_name:keyring_name
```

