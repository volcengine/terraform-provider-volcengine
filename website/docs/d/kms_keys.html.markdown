---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_keys"
sidebar_current: "docs-volcengine-datasource-kms_keys"
description: |-
  Use this data source to query detailed information of kms keys
---
# volcengine_kms_keys
Use this data source to query detailed information of kms keys
## Example Usage
```hcl
data "volcengine_kms_keys" "default" {
  keyring_id          = "7a358829-bd5a-4763-ba77-7500ecxxxxxx"
  key_name            = ["mrk-tf-key-mod", "mrk-tf-key"]
  key_spec            = ["SYMMETRIC_256"]
  description         = ["tf-test"]
  key_state           = ["Enable"]
  key_usage           = ["ENCRYPT_DECRYPT"]
  protection_level    = ["SOFTWARE"]
  rotate_state        = ["Enable"]
  origin              = ["CloudKMS"]
  creation_date_range = ["2025-06-01 19:48:06", "2025-06-04 19:48:06"]
  update_date_range   = ["2025-06-01 19:48:06", "2025-06-04 19:48:06"]
  tags {
    key    = "tf-k1"
    values = ["tf-v1"]
  }
}
```
## Argument Reference
The following arguments are supported:
* `creation_date_range` - (Optional) The creation date of the keyring.
* `description` - (Optional) The description of the key.
* `key_name` - (Optional) The name of the key.
* `key_spec` - (Optional) The algorithm used in the key.
* `key_state` - (Optional) The state of the key.
* `key_usage` - (Optional) The usage of the key.
* `keyring_id` - (Optional) Query the Key ring that meets the specified conditions, which is composed of key-value pairs.
* `keyring_name` - (Optional) Query the Key ring that meets the specified conditions, which is composed of key-value pairs.
* `name_regex` - (Optional) A Name Regex of Resource.
* `origin` - (Optional) The origin of the key.
* `output_file` - (Optional) File name where to save data source results.
* `protection_level` - (Optional) The protection level of the key.
* `rotate_state` - (Optional) The state of the rotate.
* `tags` - (Optional) A list of tags.
* `update_date_range` - (Optional) The update date of the keyring.

The `tags` object supports the following:

* `key` - (Required) The key of the tag.
* `values` - (Required) The values of the tag.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `keys` - Master key list information.
    * `creation_date` - The date when the keyring was created.
    * `description` - The description of the key.
    * `id` - The unique ID of the key.
    * `key_material_expire_time` - The time when the key material will expire.
    * `key_name` - The name of the key.
    * `key_spec` - The algorithm used in the key.
    * `key_state` - The state of the key.
    * `key_usage` - The usage of the key.
    * `last_rotation_time` - The last time the key was rotated.
    * `multi_region_configuration` - The configuration of Multi-region key.
        * `multi_region_key_type` - The type of the multi-region key.
        * `primary_key` - Trn and region id of the primary multi-region key.
            * `region` - The region id of multi-region key.
            * `trn` - The trn of multi-region key.
        * `replica_keys` - Trn and region id of replica multi-region keys.
            * `region` - The region id of multi-region key.
            * `trn` - The trn of multi-region key.
    * `multi_region` - Whether it is the master key of the Multi-region type.
    * `origin` - The origin of the key.
    * `protection_level` - The protection level of the key.
    * `rotation_state` - The rotation configuration of the key.
    * `schedule_delete_time` - The time when the key will be deleted.
    * `schedule_rotation_time` - The next time the key will be rotated.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `trn` - The name of the resource.
    * `update_date` - The date when the keyring was updated.
* `total_count` - The total count of query.


