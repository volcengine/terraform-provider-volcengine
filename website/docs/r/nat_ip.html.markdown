---
subcategory: "NAT"
layout: "volcengine"
page_title: "Volcengine: volcengine_nat_ip"
sidebar_current: "docs-volcengine-resource-nat_ip"
description: |-
  Provides a resource to manage nat ip
---
# volcengine_nat_ip
Provides a resource to manage nat ip
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

resource "volcengine_nat_gateway" "intranet_nat_gateway" {
  vpc_id           = volcengine_vpc.foo.id
  subnet_id        = volcengine_subnet.foo.id
  nat_gateway_name = "acc-test-intranet_ng"
  description      = "acc-test"
  network_type     = "intranet"
  billing_type     = "PostPaidByUsage"
  project_name     = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_nat_ip" "foo" {
  nat_gateway_id     = volcengine_nat_gateway.intranet_nat_gateway.id
  nat_ip_name        = "acc-test-nat-ip"
  nat_ip_description = "acc-test"
  nat_ip             = "172.16.0.3"
}
```
## Argument Reference
The following arguments are supported:
* `nat_gateway_id` - (Required, ForceNew) The id of the nat gateway to which the Nat Ip belongs.
* `nat_ip_description` - (Optional) The description of the Nat Ip.
* `nat_ip_name` - (Optional) The name of the Nat Ip.
* `nat_ip` - (Optional, ForceNew) The ip address of the Nat Ip.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `is_default` - Whether the Ip is the default Nat Ip.
* `status` - The status of the Nat Ip.
* `using_status` - The using status of the Nat Ip.


## Import
NatIp can be imported using the id, e.g.
```
$ terraform import volcengine_nat_ip.default resource_id
```

