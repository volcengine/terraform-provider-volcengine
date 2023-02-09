---
subcategory: "VPN(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_vpn_gateways"
sidebar_current: "docs-volcengine-datasource-vpn_gateways"
description: |-
  Use this data source to query detailed information of vpn gateways
---
# volcengine_vpn_gateways
Use this data source to query detailed information of vpn gateways
## Example Usage
```hcl
data "volcengine_vpn_gateways" "default" {
  ids = ["vgw-2c012ea9fm5mo2dx0efxg46qi"]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of VPN gateway ids.
* `ip_address` - (Optional) A IP address of the VPN gateway.
* `name_regex` - (Optional) A Name Regex of VPN gateway.
* `output_file` - (Optional) File name where to save data source results.
* `subnet_id` - (Optional) A subnet ID of the VPN gateway.
* `tags` - (Optional) Tags.
* `vpc_id` - (Optional) A VPC ID of the VPN gateway.
* `vpn_gateway_names` - (Optional) A list of VPN gateway names.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of VPN gateway query.
* `vpn_gateways` - The collection of VPN gateway query.
    * `account_id` - The account ID of the VPN gateway.
    * `bandwidth` - The bandwidth of the VPN gateway.
    * `billing_type` - The BillingType of the VPN gateway.
    * `business_status` - The business status of the VPN gateway.
    * `connection_count` - The connection count of the VPN gateway.
    * `creation_time` - The create time of VPN gateway.
    * `deleted_time` - The deleted time of the VPN gateway.
    * `description` - The description of the VPN gateway.
    * `expired_time` - The expired time of the VPN gateway.
    * `id` - The ID of the VPN gateway.
    * `ip_address` - The IP address of the VPN gateway.
    * `lock_reason` - The lock reason of the VPN gateway.
    * `route_count` - The route count of the VPN gateway.
    * `status` - The status of the VPN gateway.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `update_time` - The update time of VPN gateway.
    * `vpc_id` - The VPC ID of the VPN gateway.
    * `vpn_gateway_id` - The ID of the VPN gateway.
    * `vpn_gateway_name` - The name of the VPN gateway.


