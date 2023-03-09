---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_vpc_ipv6_address_bandwidth"
sidebar_current: "docs-volcengine-resource-vpc_ipv6_address_bandwidth"
description: |-
  Provides a resource to manage vpc ipv6 address bandwidth
---
# volcengine_vpc_ipv6_address_bandwidth
Provides a resource to manage vpc ipv6 address bandwidth
## Example Usage
```hcl
data "volcengine_ecs_instances" "dataEcs" {
  ids = ["i-yca7nb3ozzl8izx5c64d"]
}

data "volcengine_vpc_ipv6_addresses" "dataIpv6" {
  associated_instance_id = data.volcengine_ecs_instances.dataEcs.instances.0.instance_id
}

resource "volcengine_vpc_ipv6_address_bandwidth" "foo" {
  ipv6_address = data.volcengine_vpc_ipv6_addresses.dataIpv6.ipv6_addresses.0.ipv6_address
  billing_type = 3
  bandwidth    = 5
}
```
## Argument Reference
The following arguments are supported:
* `billing_type` - (Required, ForceNew) BillingType of the Ipv6 bandwidth. Valid values: 3(Pay by Traffic).
* `ipv6_address` - (Required, ForceNew) Ipv6 address.
* `bandwidth` - (Optional) Peek bandwidth of the Ipv6 address. Valid values: 1 to 200. Unit: Mbit/s.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `allocation_id` - The ID of the Ipv6AddressBandwidth.
* `business_status` - The BusinessStatus of the Ipv6AddressBandwidth.
* `creation_time` - Creation time of the Ipv6AddressBandwidth.
* `delete_time` - Delete time of the Ipv6AddressBandwidth.
* `instance_id` - The ID of the associated instance.
* `instance_type` - The type of the associated instance.
* `isp` - The ISP of the Ipv6AddressBandwidth.
* `lock_reason` - The BusinessStatus of the Ipv6AddressBandwidth.
* `network_type` - The network type of the Ipv6AddressBandwidth.
* `overdue_time` - Overdue time of the Ipv6AddressBandwidth.
* `status` - The status of the Ipv6AddressBandwidth.
* `update_time` - Update time of the Ipv6AddressBandwidth.


## Import
Ipv6AddressBandwidth can be imported using the id, e.g.
```
$ terraform import volcengine_vpc_ipv6_address_bandwidth.default eip-2fede9fsgnr4059gp674m6ney
```

