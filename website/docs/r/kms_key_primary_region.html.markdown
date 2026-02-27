---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_key_primary_region"
sidebar_current: "docs-volcengine-resource-kms_key_primary_region"
description: |-
  Provides a resource to manage kms key primary region
---
# volcengine_kms_key_primary_region
Provides a resource to manage kms key primary region
## Example Usage
```hcl
provider "volcengine" {
  region = "cn-beijing"
}
resource "volcengine_kms_key_primary_region" "primary" {
  keyring_name = "test"
  key_name     = "mrk-Tf-test"
  # Note: The primary region is switched from cn-beijing to cn-shanghai, so cn-beijing will become a replica key 
  # and will no longer support operations such as key rotation and switching the primary region.
  # To continue operating on the key, need to switch the region to cn-shanghai.
  primary_region = "cn-shanghai"
}
```
## Argument Reference
The following arguments are supported:
* `primary_region` - (Required, ForceNew) The new primary region.
* `key_id` - (Optional, ForceNew) The id of the key. When key_id is not specified, both keyring_name and key_name must be specified.
* `key_name` - (Optional, ForceNew) The name of the key. Note: Only multi-region keys support updating primary region.
* `keyring_name` - (Optional, ForceNew) The name of the keyring.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
KmsKeyPrimaryRegion can be imported using the id, e.g.
```
$ terraform import volcengine_kms_key_primary_region.default key_id
or
$ terraform import volcengine_kms_key_primary_region.default key_name:keyring_name
```

