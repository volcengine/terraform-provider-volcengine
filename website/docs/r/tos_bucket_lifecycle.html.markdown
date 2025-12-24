---
subcategory: "TOS(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_tos_bucket_lifecycle"
sidebar_current: "docs-volcengine-resource-tos_bucket_lifecycle"
description: |-
  Provides a resource to manage tos bucket lifecycle
---
# volcengine_tos_bucket_lifecycle
Provides a resource to manage tos bucket lifecycle
## Example Usage
```hcl
resource "volcengine_tos_bucket_lifecycle" "foo" {
  bucket_name = "tflybtest5"
  rules {
    id     = "rule1"
    status = "Enabled"
    prefix = "documents/"

    expiration {
      days = 122
    }

    tags {
      key   = "example1"
      value = "example-value1"
    }

    tags {
      key   = "example2"
      value = "example-value2"
    }

    filter {
      object_size_greater_than   = 1024
      object_size_less_than      = 10485760
      greater_than_include_equal = "Enabled"
      less_than_include_equal    = "Disabled"
    }
    non_current_version_expiration {
      non_current_days = 90
    }
    non_current_version_transitions {
      non_current_days = 30
      storage_class    = "IA"
    }
    non_current_version_transitions {
      non_current_days = 31
      storage_class    = "ARCHIVE"
    }


    transitions {
      days          = 7
      storage_class = "IA"
    }

    transitions {
      days          = 30
      storage_class = "ARCHIVE"
    }
  }

  rules {
    id     = "rule2"
    status = "Enabled"
    prefix = "logs/"

    expiration {
      days = 90
    }

    non_current_version_expiration {
      non_current_days = 30
    }

    non_current_version_transitions {
      non_current_days = 7
      storage_class    = "IA"
    }
  }

  rules {
    id     = "rule3"
    status = "Disabled"
    prefix = "temp/"

    expiration {
      date = "2025-12-31T00:00:00.000Z"
    }

    abort_incomplete_multipart_upload {
      days_after_initiation = 1
    }
  }
}
```
## Argument Reference
The following arguments are supported:
* `bucket_name` - (Required, ForceNew) The name of the TOS bucket.
* `rules` - (Required) The lifecycle rules of the bucket.

The `abort_incomplete_multipart_upload` object supports the following:

* `days_after_initiation` - (Required) The number of days after initiation when the incomplete multipart upload should be aborted.

The `expiration` object supports the following:

* `date` - (Optional) The date when the rule takes effect. Format: 2023-01-01T00:00:00.000Z.
* `days` - (Optional) The number of days after object creation when the rule takes effect.

The `filter` object supports the following:

* `greater_than_include_equal` - (Optional) Whether to enable equal conditions. The value can only be "Enabled" or "Disabled". If not configured, it will default to "Disabled".
* `less_than_include_equal` - (Optional) Whether to enable equal conditions. The value can only be "Enabled" or "Disabled". If not configured, it will default to "Disabled".
* `object_size_greater_than` - (Optional) The minimum object size in bytes for the rule to apply.
* `object_size_less_than` - (Optional) The maximum object size in bytes for the rule to apply.

The `non_current_version_expiration` object supports the following:

* `non_current_days` - (Required) The number of days after object creation when the non-current version expiration takes effect.

The `non_current_version_transitions` object supports the following:

* `non_current_days` - (Required) The number of days after object creation when the non-current version transition takes effect.
* `storage_class` - (Required) The storage class to transition to. Valid values: IA, ARCHIVE, COLD_ARCHIVE.

The `rules` object supports the following:

* `status` - (Required) The status of the lifecycle rule. Valid values: Enabled, Disabled.
* `abort_incomplete_multipart_upload` - (Optional) The abort incomplete multipart upload configuration of the lifecycle rule.
* `expiration` - (Optional) The expiration configuration of the lifecycle rule.
* `filter` - (Optional) The filter configuration of the lifecycle rule.
* `id` - (Optional) The ID of the lifecycle rule.
* `non_current_version_expiration` - (Optional) The non-current version expiration configuration of the lifecycle rule.
* `non_current_version_transitions` - (Optional) The non-current version transition configuration of the lifecycle rule.
* `prefix` - (Optional) The prefix of the lifecycle rule.
* `tags` - (Optional) The tag filters.
* `transitions` - (Optional) The transition configuration of the lifecycle rule.

The `tags` object supports the following:

* `key` - (Required) The key of the tag.
* `value` - (Required) The value of the tag.

The `transitions` object supports the following:

* `date` - (Optional) The date when the transition takes effect. Format: 2023-01-01T00:00:00.000Z.
* `days` - (Optional) The number of days after object creation when the transition takes effect.
* `storage_class` - (Optional) The storage class to transition to. Valid values: IA, ARCHIVE, COLD_ARCHIVE.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
TosBucketLifecycle can be imported using the bucketName, e.g.
```
$ terraform import volcengine_tos_bucket_lifecycle.default bucket_name
```

