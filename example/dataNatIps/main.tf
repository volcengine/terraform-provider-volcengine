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

data "volcengine_nat_ips" "foo" {
  nat_gateway_id = volcengine_nat_gateway.intranet_nat_gateway.id
}
