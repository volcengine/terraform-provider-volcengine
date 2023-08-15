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

resource "volcengine_scaling_group" "foo" {
  count                     = 3
  scaling_group_name        = "acc-test-scaling-group-${count.index}"
  subnet_ids                = [volcengine_subnet.foo.id]
  multi_az_policy           = "BALANCE"
  desire_instance_number    = 0
  min_instance_number       = 0
  max_instance_number       = 10
  instance_terminate_policy = "OldestInstance"
  default_cooldown          = 30

  tags {
    key   = "k2"
    value = "v2"
  }

  tags {
    key   = "k1"
    value = "v1"
  }
}