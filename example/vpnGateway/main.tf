resource "vestack_vpn_gateway" "foo" {
  vpc_id = "vpc-2bysvq1xx543k2dx0eeulpeiv"
  subnet_id = "subnet-2d68bh74345q858ozfekrm8fj"
  bandwidth = 20
  vpn_gateway_name = "tf-test"
  description = "tf-test"
  period = 2
}