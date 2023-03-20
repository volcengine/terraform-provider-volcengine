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
## Example Usage
```hcl
resource "volcengine_volume" "foo" {
  volume_name        = "terraform-test"
  zone_id            = "cn-xx-a"
  volume_type        = "ESSD_PL0"
  kind               = "data"
  size               = 40
  volume_charge_type = "PostPaid"
  project_name       = "default"
}

resource "volcengine_volume_attach" "foo" {
  volume_id   = volcengine_volume.foo.id
  instance_id = "i-yc8pfhbafwijutv6s1fv"
}

resource "volcengine_volume" "foo2" {
  volume_name        = "terraform-test3"
  zone_id            = "cn-beijing-b"
  volume_type        = "ESSD_PL0"
  kind               = "data"
  size               = 40
  volume_charge_type = "PrePaid"
  instance_id        = "i-yc8pfhbafwijutv6s1fv"
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
* `project_name` - (Optional) The ProjectName of the Volume.
* `volume_charge_type` - (Optional) The charge type of the Volume, the value is `PostPaid` or `PrePaid`. The `PrePaid` volume cannot be detached. Cannot convert `PrePaid` volume to `PostPaid`.Please note that `PrePaid` type needs to ask the system administrator to apply for a whitelist.

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

