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
## Notice
When Destroy this resource,If the resource charge type is PrePaid,Please unsubscribe the resource 
in  [Volcengine Console](https://console.volcengine.com/finance/unsubscribe/),when complete console operation,yon can
use 'terraform state rm ${resourceId}' to remove.
## Example Usage
```hcl
resource "volcengine_scaling_configuration" "foo" {
  scaling_configuration_name    = "tf-test"
  scaling_group_id              = "scg-ycinx27x25gh9y31p0fy"
  image_id                      = "image-ycgud4t4hxgso0e27bdl"
  instance_types                = ["ecs.g2i.large"]
  instance_name                 = "tf-test"
  instance_description          = ""
  host_name                     = ""
  password                      = ""
  key_pair_name                 = "tf-keypair"
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
  security_group_ids = ["sg-2fepz3c793g1s59gp67y21r34"]
  eip_bandwidth      = 10
  eip_isp            = "ChinaMobile"
  eip_billing_type   = "PostPaidByBandwidth"
  user_data          = "IyEvYmluL2Jhc2gKZWNobyAidGVzdCI="
  tags {
    key   = "tf-key1"
    value = "tf-value1"
  }
  tags {
    key   = "tf-key2"
    value = "tf-value2"
  }
  project_name   = "default"
  hpc_cluster_id = ""
  spot_strategy  = "NoSpot"
}
```
## Argument Reference
The following arguments are supported:
* `image_id` - (Required) The ECS image id which the scaling configuration set.
* `instance_name` - (Required) The ECS instance name which the scaling configuration set.
* `instance_types` - (Required) The list of the ECS instance type which the scaling configuration set. The maximum number of instance types is 10.
* `scaling_configuration_name` - (Required) The name of the scaling configuration.
* `scaling_group_id` - (Required, ForceNew) The id of the scaling group to which the scaling configuration belongs.
* `security_group_ids` - (Required) The list of the security group id of the networkInterface which the scaling configuration set. A maximum of 5 security groups can be bound at the same time, and the value ranges from 1 to 5.
* `volumes` - (Required) The list of volume of the scaling configuration. The number of supported volumes ranges from 1 to 15.
* `eip_bandwidth` - (Optional) The EIP bandwidth which the scaling configuration set. When the value of Eip.BillingType is PostPaidByBandwidth, the value is 1 to 500. When the value of Eip.BillingType is PostPaidByTraffic, the value is 1 to 200.
* `eip_billing_type` - (Optional) The EIP billing type which the scaling configuration set. Valid values: PostPaidByBandwidth, PostPaidByTraffic.
* `eip_isp` - (Optional) The EIP ISP which the scaling configuration set. Valid values: BGP, ChinaMobile, ChinaUnicom, ChinaTelecom.
* `host_name` - (Optional) The ECS hostname which the scaling configuration set.
* `hpc_cluster_id` - (Optional) The ID of the HPC cluster to which the instance belongs. Valid only when InstanceTypes.N specifies High Performance Computing GPU Type.
* `instance_description` - (Optional) The ECS instance description which the scaling configuration set.
* `key_pair_name` - (Optional) The ECS key pair name which the scaling configuration set.
* `password` - (Optional) The ECS password which the scaling configuration set.
* `project_name` - (Optional) The project to which the instance created by the scaling configuration belongs.
* `security_enhancement_strategy` - (Optional) The Ecs security enhancement strategy which the scaling configuration set. Valid values: Active, InActive.
* `spot_strategy` - (Optional) The preemption policy of the instance. Valid Value: NoSpot (default), SpotAsPriceGo.
* `tags` - (Optional) The label of the instance created by the scaling configuration. Up to 20 tags are supported.
* `user_data` - (Optional) The ECS user data which the scaling configuration set.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

The `volumes` object supports the following:

* `size` - (Required) The size of volume. System disk value range: 10 - 500. The value range of the data disk: 10 - 8192.
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

