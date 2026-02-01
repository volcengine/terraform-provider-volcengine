---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_vpc_ipv6_gateways"
sidebar_current: "docs-volcengine-datasource-vpc_ipv6_gateways"
description: |-
  Use this data source to query detailed information of vpc ipv6 gateways
---
# volcengine_vpc_ipv6_gateways
Use this data source to query detailed information of vpc ipv6 gateways
## Example Usage
```hcl
data "volcengine_vpc_ipv6_gateways" "default" {
  ids = ["ipv6gw-12bcapllb5ukg17q7y2sd3thx"]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) The ID list of the Ipv6Gateways.
* `name_regex` - (Optional) A Name Regex of the Ipv6Gateway.
* `name` - (Optional) The name of the Ipv6Gateway.
* `output_file` - (Optional) File name where to save data source results.
* `vpc_ids` - (Optional) The ID list of the VPC which the Ipv6Gateway belongs to.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `ipv6_gateways` - The collection of Ipv6Gateway query.
    * `creation_time` - Creation time of the Ipv6Gateway.
    * `description` - The description of the Ipv6Gateway.
    * `id` - The ID of the Ipv6Gateway.
    * `ipv6_gateway_id` - The ID of the Ipv6Gateway.
    * `name` - The Name of the Ipv6Gateway.
    * `status` - The Status of the Ipv6Gateway.
    * `update_time` - Update time of the Ipv6Gateway.
    * `vpc_id` - The id of the VPC which the Ipv6Gateway belongs to.
* `total_count` - The total count of Ipv6Gateway query.


