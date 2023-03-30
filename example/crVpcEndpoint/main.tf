resource "volcengine_cr_vpc_endpoint" "foo" {
  registry = "enterprise-1"
  vpcs {
    vpc_id = "vpc-3resbfzl3xgjk5zsk2iuq3vhk"
    account_id = 000000
  }
  vpcs {
    vpc_id = "vpc-3red9li8dd8g05zsk2iadytvy"
    subnet_id = "subnet-2d62do4697i8058ozfdszxl30"
  }

}