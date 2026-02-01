---
subcategory: "ECS"
layout: "volcengine"
page_title: "Volcengine: volcengine_ecs_key_pair_associate"
sidebar_current: "docs-volcengine-resource-ecs_key_pair_associate"
description: |-
  Provides a resource to manage ecs key pair associate
---
# volcengine_ecs_key_pair_associate
Provides a resource to manage ecs key pair associate
## Example Usage
```hcl
resource "volcengine_ecs_key_pair" "foo" {
  key_pair_name = "acc-test-key-name"
  description   = "acc-test"
}

data "volcengine_zones" "foo" {
}

data "volcengine_images" "foo" {
  os_type          = "Linux"
  visibility       = "public"
  instance_type_id = "ecs.g1.large"
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
  vpc_id              = volcengine_vpc.foo.id
  security_group_name = "acc-test-security-group"
}

resource "volcengine_ecs_instance" "foo" {
  image_id             = data.volcengine_images.foo.images[0].image_id
  instance_type        = "ecs.g1.large"
  instance_name        = "acc-test-ecs-name"
  password             = "your password"
  instance_charge_type = "PostPaid"
  system_volume_type   = "ESSD_PL0"
  system_volume_size   = 40
  subnet_id            = volcengine_subnet.foo.id
  security_group_ids   = [volcengine_security_group.foo.id]
}

resource "volcengine_ecs_key_pair_associate" "foo" {
  instance_id = volcengine_ecs_instance.foo.id
  key_pair_id = volcengine_ecs_key_pair.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required, ForceNew) The ID of ECS Instance.
* `key_pair_id` - (Required, ForceNew) The ID of ECS KeyPair Associate.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
ECS key pair associate can be imported using the id, e.g.

After binding the key pair, the instance needs to be restarted for the key pair to take effect.

After the key pair is bound, the password login method will automatically become invalid. If your instance has been set for password login, after the key pair is bound, you will no longer be able to use the password login method.

```
$ terraform import volcengine_ecs_key_pair_associate.default kp-ybti5tkpkv2udbfolrft:i-mizl7m1kqccg5smt1bdpijuj
```

