---
subcategory: "ALB"
layout: "volcengine"
page_title: "Volcengine: volcengine_alb_tls_access_log"
sidebar_current: "docs-volcengine-resource-alb_tls_access_log"
description: |-
  Provides a resource to manage alb tls access log
---
# volcengine_alb_tls_access_log
Provides a resource to manage alb tls access log
## Example Usage
```hcl
# Enable ALB TLS Access Log (TLS Topic)
resource "volcengine_alb_tls_access_log" "default" {
  load_balancer_id = "alb-bdchexlt87pc8dv40nbr6mu7"
  topic_id         = "a63a5016-3a68-4723-a754-235a09653ce8"
  project_id       = "3746fa99-3eda-42ab-b2c2-a0bf5d6b26ac"
}
```
## Argument Reference
The following arguments are supported:
* `load_balancer_id` - (Required, ForceNew) The ID of the LoadBalancer.
* `project_id` - (Required, ForceNew) The project ID of the Topic.
* `topic_id` - (Required, ForceNew) The ID of the Topic.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
The AlbTlsAccessLog is not support import.

