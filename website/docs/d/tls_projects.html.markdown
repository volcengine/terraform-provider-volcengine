---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_projects"
sidebar_current: "docs-volcengine-datasource-tls_projects"
description: |-
  Use this data source to query detailed information of tls projects
---
# volcengine_tls_projects
Use this data source to query detailed information of tls projects
## Example Usage
```hcl
data "volcengine_tls_projects" "default" {
  #project_id = "e020c978-4f05-40e1-9167-0113d3ef****"
}
```
## Argument Reference
The following arguments are supported:
* `iam_project_name` - (Optional) The IAM project name of the tls project.
* `is_full_name` - (Optional) Whether to match accurately when filtering based on ProjectName.
* `name_regex` - (Optional) A Name Regex of tls project.
* `output_file` - (Optional) File name where to save data source results.
* `project_id` - (Optional) The id of tls project. This field supports fuzzy queries. It is not supported to specify both ProjectName and ProjectId at the same time.
* `project_name` - (Optional) The name of tls project. This field supports fuzzy queries. It is not supported to specify both ProjectName and ProjectId at the same time.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `tls_projects` - The collection of tls project query.
    * `create_time` - The create time of the tls project.
    * `description` - The description of the tls project.
    * `iam_project_name` - The IAM project name of the tls project.
    * `id` - The ID of the tls project.
    * `inner_net_domain` - The inner net domain of the tls project.
    * `project_id` - The ID of the tls project.
    * `project_name` - The name of the tls project.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `topic_count` - The count of topics in the tls project.
* `total_count` - The total count of tls project query.


