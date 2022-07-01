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
  zone_id              = "cn-beijing-a"
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
* `system_volume_type` - (Required, ForceNew) The type of system volume.
* `zone_id` - (Required, ForceNew) The available zone ID of ECS instance.
* `auto_renew_period` - (Optional, ForceNew) The auto renew period of ECS instance.Only effective when instance_charge_type is PrePaid. Default is 1.
* `auto_renew` - (Optional, ForceNew) The auto renew flag of ECS instance.Only effective when instance_charge_type is PrePaid. Default is true.
* `data_volumes` - (Optional) The data volume collection of  ECS instance.
* `description` - (Optional) The description of ECS instance.
* `host_name` - (Optional, ForceNew) The host name of ECS instance.
* `hpc_cluster_id` - (Optional, ForceNew) The hpc cluster ID of ECS instance.
* `include_data_volumes` - (Optional) The include data volumes flag of ECS instance.Only effective when change instance charge type.include_data_volumes.
* `instance_charge_type` - (Optional) The charge type of ECS instance.
* `instance_name` - (Optional) The name of ECS instance.
* `key_pair_name` - (Optional, ForceNew) The ssh key name of ECS instance.
* `password` - (Optional) The password of ECS instance.
* `period` - (Optional) The period of ECS instance.Only effective when instance_charge_type is PrePaid. Default is 12. Unit is Month.
* `secondary_network_interfaces` - (Optional) The secondary networkInterface detail collection of ECS instance.
* `security_enhancement_strategy` - (Optional, ForceNew) The security enhancement strategy of ECS instance.Default is true.
* `user_data` - (Optional) The user data of ECS instance.

The `data_volumes` object supports the following:

* `size` - (Required, ForceNew) The size of volume.
* `volume_type` - (Required, ForceNew) The type of volume.
* `delete_with_instance` - (Optional, ForceNew) The delete with instance flag of volume.

The `secondary_network_interfaces` object supports the following:

* `security_group_ids` - (Required, ForceNew) The security group ID set of secondary networkInterface.
* `subnet_id` - (Required, ForceNew) The subnet ID of secondary networkInterface.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `cpus` - The number of ECS instance CPU cores.
* `created_at` - The create time of ECS instance.
* `instance_id` - The ID of ECS instance.
* `key_pair_id` - The ssh key ID of ECS instance.
* `memory_size` - The memory size of ECS instance.
* `network_interface_id` - The ID of primary networkInterface.
* `os_name` - The os name of ECS instance.
* `os_type` - The os type of ECS instance.
* `status` - The status of ECS instance.
* `stopped_mode` - The stop mode of ECS instance.
* `system_volume_id` - The ID of system volume.
* `updated_at` - The update time of ECS instance.
* `vpc_id` - The VPC ID of ECS instance.


## Import
ECS Instance can be imported using the id, e.g.
```
$ terraform import volcengine_ecs_instance.default i-mizl7m1kqccg5smt1bdpijuj
```

