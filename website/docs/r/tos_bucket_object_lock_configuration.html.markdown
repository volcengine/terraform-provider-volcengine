---
subcategory: "TOS(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_tos_bucket_object_lock_configuration"
sidebar_current: "docs-volcengine-resource-tos_bucket_object_lock_configuration"
description: |-
  Provides a resource to manage tos bucket object lock configuration
---
# volcengine_tos_bucket_object_lock_configuration
Provides a resource to manage tos bucket object lock configuration
## Example Usage
```hcl
resource "volcengine_tos_bucket_object_lock_configuration" "foo" {
  bucket_name = "tflyb7"

  rule {
    default_retention {
      mode = "COMPLIANCE"
      days = 31
    }
  }
}
```
## Argument Reference
The following arguments are supported:
* `bucket_name` - (Required, ForceNew) The name of the TOS bucket.
* `rule` - (Optional) The object lock rule configuration.

The `default_retention` object supports the following:

* `mode` - (Required) The default retention mode. Valid values: COMPLIANCE, GOVERNANCE.
* `days` - (Optional) The number of days for the default retention period.
* `years` - (Optional) The number of years for the default retention period.

The `rule` object supports the following:

* `default_retention` - (Required) The default retention configuration.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
TosBucketObjectLockConfiguration can be imported using the bucketName, e.g.
```
$ terraform import volcengine_tos_bucket_object_lock_configuration.default bucket_name
```

