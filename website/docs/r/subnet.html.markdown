---
subcategory: "VPC"
layout: "vestack"
page_title: "Vestack: vestack_subnet"
sidebar_current: "docs-vestack-resource-subnet"
description: |-
  Provides a resource to manage subnet
---
# vestack_subnet
Provides a resource to manage subnet
## Example Usage
```hcl
resource "vestack_subnet" "foo" {
  subnet_name = "subnet-test-2"
  cidr_block  = "192.168.1.0/24"
  zone_id     = "cn-beijing"
  vpc_id      = "vpc-2749wnlhro3y87fap8u5ztvt5"
}
```
## Argument Reference
The following arguments are supported:
* `cidr_block` - (Required, ForceNew) A network address block which should be a subnet of the three internal network segments (10.0.0.0/16, 172.16.0.0/12 and 192.168.0.0/16).
* `vpc_id` - (Required, ForceNew) Id of the VPC.
* `zone_id` - (Required, ForceNew) Id of the Zone.
* `description` - (Optional) The description of the Subnet.
* `subnet_name` - (Optional) The name of the Subnet.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `creation_time` - Creation time of Subnet.
* `status` - Status of Subnet.


## Import
Subnet can be imported using the id, e.g.
```
$ terraform import vestack_subnet.default subnet-274oj9a8rs9a87fap8sf9515b
```

