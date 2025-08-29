---
subcategory: "TOS(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_tos_bucket_encryption"
sidebar_current: "docs-volcengine-resource-tos_bucket_encryption"
description: |-
  Provides a resource to manage tos bucket encryption
---
# volcengine_tos_bucket_encryption
Provides a resource to manage tos bucket encryption
## Example Usage
```hcl
resource "volcengine_tos_bucket" "foo" {
  bucket_name   = "tf-acc-test-bucket"
  public_acl    = "private"
  az_redundancy = "multi-az"
  project_name  = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_kms_keyring" "foo" {
  keyring_name = "acc-test-keyring"
  description  = "acc-test"
  project_name = "default"
}

resource "volcengine_tos_bucket_encryption" "foo" {
  bucket_name = volcengine_tos_bucket.foo.id
  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm       = "kms"
      kms_data_encryption = "AES256"
      kms_master_key_id   = volcengine_kms_keyring.foo.id
    }
  }
}
```
## Argument Reference
The following arguments are supported:
* `bucket_name` - (Required, ForceNew) The name of the bucket.
* `rule` - (Required) The rule of the bucket encryption.

The `apply_server_side_encryption_by_default` object supports the following:

* `sse_algorithm` - (Required) The server side encryption algorithm. Valid values: `kms`, `AES256`, `SM4`.
* `kms_data_encryption` - (Optional) The kms data encryption. Valid values: `AES256`, `SM4`. Default is `AES256`.
* `kms_master_key_id` - (Optional) The kms master key id. This field is required when `sse_algorithm` is `kms`. The format is `trn:kms:<region>:<accountID>:keyrings/<keyring>/keys/<key>`.

The `rule` object supports the following:

* `apply_server_side_encryption_by_default` - (Required) The server side encryption configuration.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
TosBucketEncryption can be imported using the id, e.g.
```
$ terraform import volcengine_tos_bucket_encryption.default resource_id
```

