---
subcategory: "VEFAAS"
layout: "volcengine"
page_title: "Volcengine: volcengine_vefaas_function"
sidebar_current: "docs-volcengine-resource-vefaas_function"
description: |-
  Provides a resource to manage vefaas function
---
# volcengine_vefaas_function
Provides a resource to manage vefaas function
## Example Usage
```hcl
resource "volcengine_vefaas_function" "foo" {
  name            = "project-1"
  runtime         = "golang/v1"
  description     = "123131231"
  exclusive_mode  = false
  request_timeout = 30
}
```
## Argument Reference
The following arguments are supported:
* `name` - (Required, ForceNew) The name of Function.
* `runtime` - (Required, ForceNew) The runtime of Function.
* `command` - (Optional, ForceNew) The custom startup command for the instance.
* `cpu_strategy` - (Optional, ForceNew) Function CPU charging policy.
* `description` - (Optional) The description of Function.
* `envs` - (Optional) Function environment variable.
* `exclusive_mode` - (Optional) Exclusive mode switch.
* `initializer_sec` - (Optional) Function to initialize timeout configuration.
* `max_concurrency` - (Optional) Maximum concurrency of a single instance.
* `memory_mb` - (Optional) Maximum memory for a single instance.
* `nas_storage` - (Optional) The configuration of file storage NAS mount.
* `request_timeout` - (Optional) Request timeout (in seconds).
* `source_access_config` - (Optional) Access configuration for the image repository.
* `source_type` - (Optional) Code Source type, supports tos, zip, image (whitelist accounts support native/v1 custom images).
* `source` - (Optional) Code source.
* `tls_config` - (Optional) Function log configuration.
* `tos_mount_config` - (Optional) The configuration of Object Storage TOS mount.
* `vpc_config` - (Optional) The configuration of VPC.

The `credentials` object supports the following:

* `access_key_id` - (Required) The AccessKey ID (AK) of the Volcano Engine account.
* `secret_access_key` - (Required) The Secret Access Key (SK) of the Volcano Engine account.

The `envs` object supports the following:

* `key` - (Required) The Key of the environment variable.
* `value` - (Required) The Value of the environment variable.

The `mount_points` object supports the following:

* `bucket_name` - (Required) TOS bucket.
* `bucket_path` - (Required) The mounted TOS Bucket path.
* `endpoint` - (Required) TOS Access domain name.
* `local_mount_path` - (Required) Function local mount directory.
* `read_only` - (Optional) Function local directory access permissions. After mounting the TOS Bucket, whether the function local mount directory has read-only permissions.

The `nas_configs` object supports the following:

* `file_system_id` - (Required) The ID of NAS file system.
* `local_mount_path` - (Required) The directory of Function local mount.
* `mount_point_id` - (Required) The ID of NAS mount point.
* `remote_path` - (Required) Remote directory of the file system.
* `gid` - (Optional) User groups in the file system. Customization is not supported yet. If this parameter is provided, the parameter value is 1000 (consistent with the function run user gid).
* `uid` - (Optional) Users in the file system do not support customization yet. If this parameter is provided, its value can only be 1000 (consistent with the function run user uid).

The `nas_storage` object supports the following:

* `enable_nas` - (Required) Whether to enable NAS storage mounting.
* `nas_configs` - (Optional) The configuration of NAS.

The `source_access_config` object supports the following:

* `password` - (Required) The image repository password.
* `username` - (Required) Mirror repository username.

The `tls_config` object supports the following:

* `enable_log` - (Required) TLS log function switch.
* `tls_project_id` - (Optional) The project ID of TLS log topic.
* `tls_topic_id` - (Optional) The topic ID of TLS log topic.

The `tos_mount_config` object supports the following:

* `enable_tos` - (Required) Whether to enable TOS storage mounting.
* `credentials` - (Optional) After enabling TOS, you need to provide an AKSK with access rights to the TOS domain name.
* `mount_points` - (Optional) After enabling TOS, you need to provide a TOS storage configuration list, with a maximum of 5 items.

The `vpc_config` object supports the following:

* `enable_vpc` - (Required) Whether the function enables private network access.
* `enable_shared_internet_access` - (Optional) Function access to the public network switch.
* `security_group_ids` - (Optional) The ID of security group.
* `subnet_ids` - (Optional) The ID of subnet.
* `vpc_id` - (Optional) The ID of VPC.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `code_size_limit` - Maximum code package size.
* `code_size` - The size of code package.
* `creation_time` - The creation time of the function.
* `last_update_time` - The last update time of the function.
* `owner` - The owner of Function.
* `port` - Custom listening port for the instance.
* `source_location` - Maximum code package size.
* `triggers_count` - The number of triggers for this Function.


## Import
VefaasFunction can be imported using the id, e.g.
```
$ terraform import volcengine_vefaas_function.default resource_id
```

