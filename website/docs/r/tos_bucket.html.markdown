---
subcategory: "TOS(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_tos_bucket"
sidebar_current: "docs-volcengine-resource-tos_bucket"
description: |-
  Provides a resource to manage tos bucket
---
# volcengine_tos_bucket
Provides a resource to manage tos bucket
## Example Usage
```hcl
resource "volcengine_tos_bucket" "default" {
  bucket_name = "test-xym-1"
  #  storage_class ="IA"
  public_acl     = "private"
  enable_version = true
  account_acl {
    account_id = "1"
    permission = "READ"
  }
  account_acl {
    account_id = "2001"
    permission = "WRITE_ACP"
  }
}
```
## Argument Reference
The following arguments are supported:
* `bucket_name` - (Required, ForceNew) The name of the bucket.
* `account_acl` - (Optional) The user set of grant full control.
* `enable_version` - (Optional) The flag of enable tos version.
* `public_acl` - (Optional) The public acl control of object.Valid value is private|public-read|public-read-write|authenticated-read|bucket-owner-read.
* `storage_class` - (Optional, ForceNew) The storage type of the object.Valid value is STANDARD|IA|ARCHIVE_FR.Default is STANDARD.

The `account_acl` object supports the following:

* `account_id` - (Required) The accountId to control.
* `permission` - (Required) The permission to control.Valid value is FULL_CONTROL|READ|READ_ACP|WRITE|WRITE_ACP.
* `acl_type` - (Optional) The acl type to control.Valid value is CanonicalUser.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Tos Bucket can be imported using the id, e.g.
```
$ terraform import volcengine_tos_bucket.default region:bucketName
```

