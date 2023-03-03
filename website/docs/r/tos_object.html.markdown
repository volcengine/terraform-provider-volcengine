---
subcategory: "TOS(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_tos_object"
sidebar_current: "docs-volcengine-resource-tos_object"
description: |-
  Provides a resource to manage tos object
---
# volcengine_tos_object
Provides a resource to manage tos object
## Example Usage
```hcl
resource "volcengine_tos_object" "default" {
  bucket_name = "test-xym-1"
  object_name = "demo_xym"
  file_path   = "/Users/bytedance/Work/Go/build/test.txt"
  #  storage_class ="IA"
  public_acl = "private"
  encryption = "AES256"
  #content_type = "text/plain"
  account_acl {
    account_id = "1"
    permission = "READ"
  }
  account_acl {
    account_id = "2001"
    permission = "WRITE_ACP"
  }
  #  lifecycle {
  #    ignore_changes = ["file_path"]
  #  }
}
```
## Argument Reference
The following arguments are supported:
* `bucket_name` - (Required, ForceNew) The name of the bucket.
* `file_path` - (Required) The file path for upload.
* `object_name` - (Required, ForceNew) The name of the object.
* `account_acl` - (Optional) The user set of grant full control.
* `content_md5` - (Optional) The file md5 sum (32-bit hexadecimal string) for upload.
* `content_type` - (Optional, ForceNew) The content type of the object.
* `encryption` - (Optional, ForceNew) The encryption of the object.Valid value is AES256.
* `public_acl` - (Optional) The public acl control of object.Valid value is private|public-read|public-read-write|authenticated-read|bucket-owner-read.
* `storage_class` - (Optional, ForceNew) The storage type of the object.Valid value is STANDARD|IA.

The `account_acl` object supports the following:

* `account_id` - (Required) The accountId to control.
* `permission` - (Required) The permission to control.Valid value is FULL_CONTROL|READ|READ_ACP|WRITE|WRITE_ACP.
* `acl_type` - (Optional) The acl type to control.Valid value is CanonicalUser.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `enable_version` - The flag of enable tos version.
* `version_ids` - The version ids of the object if exist.


## Import
TOS Object can be imported using the id, e.g.
```
$ terraform import volcengine_tos_object.default bucketName:objectName
```

