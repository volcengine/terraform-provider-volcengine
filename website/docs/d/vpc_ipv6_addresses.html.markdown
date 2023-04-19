---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_vpc_ipv6_addresses"
sidebar_current: "docs-volcengine-datasource-vpc_ipv6_addresses"
description: |-
  Use this data source to query detailed information of vpc ipv6 addresses
---
# volcengine_vpc_ipv6_addresses
Use this data source to query detailed information of vpc ipv6 addresses
## Example Usage
```hcl
data "volcengine_vpc_ipv6_addresses" "default" {
  associated_instance_id = "i-yca53yuhj6gh9zl53kav"
}
```
## Argument Reference
The following arguments are supported:
* `associated_instance_id` - (Optional) The ID of the ECS instance that is assigned the IPv6 address.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `ipv6_addresses` - The collection of Ipv6Address query.
    * `ipv6_address` - The IPv6 address.
* `total_count` - The total count of Ipv6Address query.


