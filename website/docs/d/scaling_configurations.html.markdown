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
data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_security_group" "foo" {
  security_group_name = "acc-test-security-group"
  vpc_id              = volcengine_vpc.foo.id
}

data "volcengine_images" "foo" {
  os_type          = "Linux"
  visibility       = "public"
  instance_type_id = "ecs.g1.large"
}

resource "volcengine_scaling_group" "foo" {
  scaling_group_name        = "acc-test-scaling-group"
  subnet_ids                = [volcengine_subnet.foo.id]
  multi_az_policy           = "BALANCE"
  desire_instance_number    = 0
  min_instance_number       = 0
  max_instance_number       = 1
  instance_terminate_policy = "OldestInstance"
  default_cooldown          = 10
}

resource "volcengine_scaling_configuration" "foo" {
  count                      = 3
  image_id                   = data.volcengine_images.foo.images[0].image_id
  instance_name              = "acc-test-instance"
  instance_types             = ["ecs.g1.large"]
  password                   = "93f0cb0614Aab12"
  scaling_configuration_name = "acc-test-scaling-config-${count.index}"
  scaling_group_id           = volcengine_scaling_group.foo.id
  security_group_ids         = [volcengine_security_group.foo.id]
  volumes {
    volume_type          = "ESSD_PL0"
    size                 = 50
    delete_with_instance = true
  }
}

data "volcengine_scaling_configurations" "foo" {
  ids = volcengine_scaling_configuration.foo[*].id
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
    * `hpc_cluster_id` - The ID of the HPC cluster to which the instance belongs. Valid only when InstanceTypes.N specifies High Performance Computing GPU Type.
    * `id` - The id of the scaling configuration.
    * `image_id` - The ECS image id which the scaling configuration set.
    * `instance_description` - The ECS instance description which the scaling configuration set.
    * `instance_name` - The ECS instance name which the scaling configuration set.
    * `instance_types` - The list of the ECS instance type which the scaling configuration set.
    * `ipv6_address_count` - Assign IPv6 address to instance network card. Possible values:
0: Do not assign IPv6 address.
1: Assign IPv6 address and the system will automatically assign an IPv6 subnet for you.
    * `key_pair_name` - The ECS key pair name which the scaling configuration set.
    * `lifecycle_state` - The lifecycle state of the scaling configuration.
    * `project_name` - The project to which the instance created by the scaling configuration belongs.
    * `scaling_configuration_id` - The id of the scaling configuration.
    * `scaling_configuration_name` - The name of the scaling configuration.
    * `scaling_group_id` - The id of the scaling group to which the scaling configuration belongs.
    * `security_enhancement_strategy` - The Ecs security enhancement strategy which the scaling configuration set.
    * `security_group_ids` - The list of the security group id of the networkInterface which the scaling configuration set.
    * `spot_strategy` - The preemption policy of the instance. Valid Value: NoSpot (default), SpotAsPriceGo.
    * `tags` - The label of the instance created by the scaling configuration.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `updated_at` - The create time of the scaling configuration.
    * `user_data` - The ECS user data which the scaling configuration set.
    * `volumes` - The list of volume of the scaling configuration.
        * `delete_with_instance` - The delete with instance flag of volume.
        * `size` - The size of volume.
        * `volume_type` - The type of volume.
* `total_count` - The total count of scaling configuration query.


