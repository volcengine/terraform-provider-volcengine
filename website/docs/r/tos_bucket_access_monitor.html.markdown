---
subcategory: "TOS(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_tos_bucket_access_monitor"
sidebar_current: "docs-volcengine-resource-tos_bucket_access_monitor"
description: |-
  Provides a resource to manage tos bucket access monitor
---
# volcengine_tos_bucket_access_monitor
Provides a resource to manage tos bucket access monitor
## Example Usage
```hcl
# Example usage of volcengine_tos_bucket_access_monitor resource

resource "volcengine_tos_bucket_access_monitor" "foo" {
  bucket_name = "tflyb1"
}
```
## Argument Reference
The following arguments are supported:
* `bucket_name` - (Required, ForceNew) The name of the TOS bucket.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
TosBucketAccessMonitor can be imported using the bucketName, e.g.
```
$ terraform import volcengine_tos_bucket_access_monitor.default bucket_name
```

