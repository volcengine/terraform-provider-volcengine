---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_asymmetric_ciphertexts"
sidebar_current: "docs-volcengine-datasource-kms_asymmetric_ciphertexts"
description: |-
  Use this data source to query detailed information of kms asymmetric ciphertexts
---
# volcengine_kms_asymmetric_ciphertexts
Use this data source to query detailed information of kms asymmetric ciphertexts
## Example Usage
```hcl
data "volcengine_kms_asymmetric_ciphertexts" "encrypt1" {
  key_id    = "9601e1af-ad69-42df-****-eaf10ce6a3e9"
  plaintext = "VGhpcyBpcyBhIHBsYWludGV4dCBleGFtcGxlLg=="
  algorithm = "RSAES_OAEP_SHA_256"
}
data "volcengine_kms_asymmetric_ciphertexts" "encrypt2" {
  keyring_name = "Tf-test"
  key_name     = "ec-sm2"
  plaintext    = "VGhpcyBpcyBhIHBsYWludGV4dCBleGFtcGxlLg=="
  algorithm    = "SM2PKE"
}
```
## Argument Reference
The following arguments are supported:
* `algorithm` - (Required) The encryption algorithm. valid values: `RSAES_OAEP_SHA_256`, `SM2PKE`.
* `plaintext` - (Required) The plaintext to be encrypted, Base64 encoded.
* `key_id` - (Optional) The id of the key. When key_id is not specified, both keyring_name and key_name must be specified.
* `key_name` - (Optional) The name of the key.
* `keyring_name` - (Optional) The name of the keyring.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `ciphertext_info` - The information about the ciphertext.
    * `ciphertext_blob` - The ciphertext, Base64 encoded. The plaintext gets re-encrypted on each apply, resulting in a changed ciphertext_blob. If a stable ciphertext is needed use the `volcengine_kms_asymmetric_ciphertext` resource.
* `total_count` - The total count of query.


