---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_re_encrypt"
sidebar_current: "docs-volcengine-resource-kms_re_encrypt"
description: |-
  Provides a resource to manage kms re encrypt
---
# volcengine_kms_re_encrypt
Provides a resource to manage kms re encrypt
## Example Usage
```hcl
resource "volcengine_kms_ciphertext" "encrypt_stable" {
  key_id    = "c44870c3-f33b-421a-****-a2bba37c993e"
  plaintext = "VGhpcyBpcyBhIHBsYWludGV4dCBleGFtcGxlLg=="
}

resource "volcengine_kms_re_encrypt" "re_encrypt_stable" {
  new_key_id             = "33e6ae1f-62f6-415a-****-579f526274cc"
  source_ciphertext_blob = volcengine_kms_ciphertext.encrypt_stable.ciphertext_blob
}
```
## Argument Reference
The following arguments are supported:
* `source_ciphertext_blob` - (Required, ForceNew) The source ciphertext, Base64 encoded.
* `new_encryption_context` - (Optional, ForceNew) The new encryption context JSON string.
* `new_key_id` - (Optional, ForceNew) The new key id. When new_key_id is not specified, both new_keyring_name and new_key_name must be specified.
* `new_key_name` - (Optional, ForceNew) The new key name.
* `new_keyring_name` - (Optional, ForceNew) The new keyring name.
* `old_encryption_context` - (Optional, ForceNew) The old encryption context JSON string.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `ciphertext_blob` - The re-encrypted ciphertext, Base64 encoded. The data stays stable across applies. If a changing ciphertext is needed use the `volcengine_kms_re_encrypts` data source.


## Import
The KmsReEncrypt is not support import.

