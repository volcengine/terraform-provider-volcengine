---
subcategory: "EBS"
layout: "volcengine"
page_title: "Volcengine: volcengine_volume"
sidebar_current: "docs-volcengine-resource-volume"
description: |-
  Provides a resource to manage volume
---
# volcengine_volume
Provides a resource to manage volume
## Notice
When Destroy this resource,If the resource charge type is PrePaid,Please unsubscribe the resource 
in  [Volcengine Console](https://console.volcengine.com/finance/unsubscribe/),when complete console operation,yon can
use 'terraform state rm ${resourceId}' to remove.
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
  instance_name        = "acc-test-ecs"
  description          = "acc-test"
  host_name            = "tf-acc-test"
  image_id             = data.volcengine_images.foo.images[0].image_id
  instance_type        = "ecs.g1.large"
  password             = "93f0cb0614Aab12"
  instance_charge_type = "PrePaid"
  period               = 1
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

resource "volcengine_volume" "PreVolume" {
  volume_name          = "acc-test-volume"
  volume_type          = "ESSD_PL0"
  description          = "acc-test"
  kind                 = "data"
  size                 = 40
  zone_id              = data.volcengine_zones.foo.zones[0].id
  volume_charge_type   = "PrePaid"
  instance_id          = volcengine_ecs_instance.foo.id
  project_name         = "default"
  delete_with_instance = true
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_volume" "PostVolume" {
  volume_name        = "acc-test-volume"
  volume_type        = "ESSD_PL0"
  description        = "acc-test"
  kind               = "data"
  size               = 40
  zone_id            = data.volcengine_zones.foo.zones[0].id
  volume_charge_type = "PostPaid"
  project_name       = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `kind` - (Required, ForceNew) The kind of Volume, the value is `data`.
* `size` - (Required) The size of Volume.
* `volume_name` - (Required) The name of Volume.
* `volume_type` - (Required, ForceNew) The type of Volume, the value is `PTSSD` or `ESSD_PL0` or `ESSD_PL1` or `ESSD_PL2` or `ESSD_FlexPL`.
* `zone_id` - (Required, ForceNew) The id of the Zone.
* `delete_with_instance` - (Optional) Delete Volume with Attached Instance.
* `description` - (Optional) The description of the Volume.
* `instance_id` - (Optional, ForceNew) The ID of the instance to which the created volume is automatically attached. Please note this field needs to ask the system administrator to apply for a whitelist.
When use this field to attach ecs instance, the attached volume cannot be deleted by terraform, please use `terraform state rm volcengine_volume.resource_name` command to remove it from terraform state file and management.
* `project_name` - (Optional) The ProjectName of the Volume.
* `tags` - (Optional) Tags.
* `volume_charge_type` - (Optional) The charge type of the Volume, the value is `PostPaid` or `PrePaid`. The `PrePaid` volume cannot be detached. Please note that `PrePaid` type needs to ask the system administrator to apply for a whitelist.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `created_at` - Creation time of Volume.
* `status` - Status of Volume.
* `trade_status` - Status of Trade.


## Import
Volume can be imported using the id, e.g.
```
$ terraform import volcengine_volume.default vol-mizl7m1kqccg5smt1bdpijuj
```

