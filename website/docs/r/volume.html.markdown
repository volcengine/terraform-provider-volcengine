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
  volume_name = "terraform-test"
  zone_id     = "cn-beijing-a"
  volume_type = "PTSSD"
  kind        = "data"
  size        = 40
}
```
## Argument Reference
The following arguments are supported:
* `kind` - (Required, ForceNew) The kind of Volume.
* `size` - (Required) The size of Volume.
* `volume_name` - (Required) The name of Volume.
* `volume_type` - (Required, ForceNew) The type of Volume, the value is `PTSSD` or `ESSD_PL0` or `ESSD_PL1` or `ESSD_PL2` or `ESSD_FlexPL`.
* `zone_id` - (Required, ForceNew) The id of the Zone.
* `delete_with_instance` - (Optional) Delete Volume with Attached Instance.
* `description` - (Optional) The description of the Volume.
* `volume_charge_type` - (Optional, ForceNew) The charge type of the Volume.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `billing_type` - Billing type of Volume.
* `created_at` - Creation time of Volume.
* `pay_type` - Pay type of Volume.
* `status` - Status of Volume.
* `trade_status` - Status of Trade.


## Import
Volume can be imported using the id, e.g.
```
$ terraform import volcengine_volume.default vol-mizl7m1kqccg5smt1bdpijuj
```

