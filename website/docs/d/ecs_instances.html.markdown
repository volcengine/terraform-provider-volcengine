---
subcategory: "ECS"
layout: "volcengine"
page_title: "Volcengine: volcengine_ecs_instances"
sidebar_current: "docs-volcengine-datasource-ecs_instances"
description: |-
  Use this data source to query detailed information of ecs instances
---
# volcengine_ecs_instances
Use this data source to query detailed information of ecs instances
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

resource "volcengine_ecs_instance" "foo" {
  instance_name        = "acc-test-ecs-${count.index}"
  description          = "acc-test"
  host_name            = "tf-acc-test"
  image_id             = data.volcengine_images.foo.images[0].image_id
  instance_type        = "ecs.g1.large"
  password             = "93f0cb0614Aab12"
  instance_charge_type = "PostPaid"
  system_volume_type   = "ESSD_PL0"
  system_volume_size   = 40
  data_volumes {
    volume_type          = "ESSD_PL0"
    size                 = 50
    delete_with_instance = true
  }
  subnet_id          = volcengine_subnet.foo.id
  security_group_ids = [volcengine_security_group.foo.id]
  project_name       = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
  count = 2
}

data "volcengine_ecs_instances" "foo" {
  ids = volcengine_ecs_instance.foo[*].id
}
```
## Argument Reference
The following arguments are supported:
* `deployment_set_ids` - (Optional) A list of DeploymentSet IDs.
* `eip_addresses` - (Optional) A list of Eip addresses.
* `hpc_cluster_id` - (Optional) The hpc cluster ID of ECS instance.
* `ids` - (Optional) A list of ECS instance IDs.
* `instance_charge_type` - (Optional) The charge type of ECS instance.
* `instance_name` - (Optional) The name of ECS instance. This field support fuzzy query.
* `instance_type_families` - (Optional) A list of instance type families.
* `instance_type_ids` - (Optional) A list of instance type IDs.
* `ipv6_addresses` - (Optional) A list of ipv6 addresses.
* `key_pair_name` - (Optional) The key pair name of ECS instance.
* `name_regex` - (Optional) A Name Regex of ECS instance.
* `output_file` - (Optional) File name where to save data source results.
* `primary_ip_address` - (Optional) The primary ip address of ECS instance.
* `project_name` - (Optional) The ProjectName of ECS instance.
* `status` - (Optional) The status of ECS instance.
* `tags` - (Optional) Tags.
* `vpc_id` - (Optional) The VPC ID of ECS instance.
* `zone_id` - (Optional) The available zone ID of ECS instance.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `instances` - The collection of ECS instance query.
    * `cpus` - The number of ECS instance CPU cores.
    * `created_at` - The create time of ECS instance.
    * `deployment_set_id` - The ID of DeploymentSet.
    * `description` - The description of ECS instance.
    * `gpu_devices` - The GPU device info of Instance.
        * `count` - The Count of GPU device.
        * `encrypted_memory_size` - The Encrypted Memory Size of GPU device.
        * `memory_size` - The Memory Size of GPU device.
        * `product_name` - The Product Name of GPU device.
    * `host_name` - The host name of ECS instance.
    * `image_id` - The image ID of ECS instance.
    * `instance_charge_type` - The charge type of ECS instance.
    * `instance_id` - The ID of ECS instance.
    * `instance_name` - The name of ECS instance.
    * `instance_type` - The spec type of ECS instance.
    * `ipv6_address_count` - The number of IPv6 addresses of the ECS instance.
    * `ipv6_addresses` - The  IPv6 address list of the ECS instance.
    * `is_gpu` - The Flag of GPU instance.If the instance is GPU,The flag is true.
    * `key_pair_id` - The ssh key ID of ECS instance.
    * `key_pair_name` - The ssh key name of ECS instance.
    * `memory_size` - The memory size of ECS instance.
    * `network_interfaces` - The networkInterface detail collection of ECS instance.
        * `mac_address` - The mac address of networkInterface.
        * `network_interface_id` - The ID of networkInterface.
        * `primary_ip_address` - The private ip address of networkInterface.
        * `subnet_id` - The subnet ID of networkInterface.
        * `type` - The type of networkInterface.
        * `vpc_id` - The ID of networkInterface.
    * `os_name` - The os name of ECS instance.
    * `os_type` - The os type of ECS instance.
    * `project_name` - The ProjectName of ECS instance.
    * `spot_price_limit` - The spot price limit of ECS instance.
    * `spot_strategy` - The spot strategy of ECS instance.
    * `status` - The status of ECS instance.
    * `stopped_mode` - The stop mode of ECS instance.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `updated_at` - The update time of ECS instance.
    * `volumes` - The volume detail collection of volume.
        * `delete_with_instance` - The delete with instance flag of volume.
        * `size` - The size of volume.
        * `volume_id` - The ID of volume.
        * `volume_name` - The Name of volume.
        * `volume_type` - The type of volume.
    * `vpc_id` - The VPC ID of ECS instance.
    * `zone_id` - The available zone ID of ECS instance.
* `total_count` - The total count of ECS instance query.


