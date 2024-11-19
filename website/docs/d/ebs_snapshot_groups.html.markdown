---
subcategory: "EBS"
layout: "volcengine"
page_title: "Volcengine: volcengine_ebs_snapshot_groups"
sidebar_current: "docs-volcengine-datasource-ebs_snapshot_groups"
description: |-
  Use this data source to query detailed information of ebs snapshot groups
---
# volcengine_ebs_snapshot_groups
Use this data source to query detailed information of ebs snapshot groups
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

data "volcengine_ebs_snapshot_groups" "foo" {
  ids = [volcengine_ebs_snapshot_group.foo.id]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of snapshot group IDs.
* `instance_id` - (Optional) The instance id of snapshot group.
* `name_regex` - (Optional) A Name Regex of Resource.
* `name` - (Optional) The name of snapshot group.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name of snapshot group.
* `status` - (Optional) A list of snapshot group status. Valid values: `creating`, `available`, `failed`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `snapshot_groups` - The collection of query.
    * `creation_time` - The creation time of the snapshot group.
    * `description` - The description of the snapshot group.
    * `id` - The id of the snapshot group.
    * `image_id` - The image id of the snapshot group.
    * `instance_id` - The instance id of the snapshot group.
    * `name` - The name of the snapshot group.
    * `project_name` - The project name of the snapshot group.
    * `snapshot_group_id` - The id of the snapshot group.
    * `snapshots` - The snapshots of the snapshot group.
        * `creation_time` - The creation time of the snapshot.
        * `description` - The description of the snapshot.
        * `image_id` - The image id of the snapshot.
        * `progress` - The progress of the snapshot.
        * `project_name` - The id of the snapshot.
        * `retention_days` - The id of the snapshot.
        * `snapshot_id` - The id of the snapshot.
        * `snapshot_name` - The name of the snapshot.
        * `snapshot_type` - The type of the snapshot.
        * `status` - The status of the snapshot.
        * `tags` - Tags.
            * `key` - The Key of Tags.
            * `value` - The Value of Tags.
        * `volume_id` - The volume id of the snapshot.
        * `volume_kind` - The volume kind of the snapshot.
        * `volume_name` - The volume name of the snapshot.
        * `volume_size` - The volume size of the snapshot.
        * `volume_status` - The volume status of the snapshot.
        * `volume_type` - The volume type of the snapshot.
        * `zone_id` - The zone id of the snapshot.
    * `status` - The status of the snapshot group.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
* `total_count` - The total count of query.


