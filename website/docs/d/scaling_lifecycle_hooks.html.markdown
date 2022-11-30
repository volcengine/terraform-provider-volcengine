---
subcategory: "AUTOSCALING"
layout: "volcengine"
page_title: "Volcengine: volcengine_scaling_lifecycle_hooks"
sidebar_current: "docs-volcengine-datasource-scaling_lifecycle_hooks"
description: |-
  Use this data source to query detailed information of scaling lifecycle hooks
---
# volcengine_scaling_lifecycle_hooks
Use this data source to query detailed information of scaling lifecycle hooks
## Example Usage
```hcl
data "volcengine_scaling_lifecycle_hooks" "default" {
  scaling_group_id = "scg-ybru8pazhgl8j1di4tyd"
}
```
## Argument Reference
The following arguments are supported:
* `scaling_group_id` - (Required) An id of scaling group id.
* `ids` - (Optional) A list of lifecycle hook ids.
* `lifecycle_hook_names` - (Optional) A list of lifecycle hook names.
* `name_regex` - (Optional) A Name Regex of lifecycle hook.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `lifecycle_hooks` - The collection of lifecycle hook query.
    * `id` - The id of the lifecycle hook.
    * `lifecycle_hook_id` - The id of the lifecycle hook.
    * `lifecycle_hook_name` - The name of the lifecycle hook.
    * `lifecycle_hook_policy` - The policy of the lifecycle hook.
    * `lifecycle_hook_timeout` - The timeout of the lifecycle hook.
    * `lifecycle_hook_type` - The type of the lifecycle hook.
    * `scaling_group_id` - The id of the scaling group.
* `total_count` - The total count of lifecycle hook query.


