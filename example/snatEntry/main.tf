data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block = "172.16.0.0/24"
  zone_id = data.volcengine_zones.foo.zones[0].id
  vpc_id = volcengine_vpc.foo.id
}

resource "volcengine_nat_gateway" "foo" {
  vpc_id = volcengine_vpc.foo.id
  subnet_id = volcengine_subnet.foo.id
  spec = "Small"
  nat_gateway_name = "acc-test-ng"
  description = "acc-test"
  billing_type = "PostPaid"
  project_name = "default"
  tags {
    key = "k1"
    value = "v1"
  }
}

resource "volcengine_eip_address" "foo" {
  name = "acc-test-eip"
  description = "acc-test"
  bandwidth = 1
  billing_type = "PostPaidByBandwidth"
  isp = "BGP"
}

resource "volcengine_eip_associate" "foo" {
  allocation_id = volcengine_eip_address.foo.id
  instance_id = volcengine_nat_gateway.foo.id
  instance_type = "Nat"
}

resource "volcengine_snat_entry" "foo" {
  snat_entry_name = "acc-test-snat-entry"
  nat_gateway_id = volcengine_nat_gateway.foo.id
  eip_id = volcengine_eip_address.foo.id
  source_cidr = "172.16.0.0/24"
  depends_on = [volcengine_eip_associate.foo]
}
