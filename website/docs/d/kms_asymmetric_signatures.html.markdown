---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_asymmetric_signatures"
sidebar_current: "docs-volcengine-datasource-kms_asymmetric_signatures"
description: |-
  Use this data source to query detailed information of kms asymmetric signatures
---
# volcengine_kms_asymmetric_signatures
Use this data source to query detailed information of kms asymmetric signatures
## Example Usage
```hcl
data "volcengine_kms_asymmetric_signatures" "sign1" {
  key_id       = "516274b3-0cba-4fad-****-c8355e3e8213"
  message      = "VGhpcyBpcyBhIG1lc3NhZ2UgZXhhbXBsZS4="
  message_type = "RAW"
  algorithm    = "RSA_PSS_SHA_256"
}

data "volcengine_kms_asymmetric_signatures" "sign2" {
  key_id       = "516274b3-0cba-4fad-****-c8355e3e8213"
  message      = "KsFMwOobjOMHfYaPl2IgXX6tzziiT+SucmfmXTo2f6U="
  message_type = "DIGEST"
  algorithm    = "RSA_PSS_SHA_256"
}
```
## Argument Reference
The following arguments are supported:
* `algorithm` - (Required) The signing algorithm. valid values: `RSA_PSS_SHA_256`, `RSA_PKCS1_SHA_256`, `RSA_PSS_SHA_384`, `RSA_PKCS1_SHA_384`, `RSA_PSS_SHA_512`, `RSA_PKCS1_SHA_512`.
* `message` - (Required) The message to be signed, Base64 encoded.
* `key_id` - (Optional) The id of the key. When key_id is not specified, both keyring_name and key_name must be specified.
* `key_name` - (Optional) The name of the key.
* `keyring_name` - (Optional) The name of the keyring.
* `message_type` - (Optional) The type of message. Valid values: RAW or DIGEST. When message_type is DIGEST, KMS does not process the message digest of the original data source, it will sign directly with the private key.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `signature_info` - The information about the signature.
    * `signature` - The signature, Base64 encoded. The signature gets re-signed on each apply, resulting in a changed signature. If a stable signature is needed use the `volcengine_kms_asymmetric_signature` resource.
* `total_count` - The total count of query.


