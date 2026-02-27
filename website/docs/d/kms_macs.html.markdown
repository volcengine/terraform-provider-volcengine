---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_macs"
sidebar_current: "docs-volcengine-datasource-kms_macs"
description: |-
  Use this data source to query detailed information of kms macs
---
# volcengine_kms_macs
Use this data source to query detailed information of kms macs
## Example Usage
```hcl
data "volcengine_kms_macs" "mac" {
  key_id        = "68093dd1-d1a9-44ce-****-5a88c4bc31ab"
  message       = "VGhpcyBpcyBhIHRlc3QgTWVzc2FnZS4="
  mac_algorithm = "HMAC_SHA_256"
}

data "volcengine_kms_mac_verifications" "verify" {
  key_id        = "68093dd1-d1a9-44ce-****-5a88c4bc31ab"
  message       = "VGhpcyBpcyBhIHRlc3QgTWVzc2FnZS4="
  mac_algorithm = "HMAC_SHA_256"
  mac           = "Vm0D9fk6uDRZD6k9QZE9+d9gpgy6ESSPt0bfaA2p05w="
}
```
## Argument Reference
The following arguments are supported:
* `mac_algorithm` - (Required) The MAC algorithm. Valid values: `HMAC_SM3`, `HMAC_SHA_256`.
* `message` - (Required) The message, Base64 encoded. Generate a Hash-based Message Authentication Code (HMAC) for a message using an HMAC KMS key and a MAC algorithm supported by the key.
* `key_id` - (Optional) The id of the key. When key_id is not specified, both keyring_name and key_name must be specified.
* `key_name` - (Optional) The name of the key.
* `keyring_name` - (Optional) The name of the keyring.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `mac_info` - The MAC info.
    * `key_id` - The key id.
    * `mac` - The MAC result, Base64 encoded.
* `total_count` - The total count of query.


