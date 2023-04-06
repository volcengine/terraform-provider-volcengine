---
subcategory: "ECS"
layout: "volcengine"
page_title: "Volcengine: volcengine_ecs_instance"
sidebar_current: "docs-volcengine-resource-ecs_instance"
description: |-
  Provides a resource to manage ecs instance
---
# volcengine_ecs_instance
Provides a resource to manage ecs instance
## Example Usage
```hcl
resource "volcengine_vpc" "foo" {
  vpc_name   = "tf-test-2"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo1" {
  subnet_name = "subnet-test-1"
  cidr_block  = "172.16.1.0/24"
  zone_id     = "cn-beijing-a"
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_security_group" "foo1" {
  depends_on = [volcengine_subnet.foo1]
  vpc_id     = volcengine_vpc.foo.id
}

resource "volcengine_ecs_instance" "default" {
  image_id             = "image-aagd56zrw2jtdro3bnrl"
  instance_type        = "ecs.g1.large"
  instance_name        = "xym-tf-test-2"
  description          = "xym-tf-test-desc-1"
  password             = "93f0cb0614Aab12"
  instance_charge_type = "PostPaid"
  system_volume_type   = "PTSSD"
  system_volume_size   = 60
  subnet_id            = volcengine_subnet.foo1.id
  security_group_ids   = [volcengine_security_group.foo1.id]
  data_volumes {
    volume_type          = "PTSSD"
    size                 = 100
    delete_with_instance = true
  }
  deployment_set_id = ""
  #  secondary_network_interfaces {
  #    subnet_id = volcengine_subnet.foo1.id
  #    security_group_ids = [volcengine_security_group.foo1.id]
  #  }
}
```
## Argument Reference
The following arguments are supported:
* `image_id` - (Required) The Image ID of ECS instance.
* `instance_type` - (Required) The instance type of ECS instance.
* `security_group_ids` - (Required) The security group ID set of primary networkInterface.
* `subnet_id` - (Required, ForceNew) The subnet ID of primary networkInterface.
* `system_volume_size` - (Required) The size of system volume.
* `system_volume_type` - (Required, ForceNew) The type of system volume, the value is `PTSSD` or `ESSD_PL0` or `ESSD_PL1` or `ESSD_PL2` or `ESSD_FlexPL`.
* `auto_renew_period` - (Optional, ForceNew) The auto renew period of ECS instance.Only effective when instance_charge_type is PrePaid. Default is 1.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `auto_renew` - (Optional, ForceNew) The auto renew flag of ECS instance.Only effective when instance_charge_type is PrePaid. Default is true.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `data_volumes` - (Optional) The data volumes collection of  ECS instance.
* `deployment_set_id` - (Optional) The ID of Ecs Deployment Set.
* `description` - (Optional) The description of ECS instance.
* `host_name` - (Optional, ForceNew) The host name of ECS instance.
* `hpc_cluster_id` - (Optional, ForceNew) The hpc cluster ID of ECS instance.
* `include_data_volumes` - (Optional) The include data volumes flag of ECS instance.Only effective when change instance charge type.include_data_volumes.
* `instance_charge_type` - (Optional) The charge type of ECS instance, the value can be `PrePaid` or `PostPaid`.
* `instance_name` - (Optional) The name of ECS instance.
* `keep_image_credential` - (Optional) Whether to keep the mirror settings. Only custom images and shared images support this field.
 When the value of this field is true, the Password and KeyPairName cannot be specified.
 When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `key_pair_name` - (Optional, ForceNew) The ssh key name of ECS instance.
* `password` - (Optional) The password of ECS instance.
* `period` - (Optional) The period of ECS instance.Only effective when instance_charge_type is PrePaid. Default is 12. Unit is Month.
* `project_name` - (Optional) The ProjectName of the VPC.
* `secondary_network_interfaces` - (Optional) The secondary networkInterface detail collection of ECS instance.
* `security_enhancement_strategy` - (Optional, ForceNew) The security enhancement strategy of ECS instance. The value can be Active or InActive. Default is Active.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `spot_strategy` - (Optional, ForceNew) The spot strategy will autoremove instance in some conditions.Please make sure you can maintain instance lifecycle before auto remove.The spot strategy of ECS instance, the value can be `NoSpot` or `SpotAsPriceGo`.
* `tags` - (Optional) Tags.
* `user_data` - (Optional) The user data of ECS instance, this field must be encrypted with base64.
* `zone_id` - (Optional, ForceNew) The available zone ID of ECS instance.

The `data_volumes` object supports the following:

* `size` - (Required, ForceNew) The size of volume.
* `volume_type` - (Required, ForceNew) The type of volume, the value is `PTSSD` or `ESSD_PL0` or `ESSD_PL1` or `ESSD_PL2` or `ESSD_FlexPL`.
* `delete_with_instance` - (Optional, ForceNew) The delete with instance flag of volume.

The `secondary_network_interfaces` object supports the following:

* `security_group_ids` - (Required, ForceNew) The security group ID set of secondary networkInterface.
* `subnet_id` - (Required, ForceNew) The subnet ID of secondary networkInterface.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `cpus` - The number of ECS instance CPU cores.
* `created_at` - The create time of ECS instance.
* `gpu_devices` - The GPU device info of Instance.
    * `count` - The Count of GPU device.
    * `encrypted_memory_size` - The Encrypted Memory Size of GPU device.
    * `memory_size` - The Memory Size of GPU device.
    * `product_name` - The Product Name of GPU device.
* `instance_id` - The ID of ECS instance.
* `is_gpu` - The Flag of GPU instance.If the instance is GPU,The flag is true.
* `key_pair_id` - The ssh key ID of ECS instance.
* `memory_size` - The memory size of ECS instance.
* `network_interface_id` - The ID of primary networkInterface.
* `os_name` - The os name of ECS instance.
* `os_type` - The os type of ECS instance.
* `primary_ip_address` - The private ip address of primary networkInterface.
* `status` - The status of ECS instance.
* `stopped_mode` - The stop mode of ECS instance.
* `system_volume_id` - The ID of system volume.
* `updated_at` - The update time of ECS instance.
* `vpc_id` - The VPC ID of ECS instance.


## Import
ECS Instance can be imported using the id, e.g.
If Import,The data_volumes is sort by volume name
```
$ terraform import volcengine_ecs_instance.default i-mizl7m1kqccg5smt1bdpijuj
```

