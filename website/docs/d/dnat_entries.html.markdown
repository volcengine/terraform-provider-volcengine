---
subcategory: "NAT"
layout: "volcengine"
page_title: "Volcengine: volcengine_dnat_entries"
sidebar_current: "docs-volcengine-datasource-dnat_entries"
description: |-
  Use this data source to query detailed information of dnat entries
---
# volcengine_dnat_entries
Use this data source to query detailed information of dnat entries
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

data "volcengine_dnat_entries" "foo" {
  ids = [volcengine_dnat_entry.foo.id]
}
```
## Argument Reference
The following arguments are supported:
* `dnat_entry_name` - (Optional) The name of the DNAT entry.
* `external_ip` - (Optional) Provides the public IP address for public network access.
* `external_port` - (Optional) The port or port segment that receives requests from the public network. If InternalPort is passed into the port segment, ExternalPort must also be passed into the port segment.
* `ids` - (Optional) A list of DNAT entry ids.
* `internal_ip` - (Optional) Provides the internal IP address.
* `internal_port` - (Optional) The port or port segment on which the cloud server instance provides services to the public network.
* `nat_gateway_id` - (Optional) The id of the NAT gateway.
* `output_file` - (Optional) File name where to save data source results.
* `protocol` - (Optional) The network protocol.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `dnat_entries` - List of DNAT entries.
    * `dnat_entry_id` - The ID of the DNAT entry.
    * `dnat_entry_name` - The name of the DNAT entry.
    * `external_ip` - Provides the public IP address for public network access.
    * `external_port` - The port or port segment that receives requests from the public network. If InternalPort is passed into the port segment, ExternalPort must also be passed into the port segment.
    * `id` - The ID of the DNAT entry.
    * `internal_ip` - Provides the internal IP address.
    * `internal_port` - The port or port segment on which the cloud server instance provides services to the public network.
    * `nat_gateway_id` - The ID of the NAT gateway.
    * `protocol` - The network protocol.
    * `status` - The network status.
* `total_count` - The total count of snat entries query.


