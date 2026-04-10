---
subcategory: "VPN"
layout: "volcengine"
page_title: "Volcengine: volcengine_customer_gateways"
sidebar_current: "docs-volcengine-datasource-customer_gateways"
description: |-
  Use this data source to query detailed information of customer gateways
---
**❗Notice:**
The current provider is no longer being maintained. We recommend that you use the [volcenginecc](https://registry.terraform.io/providers/volcengine/volcenginecc/latest/docs) instead.
# volcengine_customer_gateways
Use this data source to query detailed information of customer gateways
## Example Usage
```hcl
resource "volcengine_customer_gateway" "foo" {
  ip_address            = "192.0.1.3"
  customer_gateway_name = "acc-test"
  description           = "acc-test"
  project_name          = "default"
}
data "volcengine_customer_gateways" "foo" {
  ids = [volcengine_customer_gateway.foo.id]
}
```
## Argument Reference
The following arguments are supported:
* `customer_gateway_names` - (Optional) A list of customer gateway names.
* `ids` - (Optional) A list of customer gateway ids.
* `ip_address` - (Optional) A IP address of the customer gateway.
* `ip_version` - (Optional) The IP version of the customer gateway. Valid value: ipv4, ipv6.
* `name_regex` - (Optional) A Name Regex of customer gateway.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name of the VPN customer gateway.
* `status` - (Optional) The status of the customer gateway. Valid value: Creating, Deleting, Pending, Available.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

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
    * `ip_version` - The IP version of the customer gateway.
    * `project_name` - The project name of the VPN customer gateway.
    * `status` - The status of the customer gateway.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `update_time` - The update time of customer gateway.
* `total_count` - The total count of customer gateway query.


