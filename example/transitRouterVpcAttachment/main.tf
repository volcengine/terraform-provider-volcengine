resource "volcengine_transit_router" "foo" {
  transit_router_name = "test-tf-acc"
  description         = "test-tf-acc"
}

data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name     = "acc-test-vpc-acc"
  cidr_block   = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  vpc_id      = volcengine_vpc.foo.id
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  subnet_name = "acc-test-subnet"
}

resource "volcengine_subnet" "foo2" {
  vpc_id      = volcengine_vpc.foo.id
  cidr_block  = "172.16.255.0/24"
  zone_id     = data.volcengine_zones.foo.zones[1].id
  subnet_name = "acc-test-subnet2"
}


resource "volcengine_transit_router_vpc_attachment" "foo" {
  transit_router_id = volcengine_transit_router.foo.id
  vpc_id            = volcengine_vpc.foo.id
  attach_points {
    subnet_id = volcengine_subnet.foo.id
    zone_id   = "cn-beijing-a"
  }
  attach_points {
    subnet_id = volcengine_subnet.foo2.id
    zone_id   = "cn-beijing-b"
  }
  transit_router_attachment_name = "tf-test-acc-name1"
  description = "tf-test-acc-description"
}
