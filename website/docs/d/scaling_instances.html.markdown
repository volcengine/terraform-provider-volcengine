---
subcategory: "AUTOSCALING"
layout: "volcengine"
page_title: "Volcengine: volcengine_scaling_instances"
sidebar_current: "docs-volcengine-datasource-scaling_instances"
description: |-
  Use this data source to query detailed information of scaling instances
---
# volcengine_scaling_instances
Use this data source to query detailed information of scaling instances
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

resource "volcengine_ecs_key_pair" "foo" {
  description   = "acc-test-2"
  key_pair_name = "acc-test-key-pair-name"
}

resource "volcengine_ecs_launch_template" "foo" {
  description          = "acc-test-desc"
  eip_bandwidth        = 200
  eip_billing_type     = "PostPaidByBandwidth"
  eip_isp              = "BGP"
  host_name            = "acc-hostname"
  image_id             = data.volcengine_images.foo.images[0].image_id
  instance_charge_type = "PostPaid"
  instance_name        = "acc-instance-name"
  instance_type_id     = "ecs.g1.large"
  key_pair_name        = volcengine_ecs_key_pair.foo.key_pair_name
  launch_template_name = "acc-test-template"
  network_interfaces {
    subnet_id          = volcengine_subnet.foo.id
    security_group_ids = [volcengine_security_group.foo.id]
  }
  volumes {
    volume_type          = "ESSD_PL0"
    size                 = 50
    delete_with_instance = true
  }
}

resource "volcengine_scaling_group" "foo" {
  scaling_group_name        = "acc-test-scaling-group"
  subnet_ids                = [volcengine_subnet.foo.id]
  multi_az_policy           = "BALANCE"
  desire_instance_number    = -1
  min_instance_number       = 0
  max_instance_number       = 10
  instance_terminate_policy = "OldestInstance"
  default_cooldown          = 10
  launch_template_id        = volcengine_ecs_launch_template.foo.id
  launch_template_version   = "Default"
}

resource "volcengine_scaling_group_enabler" "foo" {
  scaling_group_id = volcengine_scaling_group.foo.id
}

resource "volcengine_ecs_instance" "foo" {
  count                = 3
  instance_name        = "acc-test-ecs-${count.index}"
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
}

resource "volcengine_scaling_instance_attachment" "foo" {
  count            = length(volcengine_ecs_instance.foo)
  instance_id      = volcengine_ecs_instance.foo[count.index].id
  scaling_group_id = volcengine_scaling_group.foo.id
  entrusted        = true

  depends_on = [
    volcengine_scaling_group_enabler.foo
  ]
}

data "volcengine_scaling_instances" "foo" {
  scaling_group_id = volcengine_scaling_group.foo.id
  ids              = volcengine_scaling_instance_attachment.foo[*].instance_id
}
```
## Argument Reference
The following arguments are supported:
* `scaling_group_id` - (Required) The id of the scaling group.
* `creation_type` - (Optional) The creation type of the instances. Valid values: AutoCreated, Attached.
* `ids` - (Optional) A list of instance ids.
* `output_file` - (Optional) File name where to save data source results.
* `scaling_configuration_id` - (Optional) The id of the scaling configuration id.
* `status` - (Optional) The status of instances. Valid values: Init, Pending, Pending:Wait, InService, Error, Removing, Removing:Wait, Stopped, Protected.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `scaling_instances` - The collection of scaling instances.
    * `created_time` - The time when the instance was added to the scaling group.
    * `creation_type` - The creation type of the instance. Valid values: AutoCreated, Attached.
    * `entrusted` - Whether to host the instance to a scaling group.
    * `id` - The id of the scaling instance.
    * `instance_id` - The id of the scaling instance.
    * `scaling_configuration_id` - The id of the scaling configuration.
    * `scaling_group_id` - The id of the scaling group.
    * `scaling_policy_id` - The id of the scaling policy.
    * `status` - The status of instances.
* `total_count` - The total count of scaling instances query.


