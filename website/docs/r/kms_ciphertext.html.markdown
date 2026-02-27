---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_ciphertext"
sidebar_current: "docs-volcengine-resource-kms_ciphertext"
description: |-
  Provides a resource to manage kms ciphertext
---
# volcengine_kms_ciphertext
Provides a resource to manage kms ciphertext
## Example Usage
```hcl
resource "volcengine_kms_ciphertext" "encrypt_stable" {
  key_id    = "c44870c3-f33b-421a-****-a2bba37c993e"
  plaintext = "VGhpcyBpcyBhIHBsYWludGV4dCBleGFtcGxlLg=="
}
```
## Argument Reference
The following arguments are supported:
* `plaintext` - (Required, ForceNew) The plaintext to be symmetrically encrypted, Base64 encoded.
* `encryption_context` - (Optional, ForceNew) The JSON string of key/value pairs.
* `key_id` - (Optional, ForceNew) The id of the key. When key_id is not specified, both keyring_name and key_name must be specified.
* `key_name` - (Optional, ForceNew) The name of the key.
* `keyring_name` - (Optional, ForceNew) The name of the keyring.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `ciphertext_blob` - The ciphertext, Base64 encoded. The produced ciphertext_blob stays stable across applies. If the plaintext should be re-encrypted on each apply use the `volcengine_kms_ciphertexts` data source.


## Import
The KmsCiphertext is not support import.

