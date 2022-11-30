---
subcategory: "AUTOSCALING"
layout: "volcengine"
page_title: "Volcengine: volcengine_scaling_configuration"
sidebar_current: "docs-volcengine-resource-scaling_configuration"
description: |-
  Provides a resource to manage scaling configuration
---
# volcengine_scaling_configuration
Provides a resource to manage scaling configuration
## Example Usage
```hcl
resource "volcengine_scaling_configuration" "foo" {
  scaling_configuration_name    = "tf-test"
  scaling_group_id              = "scg-ybru8pazhgl8j1di4tyd"
  image_id                      = "image-ybpbrfay1gl8j1srwwyz"
  instance_types                = ["ecs.g1.4xlarge"]
  instance_name                 = "tf-test"
  instance_description          = ""
  host_name                     = ""
  password                      = ""
  key_pair_name                 = "renhuaxi"
  security_enhancement_strategy = "InActive"
  volumes {
    volume_type          = "ESSD_PL0"
    size                 = 20
    delete_with_instance = false
  }
  volumes {
    volume_type          = "ESSD_PL0"
    size                 = 20
    delete_with_instance = true
  }
  security_group_ids = ["sg-2ff4fhdtlo8ao59gp67iiq9o3"]
  eip_bandwidth      = 0
  eip_isp            = "ChinaMobile"
  eip_billing_type   = "PostPaidByBandwidth"
  user_data          = "IyEvYmluL2Jhc2gKZWNobyAidGVzdCI="
}
```
## Argument Reference
The following arguments are supported:
* `image_id` - (Required) The ECS image id which the scaling configuration set.
* `instance_name` - (Required) The ECS instance name which the scaling configuration set.
* `instance_types` - (Required) The list of the ECS instance type which the scaling configuration set.
* `scaling_configuration_name` - (Required) The name of the scaling configuration.
* `scaling_group_id` - (Required, ForceNew) The id of the scaling group to which the scaling configuration belongs.
* `security_group_ids` - (Required) The list of the security group id of the networkInterface which the scaling configuration set.
* `volumes` - (Required) The list of volume of the scaling configuration.
* `eip_bandwidth` - (Optional) The EIP bandwidth which the scaling configuration set.
* `eip_billing_type` - (Optional) The EIP billing type which the scaling configuration set. Valid values: PostPaidByBandwidth, PostPaidByTraffic.
* `eip_isp` - (Optional) The EIP ISP which the scaling configuration set. Valid values: BGP, ChinaMobile, ChinaUnicom, ChinaTelecom.
* `host_name` - (Optional) The ECS hostname which the scaling configuration set.
* `instance_description` - (Optional) The ECS instance description which the scaling configuration set.
* `key_pair_name` - (Optional) The ECS key pair name which the scaling configuration set.
* `password` - (Optional) The ECS password which the scaling configuration set.
* `security_enhancement_strategy` - (Optional) The Ecs security enhancement strategy which the scaling configuration set. Valid values: Active, InActive.
* `user_data` - (Optional) The ECS user data which the scaling configuration set.

The `volumes` object supports the following:

* `size` - (Required) The size of volume.
* `volume_type` - (Required) The type of volume.
* `delete_with_instance` - (Optional) The delete with instance flag of volume. Valid values: true, false. Default value: true.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `created_at` - The create time of the scaling configuration.
* `lifecycle_state` - The lifecycle state of the scaling configuration.
* `scaling_configuration_id` - The id of the scaling configuration.
* `updated_at` - The create time of the scaling configuration.


## Import
ScalingConfiguration can be imported using the id, e.g.
```
$ terraform import volcengine_scaling_configuration.default scc-ybkuck3mx8cm9tm5yglz
```

