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

resource "volcengine_clb" "foo" {
  type = "public"
  subnet_id = volcengine_subnet.foo.id
  load_balancer_spec = "small_1"
  description = "acc-test-demo"
  load_balancer_name = "acc-test-clb-${count.index}"
  load_balancer_billing_type = "PostPaid"
  eip_billing_config {
    isp = "BGP"
    eip_billing_type = "PostPaidByBandwidth"
    bandwidth = 1
  }
  tags {
    key = "k1"
    value = "v1"
  }
  count = 3
}

data "volcengine_clbs" "foo"{
  ids = volcengine_clb.foo[*].id
}