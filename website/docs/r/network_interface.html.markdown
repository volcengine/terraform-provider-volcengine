---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_network_interface"
sidebar_current: "docs-volcengine-resource-network_interface"
description: |-
  Provides a resource to manage network interface
---
# volcengine_network_interface
Provides a resource to manage network interface
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

resource "volcengine_security_group" "foo" {
  security_group_name = "acc-test-sg"
  vpc_id              = volcengine_vpc.foo.id
}

resource "volcengine_network_interface" "foo" {
  network_interface_name = "acc-test-eni"
  description            = "acc-test"
  subnet_id              = volcengine_subnet.foo.id
  security_group_ids     = [volcengine_security_group.foo.id]
  primary_ip_address     = "172.16.0.253"
  port_security_enabled  = false
  private_ip_address     = ["172.16.0.2"]
  project_name           = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `security_group_ids` - (Required) The list of the security group id to which the secondary ENI belongs.
* `subnet_id` - (Required, ForceNew) The id of the subnet to which the ENI is connected.
* `description` - (Optional) The description of the ENI.
* `ipv6_address_count` - (Optional) The number of IPv6 addresses to be automatically assigned from within the CIDR block of the subnet that hosts the ENI. Valid values: 0 to 10.
 You cannot specify both the ipv6_addresses and ipv6_address_count parameters.
* `ipv6_addresses` - (Optional) One or more IPv6 addresses selected from within the CIDR block of the subnet that hosts the ENI. Support up to 10.
 You cannot specify both the ipv6_addresses and ipv6_address_count parameters.
* `network_interface_name` - (Optional) The name of the ENI.
* `port_security_enabled` - (Optional) Set port security enable or disable.
* `primary_ip_address` - (Optional, ForceNew) The primary IP address of the ENI.
* `private_ip_address` - (Optional) The list of private ip address. This field conflicts with `secondary_private_ip_address_count`.
* `project_name` - (Optional) The ProjectName of the ENI.
* `secondary_private_ip_address_count` - (Optional) The count of secondary private ip address. This field conflicts with `private_ip_address`.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `status` - The status of the ENI.


## Import
Network interface can be imported using the id, e.g.
```
$ terraform import volcengine_network_interface.default eni-bp1fgnh68xyz9****
```

