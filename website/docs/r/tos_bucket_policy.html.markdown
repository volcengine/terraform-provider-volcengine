---
subcategory: "TOS(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_tos_bucket_policy"
sidebar_current: "docs-volcengine-resource-tos_bucket_policy"
description: |-
  Provides a resource to manage tos bucket policy
---
# volcengine_tos_bucket_policy
Provides a resource to manage tos bucket policy
## Example Usage
```hcl
resource "volcengine_tos_bucket_policy" "default" {
  bucket_name = "bucket-20230418"
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
          "trn:tos:::bucket-20230418"
        ]
      }
    ]
  })
}
```
## Argument Reference
The following arguments are supported:
* `bucket_name` - (Required, ForceNew) The name of the bucket.
* `policy` - (Required) The policy document. This is a JSON formatted string. For more information about building Volcengine IAM policy documents with Terraform, see the  [Volcengine IAM Policy Document Guide](https://www.volcengine.com/docs/6349/102127).

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Tos Bucket can be imported using the id, e.g.
```
$ terraform import volcengine_tos_bucket_policy.default bucketName:policy
```

