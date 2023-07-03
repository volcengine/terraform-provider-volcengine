---
subcategory: "ECS"
layout: "volcengine"
page_title: "Volcengine: volcengine_ecs_launch_templates"
sidebar_current: "docs-volcengine-datasource-ecs_launch_templates"
description: |-
  Use this data source to query detailed information of ecs launch templates
---
# volcengine_ecs_launch_templates
Use this data source to query detailed information of ecs launch templates
## Example Usage
```hcl
data "volcengine_ecs_launch_templates" "default" {

}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of launch template ids.
* `launch_template_names` - (Optional) A list of launch template names.
* `name_regex` - (Optional) A Name Regex of scaling policy.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `launch_templates` - The collection of launch templates.
    * `created_at` - The created time of the launch template.
    * `default_version_number` - The default version of the launch template.
    * `description` - The description of the instance.
    * `eip_bandwidth` - The EIP bandwidth which the scaling configuration set.
    * `eip_billing_type` - The EIP billing type which the scaling configuration set. Valid values: PostPaidByBandwidth, PostPaidByTraffic.
    * `eip_isp` - The EIP ISP which the scaling configuration set. Valid values: BGP, ChinaMobile, ChinaUnicom, ChinaTelecom.
    * `host_name` - The host name of the instance.
    * `hpc_cluster_id` - The hpc cluster id.
    * `id` - The id of the launch template.
    * `image_id` - The image id.
    * `instance_charge_type` - The charge type of the instance and volume.
    * `instance_name` - The name of the instance.
    * `key_pair_name` - When you log in to the instance using the SSH key pair, enter the name of the key pair.
    * `latest_version_number` - The latest version of the launch template.
    * `launch_template_id` - The id of the launch template.
    * `launch_template_name` - The name of the launch template.
    * `network_interfaces` - The list of network interfaces.
        * `security_group_ids` - The security group ID associated with the NIC.
        * `subnet_id` - The private network subnet ID of the instance, when creating the instance, supports binding the secondary NIC at the same time.
    * `security_enhancement_strategy` - Whether to open the security reinforcement.
    * `suffix_index` - The index of the ordered suffix.
    * `unique_suffix` - Indicates whether the ordered suffix is automatically added to Hostname and InstanceName when multiple instances are created.
    * `updated_at` - The updated time of the launch template.
    * `version_description` - The latest version description of the launch template.
    * `volumes` - The list of volume of the scaling configuration.
        * `delete_with_instance` - The delete with instance flag of volume. Valid values: true, false. Default value: true.
        * `size` - The size of volume.
        * `volume_type` - The type of volume.
    * `vpc_id` - The vpc id.
    * `zone_id` - The zone ID of the instance.
* `total_count` - The total count of scaling policy query.


