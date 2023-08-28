---
subcategory: "CLOUDFS"
layout: "volcengine"
page_title: "Volcengine: volcengine_cloudfs_file_system"
sidebar_current: "docs-volcengine-resource-cloudfs_file_system"
description: |-
  Provides a resource to manage cloudfs file system
---
# volcengine_cloudfs_file_system
Provides a resource to manage cloudfs file system
## Example Usage
```hcl
resource "volcengine_cloudfs_file_system" "foo" {
  fs_name    = "tffile"
  zone_id    = "cn-beijing-b"
  cache_plan = "T2"
  mode       = "HDFS_MODE"
  read_only  = true

  subnet_id          = "subnet-13fca1crr5d6o3n6nu46cyb5m"
  security_group_id  = "sg-rrv1klfg5s00v0x578mx14m"
  cache_capacity_tib = 10
  vpc_route_enabled  = true

  tos_bucket = "tfacc"
  tos_prefix = "pre/"
}


resource "volcengine_cloudfs_file_system" "foo1" {
  fs_name    = "tffileu"
  zone_id    = "cn-beijing-b"
  cache_plan = "T2"
  mode       = "ACC_MODE"
  read_only  = true

  subnet_id          = "subnet-13fca1crr5d6o3n6nu46cyb5m"
  security_group_id  = "sg-rrv1klfg5s00v0x578mx14m"
  cache_capacity_tib = 15
  vpc_route_enabled  = false

  tos_bucket = "tfacc"
}
```
## Argument Reference
The following arguments are supported:
* `cache_plan` - (Required) The cache plan. The value can be `DISABLED` or `T2` or `T4`. When expanding the cache size, the cache plan should remain the same. For data lakes, cache must be enabled.
* `fs_name` - (Required, ForceNew) The name of file system.
* `mode` - (Required, ForceNew) The mode of file system. The value can be `HDFS_MODE` or `ACC_MODE`.
* `zone_id` - (Required, ForceNew) The id of zone.
* `cache_capacity_tib` - (Optional) The capacity of cache. This parameter is required when cache acceleration is enabled.
* `read_only` - (Optional, ForceNew) Whether the Namespace created automatically when mounting the TOS Bucket is read-only. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `security_group_id` - (Optional) The id of security group. This parameter is required when cache acceleration is enabled.
* `subnet_id` - (Optional) The id of subnet. This parameter is required when cache acceleration is enabled.
* `tos_account_id` - (Optional, ForceNew) When a data lake scenario instance chooses to associate a bucket under another account, you need to set the ID of the account. When importing resources, this attribute will not be imported.
* `tos_ak` - (Optional, ForceNew) The tos ak. When the data lake scenario chooses to associate buckets under other accounts, need to set the Access Key ID of the account. When importing resources, this attribute will not be imported.
* `tos_bucket` - (Optional, ForceNew) The tos bucket. When importing ACC_MODE resources, this attribute will not be imported.
* `tos_prefix` - (Optional, ForceNew) The tos prefix. Must not start with /, but must end with /, such as prefix/. When it is empty, it means the root path. When importing ACC_MODE resources, this attribute will not be imported.
* `tos_sk` - (Optional, ForceNew) The tos sk. When the data lake scenario chooses to associate buckets under other accounts, need to set the Secret Access Key of the account. When importing resources, this attribute will not be imported.
* `vpc_route_enabled` - (Optional) Whether enable all vpc route.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `access_id` - The default vpc access id.
* `created_time` - The creation time.
* `mount_point` - The point mount.
* `status` - Status of file system.
* `vpc_id` - The id of vpc.


## Import
CloudFileSystem can be imported using the FsName, e.g.
```
$ terraform import volcengine_cloudfs_file_system.default tfname
```

