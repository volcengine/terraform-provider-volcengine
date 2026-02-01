---
subcategory: "ALB"
layout: "volcengine"
page_title: "Volcengine: volcengine_alb_access_log"
sidebar_current: "docs-volcengine-resource-alb_access_log"
description: |-
  Provides a resource to manage alb access log
---
# volcengine_alb_access_log
Provides a resource to manage alb access log
## Example Usage
```hcl
# Enable ALB Access Log (TOS Bucket)
resource "volcengine_alb_access_log" "default" {
  load_balancer_id = "alb-bdchexlt87pc8dv40nbr6mu7"
  bucket_name      = "tos-buket"
}
```
## Argument Reference
The following arguments are supported:
* `bucket_name` - (Required, ForceNew) The name of the TOS bucket for storing access logs.
* `load_balancer_id` - (Required, ForceNew) The ID of the LoadBalancer.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
The AlbAccessLog is not support import.

