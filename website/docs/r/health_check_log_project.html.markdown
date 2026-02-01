---
subcategory: "CLB"
layout: "volcengine"
page_title: "Volcengine: volcengine_health_check_log_project"
sidebar_current: "docs-volcengine-resource-health_check_log_project"
description: |-
  Provides a resource to manage health check log project
---
# volcengine_health_check_log_project
Provides a resource to manage health check log project
## Example Usage
```hcl
resource "volcengine_health_check_log_project" "default" {
  # No additional parameters are required for creating the health check log project
}
```
## Argument Reference
The following arguments are supported:


## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `log_project_id` - The ID of the health check log project.


## Import
HealthCheckLogProject can be imported using the id, e.g.
```
$ terraform import volcengine_health_check_log_project.default log_project_id(e.g. b8e16846-fb31-4a2c-a8c1-171434d41d15)
```

