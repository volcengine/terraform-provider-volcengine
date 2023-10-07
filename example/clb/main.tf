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

# ipv4 public clb
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

# ipv4 private clb
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

# ipv6 private clb
resource "volcengine_vpc" "vpc_ipv6" {
  vpc_name = "acc-test-vpc-ipv6"
  cidr_block = "172.16.0.0/16"
  enable_ipv6 = true
}

resource "volcengine_subnet" "subnet_ipv6" {
  subnet_name = "acc-test-subnet-ipv6"
  cidr_block = "172.16.0.0/24"
  zone_id = data.volcengine_zones.foo.zones[1].id
  vpc_id = volcengine_vpc.vpc_ipv6.id
  ipv6_cidr_block = 1
}

resource "volcengine_clb" "private_clb_ipv6" {
  type = "private"
  subnet_id = volcengine_subnet.subnet_ipv6.id
  load_balancer_name = "acc-test-clb-ipv6"
  load_balancer_spec = "small_1"
  description = "acc-test-demo"
  project_name = "default"
  address_ip_version = "DualStack"
}

resource "volcengine_vpc_ipv6_gateway" "ipv6_gateway" {
  vpc_id = volcengine_vpc.vpc_ipv6.id
  name = "acc-test-ipv6-gateway"
}

resource "volcengine_vpc_ipv6_address_bandwidth" "foo" {
  ipv6_address = volcengine_clb.private_clb_ipv6.eni_ipv6_address
  billing_type = "PostPaidByBandwidth"
  bandwidth = 5
  depends_on = [volcengine_vpc_ipv6_gateway.ipv6_gateway]
}
