---
subcategory: "TOS(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_tos_bucket_realtime_log"
sidebar_current: "docs-volcengine-resource-tos_bucket_realtime_log"
description: |-
  Provides a resource to manage tos bucket realtime log
---
# volcengine_tos_bucket_realtime_log
Provides a resource to manage tos bucket realtime log
## Example Usage
```hcl
// When deleting this resource, the tls related resources such as project and topic will not be automatically deleted
resource "volcengine_tos_bucket_realtime_log" "foo" {
  bucket_name = "terraform-demo"
  role        = "TOSLogArchiveTLSRole"
  access_log_configuration {
    ttl = 6
  }
}
```
## Argument Reference
The following arguments are supported:
* `access_log_configuration` - (Required) The export schedule of the bucket inventory.
* `bucket_name` - (Required, ForceNew) The name of the bucket.
* `role` - (Required, ForceNew) The role name used to grant TOS access to create resources such as projects and topics, and write logs to the TLS logging service. You can use the default TOS role `TOSLogArchiveTLSRole`.

The `access_log_configuration` object supports the following:

* `ttl` - (Optional) The TLS log retention duration. Unit in days. Valid values range is 1~3650. default is 7.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
TosBucketRealtimeLog can be imported using the bucket_name, e.g.
```
$ terraform import volcengine_tos_bucket_realtime_log.default resource_id
```

