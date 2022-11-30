---
subcategory: "AUTOSCALING"
layout: "volcengine"
page_title: "Volcengine: volcengine_scaling_instance_attachment"
sidebar_current: "docs-volcengine-resource-scaling_instance_attachment"
description: |-
  Provides a resource to manage scaling instance attachment
---
# volcengine_scaling_instance_attachment
Provides a resource to manage scaling instance attachment
## Example Usage
```hcl
resource "volcengine_scaling_instance_attachment" "foo" {
  scaling_group_id = "scg-yc23rtcea88hcchybf8g"
  instance_id      = "i-yc23soxj50gsnz7rxnjp"
  delete_type      = "Remove"
  entrusted        = true
  detach_option    = "none"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required, ForceNew) The id of the instance.
* `scaling_group_id` - (Required, ForceNew) The id of the scaling group.
* `delete_type` - (Optional) The type of delete activity. Valid values: Remove, Detach. Default value is Remove.
* `detach_option` - (Optional) Whether to cancel the association of the instance with the load balancing and public network IP. Valid values: both, none. Default value is both.
* `entrusted` - (Optional) Whether to host the instance to a scaling group. Default value is false.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Scaling instance attachment can be imported using the scaling_group_id and instance_id, e.g.
```
$ terraform import volcengine_scaling_instance_attachment.default scg-mizl7m1kqccg5smt1bdpijuj:i-l8u2ai4j0fauo6mrpgk8
```

