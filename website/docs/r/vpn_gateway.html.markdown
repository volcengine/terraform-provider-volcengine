---
subcategory: "VPN"
layout: "volcengine"
page_title: "Volcengine: volcengine_vpn_gateway"
sidebar_current: "docs-volcengine-resource-vpn_gateway"
description: |-
  Provides a resource to manage vpn gateway
---
# volcengine_vpn_gateway
Provides a resource to manage vpn gateway
## Notice
When Destroy this resource,If the resource charge type is PrePaid,Please unsubscribe the resource 
in  [Volcengine Console](https://console.volcengine.com/finance/unsubscribe/),when complete console operation,yon can
use 'terraform state rm ${resourceId}' to remove.
## Example Usage
```hcl
resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = "cn-beijing-a"
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_vpn_gateway" "foo" {
  vpc_id           = volcengine_vpc.foo.id
  subnet_id        = volcengine_subnet.foo.id
  bandwidth        = 50
  vpn_gateway_name = "acc-test1"
  description      = "acc-test1"
  period           = 7
  project_name     = "default"
}
```
## Argument Reference
The following arguments are supported:
* `bandwidth` - (Required) The bandwidth of the VPN gateway. Unit: Mbps. Values: 5, 10, 20, 50, 100, 200, 500.
* `subnet_id` - (Required, ForceNew) The ID of the subnet where you want to create the VPN gateway.
* `vpc_id` - (Required, ForceNew) The ID of the VPC where you want to create the VPN gateway.
* `billing_type` - (Optional, ForceNew) The BillingType of the VPN gateway. Valid values: `PrePaid`, `PostPaid`.
* `description` - (Optional) The description of the VPN gateway.
* `ipsec_enabled` - (Optional) Whether ipsec is enabled. Default is true.
* `period` - (Optional) The Period of the VPN gateway. Default value is 12. This parameter is only useful when creating vpn gateway. Default period unit is Month.
Value range: 1~9, 12, 24, 36. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `project_name` - (Optional) The project name of the VPN gateway.
* `ssl_enabled` - (Optional) Whether ssl is enabled. Default is false.
* `ssl_max_connections` - (Optional) The max connections of ssl. This parameter can only be passed in when ssl_enabled is true. Default is 5.
* `tags` - (Optional) Tags.
* `vpn_gateway_name` - (Optional) The name of the VPN gateway.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `account_id` - The account ID of the VPN gateway.
* `business_status` - The business status of the VPN gateway.
* `connection_count` - The connection count of the VPN gateway.
* `creation_time` - The create time of VPN gateway.
* `deleted_time` - The deleted time of the VPN gateway.
* `expired_time` - The expired time of the VPN gateway.
* `ip_address` - The IP address of the VPN gateway.
* `lock_reason` - The lock reason of the VPN gateway.
* `renew_type` - The renew type of the VPN gateway.
* `route_count` - The route count of the VPN gateway.
* `status` - The status of the VPN gateway.
* `update_time` - The update time of VPN gateway.
* `vpn_gateway_id` - The ID of the VPN gateway.


## Import
VpnGateway can be imported using the id, e.g.
```
$ terraform import volcengine_vpn_gateway.default vgw-273zkshb2qayo7fap8t2****
```

