---
subcategory: "NAT"
layout: "volcengine"
page_title: "Volcengine: volcengine_snat_entries"
sidebar_current: "docs-volcengine-datasource-snat_entries"
description: |-
  Use this data source to query detailed information of snat entries
---
# volcengine_snat_entries
Use this data source to query detailed information of snat entries
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

resource "volcengine_snat_entry" "foo1" {
  snat_entry_name = "acc-test-snat-entry"
  nat_gateway_id  = volcengine_nat_gateway.foo.id
  eip_id          = volcengine_eip_address.foo.id
  source_cidr     = "172.16.0.0/24"
  depends_on      = ["volcengine_eip_associate.foo"]
}

resource "volcengine_snat_entry" "foo2" {
  snat_entry_name = "acc-test-snat-entry"
  nat_gateway_id  = volcengine_nat_gateway.foo.id
  eip_id          = volcengine_eip_address.foo.id
  source_cidr     = "172.16.0.0/16"
  depends_on      = ["volcengine_eip_associate.foo"]
}

data "volcengine_snat_entries" "foo" {
  ids = [volcengine_snat_entry.foo1.id, volcengine_snat_entry.foo2.id]
}
```
## Argument Reference
The following arguments are supported:
* `eip_id` - (Optional) An id of the public ip address used by the SNAT entry.
* `ids` - (Optional) A list of SNAT entry ids.
* `nat_gateway_id` - (Optional) An id of the nat gateway to which the entry belongs.
* `output_file` - (Optional) File name where to save data source results.
* `snat_entry_name` - (Optional) A name of SNAT entry.
* `source_cidr` - (Optional) The SourceCidr of SNAT entry.
* `subnet_id` - (Optional) An id of the subnet that is required to access the Internet.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `snat_entries` - The collection of snat entries.
    * `eip_address` - The public ip address used by the SNAT entry.
    * `eip_id` - The id of the public ip address used by the SNAT entry.
    * `id` - The id of the SNAT entry.
    * `nat_gateway_id` - The id of the nat gateway to which the entry belongs.
    * `snat_entry_id` - The id of the SNAT entry.
    * `snat_entry_name` - The name of the SNAT entry.
    * `source_cidr` - The SourceCidr of the SNAT entry.
    * `status` - The status of the SNAT entry.
    * `subnet_id` - The id of the subnet that is required to access the internet.
* `total_count` - The total count of snat entries query.


