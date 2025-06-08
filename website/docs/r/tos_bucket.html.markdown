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
# create tos bucket
resource "volcengine_tos_bucket" "foo" {
  bucket_name = "tf-acc-test-bucket"
  #  storage_class        = "IA"
  public_acl           = "private"
  az_redundancy        = "multi-az"
  enable_version       = true
  bucket_acl_delivered = true
  account_acl {
    account_id = "1"
    permission = "READ"
  }
  account_acl {
    account_id = "2001"
    permission = "WRITE_ACP"
  }
  project_name = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

# create tos bucket policy
resource "volcengine_tos_bucket_policy" "foo" {
  bucket_name = volcengine_tos_bucket.foo.id
  policy = jsonencode({
    Statement = [
      {
        Sid    = "test"
        Effect = "Allow"
        Principal = [
          "AccountId/subUserName"
        ]
        Action = [
          "tos:List*"
        ]
        Resource = [
          "trn:tos:::${volcengine_tos_bucket.foo.id}"
        ]
      }
    ]
  })
}
```
## Argument Reference
The following arguments are supported:
* `bucket_name` - (Required, ForceNew) The name of the bucket.
* `account_acl` - (Optional) The user set of grant full control.
* `az_redundancy` - (Optional, ForceNew) The AZ redundancy of the Tos Bucket. Default is `single-az`. Valid values: `single-az`, `multi-az`.
* `bucket_acl_delivered` - (Optional) Whether to enable the default inheritance bucket ACL function for objects. Default is false.
* `enable_version` - (Optional) The flag of enable tos version.
* `project_name` - (Optional) The ProjectName of the Tos Bucket. Default is `default`.
* `public_acl` - (Optional) The public acl control of object.Valid value is private|public-read|public-read-write|authenticated-read|bucket-owner-read.
* `storage_class` - (Optional, ForceNew) The storage type of the object.Valid value is STANDARD|IA|ARCHIVE_FR.Default is STANDARD.
* `tags` - (Optional) Tos Bucket Tags.

The `account_acl` object supports the following:

* `account_id` - (Required) The accountId to control.
* `permission` - (Required) The permission to control.Valid value is FULL_CONTROL|READ|READ_ACP|WRITE|WRITE_ACP.
* `acl_type` - (Optional) The acl type to control.Valid value is CanonicalUser.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `creation_date` - The create date of the TOS bucket.
* `extranet_endpoint` - The extranet endpoint of the TOS bucket.
* `intranet_endpoint` - The intranet endpoint the TOS bucket.
* `location` - The location of the TOS bucket.


## Import
Tos Bucket can be imported using the id, e.g.
```
$ terraform import volcengine_tos_bucket.default bucketName
```

