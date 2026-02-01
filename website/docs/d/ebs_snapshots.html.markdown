---
subcategory: "EBS"
layout: "volcengine"
page_title: "Volcengine: volcengine_ebs_snapshots"
sidebar_current: "docs-volcengine-datasource-ebs_snapshots"
description: |-
  Use this data source to query detailed information of ebs snapshots
---
# volcengine_ebs_snapshots
Use this data source to query detailed information of ebs snapshots
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
  count = 2
}

data "volcengine_ebs_snapshots" "foo" {
  ids = volcengine_ebs_snapshot.foo[*].id
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of snapshot IDs.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name of snapshot.
* `snapshot_status` - (Optional) A list of snapshot status.
* `tags` - (Optional) Tags.
* `zone_id` - (Optional) The zone id of snapshot.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `snapshots` - The collection of query.
    * `creation_time` - The creation time of the snapshot.
    * `description` - The description of the snapshot.
    * `id` - The id of the snapshot.
    * `project_name` - The project name of the snapshot.
    * `retention_days` - The retention days of the snapshot.
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
* `total_count` - The total count of query.


