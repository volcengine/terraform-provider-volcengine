---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_key"
sidebar_current: "docs-volcengine-resource-kms_key"
description: |-
  Provides a resource to manage kms key
---
# volcengine_kms_key
Provides a resource to manage kms key
## Example Usage
```hcl
resource "volcengine_kms_keyring" "foo" {
  keyring_name = "tf-test"
  description  = "tf-test"
  project_name = "default"
}

resource "volcengine_kms_key" "foo" {
  keyring_name = volcengine_kms_keyring.foo.keyring_name
  key_name     = "mrk-tf-key-mod"
  description  = "tf test key-mod"
  tags {
    key   = "tfkey3"
    value = "tfvalue3"
  }
}
```
## Argument Reference
The following arguments are supported:
* `key_name` - (Required) The name of the CMK.
* `keyring_name` - (Required) The name of the keyring.
* `description` - (Optional) The description of the key.
* `key_spec` - (Optional) The type of the keys.
* `key_usage` - (Optional) The usage of the key.
* `multi_region` - (Optional) Whether it is the master key of the Multi-region type.
* `origin` - (Optional) The origin of the key.
* `pending_window_in_days` - (Optional) The pre-deletion cycle of the key.
* `protection_level` - (Optional) The protection level of the key.
* `rotate_state` - (Optional) The rotation state of the key.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `creation_date` - The date when the keyring was created.
* `key_material_expire_time` - The time when the key material will expire.
* `last_rotation_time` - The last time the key was rotated.
* `multi_region_configuration` - The configuration of Multi-region key.
    * `multi_region_key_type` - The type of the multi-region key.
    * `primary_key` - Trn and region id of the primary multi-region key.
        * `region` - The region id of multi-region key.
        * `trn` - The trn of multi-region key.
    * `replica_keys` - Trn and region id of replica multi-region keys.
        * `region` - The region id of multi-region key.
        * `trn` - The trn of multi-region key.
* `rotation_state` - The rotation configuration of the key.
* `schedule_delete_time` - The time when the key will be deleted.
* `schedule_rotation_time` - The next time the key will be rotated.
* `state` - The state of the key.
* `trn` - The name of the resource.
* `update_date` - The date when the keyring was updated.


## Import
KmsKey can be imported using the id, e.g.
```
$ terraform import volcengine_kms_key.default resource_id
```

