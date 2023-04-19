---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_vpc_ipv6_address_bandwidths"
sidebar_current: "docs-volcengine-datasource-vpc_ipv6_address_bandwidths"
description: |-
  Use this data source to query detailed information of vpc ipv6 address bandwidths
---
# volcengine_vpc_ipv6_address_bandwidths
Use this data source to query detailed information of vpc ipv6 address bandwidths
## Example Usage
```hcl
data "volcengine_vpc_ipv6_address_bandwidths" "default" {
  ids = ["eip-in2y2duvtlhc8gbssyfnhfre"]
}
```
## Argument Reference
The following arguments are supported:
* `associated_instance_id` - (Optional) The ID of the associated instance.
* `associated_instance_type` - (Optional) The type of the associated instance.
* `ids` - (Optional) Allocation IDs of the Ipv6 address width.
* `ipv6_addresses` - (Optional) The ipv6 addresses.
* `isp` - (Optional) ISP of the ipv6 address.
* `network_type` - (Optional) The network type of the ipv6 address.
* `output_file` - (Optional) File name where to save data source results.
* `vpc_id` - (Optional) The ID of Vpc the ipv6 address in.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `ipv6_address_bandwidths` - The collection of Ipv6AddressBandwidth query.
    * `allocation_id` - The ID of the Ipv6AddressBandwidth.
    * `bandwidth` - Peek bandwidth of the Ipv6 address.
    * `billing_type` - BillingType of the Ipv6 bandwidth.
    * `business_status` - The BusinessStatus of the Ipv6AddressBandwidth.
    * `creation_time` - Creation time of the Ipv6AddressBandwidth.
    * `delete_time` - Delete time of the Ipv6AddressBandwidth.
    * `id` - The ID of the Ipv6AddressBandwidth.
    * `instance_id` - The ID of the associated instance.
    * `instance_type` - The type of the associated instance.
    * `ipv6_address` - The IPv6 address.
    * `isp` - The ISP of the Ipv6AddressBandwidth.
    * `lock_reason` - The BusinessStatus of the Ipv6AddressBandwidth.
    * `network_type` - The network type of the Ipv6AddressBandwidth.
    * `overdue_time` - Overdue time of the Ipv6AddressBandwidth.
    * `status` - The status of the Ipv6AddressBandwidth.
    * `update_time` - Update time of the Ipv6AddressBandwidth.
* `total_count` - The total count of Ipv6AddressBandwidth query.


