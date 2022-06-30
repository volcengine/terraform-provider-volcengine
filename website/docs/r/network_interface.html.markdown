---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_network_interface"
sidebar_current: "docs-volcengine-resource-network_interface"
description: |-
  Provides a resource to manage network interface
---
# volcengine_network_interface
Provides a resource to manage network interface
## Example Usage
```hcl
resource "volcengine_network_interface" "foo" {
  subnet_id              = "subnet-2744ht7fhjthc7fap8tm10eqg"
  security_group_ids     = ["sg-2744hspo7jbpc7fap8t7lef1p"]
  primary_ip_address     = "192.168.0.253"
  network_interface_name = "tf-test-up"
  description            = "tf-test-up"
  port_security_enabled  = false
}
```
## Argument Reference
The following arguments are supported:
* `security_group_ids` - (Required) The list of the security group id to which the secondary ENI belongs.
* `subnet_id` - (Required, ForceNew) The id of the subnet to which the ENI is connected.
* `description` - (Optional) The description of the ENI.
* `network_interface_name` - (Optional) The name of the ENI.
* `port_security_enabled` - (Optional) Set port security enable or disable.
* `primary_ip_address` - (Optional, ForceNew) The primary IP address of the ENI.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `status` - The status of the ENI.


## Import
Network interface can be imported using the id, e.g.
```
$ terraform import volcengine_network_interface.default eni-bp1fgnh68xyz9****
```

