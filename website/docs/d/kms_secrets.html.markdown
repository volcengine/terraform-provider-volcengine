---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_secrets"
sidebar_current: "docs-volcengine-datasource-kms_secrets"
description: |-
  Use this data source to query detailed information of kms secrets
---
# volcengine_kms_secrets
Use this data source to query detailed information of kms secrets
## Example Usage
```hcl
data "volcengine_kms_secrets" "default" {
  secret_name = ["5r3", "5r", "tf"]
  description = ["tf-1", "tf-2"]
}
```
## Argument Reference
The following arguments are supported:
* `creation_date_range` - (Optional) The creation date of the secret.
* `description` - (Optional) The description of the secret.
* `managed_state` - (Optional) The state of the managed.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `owning_service` - (Optional) The cloud service that owns the secret.
* `project_name` - (Optional) The name of the project to which the secret belongs.
* `rotation_state` - (Optional) The state of the rotation.
* `secret_name` - (Optional) The name of the secret.
* `secret_state` - (Optional) The state of the secret.
* `secret_type` - (Optional) The type of the secret.
* `trn` - (Optional) The trn of the secret.
* `update_date_range` - (Optional) The update date of the secret.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `secrets` - The information about the secret.
    * `creation_date` - The date when the keyring was created.
    * `description` - The description of the secret.
    * `encryption_key` - The TRN of the KMS key used to encrypt the secret value.
    * `extended_config` - The extended configurations of the secret.
    * `id` - The unique ID of the secret. The value is in the UUID format.
    * `last_rotation_time` - The last time the secret was rotated.
    * `managed` - Indicates whether the secret is hosted.
    * `owning_service` - The cloud service that owns the secret.
    * `project_name` - The project name of the secret.
    * `rotation_interval` - The interval at which automatic rotation is performed.
    * `rotation_state` - The rotation state of the secret.
    * `schedule_delete_time` - The time when the secret will be deleted.
    * `schedule_rotation_time` - The next time the secret will be rotated.
    * `secret_name` - The name of the secret.
    * `secret_state` - The state of secret.
    * `secret_type` - The type of the secret.
    * `trn` - The information about the tenant resource name (TRN).
    * `uid` - The tenant ID of the secret.
    * `update_date` - The date when the keyring was updated.
* `total_count` - The total count of query.


