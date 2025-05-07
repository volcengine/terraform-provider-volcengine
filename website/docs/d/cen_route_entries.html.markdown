---
subcategory: "CEN"
layout: "volcengine"
page_title: "Volcengine: volcengine_cen_route_entries"
sidebar_current: "docs-volcengine-datasource-cen_route_entries"
description: |-
  Use this data source to query detailed information of cen route entries
---
# volcengine_cen_route_entries
Use this data source to query detailed information of cen route entries
## Example Usage
```hcl
data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc-rn"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet-rn"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_cen_attach_instance.foo.instance_id
}

resource "volcengine_nat_gateway" "foo" {
  vpc_id           = volcengine_vpc.foo.id
  subnet_id        = volcengine_subnet.foo.id
  spec             = "Small"
  nat_gateway_name = "acc-test-nat-rn"
}

resource "volcengine_route_entry" "foo" {
  route_table_id         = volcengine_vpc.foo.route_table_ids[0]
  destination_cidr_block = "172.16.1.0/24"
  next_hop_type          = "NatGW"
  next_hop_id            = volcengine_nat_gateway.foo.id
  route_entry_name       = "acc-test-route-entry"
}

resource "volcengine_cen" "foo" {
  cen_name     = "acc-test-cen"
  description  = "acc-test"
  project_name = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_cen_attach_instance" "foo" {
  cen_id             = volcengine_cen.foo.id
  instance_id        = volcengine_vpc.foo.id
  instance_region_id = "cn-beijing"
  instance_type      = "VPC"
}

resource "volcengine_cen_route_entry" "foo" {
  cen_id                 = volcengine_cen.foo.id
  destination_cidr_block = volcengine_route_entry.foo.destination_cidr_block
  instance_type          = "VPC"
  instance_region_id     = "cn-beijing"
  instance_id            = volcengine_cen_attach_instance.foo.instance_id
}

data "volcengine_cen_route_entries" "foo" {
  cen_id                 = volcengine_cen.foo.id
  destination_cidr_block = volcengine_cen_route_entry.foo.destination_cidr_block
}
```
## Argument Reference
The following arguments are supported:
* `cen_id` - (Required) A cen ID.
* `destination_cidr_block` - (Optional) A destination cidr block.
* `instance_id` - (Optional) An instance ID.
* `instance_region_id` - (Optional) An instance region ID.
* `instance_type` - (Optional) An instance type.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `cen_route_entries` - The collection of cen route entry query.
    * `as_path` - The AS path of the cen route entry.
    * `cen_id` - The cen ID of the cen route entry.
    * `destination_cidr_block` - The destination cidr block of the cen route entry.
    * `instance_id` - The instance id of the next hop of the cen route entry.
    * `instance_region_id` - The instance region id of the next hop of the cen route entry.
    * `instance_type` - The instance type of the next hop of the cen route entry.
    * `publish_status` - The publish status of the cen route entry.
    * `status` - The status of the cen route entry.
* `total_count` - The total count of cen route entry.


