---
subcategory: "VPN(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_customer_gateways"
sidebar_current: "docs-volcengine-datasource-customer_gateways"
description: |-
  Use this data source to query detailed information of customer gateways
---
# volcengine_customer_gateways
Use this data source to query detailed information of customer gateways
## Example Usage
```hcl
data "volcengine_customer_gateways" "foo" {
  ids = ["cgw-2d68c4zglycjk58ozfe96norh"]
}
```
## Argument Reference
The following arguments are supported:
* `customer_gateway_names` - (Optional) A list of customer gateway names.
* `ids` - (Optional) A list of customer gateway ids.
* `ip_address` - (Optional) A IP address of the customer gateway.
* `name_regex` - (Optional) A Name Regex of customer gateway.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `customer_gateways` - The collection of customer gateway query.
    * `account_id` - The account ID of the customer gateway.
    * `connection_count` - The connection count of the customer gateway.
    * `creation_time` - The create time of customer gateway.
    * `customer_gateway_id` - The ID of the customer gateway.
    * `customer_gateway_name` - The name of the customer gateway.
    * `description` - The description of the customer gateway.
    * `id` - The ID of the customer gateway.
    * `ip_address` - The IP address of the customer gateway.
    * `status` - The status of the customer gateway.
    * `update_time` - The update time of customer gateway.
* `total_count` - The total count of customer gateway query.


