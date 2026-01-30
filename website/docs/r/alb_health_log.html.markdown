---
subcategory: "ALB"
layout: "volcengine"
page_title: "Volcengine: volcengine_alb_health_log"
sidebar_current: "docs-volcengine-resource-alb_health_log"
description: |-
  Provides a resource to manage alb health log
---
# volcengine_alb_health_log
Provides a resource to manage alb health log
## Example Usage
```hcl
# Enable health check log collection
resource "volcengine_alb_health_log" "example" {
  load_balancer_id = "alb-bdchexlt87pc8dv40nbr6mu7"
  topic_id         = "cd507e58-64d2-48e3-9e98-f384430d773a"
  project_id       = "29018d87-858b-4d24-bb8e-5ac958fa5ca5"
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
The AlbHealthLog is not support import.

