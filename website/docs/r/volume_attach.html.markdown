---
subcategory: "EBS"
layout: "vestack"
page_title: "Vestack: vestack_volume_attach"
sidebar_current: "docs-vestack-resource-volume_attach"
description: |-
  Provides a resource to manage volume attach
---
# vestack_volume_attach
Provides a resource to manage volume attach
## Example Usage
```hcl
resource "vestack_volume_attach" "foo" {
  volume_id   = "vol-3tzl52wubz3b9fciw7ev"
  instance_id = "i-4ay59ww7dq8dt9c29hd4"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required, ForceNew) The Id of Instance.
* `volume_id` - (Required, ForceNew) The Id of Volume.
* `delete_with_instance` - (Optional) Delete Volume with Attached Instance.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `created_at` - Creation time of Volume.
* `status` - Status of Volume.
* `updated_at` - Update time of Volume.


## Import
VolumeAttach can be imported using the id, e.g.
```
$ terraform import vestack_volume_attach.default vol-abc12345:i-abc12345
```

