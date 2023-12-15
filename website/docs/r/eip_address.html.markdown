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
## Notice
When Destroy this resource,If the resource charge type is PrePaid,Please unsubscribe the resource 
in  [Volcengine Console](https://console.volcengine.com/finance/unsubscribe/),when complete console operation,yon can
use 'terraform state rm ${resourceId}' to remove.
## Example Usage
```hcl
resource "volcengine_eip_address" "foo" {
  billing_type = "PostPaidByBandwidth"
  bandwidth    = 1
  isp          = "ChinaUnicom"
  name         = "acc-eip"
  description  = "acc-test"
  project_name = "default"
}
```
## Argument Reference
The following arguments are supported:
* `billing_type` - (Required) The billing type of the EIP Address. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic` or `PrePaid`.
* `bandwidth` - (Optional) The peek bandwidth of the EIP.
* `description` - (Optional) The description of the EIP.
* `isp` - (Optional, ForceNew) The ISP of the EIP, the value can be `BGP` or `ChinaMobile` or `ChinaUnicom` or `ChinaTelecom` or `SingleLine_BGP` or `Static_BGP`.
* `name` - (Optional) The name of the EIP Address.
* `period` - (Optional) The period of the EIP Address, the valid value range in 1~9 or 12 or 36. Default value is 12. The period unit defaults to `Month`.This field is only effective when creating a PrePaid Eip or changing the billing_type from PostPaid to PrePaid.
* `project_name` - (Optional) The ProjectName of the EIP.
* `security_protection_types` - (Optional, ForceNew) Security protection types for public IP addresses. Parameter - N: Indicates the number of security protection types, currently only supports taking 1. Value: `AntiDDoS_Enhanced` or left blank.If the value is `AntiDDoS_Enhanced`, then will create an eip with enhanced protection,(can be added to DDoS native protection (enterprise version) instance). If left blank, it indicates an eip with basic protection.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `deleted_time` - The deleted time of the EIP.
* `eip_address` - The ip address of the EIP.
* `expired_time` - The expired time of the EIP.
* `overdue_time` - The overdue time of the EIP.
* `status` - The status of the EIP.


## Import
Eip address can be imported using the id, e.g.
```
$ terraform import volcengine_eip_address.default eip-274oj9a8rs9a87fap8sf9515b
```

