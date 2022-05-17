---
subcategory: "VPC"
layout: "vestack"
page_title: "Vestack: vestack_network_interface_attach"
sidebar_current: "docs-vestack-resource-network_interface_attach"
description: |-
  Provides a resource to manage network interface attach
---
# vestack_network_interface_attach
Provides a resource to manage network interface attach
## Example Usage
```hcl
resource "vestack_network_interface_attach" "foo" {
  network_interface_id = "eni-274ecj646ylts7fap8t6xbba1"
  instance_id          = "i-72q20hi6s082wcafdem8"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required, ForceNew) The id of the instance to which the ENI is bound.
* `network_interface_id` - (Required, ForceNew) The id of the ENI.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Network interface attach can be imported using the network_interface_id:instance_id.
```
$ terraform import vestack_network_interface_attach.default eni-bp1fg655nh68xyz9***:i-wijfn35c****
```

