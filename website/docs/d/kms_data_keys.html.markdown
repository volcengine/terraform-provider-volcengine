---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_data_keys"
sidebar_current: "docs-volcengine-datasource-kms_data_keys"
description: |-
  Use this data source to query detailed information of kms data keys
---
# volcengine_kms_data_keys
Use this data source to query detailed information of kms data keys
## Example Usage
```hcl
data "volcengine_kms_data_keys" "data_key" {
  key_id          = "c44870c3-f33b-421a-****-a2bba37c993e"
  number_of_bytes = 1024
}

data "volcengine_kms_plaintexts" "default" {
  ciphertext_blob = data.volcengine_kms_data_keys.data_key.data_key_info[0].ciphertext_blob
}
```
## Argument Reference
The following arguments are supported:
* `encryption_context` - (Optional) The JSON string of key/value pairs.
* `key_id` - (Optional) The id of the key. When key_id is not specified, both keyring_name and key_name must be specified.
* `key_name` - (Optional) The name of the key. Only symmetric key is supported.
* `keyring_name` - (Optional) The name of the keyring.
* `number_of_bytes` - (Optional) The length of data key to generate.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `data_key_info` - The data key info.
    * `ciphertext_blob` - The generated ciphertext, Base64 encoded.
    * `plaintext` - The generated plaintext, Base64 encoded.
* `total_count` - The total count of query.


