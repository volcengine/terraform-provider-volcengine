resource "volcengine_subnet" "foo" {
  subnet_name = "subnet-test-2"
  cidr_block = "192.168.1.0/24"
  zone_id = "cn-beijing"
  vpc_id = "vpc-2749wnlhro3y87fap8u5ztvt5"
}