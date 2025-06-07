data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc-rn"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet-rn"
  cidr_block = "172.16.0.0/24"
  zone_id = data.volcengine_zones.foo.zones[0].id
  vpc_id = volcengine_vpc.foo.id
}

resource "volcengine_nat_gateway" "foo" {
  vpc_id = volcengine_vpc.foo.id
  subnet_id = volcengine_subnet.foo.id
  spec = "Small"
  nat_gateway_name = "acc-test-nat-rn"
}

resource "volcengine_route_table" "foo" {
  vpc_id = volcengine_vpc.foo.id
  route_table_name = "acc-test-route-table"
}

resource "volcengine_route_entry" "foo" {
  route_table_id = volcengine_route_table.foo.id
  destination_cidr_block = "172.16.1.0/24"
  next_hop_type = "NatGW"
  next_hop_id = volcengine_nat_gateway.foo.id
  route_entry_name = "acc-test-route-entry"
}
