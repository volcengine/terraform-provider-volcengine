---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_mac_verifications"
sidebar_current: "docs-volcengine-datasource-kms_mac_verifications"
description: |-
  Use this data source to query detailed information of kms mac verifications
---
# volcengine_kms_mac_verifications
Use this data source to query detailed information of kms mac verifications
## Example Usage
```hcl
data "volcengine_kms_macs" "mac" {
  key_id        = "68093dd1-d1a9-44ce-832a****-5a88c4bc31ab"
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
* `mac` - (Required) The MAC to verify, Base64 encoded. Verify the Hash-based Message Authentication Code (HMAC), HMAC KMS key, and MAC algorithm for the specified message.
* `message` - (Required) The message to verify, Base64 encoded.
* `key_id` - (Optional) The id of the key. When key_id is not specified, both keyring_name and key_name must be specified.
* `key_name` - (Optional) The name of the key.
* `keyring_name` - (Optional) The name of the keyring.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `mac_verification_info` - The MAC verification info.
    * `key_id` - The key id.
    * `mac_valid` - Whether the MAC is valid.
* `total_count` - The total count of query.


