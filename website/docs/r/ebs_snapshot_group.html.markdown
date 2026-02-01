---
subcategory: "EBS"
layout: "volcengine"
page_title: "Volcengine: volcengine_ebs_snapshot_group"
sidebar_current: "docs-volcengine-resource-ebs_snapshot_group"
description: |-
  Provides a resource to manage ebs snapshot group
---
# volcengine_ebs_snapshot_group
Provides a resource to manage ebs snapshot group
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
  instance_type_id = "ecs.g3il.large"
}

resource "volcengine_ecs_instance" "foo" {
  instance_name        = "acc-test-ecs"
  description          = "acc-test"
  host_name            = "tf-acc-test"
  image_id             = data.volcengine_images.foo.images[0].image_id
  instance_type        = "ecs.g3il.large"
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

resource "volcengine_volume" "foo" {
  volume_name        = "acc-test-volume"
  volume_type        = "ESSD_PL0"
  description        = "acc-test"
  kind               = "data"
  size               = 500
  zone_id            = data.volcengine_zones.foo.zones[0].id
  volume_charge_type = "PostPaid"
  project_name       = "default"
}

resource "volcengine_volume_attach" "foo" {
  instance_id = volcengine_ecs_instance.foo.id
  volume_id   = volcengine_volume.foo.id
}

resource "volcengine_ebs_snapshot_group" "foo" {
  volume_ids   = [volcengine_ecs_instance.foo.system_volume_id, volcengine_volume.foo.id]
  instance_id  = volcengine_ecs_instance.foo.id
  name         = "acc-test-snapshot-group"
  description  = "acc-test"
  project_name = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
  depends_on = [volcengine_volume_attach.foo]
}
```
## Argument Reference
The following arguments are supported:
* `volume_ids` - (Required, ForceNew) The volume id of the snapshot group. The status of the volume must be `attached`.If multiple volumes are specified, they need to be attached to the same ECS instance.
* `description` - (Optional) The instance id of the snapshot group.
* `instance_id` - (Optional, ForceNew) The instance id of the snapshot group.
* `name` - (Optional) The name of the snapshot group.
* `project_name` - (Optional) The project name of the snapshot group.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `creation_time` - The creation time of the snapshot group.
* `image_id` - The image id of the snapshot group.
* `status` - The status of the snapshot group.


## Import
EbsSnapshotGroup can be imported using the id, e.g.
```
$ terraform import volcengine_ebs_snapshot_group.default resource_id
```

