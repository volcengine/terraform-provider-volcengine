resource "vestack_vpc" "foo" {
  vpc_name = "tf-test-2"
  cidr_block = "172.16.0.0/16"
  dns_servers = ["8.8.8.8","114.114.114.114"]
}