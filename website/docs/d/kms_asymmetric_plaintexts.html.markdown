---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_asymmetric_plaintexts"
sidebar_current: "docs-volcengine-datasource-kms_asymmetric_plaintexts"
description: |-
  Use this data source to query detailed information of kms asymmetric plaintexts
---
# volcengine_kms_asymmetric_plaintexts
Use this data source to query detailed information of kms asymmetric plaintexts
## Example Usage
```hcl
resource "volcengine_kms_asymmetric_ciphertext" "encrypt_stable" {
  key_id    = "9601e1af-ad69-42df-****-eaf10ce6a3e9"
  plaintext = "VGhpcyBpcyBhIHBsYWludGV4dCBleGFtcGxLg=="
  algorithm = "RSAES_OAEP_SHA_256"
}

data "volcengine_kms_asymmetric_plaintexts" "decrypt" {
  key_id          = volcengine_kms_asymmetric_ciphertext.encrypt_stable.key_id
  ciphertext_blob = volcengine_kms_asymmetric_ciphertext.encrypt_stable.ciphertext_blob
  algorithm       = "RSAES_OAEP_SHA_256"
}
```
## Argument Reference
The following arguments are supported:
* `algorithm` - (Required) The encryption algorithm. valid values: `RSAES_OAEP_SHA_256`, `SM2PKE`.
* `ciphertext_blob` - (Required) The ciphertext to be decrypted, Base64 encoded.
* `key_id` - (Optional) The id of the key. When key_id is not specified, both keyring_name and key_name must be specified.
* `key_name` - (Optional) The name of the key.
* `keyring_name` - (Optional) The name of the keyring.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `plaintext_info` - The decrypted plaintext.
    * `plaintext` - The decrypted plaintext, Base64 encoded.
* `total_count` - The total count of query.


