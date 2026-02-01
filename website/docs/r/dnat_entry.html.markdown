---
subcategory: "NAT"
layout: "volcengine"
page_title: "Volcengine: volcengine_dnat_entry"
sidebar_current: "docs-volcengine-resource-dnat_entry"
description: |-
  Provides a resource to manage dnat entry
---
# volcengine_dnat_entry
Provides a resource to manage dnat entry
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
  nat_gateway_name = "acc-test-ng"
  description      = "acc-test"
  billing_type     = "PostPaid"
  project_name     = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_eip_address" "foo" {
  name         = "acc-test-eip"
  description  = "acc-test"
  bandwidth    = 1
  billing_type = "PostPaidByBandwidth"
  isp          = "BGP"
}

resource "volcengine_eip_associate" "foo" {
  allocation_id = volcengine_eip_address.foo.id
  instance_id   = volcengine_nat_gateway.foo.id
  instance_type = "Nat"
}

resource "volcengine_dnat_entry" "foo" {
  dnat_entry_name = "acc-test-dnat-entry"
  external_ip     = volcengine_eip_address.foo.eip_address
  external_port   = 80
  internal_ip     = "172.16.0.10"
  internal_port   = 80
  nat_gateway_id  = volcengine_nat_gateway.foo.id
  protocol        = "tcp"
  depends_on      = [volcengine_eip_associate.foo]
}
```
## Argument Reference
The following arguments are supported:
* `external_ip` - (Required) Provides the public IP address for public network access.
* `external_port` - (Required) The port or port segment that receives requests from the public network. If InternalPort is passed into the port segment, ExternalPort must also be passed into the port segment.
* `internal_ip` - (Required) Provides the internal IP address.
* `internal_port` - (Required) The port or port segment on which the cloud server instance provides services to the public network.
* `nat_gateway_id` - (Required, ForceNew) The id of the nat gateway to which the entry belongs.
* `protocol` - (Required) The network protocol. Valid values: `tcp`, `udp`.
* `dnat_entry_name` - (Optional) The name of the DNAT rule.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `dnat_entry_id` - The id of the DNAT rule.


## Import
Dnat entry can be imported using the id, e.g.
```
$ terraform import volcengine_dnat_entry.default dnat-3fvhk47kf56****
```

