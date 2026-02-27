---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_key_enable"
sidebar_current: "docs-volcengine-resource-kms_key_enable"
description: |-
  Provides a resource to manage kms key enable
---
# volcengine_kms_key_enable
Provides a resource to manage kms key enable
## Example Usage
```hcl
resource "volcengine_kms_key_enable" "foo" {
  key_id = "0e5a256d-d075-44b1-bcd2-09efafxxxxxx"
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
KmsKeyEnable can be imported using the id, e.g.
```
$ terraform import volcengine_kms_key_enable.default resource_id
or
$ terraform import volcengine_kms_key_enable.default key_name:keyring_name
```

