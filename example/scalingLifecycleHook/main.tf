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
  scaling_group_name        = "acc-test-scaling-group-lifecycle"
  subnet_ids                = [volcengine_subnet.foo.id]
  multi_az_policy           = "BALANCE"
  desire_instance_number    = 0
  min_instance_number       = 0
  max_instance_number       = 1
  instance_terminate_policy = "OldestInstance"
  default_cooldown          = 10
}

resource "volcengine_scaling_lifecycle_hook" "foo" {
  lifecycle_hook_name    = "acc-test-lifecycle"
  lifecycle_hook_policy  = "CONTINUE"
  lifecycle_hook_timeout = 30
  lifecycle_hook_type    = "SCALE_IN"
  scaling_group_id       = volcengine_scaling_group.foo.id
}