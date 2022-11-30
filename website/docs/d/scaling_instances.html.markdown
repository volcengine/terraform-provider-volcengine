---
subcategory: "AUTOSCALING"
layout: "volcengine"
page_title: "Volcengine: volcengine_scaling_instances"
sidebar_current: "docs-volcengine-datasource-scaling_instances"
description: |-
  Use this data source to query detailed information of scaling instances
---
# volcengine_scaling_instances
Use this data source to query detailed information of scaling instances
## Example Usage
```hcl
data "volcengine_scaling_instances" "default" {
  scaling_group_id         = "scg-ybtawtznszgh9yv8agcp"
  ids                      = ["i-ybzl4u5uogl8j07hgcbg", "i-ybyncrcpzpgh9zmlct0w", "i-ybyncrcpzogh9z4ax9tv"]
  scaling_configuration_id = "scc-ybtawzucw95pkgon0wst"
  status                   = "InService"
}
```
## Argument Reference
The following arguments are supported:
* `scaling_group_id` - (Required) The id of the scaling group.
* `creation_type` - (Optional) The creation type of the instances. Valid values: AutoCreated, Attached.
* `ids` - (Optional) A list of instance ids.
* `output_file` - (Optional) File name where to save data source results.
* `scaling_configuration_id` - (Optional) The id of the scaling configuration id.
* `status` - (Optional) The status of instances. Valid values: Init, Pending, Pending:Wait, InService, Error, Removing, Removing:Wait, Stopped, Protected.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `scaling_instances` - The collection of scaling instances.
    * `created_time` - The time when the instance was added to the scaling group.
    * `creation_type` - The creation type of the instance. Valid values: AutoCreated, Attached.
    * `entrusted` - Whether to host the instance to a scaling group.
    * `id` - The id of the scaling instance.
    * `instance_id` - The id of the scaling instance.
    * `scaling_configuration_id` - The id of the scaling configuration.
    * `scaling_group_id` - The id of the scaling group.
    * `scaling_policy_id` - The id of the scaling policy.
    * `status` - The status of instances.
* `total_count` - The total count of scaling instances query.


