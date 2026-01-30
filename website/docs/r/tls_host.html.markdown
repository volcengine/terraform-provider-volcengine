---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_host"
sidebar_current: "docs-volcengine-resource-tls_host"
description: |-
  Provides a resource to manage tls host
---
# volcengine_tls_host
Provides a resource to manage tls host
## Example Usage
```hcl
resource "volcengine_tls_host_group" "foo" {
  host_group_name   = "tfgroup-ip-tf"
  host_group_type   = "IP"
  host_ip_list      = ["192.168.0.1", "192.168.0.2", "192.168.0.3"]
  auto_update       = true
  update_start_time = "00:00"
  update_end_time   = "02:00"
  service_logging   = false
  iam_project_name  = "default"
}

# 删除指定 IP
resource "volcengine_tls_host" "delete_foo" {
  host_group_id = volcengine_tls_host_group.foo.id
  ip            = "192.168.0.1"
}

# 删除异常机器
resource "volcengine_tls_host" "delete_abnormal" {
  host_group_id = volcengine_tls_host_group.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `host_group_id` - (Required, ForceNew) The id of host group.
* `ip` - (Optional, ForceNew) The ip address.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
The TlsHost is not support import.

