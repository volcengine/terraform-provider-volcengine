---
subcategory: "TOS(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_tos_bucket_logging"
sidebar_current: "docs-volcengine-resource-tos_bucket_logging"
description: |-
  Provides a resource to manage tos bucket logging
---
# volcengine_tos_bucket_logging
Provides a resource to manage tos bucket logging
## Example Usage
```hcl
resource "volcengine_tos_bucket_logging" "foo" {
  bucket_name = "tflyb7"
  logging_enabled {
    target_bucket = "tflyb78"
    target_prefix = "logs1/"
    role          = "ServiceRoleforTOSLogging"
  }
}
```
## Argument Reference
The following arguments are supported:
* `bucket_name` - (Required, ForceNew) The name of the TOS bucket.
* `logging_enabled` - (Optional) The name of the TOS bucket.

The `logging_enabled` object supports the following:

* `role` - (Optional) The role that is assumed by TOS to write log objects to the target bucket.
* `target_bucket` - (Optional) The name of the target bucket where the access logs are stored.
* `target_prefix` - (Optional) The prefix for the log object keys.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
TosBucketLogging can be imported using the bucketName, e.g.
```
$ terraform import volcengine_tos_bucket_logging.default bucket_name
```

