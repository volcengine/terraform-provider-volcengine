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
  load_balancer_name = "acc-test-clb"
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
}

resource "volcengine_clb" "public_clb" {
  type = "public"
  subnet_id = volcengine_subnet.foo.id
  load_balancer_name = "acc-test-clb-public"
  load_balancer_spec = "small_1"
  description = "acc-test-demo"
  project_name = "default"
  eip_billing_config {
    isp = "BGP"
    eip_billing_type = "PostPaidByBandwidth"
    bandwidth = 1
  }
  tags {
    key = "k1"
    value = "v1"
  }
}

resource "volcengine_clb" "private_clb" {
  type = "private"
  subnet_id = volcengine_subnet.foo.id
  load_balancer_name = "acc-test-clb-private"
  load_balancer_spec = "small_1"
  description = "acc-test-demo"
  project_name = "default"
}

resource "volcengine_eip_address" "eip" {
  billing_type = "PostPaidByBandwidth"
  bandwidth = 1
  isp = "BGP"
  name = "tf-eip"
  description = "tf-test"
  project_name = "default"
}

resource "volcengine_eip_associate" "associate" {
  allocation_id = volcengine_eip_address.eip.id
  instance_id = volcengine_clb.private_clb.id
  instance_type = "ClbInstance"
}