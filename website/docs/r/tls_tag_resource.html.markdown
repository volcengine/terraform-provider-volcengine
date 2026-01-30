---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_tag_resource"
sidebar_current: "docs-volcengine-resource-tls_tag_resource"
description: |-
  Provides a resource to manage tls tag resource
---
# volcengine_tls_tag_resource
Provides a resource to manage tls tag resource
## Example Usage
```hcl
# Example: Add tags to a TLS topic
resource "volcengine_tls_tag_resource" "foo" {
  resource_id   = "bdb87e4d-7dad-4b96-ac43-e1b09e9dc8ac"
  resource_type = "project"
  tags {
    key   = "environment"
    value = "production"
  }
  tags {
    key   = "key1"
    value = "value2"
  }
}

output "tls_tag_id" {
  value = volcengine_tls_tag_resource.foo.id
}

output "tls_tag_resource_id" {
  value = volcengine_tls_tag_resource.foo.resource_id
}

output "tls_tag_resource_type" {
  value = volcengine_tls_tag_resource.foo.resource_type
}

output "tls_tag_tags" {
  value = volcengine_tls_tag_resource.foo.tags
}
```
## Argument Reference
The following arguments are supported:
* `resource_id` - (Required, ForceNew) The ID of the resource.
* `resource_type` - (Required, ForceNew) The type of the resource. Valid values: project, topic, shipper, host_group, host, consumer_group, rule, alarm, alarm_notify_group, etl_task, import_task, schedule_sql_task, download_task, trace_instance.
* `tags` - (Required, ForceNew) Tags. The tag key must be unique within a resource, and the same tag key is not allowed to be repeated. The tag key must be 1 to 128 characters long, and can contain letters, digits, spaces, and the following special characters: _.:/=+-@. The tag value can be empty and must be 0 to 256 characters long, and can contain letters, digits, spaces, and the following special characters: _.:/=+-@.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
tls tag can be imported using the resource_id:resource_type, e.g.
```
$ terraform import volcengine_tls_tag.default resource-123456:project
```

