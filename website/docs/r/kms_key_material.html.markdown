---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_key_material"
sidebar_current: "docs-volcengine-resource-kms_key_material"
description: |-
  Provides a resource to manage kms key material
---
# volcengine_kms_key_material
Provides a resource to manage kms key material
## Example Usage
```hcl
# It is necessary to first use data volcengine_kms_key_materials to obtain the import materials, such as import_token, public_key.
# Reference document: https://www.volcengine.com/docs/6476/144950?lang=zh
resource "volcengine_kms_key_material" "default" {
  keyring_name           = "Tf-test-1"
  key_name               = "Test-3"
  key_id                 = "8798cd1e-****-4f9b-****-d51847ad53ae"
  encrypted_key_material = "***"
  import_token           = "***"
  expiration_model       = "KEY_MATERIAL_EXPIRES"
  valid_to               = 1770969621
}
```
## Argument Reference
The following arguments are supported:
* `encrypted_key_material` - (Required, ForceNew) The encrypted key material, Base64 encoded.
* `import_token` - (Required, ForceNew) The import token.
* `expiration_model` - (Optional, ForceNew) The expiration model of key material. Valid values: `KEY_MATERIAL_DOES_NOT_EXPIRE`, `KEY_MATERIAL_EXPIRES`. Default value: `KEY_MATERIAL_DOES_NOT_EXPIRE`.
* `key_id` - (Optional, ForceNew) The id of key. When key_id is not specified, both keyring_name and key_name must be specified.
* `key_name` - (Optional, ForceNew) The name of key.
* `keyring_name` - (Optional, ForceNew) The name of keyring.
* `valid_to` - (Optional, ForceNew) The valid to timestamp of key material. Required when expiration_model is KEY_MATERIAL_EXPIRES. Unit: second.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



