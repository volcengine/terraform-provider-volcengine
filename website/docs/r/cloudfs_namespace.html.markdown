---
subcategory: "CLOUDFS"
layout: "volcengine"
page_title: "Volcengine: volcengine_cloudfs_namespace"
sidebar_current: "docs-volcengine-resource-cloudfs_namespace"
description: |-
  Provides a resource to manage cloudfs namespace
---
# volcengine_cloudfs_namespace
Provides a resource to manage cloudfs namespace
## Example Usage
```hcl
resource "volcengine_cloudfs_namespace" "foo" {
  fs_name    = "tf-test-fs"
  tos_bucket = "tf-test"
  read_only  = true
}
```
## Argument Reference
The following arguments are supported:
* `fs_name` - (Required, ForceNew) The name of file system.
* `tos_bucket` - (Required, ForceNew) The name of tos bucket.
* `read_only` - (Optional, ForceNew) Whether the namespace is read-only.
* `tos_account_id` - (Optional, ForceNew) When a data lake scenario instance chooses to associate a bucket under another account, you need to set the ID of the account. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `tos_ak` - (Optional, ForceNew) The tos ak. When the data lake scenario chooses to associate buckets under other accounts, need to set the Access Key ID of the account. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `tos_prefix` - (Optional, ForceNew) The tos prefix. Must not start with /, but must end with /, such as prefix/. When it is empty, it means the root path.
* `tos_sk` - (Optional, ForceNew) The tos sk. When the data lake scenario chooses to associate buckets under other accounts, need to set the Secret Access Key of the account. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `created_time` - The creation time of the namespace.
* `is_my_bucket` - Whether the tos bucket is your own bucket.
* `ns_id` - The id of namespace.
* `service_managed` - Whether the namespace is the official service for volcengine.
* `status` - The status of the namespace.


## Import
CloudfsNamespace can be imported using the FsName:NsId, e.g.
```
$ terraform import volcengine_cloudfs_namespace.default tfname:1801439850948****
```

