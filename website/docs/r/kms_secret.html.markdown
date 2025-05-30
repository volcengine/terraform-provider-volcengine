---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_secret"
sidebar_current: "docs-volcengine-resource-kms_secret"
description: |-
  Provides a resource to manage kms secret
---
# volcengine_kms_secret
Provides a resource to manage kms secret
## Example Usage
```hcl
resource "volcengine_kms_secret" "foo" {
  secret_name  = "tf-test1"
  secret_type  = "Generic"
  description  = "tf-test"
  secret_value = "{\"dasdasd\":\"dasdasd\"}"
}
```
## Argument Reference
The following arguments are supported:
* `secret_name` - (Required, ForceNew) The name of the secret.
* `secret_type` - (Required, ForceNew) The type of the secret.
* `secret_value` - (Required, ForceNew) The value of the secret.
* `automatic_rotation` - (Optional) The rotation state of the secret.
* `description` - (Optional) The description of the secret.
* `encryption_key` - (Optional, ForceNew) The TRN of the KMS key used to encrypt the secret value.
* `extended_config` - (Optional, ForceNew) The extended configurations of the secret.
* `project_name` - (Optional) The project name of the secret.
* `rotation_interval` - (Optional) The interval at which automatic rotation is performed.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `creation_date` - The date when the secret was created.
* `last_rotation_time` - The rotation state of the secret.
* `managed` - Indicates whether the secret is hosted.
* `rotation_interval_second` - Rotation interval second.
* `rotation_state` - The rotation state of the secret.
* `schedule_delete_time` - The time when the secret will be deleted.
* `schedule_rotation_time` - The next time the secret will be rotated.
* `secret_state` - The state of secret.
* `trn` - The information about the tenant resource name (TRN).
* `uid` - The tenant ID of the secret.
* `update_date` - The date when the secret was updated.
* `uuid` - The ID of secret.


## Import
KmsSecret can be imported using the id, e.g.
```
$ terraform import volcengine_kms_secret.default resource_id
```

