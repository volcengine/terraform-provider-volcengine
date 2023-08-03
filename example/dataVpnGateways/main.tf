resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block = "172.16.0.0/24"
  zone_id = "cn-beijing-a"
  vpc_id = volcengine_vpc.foo.id
}

resource "volcengine_vpn_gateway" "foo" {
  vpc_id = volcengine_vpc.foo.id
  subnet_id = volcengine_subnet.foo.id
  bandwidth = 20
  vpn_gateway_name = "acc-test"
  description = "acc-test"
  period = 2
  project_name = "default"
}

data "volcengine_vpn_gateways" "foo"{
  ids = [volcengine_vpn_gateway.foo.id]
}