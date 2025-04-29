resource "volcengine_vpc" "foo" {
  vpc_name     = "acc-test-vpc"
  cidr_block   = "172.16.0.0/16"
  project_name = "default"
}

resource "volcengine_route_table" "foo" {
  vpc_id           = volcengine_vpc.foo.id
  route_table_name = "acc-test-route-table"
  description      = "tf-test"
  project_name     = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}
