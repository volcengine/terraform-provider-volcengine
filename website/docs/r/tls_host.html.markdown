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
resource "volcengine_tls_host" "foo" {
  host_group_id = "fbea6619-7b0c-40f3-ac7e-45c63e3f676e"
  ip            = "10.180.50.18"
}
```
## Argument Reference
The following arguments are supported:
* `host_group_id` - (Required, ForceNew) The id of host group.
* `ip` - (Required, ForceNew) The ip address.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Tls Host can be imported using the host_group_id:ip, e.g.
```
$ terraform import volcengine_tls_host.default edf051ed-3c46-49:1.1.1.1
```

