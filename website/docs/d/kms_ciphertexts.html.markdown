---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_ciphertexts"
sidebar_current: "docs-volcengine-datasource-kms_ciphertexts"
description: |-
  Use this data source to query detailed information of kms ciphertexts
---
# volcengine_kms_ciphertexts
Use this data source to query detailed information of kms ciphertexts
## Example Usage
```hcl
data "volcengine_kms_ciphertexts" "encrypted" {
  key_id    = "c44870c3-f33b-421a-****-a2bba37c993e"
  plaintext = "VGhpcyBpcyBhIHBsYWludGV4dCBleGFtcGxlLg=="
}
```
## Argument Reference
The following arguments are supported:
* `plaintext` - (Required) The plaintext to be encrypted, Base64 encoded.
* `encryption_context` - (Optional) The JSON string of key/value pairs.
* `key_id` - (Optional) The id of the key. When key_id is not specified, both keyring_name and key_name must be specified.
* `key_name` - (Optional) The name of the key.
* `keyring_name` - (Optional) The name of the keyring.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `ciphertext_info` - The information about the ciphertext.
    * `ciphertext_blob` - The ciphertext, Base64 encoded. The plaintext gets re-encrypted on each apply, resulting in a changed ciphertext_blob. If a stable ciphertext is needed use the `volcengine_kms_ciphertext` resource.
* `total_count` - The total count of query.


