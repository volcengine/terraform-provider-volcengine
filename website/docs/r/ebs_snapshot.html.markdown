---
subcategory: "EBS"
layout: "volcengine"
page_title: "Volcengine: volcengine_ebs_snapshot"
sidebar_current: "docs-volcengine-resource-ebs_snapshot"
description: |-
  Provides a resource to manage ebs snapshot
---
# volcengine_ebs_snapshot
Provides a resource to manage ebs snapshot
## Example Usage
```hcl
data "volcengine_zones" "foo" {
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

resource "volcengine_ebs_snapshot" "foo" {
  volume_id      = volcengine_volume.foo.id
  snapshot_name  = "acc-test-snapshot"
  description    = "acc-test"
  retention_days = 3
  project_name   = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `snapshot_name` - (Required) The name of the snapshot.
* `volume_id` - (Required, ForceNew) The volume id to create snapshot.
* `description` - (Optional) The description of the snapshot.
* `project_name` - (Optional) The project name of the snapshot.
* `retention_days` - (Optional) The retention days of the snapshot. Valid values: 1~65536. Not specifying this field means permanently preserving the snapshot.When modifying this field, the retention days only supports extension and not shortening. The value range is N+1~65536, where N is the retention days set during snapshot creation.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `creation_time` - The creation time of the snapshot.
* `snapshot_type` - The type of the snapshot.
* `status` - The status of the snapshot.
* `volume_kind` - The volume kind of the snapshot.
* `volume_name` - The volume name of the snapshot.
* `volume_size` - The volume size of the snapshot.
* `volume_status` - The volume status of the snapshot.
* `volume_type` - The volume type of the snapshot.
* `zone_id` - The zone id of the snapshot.


## Import
EbsSnapshot can be imported using the id, e.g.
```
$ terraform import volcengine_ebs_snapshot.default resource_id
```

