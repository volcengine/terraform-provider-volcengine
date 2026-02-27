---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_key_materials"
sidebar_current: "docs-volcengine-datasource-kms_key_materials"
description: |-
  Use this data source to query detailed information of kms key materials
---
# volcengine_kms_key_materials
Use this data source to query detailed information of kms key materials
## Example Usage
```hcl
data "volcengine_kms_key_materials" "default" {
  keyring_name       = "Tf-test-1"
  key_name           = "Test-3"
  wrapping_key_spec  = "RSA_2048"
  wrapping_algorithm = "RSAES_OAEP_SHA_256"
}
```
## Argument Reference
The following arguments are supported:
* `key_id` - (Optional) The id of key. When key_id is not specified, both keyring_name and key_name must be specified.
* `key_name` - (Optional) The name of key.
* `keyring_name` - (Optional) The name of keyring.
* `output_file` - (Optional) File name where to save data source results.
* `wrapping_algorithm` - (Optional) The wrapping algorithm. Valid values: `RSAES_OAEP_SHA_256`, `RSAES_OAEP_SHA_1`, `RSAES_PKCS1_V1_5`, `SM2PKE`. Default value: `RSAES_OAEP_SHA_256`. When the wrapping_key_spec is EC_SM2, only SM2PKE is supported.
* `wrapping_key_spec` - (Optional) The wrapping key spec. Valid values: `RSA_2048`, `EC_SM2`. Default value: `RSA_2048`. When the user's master key protection level is SOFTWARE, selecting EC_SM2 is prohibited.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `import_parameters` - The import parameters info.
    * `import_token` - The import token, Base64 encoded.
    * `key_id` - The id of key.
    * `keyring_id` - The id of keyring.
    * `public_key` - The public key used to encrypt key materials, Base64 encoded.
    * `token_expire_time` - The token expire time.
* `total_count` - The total count of query.


