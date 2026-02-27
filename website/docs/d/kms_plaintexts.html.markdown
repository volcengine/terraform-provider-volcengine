---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_plaintexts"
sidebar_current: "docs-volcengine-datasource-kms_plaintexts"
description: |-
  Use this data source to query detailed information of kms plaintexts
---
# volcengine_kms_plaintexts
Use this data source to query detailed information of kms plaintexts
## Example Usage
```hcl
resource "volcengine_kms_ciphertext" "encrypt_stable" {
  key_id    = "c44870c3-f33b-421a-****-a2bba37c993e"
  plaintext = "VGhpcyBpcyBhIHBsYWludGV4dCBleGFtcGxlLg=="
}

data "volcengine_kms_plaintexts" "decrypt" {
  ciphertext_blob = volcengine_kms_ciphertext.encrypt_stable.ciphertext_blob
}
```
## Argument Reference
The following arguments are supported:
* `ciphertext_blob` - (Required) The ciphertext to be decrypted.
* `encryption_context` - (Optional) The JSON string of key/value pairs.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `plaintext_info` - The decrypted plaintext.
    * `plaintext` - The decrypted plaintext, Base64 encoded.
* `total_count` - The total count of query.


