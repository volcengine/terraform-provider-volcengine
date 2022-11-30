---
subcategory: "AUTOSCALING"
layout: "volcengine"
page_title: "Volcengine: volcengine_scaling_configurations"
sidebar_current: "docs-volcengine-datasource-scaling_configurations"
description: |-
  Use this data source to query detailed information of scaling configurations
---
# volcengine_scaling_configurations
Use this data source to query detailed information of scaling configurations
## Example Usage
```hcl
data "volcengine_scaling_configurations" "default" {
  ids = ["scc-ybrurj4uw6gh9zecj327"]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of scaling configuration ids.
* `name_regex` - (Optional) A Name Regex of scaling configuration.
* `output_file` - (Optional) File name where to save data source results.
* `scaling_configuration_names` - (Optional) A list of scaling configuration names.
* `scaling_group_id` - (Optional) An id of scaling group.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `scaling_configurations` - The collection of scaling configuration query.
    * `created_at` - The create time of the scaling configuration.
    * `eip_bandwidth` - The EIP bandwidth which the scaling configuration set.
    * `eip_billing_type` - The EIP ISP which the scaling configuration set.
    * `eip_isp` - The EIP ISP which the scaling configuration set.
    * `host_name` - The ECS hostname which the scaling configuration set.
    * `id` - The id of the scaling configuration.
    * `image_id` - The ECS image id which the scaling configuration set.
    * `instance_description` - The ECS instance description which the scaling configuration set.
    * `instance_name` - The ECS instance name which the scaling configuration set.
    * `instance_types` - The list of the ECS instance type which the scaling configuration set.
    * `key_pair_name` - The ECS key pair name which the scaling configuration set.
    * `lifecycle_state` - The lifecycle state of the scaling configuration.
    * `scaling_configuration_id` - The id of the scaling configuration.
    * `scaling_configuration_name` - The name of the scaling configuration.
    * `scaling_group_id` - The id of the scaling group to which the scaling configuration belongs.
    * `security_enhancement_strategy` - The Ecs security enhancement strategy which the scaling configuration set.
    * `security_group_ids` - The list of the security group id of the networkInterface which the scaling configuration set.
    * `updated_at` - The create time of the scaling configuration.
    * `user_data` - The ECS user data which the scaling configuration set.
    * `volumes` - The list of volume of the scaling configuration.
        * `delete_with_instance` - The delete with instance flag of volume.
        * `size` - The size of volume.
        * `volume_type` - The type of volume.
* `total_count` - The total count of scaling configuration query.


