---
subcategory: "AUTOSCALING"
layout: "volcengine"
page_title: "Volcengine: volcengine_scaling_instance_attachment"
sidebar_current: "docs-volcengine-resource-scaling_instance_attachment"
description: |-
  Provides a resource to manage scaling instance attachment
---
# volcengine_scaling_instance_attachment
Provides a resource to manage scaling instance attachment
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
  max_instance_number       = 1
  instance_terminate_policy = "OldestInstance"
  default_cooldown          = 10
  launch_template_id        = volcengine_ecs_launch_template.foo.id
  launch_template_version   = "Default"
}

resource "volcengine_scaling_group_enabler" "foo" {
  scaling_group_id = volcengine_scaling_group.foo.id
}

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
}

resource "volcengine_scaling_instance_attachment" "foo" {
  instance_id      = volcengine_ecs_instance.foo.id
  scaling_group_id = volcengine_scaling_group.foo.id
  entrusted        = true

  depends_on = [
    volcengine_scaling_group_enabler.foo
  ]
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required, ForceNew) The id of the instance.
* `scaling_group_id` - (Required, ForceNew) The id of the scaling group.
* `delete_type` - (Optional) The type of delete activity. Valid values: Remove, Detach. Default value is Remove.
* `detach_option` - (Optional) Whether to cancel the association of the instance with the load balancing and public network IP. Valid values: both, none. Default value is both.
* `entrusted` - (Optional, ForceNew) Whether to host the instance to a scaling group. Default value is false.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Scaling instance attachment can be imported using the scaling_group_id and instance_id, e.g.
You can choose to remove or detach the instance according to the `delete_type` field.
```
$ terraform import volcengine_scaling_instance_attachment.default scg-mizl7m1kqccg5smt1bdpijuj:i-l8u2ai4j0fauo6mrpgk8
```

