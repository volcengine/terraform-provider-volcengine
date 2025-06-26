---
subcategory: "NAS"
layout: "volcengine"
page_title: "Volcengine: volcengine_nas_auto_snapshot_policy_apply"
sidebar_current: "docs-volcengine-resource-nas_auto_snapshot_policy_apply"
description: |-
  Provides a resource to manage nas auto snapshot policy apply
---
# volcengine_nas_auto_snapshot_policy_apply
Provides a resource to manage nas auto snapshot policy apply
## Example Usage
```hcl
data "volcengine_nas_zones" "foo" {

}

resource "volcengine_nas_file_system" "foo" {
  file_system_name = "acc-test-fs"
  description      = "acc-test"
  zone_id          = data.volcengine_nas_zones.foo.zones[0].id
  capacity         = 103
  project_name     = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_nas_auto_snapshot_policy" "foo" {
  auto_snapshot_policy_name = "acc-test-auto_snapshot_policy"
  repeat_weekdays           = "1,3,5,7"
  time_points               = "0,7,17"
  retention_days            = 20
}

resource "volcengine_nas_auto_snapshot_policy_apply" "foo" {
  file_system_id          = volcengine_nas_file_system.foo.id
  auto_snapshot_policy_id = volcengine_nas_auto_snapshot_policy.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `auto_snapshot_policy_id` - (Required, ForceNew) The id of auto snapshot policy.
* `file_system_id` - (Required, ForceNew) The id of file system.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
NasAutoSnapshotPolicyApply can be imported using the auto_snapshot_policy_id:file_system_id, e.g.
```
$ terraform import volcengine_nas_auto_snapshot_policy_apply.default resource_id
```

