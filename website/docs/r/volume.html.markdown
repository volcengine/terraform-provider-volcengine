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

// create PrePaid ecs instance
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

// create PrePaid data volume
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

// create PostPaid data volume
resource "volcengine_volume" "PostVolume" {
  volume_name = "acc-test-volume"
  volume_type = "ESSD_PL0"
  description = "acc-test"
  kind        = "data"
  size        = 40
  #  snapshot_id        = "snap-3vydtmc0fl3qunm4****"
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
* `volume_type` - (Required) The type of Volume, the value is `PTSSD` or `ESSD_PL0` or `ESSD_PL1` or `ESSD_PL2` or `ESSD_FlexPL`.
* `zone_id` - (Required, ForceNew) The id of the Zone.
* `delete_with_instance` - (Optional) Delete Volume with Attached Instance.
* `description` - (Optional) The description of the Volume.
* `extra_performance_iops` - (Optional) The extra IOPS performance size for volume. Unit: times per second. The valid values for `Balance` and `IOPS` is 0~50000.
* `extra_performance_throughput_mb` - (Optional) The extra Throughput performance size for volume. Unit: MB/s. The valid values for ESSD FlexPL volume is 0~650.
* `extra_performance_type_id` - (Optional) The type of extra performance for volume. The valid values for ESSD FlexPL volume are `Throughput`, `Balance`, `IOPS`. The valid value for TSSD_TL0 volume is `Throughput`.
* `instance_id` - (Optional, ForceNew) The ID of the instance to which the created volume is automatically attached. When use this field to attach ecs instance, the attached volume cannot be deleted by terraform, please use `terraform state rm volcengine_volume.resource_name` command to remove it from terraform state file and management.
* `project_name` - (Optional) The ProjectName of the Volume.
* `snapshot_id` - (Optional, ForceNew) The id of the snapshot. When creating a volume using snapshots, this field is required.
When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `tags` - (Optional) Tags.
* `volume_charge_type` - (Optional) The charge type of the Volume, the value is `PostPaid` or `PrePaid`. The `PrePaid` volume cannot be detached.

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

