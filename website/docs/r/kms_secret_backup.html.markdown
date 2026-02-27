---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_secret_backup"
sidebar_current: "docs-volcengine-resource-kms_secret_backup"
description: |-
  Provides a resource to manage kms secret backup
---
# volcengine_kms_secret_backup
Provides a resource to manage kms secret backup
## Example Usage
```hcl
resource "volcengine_kms_secret_backup" "default" {
  secret_name = "Test-1"
}
```
## Argument Reference
The following arguments are supported:
* `secret_name` - (Required, ForceNew) The name of the secret to backup.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `backup_data` - The full backup data of the secret. JSON format.
* `secret_data_key` - The ciphertext of the data key used to encrypt the secret value, Base64 encoded.
* `signature` - The signature of the backup_data. Base64 encoded.


## Import
KmsSecretBackup can be imported using the secret_name, e.g.
```
$ terraform import volcengine_kms_secret_backup.default ecs-secret-test
```

