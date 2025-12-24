---
subcategory: "TOS(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_tos_bucket_rename"
sidebar_current: "docs-volcengine-resource-tos_bucket_rename"
description: |-
  Provides a resource to manage tos bucket rename
---
# volcengine_tos_bucket_rename
Provides a resource to manage tos bucket rename
## Example Usage
```hcl
resource "volcengine_tos_bucket_rename" "default" {
  bucket_name = "tflyb78"
}
```
## Argument Reference
The following arguments are supported:
* `bucket_name` - (Required, ForceNew) The name of the TOS bucket to configure rename functionality for.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
TosBucketRename can be imported using the bucketName, e.g.
```
$ terraform import volcengine_tos_bucket_rename.default bucket_name
```

