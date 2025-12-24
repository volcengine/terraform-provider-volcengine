resource "volcengine_vpc" "foo" {
  vpc_name     = "acc-test-vpc"
  cidr_block   = "10.200.0.0/16"
  project_name = "default"
}

resource "volcengine_vpc_user_cidr_block_associate" "foo1" {
  vpc_id          = volcengine_vpc.foo.id
  user_cidr_block = "10.201.0.0/16"
}

resource "volcengine_vpc_user_cidr_block_associate" "foo2" {
  vpc_id          = volcengine_vpc.foo.id
  user_cidr_block = "10.202.0.0/16"
}

resource "volcengine_vpc_user_cidr_block_associate" "foo3" {
  vpc_id          = volcengine_vpc.foo.id
  user_cidr_block = "10.203.0.0/16"
}
