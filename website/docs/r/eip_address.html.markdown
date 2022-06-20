---
subcategory: "EIP"
layout: "vestack"
page_title: "Vestack: vestack_eip_address"
sidebar_current: "docs-vestack-resource-eip_address"
description: |-
  Provides a resource to manage eip address
---
# vestack_eip_address
Provides a resource to manage eip address
## Example Usage
```hcl
resource "vestack_eip_address" "foo" {
  billing_type = "PostPaidByBandwidth"
  bandwidth    = 1
  isp          = "BGP"
  name         = "tf-test"
  description  = "tf-test"
}
```
## Argument Reference
The following arguments are supported:
* `billing_type` - (Required, ForceNew) The billing type of the EIP Address. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic`.
* `bandwidth` - (Optional) The peek bandwidth of the EIP.
* `description` - (Optional) The description of the EIP.
* `isp` - (Optional, ForceNew) The ISP of the EIP.
* `name` - (Optional) The name of the EIP Address.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `eip_address` - The ip address of the EIP.
* `status` - The status of the EIP.


## Import
Eip address can be imported using the id, e.g.
```
$ terraform import volcstack_eip_address.default eip-274oj9a8rs9a87fap8sf9515b
```

