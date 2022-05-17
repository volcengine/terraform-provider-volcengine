---
subcategory: "CLB"
layout: "vestack"
page_title: "Vestack: vestack_clb"
sidebar_current: "docs-vestack-resource-clb"
description: |-
  Provides a resource to manage clb
---
# vestack_clb
Provides a resource to manage clb
## Example Usage
```hcl
resource "vestack_clb" "foo" {
  type               = "public"
  subnet_id          = "subnet-2744i7u9alnnk7fap8tkq8aft"
  load_balancer_spec = "small_1"
  region_id          = "cn-north-3"
  description        = "Demo"
}
```
## Argument Reference
The following arguments are supported:
* `load_balancer_spec` - (Required) The specification of the CLB.
* `region_id` - (Required, ForceNew) The region of the request.
* `subnet_id` - (Required, ForceNew) The id of the Subnet.
* `type` - (Required, ForceNew) The type of the CLB. And optional choice contains `public` or `private`.
* `description` - (Optional) The description of the CLB.
* `eni_address` - (Optional, ForceNew) The eni address of the CLB.
* `load_balancer_billing_type` - (Optional, ForceNew) The billing type of the CLB.
* `load_balancer_name` - (Optional) The name of the CLB.
* `modification_protection_reason` - (Optional) The reason of the console modification protection.
* `modification_protection_status` - (Optional) The status of the console modification protection.
* `vpc_id` - (Optional, ForceNew) The id of the VPC.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
CLB can be imported using the id, e.g.
```
$ terraform import vestack_clb.default clb-273y2ok6ets007fap8txvf6us
```

