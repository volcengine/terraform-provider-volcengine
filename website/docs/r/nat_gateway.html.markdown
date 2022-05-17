---
subcategory: "NAT"
layout: "vestack"
page_title: "Vestack: vestack_nat_gateway"
sidebar_current: "docs-vestack-resource-nat_gateway"
description: |-
  Provides a resource to manage nat gateway
---
# vestack_nat_gateway
Provides a resource to manage nat gateway
## Example Usage
```hcl
resource "vestack_nat_gateway" "foo" {
  vpc_id           = "vpc-2740cxyk9im0w7fap8u013dfe"
  subnet_id        = "subnet-2740cym8mv9q87fap8u3hfx4i"
  spec             = "Medium"
  nat_gateway_name = "tf-auto-demo-1"
  description      = "This nat gateway auto-created by terraform. "
}
```
## Argument Reference
The following arguments are supported:
* `subnet_id` - (Required, ForceNew) The ID of the Subnet.
* `vpc_id` - (Required, ForceNew) The ID of the VPC.
* `billing_type` - (Optional, ForceNew) The billing type of the NatGateway.
* `description` - (Optional) The description of the NatGateway.
* `nat_gateway_name` - (Optional) The name of the NatGateway.
* `spec` - (Optional) The specification of the NatGateway. Optional choice contains `Small`(default), `Medium`, `Large`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
NatGateway can be imported using the id, e.g.
```
$ terraform import vestack_nat_gateway.default ngw-vv3t043k05sm****
```

