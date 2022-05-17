---
subcategory: "EBS"
layout: "vestack"
page_title: "Vestack: vestack_volume"
sidebar_current: "docs-vestack-resource-volume"
description: |-
  Provides a resource to manage volume
---
# vestack_volume
Provides a resource to manage volume
## Example Usage
```hcl
resource "vestack_volume" "foo" {
  volume_name = "terraform-test"
  zone_id     = "cn-lingqiu-a"
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
* `volume_type` - (Required, ForceNew) The type of Volume.
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
$ terraform import vestack_volume.default vol-mizl7m1kqccg5smt1bdpijuj
```

