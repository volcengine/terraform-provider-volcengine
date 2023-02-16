---
subcategory: "CLB"
layout: "volcengine"
page_title: "Volcengine: volcengine_clb"
sidebar_current: "docs-volcengine-resource-clb"
description: |-
  Provides a resource to manage clb
---
# volcengine_clb
Provides a resource to manage clb
## Example Usage
```hcl
resource "volcengine_clb" "foo" {
  type               = "public"
  subnet_id          = "subnet-273xjcb6wohs07fap8sz3ihhs"
  load_balancer_spec = "small_1"
  description        = "Demo"
  load_balancer_name = "terraform-auto-create"
}
```
## Argument Reference
The following arguments are supported:
* `load_balancer_spec` - (Required) The specification of the CLB, the value can be `small_1`, `small_2`, `medium_1`, `medium_2`, `large_1`, `large_2`.
* `subnet_id` - (Required, ForceNew) The id of the Subnet.
* `type` - (Required, ForceNew) The type of the CLB. And optional choice contains `public` or `private`.
* `description` - (Optional) The description of the CLB.
* `eni_address` - (Optional, ForceNew) The eni address of the CLB.
* `load_balancer_billing_type` - (Optional, ForceNew) The billing type of the CLB, the value can be `PostPaid`.
* `load_balancer_name` - (Optional) The name of the CLB.
* `master_zone_id` - (Optional) The master zone ID of the CLB.
* `modification_protection_reason` - (Optional) The reason of the console modification protection.
* `modification_protection_status` - (Optional) The status of the console modification protection, the value can be `NonProtection` or `ConsoleProtection`.
* `project_name` - (Optional, ForceNew) The ProjectName of the CLB.
* `region_id` - (Optional, ForceNew) The region of the request.
* `slave_zone_id` - (Optional) The slave zone ID of the CLB.
* `tags` - (Optional) Tags.
* `vpc_id` - (Optional, ForceNew) The id of the VPC.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
CLB can be imported using the id, e.g.
```
$ terraform import volcengine_clb.default clb-273y2ok6ets007fap8txvf6us
```

