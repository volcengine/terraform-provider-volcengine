---
subcategory: "VEFAAS"
layout: "volcengine"
page_title: "Volcengine: volcengine_vefaas_functions"
sidebar_current: "docs-volcengine-datasource-vefaas_functions"
description: |-
  Use this data source to query detailed information of vefaas functions
---
# volcengine_vefaas_functions
Use this data source to query detailed information of vefaas functions
## Example Usage
```hcl
data "volcengine_vefaas_functions" "foo" {

}
```
## Argument Reference
The following arguments are supported:
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `items` - The collection of query.
    * `code_size_limit` - Maximum code package size.
    * `code_size` - The size of code package.
    * `command` - The custom startup command for the instance.
    * `cpu_strategy` - Function CPU charging policy.
    * `creation_time` - Creation time.
    * `description` - The description of Function.
    * `envs` - Function environment variable.
        * `key` - The Key of the environment variable.
        * `value` - The Value of the environment variable.
    * `exclusive_mode` - Exclusive mode switch.
    * `id` - The ID of Function.
    * `initializer_sec` - Function to initialize timeout configuration.
    * `instance_type` - The instance type of the function instance.
    * `last_update_time` - Update time.
    * `max_concurrency` - Maximum concurrency of a single instance.
    * `memory_mb` - Maximum memory for a single instance.
    * `name` - The name of Function.
    * `nas_storage` - The configuration of file storage NAS mount.
        * `enable_nas` - Whether to enable NAS storage mounting.
        * `nas_configs` - The configuration of NAS.
            * `file_system_id` - The ID of NAS file system.
            * `gid` - User groups in the file system. Customization is not supported yet. If this parameter is provided, the parameter value is 1000 (consistent with the function run user gid).
            * `local_mount_path` - The directory of Function local mount.
            * `mount_point_id` - The ID of NAS mount point.
            * `remote_path` - Remote directory of the file system.
            * `uid` - Users in the file system do not support customization yet. If this parameter is provided, its value can only be 1000 (consistent with the function run user uid).
    * `owner` - The owner of Function.
    * `port` - Custom listening port for the instance.
    * `request_timeout` - Request timeout (in seconds).
    * `runtime` - The runtime of Function.
    * `source_location` - The source address of the code/image.
    * `source_type` - Code Source type, supports tos, zip, image (whitelist accounts support native/v1 custom images).
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `tls_config` - Function log configuration.
        * `enable_log` - TLS log function switch.
        * `tls_project_id` - The project ID of TLS log topic.
        * `tls_topic_id` - The topic ID of TLS log topic.
    * `tos_mount_config` - The configuration of Object Storage TOS mount.
        * `credentials` - After enabling TOS, you need to provide an AKSK with access rights to the TOS domain name.
            * `access_key_id` - The AccessKey ID (AK) of the Volcano Engine account.
            * `secret_access_key` - The Secret Access Key (SK) of the Volcano Engine account.
        * `enable_tos` - Whether to enable TOS storage mounting.
        * `mount_points` - After enabling TOS, you need to provide a TOS storage configuration list, with a maximum of 5 items.
            * `bucket_name` - TOS bucket.
            * `bucket_path` - The mounted TOS Bucket path.
            * `endpoint` - TOS Access domain name.
            * `local_mount_path` - Function local mount directory.
            * `read_only` - Function local directory access permissions. After mounting the TOS Bucket, whether the function local mount directory has read-only permissions.
    * `triggers_count` - The number of triggers for this Function.
    * `vpc_config` - The configuration of VPC.
        * `enable_shared_internet_access` - Function access to the public network switch.
        * `enable_vpc` - Whether the function enables private network access.
        * `security_group_ids` - The ID of security group.
        * `subnet_ids` - The ID of subnet.
        * `vpc_id` - The ID of VPC.
* `total_count` - The total count of query.


