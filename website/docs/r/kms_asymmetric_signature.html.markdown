---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_asymmetric_signature"
sidebar_current: "docs-volcengine-resource-kms_asymmetric_signature"
description: |-
  Provides a resource to manage kms asymmetric signature
---
# volcengine_kms_asymmetric_signature
Provides a resource to manage kms asymmetric signature
## Example Usage
```hcl
resource "volcengine_kms_asymmetric_signature" "sign_stable1" {
  key_id       = "516274b3-0cba-4fad-****-c8355e3e8213"
  message      = "VGhpcyBpcyBhIG1lc3NhZ2UgZXhhbXBsZS4="
  message_type = "RAW"
  algorithm    = "RSA_PSS_SHA_256"
}

resource "volcengine_kms_asymmetric_signature" "sign_stable2" {
  key_id       = "516274b3-0cba-4fad-****-c8355e3e8213"
  message      = "KsFMwOobjOMHfYaPl2IgXX6tzziiT+SucmfmXTo2f6U="
  message_type = "DIGEST"
  algorithm    = "RSA_PSS_SHA_256"
}
```
## Argument Reference
The following arguments are supported:
* `algorithm` - (Required, ForceNew) The signing algorithm. valid values: `RSA_PSS_SHA_256`, `RSA_PKCS1_SHA_256`, `RSA_PSS_SHA_384`, `RSA_PKCS1_SHA_384`, `RSA_PSS_SHA_512`, `RSA_PKCS1_SHA_512`.
* `message` - (Required, ForceNew) The message to be signed, Base64 encoded.
* `key_id` - (Optional, ForceNew) The id of the key. When key_id is not specified, both keyring_name and key_name must be specified.
* `key_name` - (Optional, ForceNew) The name of the key.
* `keyring_name` - (Optional, ForceNew) The name of the keyring.
* `message_type` - (Optional, ForceNew) The type of message. Valid values: RAW or DIGEST. When message_type is DIGEST, KMS does not process the message digest of the original data source, it will sign directly with the private key.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `signature` - The signature, Base64 encoded. The produced signature stays stable across applies. If the message should be re-signed on each apply use the `volcengine_kms_asymmetric_signatures` data source.


## Import
The KmsAsymmetricSignature is not support import.

