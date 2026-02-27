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
  version_name = "v1.0"
}

resource "volcengine_kms_secret" "foo_ecs" {
  secret_name            = "tf-test2"
  version_name           = "v2.0"
  secret_type            = "ECS"
  description            = "tf-test ecs"
  secret_value           = "{\"UserName\":\"root\",\"Password\":\"********\"}"
  extended_config        = "{\"InstanceId\":\"i-yeehzz2tc0ygp2******\",\"SecretSubType\":\"Password\",\"CustomData\":{\"desc\":\"test\"}}"
  project_name           = "default"
  encryption_key         = "trn:kms:cn-beijing:21000******:keyrings/Tf-test/keys/Test-key1"
  automatic_rotation     = false
  force_delete           = false
  pending_window_in_days = 7
}
```
## Argument Reference
The following arguments are supported:
* `secret_name` - (Required, ForceNew) The name of the secret.
* `secret_type` - (Required, ForceNew) The type of the secret. Valid values: Generic, IAM, RDS, Redis, ECS.
* `secret_value` - (Required) The value of the secret. Only Generic type secret support modifying secret_value.
* `automatic_rotation` - (Optional) The rotation state of the secret. Only valid for IAM, RDS, Redis, ECS secrets.
* `description` - (Optional) The description of the secret.
* `encryption_key` - (Optional, ForceNew) The TRN of the KMS key used to encrypt the secret value.
* `extended_config` - (Optional, ForceNew) The extended configurations of the secret.
* `force_delete` - (Optional) Whether to delete the secret immediately. If false, the secret enters pending deletion state. Only effective when destroying resources.
* `pending_window_in_days` - (Optional) The waiting period before deletion when force_delete is false. Valid values: 7~30. Only effective when destroying resources.
* `project_name` - (Optional) The project name of the secret.
* `rotation_interval` - (Optional) The interval at which automatic rotation is performed. This parameter must be specified when automatic_rotation is true.
* `version_name` - (Optional) The version alias of the secret. Only Generic type secret support modifying version_name.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `creation_date` - The date when the secret was created.
* `last_rotation_time` - The last time the secret was rotated.
* `managed` - Indicates whether the secret is hosted.
* `owning_service` - The cloud service that owns the secret.
* `rotation_interval_second` - Rotation interval second.
* `rotation_state` - The rotation state of the secret.
* `schedule_delete_time` - The time when the secret will be deleted.
* `schedule_rotation_time` - The next time the secret will be rotated.
* `state` - The state of secret.
* `trn` - The information about the tenant resource name (TRN).
* `uid` - The tenant ID of the secret.
* `update_date` - The date when the secret was updated.
* `uuid` - The ID of secret.


## Import
KmsSecret can be imported using the id, e.g.
```
$ terraform import volcengine_kms_secret.default resource_id
```

