resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_cfw_dns_control_policy" "foo" {
  description      = "acc-test-dns-control-policy"
  destination_type = "domain"
  destination      = "www.test.com"
  source {
    vpc_id = volcengine_vpc.foo.id
    region = "cn-beijing"
  }
}
