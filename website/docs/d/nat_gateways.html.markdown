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
data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_nat_gateway" "foo" {
  vpc_id           = volcengine_vpc.foo.id
  subnet_id        = volcengine_subnet.foo.id
  spec             = "Small"
  nat_gateway_name = "acc-test-ng-${count.index}"
  description      = "acc-test"
  billing_type     = "PostPaid"
  project_name     = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
  count = 3
}

data "volcengine_nat_gateways" "foo" {
  ids = volcengine_nat_gateway.foo[*].id
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
* `tags` - (Optional) Tags.
* `vpc_id` - (Optional) The id of the VPC.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

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
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `updated_at` - The update time of the NatGateway.
    * `vpc_id` - The ID of the VPC.
* `total_count` - The total count of NatGateway query.


