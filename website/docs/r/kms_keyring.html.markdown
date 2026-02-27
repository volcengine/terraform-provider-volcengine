---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_keyring"
sidebar_current: "docs-volcengine-resource-kms_keyring"
description: |-
  Provides a resource to manage kms keyring
---
# volcengine_kms_keyring
Provides a resource to manage kms keyring
## Example Usage
```hcl
resource "volcengine_kms_keyring" "foo" {
  keyring_name = "tf-test"
  description  = "tf-test"
  project_name = "default"
}
```
## Argument Reference
The following arguments are supported:
* `keyring_name` - (Required, ForceNew) The name of the keyring. Note: the keyring can only be deleted after all keys in it have been removed.
* `description` - (Optional) The description of the keyring.
* `keyring_type` - (Optional, ForceNew) The type of the keyring.
* `project_name` - (Optional) The name of the project.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `creation_date` - The date when the keyring was created.
* `trn` - The information about the tenant resource name (TRN).
* `uid` - The tenant ID of the keyring.
* `update_date` - The date when the keyring was updated.


## Import
KmsKeyring can be imported using the id, e.g.
```
$ terraform import volcengine_kms_keyring.default resource_id
```

