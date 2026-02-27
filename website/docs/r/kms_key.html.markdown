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

resource "volcengine_kms_key" "foo1" {
  keyring_name     = volcengine_kms_keyring.foo.keyring_name
  key_name         = "Tf-test-key-1"
  rotate_state     = "Enable"
  rotate_interval  = 90
  key_spec         = "SYMMETRIC_128"
  description      = "Tf test key with SYMMETRIC_128"
  key_usage        = "ENCRYPT_DECRYPT"
  protection_level = "SOFTWARE"
  origin           = "CloudKMS"
  multi_region     = false
  #The scheduled deletion time when deleting the key
  pending_window_in_days = 30
  tags {
    key   = "tfk1"
    value = "tfv1"
  }
  tags {
    key   = "tfk2"
    value = "tfv2"
  }
}

resource "volcengine_kms_key" "foo2" {
  keyring_name = volcengine_kms_keyring.foo.keyring_name
  key_name     = "mrk-Tf-test-key-2"
  key_usage    = "ENCRYPT_DECRYPT"
  origin       = "External"
  multi_region = true
}

resource "volcengine_kms_key_material" "default" {
  keyring_name           = volcengine_kms_keyring.foo.keyring_name
  key_name               = volcengine_kms_key.foo2.key_name
  encrypted_key_material = "***"
  import_token           = "***"
  expiration_model       = "KEY_MATERIAL_EXPIRES"
  valid_to               = 1770999621
}
```
## Argument Reference
The following arguments are supported:
* `key_name` - (Required) The name of the key.
* `keyring_name` - (Required, ForceNew) The name of the keyring.
* `custom_key_store_id` - (Optional, ForceNew) The ID of the custom key store.
* `description` - (Optional) The description of the key.
* `key_spec` - (Optional, ForceNew) The type of the key. Valid values: SYMMETRIC_256, SYMMETRIC_128, RSA_2048, RSA_3072, RSA_4096, EC_P256K, EC_P256, EC_P384, EC_P521, EC_SM2. Default value: SYMMETRIC_256.
* `key_usage` - (Optional, ForceNew) The usage of the key. Valid values: ENCRYPT_DECRYPT, SIGN_VERIFY, GENERATE_VERIFY_MAC. Default value: ENCRYPT_DECRYPT.
* `multi_region` - (Optional, ForceNew) Whether it is the master key of the Multi-region type. When multi_region is true, the key name must start with "mrk-".
* `origin` - (Optional, ForceNew) The origin of the key. Valid values: CloudKMS, External, ExternalKeyStore. Default value: CloudKMS.
* `pending_window_in_days` - (Optional) The pre-deletion cycle of the key. Valid values: [7, 30]. Default value: 7.
* `protection_level` - (Optional, ForceNew) The protection level of the key. Valid values: SOFTWARE, HSM. Default value: SOFTWARE.
* `rotate_interval` - (Optional) Key rotation period, unit: days; value range: [90, 2560], required when rotate_state is Enable.
* `rotate_state` - (Optional) The rotation state of the key. Valid values: Enable, Disable. Only symmetric keys support rotation.
* `tags` - (Optional) Tags.
* `xks_key_id` - (Optional, ForceNew) The ID of the external key store.

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

