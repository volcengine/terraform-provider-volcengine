---
subcategory: "CLOUDFS"
layout: "volcengine"
page_title: "Volcengine: volcengine_cloudfs_namespaces"
sidebar_current: "docs-volcengine-datasource-cloudfs_namespaces"
description: |-
  Use this data source to query detailed information of cloudfs namespaces
---
# volcengine_cloudfs_namespaces
Use this data source to query detailed information of cloudfs namespaces
## Example Usage
```hcl
data "volcengine_cloudfs_namespaces" "default" {
  fs_name = "tf-test-fs"
  ns_id   = "1801439850948****"
}
```
## Argument Reference
The following arguments are supported:
* `fs_name` - (Required) The name of file system.
* `ns_id` - (Optional) The id of namespace.
* `output_file` - (Optional) File name where to save data source results.
* `tos_bucket` - (Optional) The name of tos bucket.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `namespaces` - The collection of query.
    * `created_time` - The creation time of the namespace.
    * `id` - The ID of the namespace.
    * `is_my_bucket` - Whether the tos bucket is your own bucket.
    * `read_only` - Whether the namespace is read-only.
    * `service_managed` - Whether the namespace is the official service for volcengine.
    * `status` - The status of the namespace.
    * `tos_bucket` - The name of the tos bucket.
    * `tos_prefix` - The tos prefix.
* `total_count` - The total count of query.


