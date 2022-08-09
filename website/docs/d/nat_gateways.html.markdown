---
subcategory: "NAT"
layout: "volcengine"
page_title: "Volcengine: volcengine_nat_gateways"
sidebar_current: "docs-volcengine-datasource-nat_gateways"
description: |-
  Use this data source to query detailed information of nat gateways
---
# volcengine_nat_gateways
Use this data source to query detailed information of nat gateways
## Example Usage
```hcl
data "volcengine_nat_gateways" "default" {
  ids = ["ngw-2743w1f6iqby87fap8tvm9kop", "ngw-274gwbqe340zk7fap8spkzo7x"]
}
```
## Argument Reference
The following arguments are supported:
* `description` - (Optional) The description of the NatGateway.
* `ids` - (Optional) The list of NatGateway IDs.
* `name_regex` - (Optional) The Name Regex of NatGateway.
* `nat_gateway_name` - (Optional) The name of the NatGateway.
* `output_file` - (Optional) File name where to save data source results.
* `spec` - (Optional) The specification of the NatGateway.
* `subnet_id` - (Optional) The id of the Subnet.
* `vpc_id` - (Optional) The id of the VPC.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `nat_gateways` - The collection of NatGateway query.
    * `billing_type` - The billing type of the NatGateway.
    * `business_status` - Whether the NatGateway is locked.
    * `creation_time` - The creation time of the NatGateway.
    * `deleted_time` - The deleted time of the NatGateway.
    * `description` - The description of the NatGateway.
    * `eip_addresses` - The eip addresses of the NatGateway.
        * `allocation_id` - The ID of Eip.
        * `eip_address` - The address of Eip.
        * `using_status` - The using status of Eip.
    * `id` - The ID of the NatGateway.
    * `lock_reason` - The reason why locking NatGateway.
    * `nat_gateway_id` - The ID of the NatGateway.
    * `nat_gateway_name` - The name of the NatGateway.
    * `network_interface_id` - The ID of the network interface.
    * `overdue_time` - The overdue time of the NatGateway.
    * `spec` - The specification of the NatGateway.
    * `status` - The status of the NatGateway.
    * `subnet_id` - The ID of the Subnet.
    * `updated_at` - The update time of the NatGateway.
    * `vpc_id` - The ID of the VPC.
* `total_count` - The total count of NatGateway query.


