---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_secret_restore"
sidebar_current: "docs-volcengine-resource-kms_secret_restore"
description: |-
  Provides a resource to manage kms secret restore
---
# volcengine_kms_secret_restore
Provides a resource to manage kms secret restore
## Example Usage
```hcl
resource "volcengine_kms_secret_backup" "example" {
  secret_name = "Test-1"
}

resource "volcengine_kms_secret_restore" "default" {
  secret_data_key = volcengine_kms_secret_backup.example.secret_data_key
  backup_data     = volcengine_kms_secret_backup.example.backup_data
  signature       = volcengine_kms_secret_backup.example.signature
}
```
## Argument Reference
The following arguments are supported:
* `backup_data` - (Required, ForceNew) The full secret data returned during backup. JSON format.
* `secret_data_key` - (Required, ForceNew) The data key ciphertext returned during backup. Base64 encoded.
* `signature` - (Required, ForceNew) The signature of the backup data returned during backup. Base64 encoded.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
The KmsSecretRestore is not support import.

