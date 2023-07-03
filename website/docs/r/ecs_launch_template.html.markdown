---
subcategory: "ECS"
layout: "volcengine"
page_title: "Volcengine: volcengine_ecs_launch_template"
sidebar_current: "docs-volcengine-resource-ecs_launch_template"
description: |-
  Provides a resource to manage ecs launch template
---
# volcengine_ecs_launch_template
Provides a resource to manage ecs launch template
## Notice
When Destroy this resource,If the resource charge type is PrePaid,Please unsubscribe the resource 
in  [Volcengine Console](https://console.volcengine.com/finance/unsubscribe/),when complete console operation,yon can
use 'terraform state rm ${resourceId}' to remove.
## Example Usage
```hcl
resource "volcengine_ecs_launch_template" "foo" {
  launch_template_name = "tf-test-zzm"
  version_description  = "testterraformaa"
  instance_type_id     = "ecs.g1.large"
  image_id             = "image-ycb26d1ryzl8j1fcxa9m"
  instance_name        = "instance-test"
  instance_charge_type = "PostPaid"
  host_name            = "instance-host-name"
  volumes {
    volume_type = "ESSD_PL0"
    size        = 20
    //delete_with_instance = false
  }

  network_interfaces {
    subnet_id          = "subnet-3tispp1nai4e8idddd"
    security_group_ids = ["sg-xxxxxxxx"]
  }
  eip_bandwidth    = 3
  eip_isp          = "BGP"
  eip_billing_type = "PostPaidByBandwidth"
  user_data        = "IyEvYmluL2Jhc2gKZWNobyAidGVzdCI="
}
```
## Argument Reference
The following arguments are supported:
* `launch_template_name` - (Required, ForceNew) The name of the launch template.
* `description` - (Optional) The description of the instance.
* `eip_bandwidth` - (Optional) The EIP bandwidth which the scaling configuration set.
* `eip_billing_type` - (Optional) The EIP billing type which the scaling configuration set. Valid values: PostPaidByBandwidth, PostPaidByTraffic.
* `eip_isp` - (Optional) The EIP ISP which the scaling configuration set. Valid values: BGP, ChinaMobile, ChinaUnicom, ChinaTelecom.
* `host_name` - (Optional) The host name of the instance.
* `hpc_cluster_id` - (Optional) The hpc cluster id.
* `image_id` - (Optional) The image ID.
* `instance_charge_type` - (Optional) The charge type of the instance and volume.
* `instance_name` - (Optional) The name of the instance.
* `instance_type_id` - (Optional) The compute type of the instance.
* `key_pair_name` - (Optional) When you log in to the instance using the SSH key pair, enter the name of the key pair.
* `network_interfaces` - (Optional) The list of network interfaces. When creating an instance, it is supported to bind auxiliary network cards at the same time. The first one is the primary network card, and the others are secondary network cards.
* `security_enhancement_strategy` - (Optional) Whether to open the security reinforcement.
* `suffix_index` - (Optional) The index of the ordered suffix.
* `unique_suffix` - (Optional) Indicates whether the ordered suffix is automatically added to Hostname and InstanceName when multiple instances are created.
* `user_data` - (Optional) Instance custom data. The set custom data must be Base64 encoded, and the size of the custom data before Base64 encoding cannot exceed 16KB.
* `version_description` - (Optional) The latest version description of the launch template.
* `volumes` - (Optional) The list of volume of the scaling configuration.
* `vpc_id` - (Optional) The vpc id.
* `zone_id` - (Optional) The zone id.

The `network_interfaces` object supports the following:

* `security_group_ids` - (Optional) The security group ID associated with the NIC.
* `subnet_id` - (Optional) The private network subnet ID of the instance, when creating the instance, supports binding the secondary NIC at the same time.

The `volumes` object supports the following:

* `delete_with_instance` - (Optional) The delete with instance flag of volume. Valid values: true, false. Default value: true.
* `size` - (Optional) The size of volume.
* `volume_type` - (Optional) The type of volume.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `launch_template_id` - The launch template id.


## Import
LaunchTemplate can be imported using the LaunchTemplateId, e.g.
When the instance launch template is modified, a new version will be created.
When the number of versions reaches the upper limit (30), the oldest version that is not the default version will be deleted.
```
$ terraform import volcengine_ecs_launch_template.default lt-ysxc16auaugh9zfy****
```

