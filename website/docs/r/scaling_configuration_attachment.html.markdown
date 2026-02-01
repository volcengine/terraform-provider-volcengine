---
subcategory: "AUTOSCALING"
layout: "volcengine"
page_title: "Volcengine: volcengine_scaling_configuration_attachment"
sidebar_current: "docs-volcengine-resource-scaling_configuration_attachment"
description: |-
  Provides a resource to manage scaling configuration attachment
---
# volcengine_scaling_configuration_attachment
Provides a resource to manage scaling configuration attachment
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
  image_id                   = data.volcengine_images.foo.images[0].image_id
  instance_name              = "acc-test-instance"
  instance_types             = ["ecs.g1.large"]
  password                   = "93f0cb0614Aab12"
  scaling_configuration_name = "acc-test-scaling-config"
  scaling_group_id           = volcengine_scaling_group.foo.id
  security_group_ids         = [volcengine_security_group.foo.id]
  volumes {
    volume_type          = "ESSD_PL0"
    size                 = 50
    delete_with_instance = true
  }
}

resource "volcengine_scaling_configuration_attachment" "foo" {
  scaling_configuration_id = volcengine_scaling_configuration.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `scaling_configuration_id` - (Required, ForceNew) The id of the scaling configuration.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Scaling Configuration attachment can be imported using the scaling_configuration_id e.g.
The launch template and scaling configuration cannot take effect at the same time.
```
$ terraform import volcengine_scaling_configuration_attachment.default enable:scc-ybrurj4uw6gh9zecj327
```

