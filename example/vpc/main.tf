resource "volcengine_vpc" "foo" {
  vpc_name = "tf-project-1"
  cidr_block = "172.16.0.0/16"
  dns_servers = ["8.8.8.8","114.114.114.114"]
  project_name = "AS_test"
}