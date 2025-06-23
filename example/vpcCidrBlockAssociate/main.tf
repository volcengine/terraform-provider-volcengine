resource "volcengine_vpc" "foo" {
  vpc_name     = "acc-test-vpc"
  cidr_block   = "192.168.0.0/20"
  project_name = "default"
}

resource "volcengine_vpc_cidr_block_associate" "foo" {
  vpc_id               = volcengine_vpc.foo.id
  secondary_cidr_block = "192.168.16.0/20"
}