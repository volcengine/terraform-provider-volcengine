---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_re_encrypts"
sidebar_current: "docs-volcengine-datasource-kms_re_encrypts"
description: |-
  Use this data source to query detailed information of kms re encrypts
---
# volcengine_kms_re_encrypts
Use this data source to query detailed information of kms re encrypts
## Example Usage
```hcl
resource "volcengine_kms_ciphertext" "encrypt_stable" {
  key_id    = "c44870c3-f33b-421a-****-a2bba37c993e"
  plaintext = "VGhpcyBpcyBhIHBsYWludGV4dCBleGFtcGxlLg=="
}

data "volcengine_kms_re_encrypts" "re_encrypt_changing" {
  new_key_id             = "33e6ae1f-62f6-415a-****-579f526274cc"
  source_ciphertext_blob = volcengine_kms_ciphertext.encrypt_stable.ciphertext_blob
}
```
## Argument Reference
The following arguments are supported:
* `source_ciphertext_blob` - (Required) The ciphertext data to be re-encrypted, Base64 encoded.
* `new_encryption_context` - (Optional) The new encryption context JSON string of key/value pairs.
* `new_key_id` - (Optional) The new key id. When new_key_id is not specified, both new_keyring_name and new_key_name must be specified.
* `new_key_name` - (Optional) The new key name.
* `new_keyring_name` - (Optional) The new keyring name.
* `old_encryption_context` - (Optional) The old encryption context JSON string of key/value pairs.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `ciphertext_info` - The information about the ciphertext.
    * `ciphertext_blob` - The re-encrypted ciphertext, Base64 encoded. The data gets re-encrypted on each apply, resulting in a changed ciphertext_blob. If a stable ciphertext is needed use the `volcengine_kms_re_encrypt` resource.
* `total_count` - The total count of query.


