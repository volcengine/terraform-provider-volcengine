data "volcengine_alb_zones" "foo"{
}

resource "volcengine_vpc" "vpc_ipv6" {
  vpc_name = "acc-test-vpc-ipv6"
  cidr_block = "172.16.0.0/16"
  enable_ipv6 = true
}

resource "volcengine_subnet" "subnet_ipv6_1" {
  subnet_name = "acc-test-subnet-ipv6-1"
  cidr_block = "172.16.1.0/24"
  zone_id = data.volcengine_alb_zones.foo.zones[0].id
  vpc_id = volcengine_vpc.vpc_ipv6.id
  ipv6_cidr_block = 1
}

resource "volcengine_subnet" "subnet_ipv6_2" {
  subnet_name = "acc-test-subnet-ipv6-2"
  cidr_block = "172.16.2.0/24"
  zone_id = data.volcengine_alb_zones.foo.zones[1].id
  vpc_id = volcengine_vpc.vpc_ipv6.id
  ipv6_cidr_block = 2
}

resource "volcengine_vpc_ipv6_gateway" "ipv6_gateway" {
  vpc_id = volcengine_vpc.vpc_ipv6.id
  name = "acc-test-ipv6-gateway"
}

resource "volcengine_alb" "alb-private" {
  address_ip_version = "IPv4"
  type = "private"
  load_balancer_name = "acc-test-alb-private"
  description = "acc-test"
  subnet_ids = [volcengine_subnet.subnet_ipv6_1.id, volcengine_subnet.subnet_ipv6_2.id]
  project_name = "default"
  delete_protection = "off"
  tags {
    key = "k1"
    value = "v1"
  }
}

resource "volcengine_alb" "alb-public" {
  address_ip_version = "DualStack"
  type = "public"
  load_balancer_name = "acc-test-alb-public"
  description = "acc-test"
  subnet_ids = [volcengine_subnet.subnet_ipv6_1.id, volcengine_subnet.subnet_ipv6_2.id]
  project_name = "default"
  delete_protection = "off"
  modification_protection_status = "NonProtection"
  modification_protection_reason = "Test modification protection"
  load_balancer_edition = "Basic"

  eip_billing_config {
    isp = "BGP"
    eip_billing_type = "PostPaidByBandwidth"
    bandwidth = 1
  }
  ipv6_eip_billing_config {
    isp = "BGP"
    billing_type = "PostPaidByBandwidth"
    bandwidth = 1
  }

  tags {
    key = "k1"
    value = "v1"
  }
  depends_on = [volcengine_vpc_ipv6_gateway.ipv6_gateway]
}

# CLone ALB instance
resource "volcengine_alb" "alb-cloned" {
  source_load_balancer_id = volcengine_alb.alb-private.id
  load_balancer_name = "acc-test-alb-cloned"
  description = "cloned from alb-private"
  subnet_ids = [volcengine_subnet.subnet_ipv6_1.id]
  type = "private"
  project_name = "default"
}

# Example of ALB network type change, private -> public
resource "volcengine_alb" "alb-type-change" {
  load_balancer_name = "acc-test-alb-type-change"
  description = "will change to public type"
  subnet_ids = [volcengine_subnet.subnet_ipv6_1.id, volcengine_subnet.subnet_ipv6_2.id]
  type = "public"
  project_name = "default"
  allocation_ids = ["eip-iinpy4k1rytc74o8curgocd7", "eip-iinpy4k1rytc74o8curgocd8"]
}
