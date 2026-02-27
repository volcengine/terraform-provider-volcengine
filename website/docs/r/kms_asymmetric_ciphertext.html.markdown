---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_asymmetric_ciphertext"
sidebar_current: "docs-volcengine-resource-kms_asymmetric_ciphertext"
description: |-
  Provides a resource to manage kms asymmetric ciphertext
---
# volcengine_kms_asymmetric_ciphertext
Provides a resource to manage kms asymmetric ciphertext
## Example Usage
```hcl
resource "volcengine_kms_asymmetric_ciphertext" "encrypt1" {
  key_id    = "9601e1af-ad69-42df-****-eaf10ce6a3e9"
  plaintext = "VGhpcyBpcyBhIHBsYWludGV4dCBleGFtcGxlLg=="
  algorithm = "RSAES_OAEP_SHA_256"
}
resource "volcengine_kms_asymmetric_ciphertext" "encrypt2" {
  keyring_name = "Tf-test"
  key_name     = "ec-sm2"
  plaintext    = "VGhpcyBpcyBhIHBsYWludGV4dCBleGFtcGxlLg=="
  algorithm    = "SM2PKE"
}
```
## Argument Reference
The following arguments are supported:
* `algorithm` - (Required, ForceNew) The encryption algorithm. valid values: `RSAES_OAEP_SHA_256`, `SM2PKE`.
* `plaintext` - (Required, ForceNew) The plaintext to be encrypted, Base64 encoded.
* `key_id` - (Optional, ForceNew) The id of the key. When key_id is not specified, both keyring_name and key_name must be specified.
* `key_name` - (Optional, ForceNew) The name of the key.
* `keyring_name` - (Optional, ForceNew) The name of the keyring.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `ciphertext_blob` - The ciphertext, Base64 encoded. The produced ciphertext_blob stays stable across applies. If the plaintext should be re-encrypted on each apply use the `volcengine_kms_asymmetric_ciphertexts` data source.


## Import
The KmsAsymmetricCiphertext is not support import.

