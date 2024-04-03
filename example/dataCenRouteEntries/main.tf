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
  route_table_id         = tolist(volcengine_vpc.foo.route_table_ids)[0]
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
