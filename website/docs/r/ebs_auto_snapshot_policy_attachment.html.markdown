---
subcategory: "EBS"
layout: "volcengine"
page_title: "Volcengine: volcengine_ebs_auto_snapshot_policy_attachment"
sidebar_current: "docs-volcengine-resource-ebs_auto_snapshot_policy_attachment"
description: |-
  Provides a resource to manage ebs auto snapshot policy attachment
---
# volcengine_ebs_auto_snapshot_policy_attachment
Provides a resource to manage ebs auto snapshot policy attachment
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

resource "volcengine_ebs_auto_snapshot_policy" "foo" {
  auto_snapshot_policy_name = "acc-test-auto-snapshot-policy"
  time_points               = [1, 5, 9]
  retention_days            = -1
  repeat_weekdays           = [2, 6]
  project_name              = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_ebs_auto_snapshot_policy_attachment" "foo" {
  auto_snapshot_policy_id = volcengine_ebs_auto_snapshot_policy.foo.id
  volume_id               = volcengine_volume.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `auto_snapshot_policy_id` - (Required, ForceNew) The id of the auto snapshot policy.
* `volume_id` - (Required, ForceNew) The id of the volume.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
EbsAutoSnapshotPolicyAttachment can be imported using the auto_snapshot_policy_id:volume_id, e.g.
```
$ terraform import volcengine_ebs_auto_snapshot_policy_attachment.default resource_id
```

