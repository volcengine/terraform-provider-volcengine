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

resource "volcengine_ha_vip" "foo" {
  ha_vip_name = "acc-test-ha-vip"
  description = "acc-test"
  subnet_id = volcengine_subnet.foo.id
  #  ip_address = "172.16.0.5"
}

resource "volcengine_eip_address" "foo" {
  billing_type = "PostPaidByTraffic"
}

resource "volcengine_eip_associate" "foo" {
  allocation_id = volcengine_eip_address.foo.id
  instance_id = volcengine_ha_vip.foo.id
  instance_type = "HaVip"
}