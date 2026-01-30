---
subcategory: "CLB"
layout: "volcengine"
page_title: "Volcengine: volcengine_health_check_log_projects"
sidebar_current: "docs-volcengine-datasource-health_check_log_projects"
description: |-
  Use this data source to query detailed information of health check log projects
---
# volcengine_health_check_log_projects
Use this data source to query detailed information of health check log projects
## Example Usage
```hcl
data "volcengine_health_check_log_projects" "example" {
}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `health_check_log_projects` - The collection of health check log projects.
    * `id` - The ID of the health check log project.
    * `log_project_id` - The ID of the health check log project.
* `total_count` - The total count of query.


