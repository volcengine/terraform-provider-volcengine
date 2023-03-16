---
subcategory: "EIP"
layout: "volcengine"
page_title: "Volcengine: volcengine_eip_address"
sidebar_current: "docs-volcengine-resource-eip_address"
description: |-
  Provides a resource to manage eip address
---
# volcengine_eip_address
Provides a resource to manage eip address
## Example Usage
```hcl
resource "volcengine_eip_address" "foo" {
  billing_type = "PostPaidByBandwidth"
  bandwidth    = 1
  isp          = "BGP"
  name         = "tf-test"
  description  = "tf-test"
}
```
## Argument Reference
The following arguments are supported:
* `billing_type` - (Required) The billing type of the EIP Address. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic`.
* `bandwidth` - (Optional) The peek bandwidth of the EIP, the value range in 1~500 for PostPaidByBandwidth, and 1~200 for PostPaidByTraffic.
* `description` - (Optional) The description of the EIP.
* `isp` - (Optional, ForceNew) The ISP of the EIP, the value can be `BGP` or `ChinaMobile` or `ChinaUnicom` or `ChinaTelecom`.
* `name` - (Optional) The name of the EIP Address.
* `project_name` - (Optional) The ProjectName of the EIP.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `eip_address` - The ip address of the EIP.
* `status` - The status of the EIP.


## Import
Eip address can be imported using the id, e.g.
```
$ terraform import volcengine_eip_address.default eip-274oj9a8rs9a87fap8sf9515b
```

