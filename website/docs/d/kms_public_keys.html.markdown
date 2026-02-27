---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_public_keys"
sidebar_current: "docs-volcengine-datasource-kms_public_keys"
description: |-
  Use this data source to query detailed information of kms public keys
---
# volcengine_kms_public_keys
Use this data source to query detailed information of kms public keys
## Example Usage
```hcl
# Obtain the public key of the specified asymmetric key
data "volcengine_kms_public_keys" "default" {
  keyring_name = "Tf-test"
  key_name     = "Test-key2"
}
```
## Argument Reference
The following arguments are supported:
* `key_id` - (Optional) The id of key. When key_id is not specified, both keyring_name and key_name must be specified.
* `key_name` - (Optional) The name of key.
* `keyring_name` - (Optional) The name of keyring.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `public_key` - The public key info.
    * `key_id` - The id of key.
    * `public_key` - The public key in PEM format.
* `total_count` - The total count of query.


