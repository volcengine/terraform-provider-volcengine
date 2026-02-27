---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_asymmetric_verifications"
sidebar_current: "docs-volcengine-datasource-kms_asymmetric_verifications"
description: |-
  Use this data source to query detailed information of kms asymmetric verifications
---
# volcengine_kms_asymmetric_verifications
Use this data source to query detailed information of kms asymmetric verifications
## Example Usage
```hcl
data "volcengine_kms_asymmetric_verifications" "verify" {
  key_id       = "516274b3-0cba-4fad-****-c8355e3e8213"
  message      = "VGhpcyBpcyBhIG1lc3NhZ2UgZXhhbXBsZS4="
  message_type = "RAW"
  algorithm    = "RSA_PSS_SHA_256"
  signature    = "UPeR+U1fq+Om6jKy/VUz1xsTMJijfJNMy/p63uozgX5wVJR7Z6anvAvA4Yw8/Z32eXpyF0fgBsz+VbDSpfIet8rg4W2PoMQJNSAeg9UG4liV4CnYyYJIPopYUhHdDFSi/K/XguqD6IosMQQwHzkK1YneEBYTKySLypYAbwiVQHOZMNbF6HkmMg+P6qK7ircuwtG5I7vFVt//Nk2Elj6s66V1luNTlab1BfmdtCCVeYyHh4tGn5s0kyhahK3eYsnOC6ZE0AO4R4SNlLnWAN6BHmZgQvGLkPP/C0gnDlouiF1I7gNx/jG8GtDB+JqvDvsGviXMLoqWK1/fyydY/OcmVdzZnG6i6NlGLFXZqLT5K2ILf/i/w/BOjQmVyV9GpJJXaCSjy7Mq0aX/DGvfHNdmEDZew3KkMTeGynJUJFYaiTV5ToFcP9x9Vw8gCDqGdEAvaoFmTTHFcOVdwGyo7n9g7sagxtIYReQPsKePnTb18QOIHUbUc7BwqAPBa7gDX0XJsuockGPULsF7SuIHFHtNxYs5UZgrLmCPn49Xw0o+bJemyBixZcMIzMnwGUS0Ew9te+5is1swdIUQFPc4KZG0ohLXyRrxQK7rSpVTWM0ggy2OWUbm/X6kMsLmDueF40SZdqibYsNpWtIUeJatAR9adG5p8NcdvvY4S0/k7KjfQtw="
}
```
## Argument Reference
The following arguments are supported:
* `algorithm` - (Required) The signing algorithm. valid values: `RSA_PSS_SHA_256`, `RSA_PKCS1_SHA_256`, `RSA_PSS_SHA_384`, `RSA_PKCS1_SHA_384`, `RSA_PSS_SHA_512`, `RSA_PKCS1_SHA_512`.
* `message` - (Required) The message to be verified, Base64 encoded.
* `signature` - (Required) The signature to be verified, Base64 encoded.
* `key_id` - (Optional) The id of the key. When key_id is not specified, both keyring_name and key_name must be specified.
* `key_name` - (Optional) The name of the key.
* `keyring_name` - (Optional) The name of the keyring.
* `message_type` - (Optional) The type of message. Valid values: RAW or DIGEST.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of query.
* `verification_info` - The verification result.
    * `signature_valid` - Whether the signature is valid.


