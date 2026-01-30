---
subcategory: "CLB"
layout: "volcengine"
page_title: "Volcengine: volcengine_access_log"
sidebar_current: "docs-volcengine-resource-access_log"
description: |-
  Provides a resource to manage access log
---
# volcengine_access_log
Provides a resource to manage access log
## Example Usage
```hcl
# Enable CLB Access Log (TOS Bucket)
resource "volcengine_access_log" "tos_example" {
  load_balancer_id = "clb-13g5i2cbg6nsw3n6nu5r*****"
  bucket_name      = "tos-bucket"
}

# Enable CLB Access Log (TLS)
resource "volcengine_access_log" "tls_example" {
  load_balancer_id = "clb-13g5i2cbg6nsw3n6nu5r*****"
  delivery_type    = "tls"
  tls_project_id   = "d8c6e4c2-8d22-****-****-9811f2067580"
  tls_topic_id     = "081aa4ff-991b-****-****-5d573dcf4ba4"
}
```
## Argument Reference
The following arguments are supported:
* `load_balancer_id` - (Required, ForceNew) The ID of the CLB instance.
* `bucket_name` - (Optional, ForceNew) The name of the TOS bucket for storing access logs. Required when delivery_type is 'tos'.
* `delivery_type` - (Optional, ForceNew) The type of log delivery. Valid values: 'tos', 'tls'. Default: 'tos'.
* `tls_project_id` - (Optional, ForceNew) The ID of the TLS project. Required when delivery_type is 'tls'.
* `tls_topic_id` - (Optional, ForceNew) The ID of the TLS topic. Required when delivery_type is 'tls'.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
The AccessLog is not support import.

