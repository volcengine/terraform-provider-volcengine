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
## Notice
When Destroy this resource,If the resource charge type is PrePaid,Please unsubscribe the resource 
in  [Volcengine Console](https://console.volcengine.com/finance/unsubscribe/),when complete console operation,yon can
use 'terraform state rm ${resourceId}' to remove.
## Example Usage
```hcl
// query available zones in current region
data "volcengine_zones" "foo" {
}

// create vpc
resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

// create subnet
resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

// create security group
resource "volcengine_security_group" "foo" {
  security_group_name = "acc-test-security-group"
  vpc_id              = volcengine_vpc.foo.id
}

// query the image_id which match the specified instance_type
data "volcengine_images" "foo" {
  os_type          = "Linux"
  visibility       = "public"
  instance_type_id = "ecs.g1.large"
}

// create ecs instance
resource "volcengine_ecs_instance" "foo" {
  instance_name        = "acc-test-ecs"
  description          = "acc-test"
  host_name            = "tf-acc-test"
  image_id             = data.volcengine_images.foo.images[0].image_id
  instance_type        = "ecs.g1.large"
  password             = "93f0cb0614Aab12"
  instance_charge_type = "PostPaid"
  system_volume_type   = "ESSD_PL0"
  system_volume_size   = 40
  subnet_id            = volcengine_subnet.foo.id
  security_group_ids   = [volcengine_security_group.foo.id]
  project_name         = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

// create ebs data volume
resource "volcengine_volume" "foo" {
  volume_name        = "acc-test-volume"
  volume_type        = "ESSD_PL0"
  description        = "acc-test"
  kind               = "data"
  size               = 40
  zone_id            = data.volcengine_zones.foo.zones[0].id
  volume_charge_type = "PostPaid"
  project_name       = "default"
}

// attach ebs data volume to ecs instance
resource "volcengine_volume_attach" "foo" {
  instance_id = volcengine_ecs_instance.foo.id
  volume_id   = volcengine_volume.foo.id
}

// create eip
resource "volcengine_eip_address" "foo" {
  billing_type = "PostPaidByTraffic"
}

// associate eip to ecs instance
resource "volcengine_eip_associate" "foo" {
  allocation_id = volcengine_eip_address.foo.id
  instance_id   = volcengine_ecs_instance.foo.id
  instance_type = "EcsInstance"
}
```
## Argument Reference
The following arguments are supported:
* `image_id` - (Required) The Image ID of ECS instance.
* `instance_type` - (Required) The instance type of ECS instance.
* `security_group_ids` - (Required) The security group ID set of primary networkInterface.
* `subnet_id` - (Required, ForceNew) The subnet ID of primary networkInterface.
* `system_volume_size` - (Required) The size of system volume. The value range of the system volume size is ESSD_PL0: 20~2048, ESSD_FlexPL: 20~2048, PTSSD: 10~500.
* `system_volume_type` - (Required, ForceNew) The type of system volume, the value is `PTSSD` or `ESSD_PL0` or `ESSD_PL1` or `ESSD_PL2` or `ESSD_FlexPL`.
* `auto_renew_period` - (Optional) The auto renew period of ECS instance.Only effective when instance_charge_type is PrePaid. Default is 1.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `auto_renew` - (Optional) The auto renew flag of ECS instance.Only effective when instance_charge_type is PrePaid. Default is true.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `cpu_options` - (Optional) The option of cpu,only support for ebm.
* `data_volumes` - (Optional) The data volumes collection of  ECS instance.
* `deployment_set_id` - (Optional) The ID of Ecs Deployment Set. This field only used to associate a deployment set to the ECS instance. Setting this field to null means disassociating the instance from the deployment set. 
The current deployment set id of the ECS instance is the `deployment_set_id_computed` field.
* `description` - (Optional) The description of ECS instance.
* `eip_address` - (Optional, ForceNew) The config of the eip which will be automatically created and assigned to this instance. `Prepaid` type eip cannot be created in this way, please use `volcengine_eip_address`.
When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `eip_id` - (Optional, ForceNew) The id of an existing Available EIP which will be automatically assigned to this instance. 
It is not recommended to use this field, it is recommended to use `volcengine_eip_associate` resource to bind EIP.
* `host_name` - (Optional, ForceNew) The host name of ECS instance.
* `hpc_cluster_id` - (Optional, ForceNew) The hpc cluster ID of ECS instance.
* `include_data_volumes` - (Optional) The include data volumes flag of ECS instance.Only effective when change instance charge type.include_data_volumes.
* `instance_charge_type` - (Optional) The charge type of ECS instance, the value can be `PrePaid` or `PostPaid`.
* `instance_name` - (Optional) The name of ECS instance.
* `ipv6_address_count` - (Optional, ForceNew) The number of IPv6 addresses to be automatically assigned from within the CIDR block of the subnet that hosts the ENI. Valid values: 1 to 10.
* `ipv6_addresses` - (Optional, ForceNew) One or more IPv6 addresses selected from within the CIDR block of the subnet that hosts the ENI. Support up to 10.
 You cannot specify both the ipv6_addresses and ipv6_address_count parameters.
* `keep_image_credential` - (Optional) Whether to keep the mirror settings. Only custom images and shared images support this field.
 When the value of this field is true, the Password and KeyPairName cannot be specified.
 When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `key_pair_name` - (Optional) The ssh key name of ECS instance. This field can be modified only when the `image_id` is modified.
* `password` - (Optional) The password of ECS instance.
* `period` - (Optional) The period of ECS instance.Only effective when instance_charge_type is PrePaid. Default is 12. Unit is Month.
* `primary_ip_address` - (Optional, ForceNew) The private ip address of primary networkInterface.
* `project_name` - (Optional) The ProjectName of the ecs instance.
* `secondary_network_interfaces` - (Optional, ForceNew) The secondary networkInterface detail collection of ECS instance.
* `security_enhancement_strategy` - (Optional, ForceNew) The security enhancement strategy of ECS instance. The value can be Active or InActive. Default is Active.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `spot_price_limit` - (Optional, ForceNew) The maximum hourly price for spot instances supports up to three decimal places. This parameter only takes effect when SpotStrategy=SpotWithPriceLimit.
* `spot_strategy` - (Optional, ForceNew) The spot strategy will autoremove instance in some conditions.Please make sure you can maintain instance lifecycle before auto remove.The spot strategy of ECS instance, values:
 NoSpot (default): indicates creating a normal pay-as-you-go instance.
SpotAsPriceGo: spot instance with system automatically bidding and following the current market price.
SpotWithPriceLimit: spot instance with a set upper limit for bidding price.
* `tags` - (Optional) Tags.
* `user_data` - (Optional) The user data of ECS instance, this field must be encrypted with base64.
* `zone_id` - (Optional, ForceNew) The available zone ID of ECS instance.

The `cpu_options` object supports the following:

* `numa_per_socket` - (Optional, ForceNew) The number of subnuma in socket, only support for ebm. `1` indicates disabling SNC/NPS function. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `threads_per_core` - (Optional, ForceNew) The per core of threads, only support for ebm. `1` indicates disabling hyper threading function.

The `data_volumes` object supports the following:

* `size` - (Required, ForceNew) The size of volume. The value range of the data volume size is ESSD_PL0: 10~32768, ESSD_FlexPL: 10~32768, PTSSD: 20~8192.
* `volume_type` - (Required, ForceNew) The type of volume, the value is `PTSSD` or `ESSD_PL0` or `ESSD_PL1` or `ESSD_PL2` or `ESSD_FlexPL`.
* `delete_with_instance` - (Optional, ForceNew) The delete with instance flag of volume.

The `eip_address` object supports the following:

* `bandwidth_mbps` - (Optional, ForceNew) The peek bandwidth of the EIP. The value range in 1~500 for PostPaidByBandwidth, and 1~200 for PostPaidByTraffic. Default is 1.
* `bandwidth_package_id` - (Optional, ForceNew) The id of the bandwidth package, indicates that the public IP address will be added to the bandwidth package.
* `charge_type` - (Optional, ForceNew) The billing type of the EIP Address. Valid values: `PayByBandwidth`, `PayByTraffic`. Default is `PayByBandwidth`.
* `isp` - (Optional, ForceNew) The ISP of the EIP. Valid values: `BGP`, `ChinaMobile`, `ChinaUnicom`, `ChinaTelecom`, `SingleLine_BGP`, `Static_BGP`.

The `secondary_network_interfaces` object supports the following:

* `security_group_ids` - (Required, ForceNew) The security group ID set of secondary networkInterface.
* `subnet_id` - (Required, ForceNew) The subnet ID of secondary networkInterface.
* `primary_ip_address` - (Optional, ForceNew) The private ip address of secondary networkInterface.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `cpus` - The number of ECS instance CPU cores.
* `created_at` - The create time of ECS instance.
* `deployment_set_id_computed` - The ID of Ecs Deployment Set. Computed field.
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

