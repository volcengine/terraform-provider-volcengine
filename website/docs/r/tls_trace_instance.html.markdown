---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_trace_instance"
sidebar_current: "docs-volcengine-resource-tls_trace_instance"
description: |-
  Provides a resource to manage tls trace instance
---
# volcengine_tls_trace_instance
Provides a resource to manage tls trace instance
## Example Usage
```hcl
# Example: Create a TLS trace instance
resource "volcengine_tls_trace_instance" "foo" {
  project_id          = "bdb87e4d-7dad-4b96-ac43-e1b09e9dc8ac"
  trace_instance_name = "tf-trace-instance-df"
  description         = "This is an example trace instance"
  backend_config {
    ttl                  = 60
    enable_hot_ttl       = true
    hot_ttl              = 30
    cold_ttl             = 30
    archive_ttl          = 0
    auto_split           = true
    max_split_partitions = 10
  }
}

output "tls_trace_instance_id" {
  value = volcengine_tls_trace_instance.foo.id
}

output "tls_trace_instance_name" {
  value = volcengine_tls_trace_instance.foo.trace_instance_name
}

output "tls_trace_instance_description" {
  value = volcengine_tls_trace_instance.foo.description
}
```
## Argument Reference
The following arguments are supported:
* `project_id` - (Required, ForceNew) The ID of the project.
* `trace_instance_name` - (Required) The name of the trace instance.
* `backend_config` - (Optional) The backend config of the trace instance.
* `description` - (Optional) The description of the trace instance.

The `backend_config` object supports the following:

* `archive_ttl` - (Optional) Archive storage duration in days.
* `auto_split` - (Optional) Whether to enable auto split.
* `cold_ttl` - (Optional) Infrequent storage duration in days.
* `enable_hot_ttl` - (Optional) Whether to enable tiered storage.
* `hot_ttl` - (Optional) Standard storage duration in days.
* `max_split_partitions` - (Optional) Max split partitions.
* `ttl` - (Optional) Total log retention time in days.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
tls trace instance can be imported using the id, e.g.
```
$ terraform import volcengine_tls_trace_instance.default instance-1234567890
```

