---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_replicate_key"
sidebar_current: "docs-volcengine-resource-kms_replicate_key"
description: |-
  Provides a resource to manage kms replicate key
---
# volcengine_kms_replicate_key
Provides a resource to manage kms replicate key
## Example Usage
```hcl
# Only create a backup key in the replica region;
# Next, managing key requires the use of resource "volcengine_kms_key".
resource "volcengine_kms_replicate_key" "replica" {
  keyring_name   = "test"
  key_name       = "mrk-Tf-Test-1"
  replica_region = "cn-shanghai"
  description    = "replica description"
  tags {
    key   = "tfk1"
    value = "tfv1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `replica_region` - (Required, ForceNew) The target region for replica key.
* `description` - (Optional, ForceNew) The description of the replicated regional key.
* `key_id` - (Optional, ForceNew) The id of the key. When key_id is not specified, both keyring_name and key_name must be specified.
* `key_name` - (Optional, ForceNew) The name of the key. Note: Only multi-region keys support replication.
* `keyring_name` - (Optional, ForceNew) The name of the keyring.
* `tags` - (Optional, ForceNew) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `replica_key_id` - The id of the replica key.


## Import
The KmsReplicateKey is not support imported.

